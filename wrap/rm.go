package wrap

import (
	"flag"
)

type Rm struct {
	Cmd *flag.FlagSet
}

func (rm *Rm) InitFlags() {
	rm.Cmd = flag.NewFlagSet("rm", flag.ExitOnError)
}

func (rm *Rm) ParseToArgs(rawArgs []string) []string {
	rm.Cmd.Parse(rawArgs)
	args := []string{"rm"}
	if rm.Cmd.NArg() > 0 {
		// add -- to make sure additional arguments are not interpreted as
		// potentially harmful flags. Here this is the container to rm
		args = append(args, "--")
		args = append(args, rm.Cmd.Args()...)
	}
	return args
}
