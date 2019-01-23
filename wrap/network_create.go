package wrap

import (
	"flag"
)

type NetworkCreate struct {
	Cmd *flag.FlagSet
}

func (c *NetworkCreate) InitFlags() {
	c.Cmd = flag.NewFlagSet("create", flag.ExitOnError)
}

func (c *NetworkCreate) ParseToArgs(rawArgs []string) []string {
	c.Cmd.Parse(rawArgs)
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
