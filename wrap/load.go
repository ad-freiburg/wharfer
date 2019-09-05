package wrap

import (
	"flag"
	"os"
)

type Load struct {
	Cmd   *flag.FlagSet
	Input string
	Quiet bool
}

func (load *Load) InitFlags() {
	const inputUsage = "Read from tar archive file, instead of STDIN"
	load.Cmd = flag.NewFlagSet("load", flag.ExitOnError)
	load.Cmd.StringVar(&load.Input, "input", "", inputUsage)
	load.Cmd.StringVar(&load.Input, "i", "", inputUsage+" (shorthand)")
	load.Cmd.BoolVar(&load.Quiet, "quiet", false, "Supress output")
	load.Cmd.BoolVar(&load.Quiet, "q", false, "Supress output (shorthand)")
}

func (load *Load) ParseToArgs(rawArgs []string) []string {
	if err := load.Cmd.Parse(rawArgs); err != nil {
		// Only returns an error if the Usage was shown
		os.Exit(0)
	}
	args := []string{"load"}
	if load.Input != "" {
		args = append(args, "--input", load.Input)
	}

	if load.Quiet {
		args = append(args, "--quiet")
	}

	if load.Cmd.NArg() > 0 {
		// add -- to make sure additional arguments are not interpreted as
		// potentially harmful flags. Here this is the PATH
		args = append(args, "--")
		args = append(args, load.Cmd.Args()...)
	}
	return args
}
