package wrap

import (
	"flag"
	"os"
)

type Kill struct {
	Cmd *flag.FlagSet
}

func (kill *Kill) InitFlags() {
	kill.Cmd = flag.NewFlagSet("kill", flag.ExitOnError)
}

func (kill *Kill) ParseToArgs(rawArgs []string) []string {
	if err := kill.Cmd.Parse(rawArgs); err != nil {
		// Only returns an error if the Usage was shown
		os.Exit(0)
	}
	args := []string{"kill"}
	if kill.Cmd.NArg() > 0 {
		// add -- to make sure additional arguments are not interpreted as
		// potentially harmful flags. Here this is the container to kill
		args = append(args, "--")
		for _, arg := range kill.Cmd.Args() {
			if !IsHexOnly(arg) {
				arg = PrependUsername(arg)
			}
			args = append(args, arg)
		}
	}
	return args
}
