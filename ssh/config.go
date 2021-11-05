package ssh

import "strings"

type ClientConfig struct {
	User     string
	HostName string
	Host     string
	Port     string
}

func (c *ClientConfig) String() string {
	sb := &strings.Builder{}
	sb.WriteString("Host ")
	sb.WriteString(c.Host)
	sb.WriteString("\n")
	sb.WriteString("  HostName ")
	sb.WriteString(c.HostName)
	sb.WriteString("\n")
	if c.User != "" {
		sb.WriteString("  User ")
		sb.WriteString(c.User)
		sb.WriteString("\n")
	}
	return sb.String()
}
