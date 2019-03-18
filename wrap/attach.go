package wrap

import (
	"flag"
)

type Attach struct {
	Cmd *flag.FlagSet
}

func (c *Attach) InitFlags() {
	c.Cmd = flag.NewFlagSet("attach", flag.ExitOnError)
}

func (c *Attach) ParseToArgs(rawArgs []string) []string {
	c.Cmd.Parse(rawArgs)
	args := []string{"attach"}

	if c.Cmd.NArg() > 0 {
		// add -- to make sure additional arguments are not interpreted as
		// potentially harmful flags. Here this is the container to rm
		args = append(args, "--")
		for _, arg := range c.Cmd.Args() {
			// If the container is given by name we prepend the user name
			if !IsHexOnly(arg) {
				arg = PrependUsername(arg)
				args = append(args, arg)
			}
		}
	}
	return args
}
