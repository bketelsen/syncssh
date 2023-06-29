/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/bketelsen/syncssh/lxd"
	"github.com/bketelsen/syncssh/ssh"
	"github.com/spf13/cobra"
)

// syncCmd represents the sync command
var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "synchronize configs",
	Long: `syncssh reads the list of LXC instances from your configured
	LXD remote and adds them to ~/.ssh/config.lxd, which you can import into
	your ssh config file.`,
	Run: func(cmd *cobra.Command, args []string) {
		/*fmt.Println("tailscale:", viper.GetBool("tailscale"))
		fmt.Println("lxd:", viper.GetBool("lxd"))
		*/
		hosts, err := lxd.GetHosts()
		if err != nil {
			panic(err)
		}
		writeConfig(hosts)
	},
}

func init() {
	rootCmd.AddCommand(syncCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// syncCmd.PersistentFlags().String("foo", "", "A help for foo")
	/*	syncCmd.PersistentFlags().BoolP("tailscale", "t", false, "sync tailscale hosts")
		syncCmd.PersistentFlags().BoolP("lxd", "l", false, "sync lxd hosts")
		viper.BindPFlag("tailscale", syncCmd.PersistentFlags().Lookup("tailscale"))
		viper.BindPFlag("lxd", syncCmd.PersistentFlags().Lookup("lxd"))
	*/
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// syncCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func writeConfig(hosts []ssh.ClientConfig) error {
	fmt.Println("writing config")
	cfgPath := filepath.Join(os.Getenv("HOME"), ".ssh", "config.lxd")

	return saveConfig(cfgPath, hosts)
}

func saveConfig(p string, nodes []ssh.ClientConfig) error {
	sb := strings.Builder{}
	for _, node := range nodes {
		sb.WriteString(node.String())
	}
	return ioutil.WriteFile(p, []byte(sb.String()), 0644)
}
