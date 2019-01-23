package wrap

import (
	"flag"
)

type NetworkRemove struct {
	Cmd *flag.FlagSet
}

func (c *NetworkRemove) InitFlags() {
	c.Cmd = flag.NewFlagSet("rm", flag.ExitOnError)
}

func (c *NetworkRemove) ParseToArgs(rawArgs []string) []string {
	c.Cmd.Parse(rawArgs)
	args := []string{"network", "rm"}
	if c.Cmd.NArg() > 0 {
		// add -- to make sure additional arguments are not interpreted as
		// potentially harmful flags.
		args = append(args, "--")
		for _, arg := range c.Cmd.Args() {
			if !IsHexOnly(arg) {
				arg = PrependUsername(arg)
			}
			args = append(args, arg)
		}
	}
	return args
}
