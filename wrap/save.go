package wrap

import (
	"flag"
	"os"
)

type Save struct {
	Cmd    *flag.FlagSet
	Output string
}

func (save *Save) InitFlags() {
	const outputUsage = "Write to a file, instead of STDOUT"
	save.Cmd = flag.NewFlagSet("save", flag.ExitOnError)
	save.Cmd.StringVar(&save.Output, "output", "", outputUsage)
	save.Cmd.StringVar(&save.Output, "o", "", outputUsage+" (shorthand)")
}

func (save *Save) ParseToArgs(rawArgs []string) []string {
	if err := save.Cmd.Parse(rawArgs); err != nil {
		// Only returns an error if the Usage was shown
		os.Exit(0)
	}
	args := []string{"save"}
	if save.Output != "" {
		args = append(args, "--output", save.Output)
	}

	if save.Cmd.NArg() > 0 {
		// add -- to make sure additional arguments are not interpreted as
		// potentially harmful flags. Here this is the container to kill
		args = append(args, "--")
		args = append(args, save.Cmd.Args()...)
	}
	return args
}
