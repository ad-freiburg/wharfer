package wrap

import (
	"flag"
	"fmt"
	"os"
)

type Attach struct {
	Cmd *flag.FlagSet
}

func (attach *Attach) InitFlags() {
	attach.Cmd = flag.NewFlagSet("attach", flag.ExitOnError)
}

func (attach *Attach) ParseToArgs(rawArgs []string) []string {
	attach.Cmd.Parse(rawArgs)
	args := []string{"attach"}

	if attach.Cmd.NArg() < 1 {
		fmt.Fprintln(os.Stderr, "Missing [CONTAINER] argument")
		os.Exit(3)
	}

	// add -- to make sure additional arguments are not interpreted as
	// potentially harmful flags. Here these are args for the entrypoint.
	args = append(args, "--")

	// The [CONTAINER] positional argument
	container := attach.Cmd.Args()[0]
	// If the container is given by name we prepend the user name
	if !IsHexOnly(container) {
		container = PrependUsername(container)
	}
	args = append(args, container)
	return args
}
