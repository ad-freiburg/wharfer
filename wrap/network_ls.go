package wrap

import (
	"flag"
	"os"
)

type NetworkList struct {
	Cmd *flag.FlagSet
}

func (c *NetworkList) InitFlags() {
	c.Cmd = flag.NewFlagSet("create", flag.ExitOnError)
}

func (c *NetworkList) ParseToArgs(rawArgs []string) []string {
	if err := c.Cmd.Parse(rawArgs); err != nil {
		// Only returns an error if the Usage was shown
		os.Exit(0)
	}
	args := []string{"network", "ls"}
	return args
}
