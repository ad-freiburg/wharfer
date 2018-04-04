package wrap

import (
	"flag"
)

type Logs struct {
	Cmd    *flag.FlagSet
	Follow bool
}

func (logs *Logs) InitFlags() {
	logs.Cmd = flag.NewFlagSet("logs", flag.ExitOnError)
	logs.Cmd.BoolVar(&logs.Follow, "f", false, "Follow output")
}

func (logs *Logs) ParseToArgs(rawArgs []string) []string {
	logs.Cmd.Parse(rawArgs)
	args := []string{"logs"}
	if logs.Follow {
		args = append(args, "-f")
	}

	if logs.Cmd.NArg() > 0 {
		// add -- to make sure additional arguments are not interpreted as
		// potentially harmful flags. Here this is the container to logs
		args = append(args, "--")
		args = append(args, logs.Cmd.Args()...)
	}
	return args
}
