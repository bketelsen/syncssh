package lxd

import (
	"encoding/json"
	"os/exec"
	"strings"

	"github.com/lxc/lxd/shared/api"

	"github.com/bketelsen/syncssh/ssh"
)

func GetHosts() ([]ssh.ClientConfig, error) {
	var configs []ssh.ClientConfig
	cmd := exec.Command("lxc", "ls", "--format", "json")
	bb, err := cmd.CombinedOutput()
	if err != nil {
		return configs, err
	}
	var instances []api.InstanceFull
	err = json.Unmarshal(bb, &instances)
	if err != nil {
		return configs, err
	}
	for _, instance := range instances {
		cmd := exec.Command("lxc", "config", "get", instance.Name, "user.sshuser")
		bb, err := cmd.CombinedOutput()
		if err != nil {
			return configs, err
		}
		user := strings.TrimSpace(string(bb))
		for netName, net := range instance.State.Network {
			if netName == "lo" || netName == "docker0" {
				continue
			}
			if netName == "tailscale0" {
				netName = "tail"
			}
			if netName == "eth0" {
				netName = "bridged"
			}

			for _, addr := range net.Addresses {
				if addr.Family == "inet" {
					configs = append(configs, ssh.ClientConfig{
						Host:     instance.Name + "." + netName,
						HostName: addr.Address,
						User:     user,
					})
				}
			}
		}
	}

	return configs, nil
}
