package wrap

import (
	"flag"
	"os"
	"strconv"
)

type Logs struct {
	Cmd    *flag.FlagSet
	Follow bool
	Tail   int
}

func (logs *Logs) InitFlags() {
	logs.Cmd = flag.NewFlagSet("logs", flag.ExitOnError)
	logs.Cmd.BoolVar(&logs.Follow, "f", false, "Follow output")
	logs.Cmd.IntVar(&logs.Tail, "tail", 0, "Only show last <num> lines")
}

func (logs *Logs) ParseToArgs(rawArgs []string) []string {
	if err := logs.Cmd.Parse(rawArgs); err != nil {
		// Only returns an error if the Usage was shown
		os.Exit(0)
	}
	args := []string{"logs"}
	if logs.Follow {
		args = append(args, "-f")
	}

	if logs.Tail > 0 {
		args = append(args, "--tail", strconv.Itoa(logs.Tail))
	}

	if logs.Cmd.NArg() > 0 {
		// add -- to make sure additional arguments are not interpreted as
		// potentially harmful flags. Here this is the container to logs
		args = append(args, "--")
		for _, arg := range logs.Cmd.Args() {
			if !IsHexOnly(arg) {
				arg = PrependUsername(arg)
			}
			args = append(args, arg)
		}
	}
	return args
}
