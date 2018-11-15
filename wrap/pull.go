package wrap

import (
	"flag"
)

type Pull struct {
	Cmd *flag.FlagSet
}

func (pull *Pull) InitFlags() {
	pull.Cmd = flag.NewFlagSet("pull", flag.ExitOnError)
}

func (pull *Pull) ParseToArgs(rawArgs []string) []string {
	pull.Cmd.Parse(rawArgs)
	args := []string{"pull"}

	if pull.Cmd.NArg() > 0 {
		// add -- to make sure additional arguments are not interpreted as
		// potentially harmful flags. Here this is the PATH
		args = append(args, "--")
		args = append(args, pull.Cmd.Args()...)
	}
	return args
}
