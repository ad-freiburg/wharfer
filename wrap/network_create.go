package wrap

import (
	"flag"
	"os"
)

type NetworkCreate struct {
	Cmd *flag.FlagSet
}

func (c *NetworkCreate) InitFlags() {
	c.Cmd = flag.NewFlagSet("network create", flag.ExitOnError)
}

func (c *NetworkCreate) ParseToArgs(rawArgs []string) []string {
	if err := c.Cmd.Parse(rawArgs); err != nil {
		// Only returns an error if the Usage was shown
		os.Exit(0)
	}
	args := []string{"network", "create"}
	if c.Cmd.NArg() > 0 {
		// add -- to make sure additional arguments are not interpreted as
		// potentially harmful flags.
		args = append(args, "--")
		for _, arg := range c.Cmd.Args() {
			args = append(args, PrependUsername(arg))
		}
	}
	return args
}
