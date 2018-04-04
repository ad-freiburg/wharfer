package wrap

import (
	"flag"
)

type Kill struct {
	Cmd *flag.FlagSet
}

func (kill *Kill) InitFlags() {
	kill.Cmd = flag.NewFlagSet("kill", flag.ExitOnError)
}

func (kill *Kill) ParseToArgs(rawArgs []string) []string {
	kill.Cmd.Parse(rawArgs)
	args := []string{"kill"}
	if kill.Cmd.NArg() > 0 {
		// add -- to make sure additional arguments are not interpreted as
		// potentially harmful flags. Here this is the container to kill
		args = append(args, "--")
		args = append(args, kill.Cmd.Args()...)
	}
	return args
}
