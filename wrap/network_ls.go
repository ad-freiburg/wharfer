package wrap

import (
	"flag"
)

type NetworkList struct {
	Cmd *flag.FlagSet
}

func (c *NetworkList) InitFlags() {
	c.Cmd = flag.NewFlagSet("create", flag.ExitOnError)
}

func (c *NetworkList) ParseToArgs(rawArgs []string) []string {
	c.Cmd.Parse(rawArgs)
	args := []string{"network", "ls"}
	return args
}
