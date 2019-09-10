package wrap

import (
	"flag"
	"os"
)

type NetworkRemove struct {
	Cmd *flag.FlagSet
}

func (c *NetworkRemove) InitFlags() {
	c.Cmd = flag.NewFlagSet("network rm", flag.ExitOnError)
}

func (c *NetworkRemove) ParseToArgs(rawArgs []string) []string {
	if err := c.Cmd.Parse(rawArgs); err != nil {
		// Only returns an error if the Usage was shown
		os.Exit(0)
	}
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
