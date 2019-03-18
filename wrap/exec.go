package wrap

import (
	"flag"
	"fmt"
	"os"
)

type Exec struct {
	Cmd            *flag.FlagSet
	Env            StringSliceFlag
	Detach         bool
	InteractiveTTY bool
}

func (exec *Exec) InitFlags() {
	exec.Cmd = flag.NewFlagSet("exec", flag.ExitOnError)
	exec.Cmd.Var(&exec.Env, "e", "Add an environment variable mapping of the form \"VARIABLE=value\"")
	exec.Cmd.BoolVar(&exec.Detach, "d", false, "Detach container after starting it. Disables interactive mode")
	exec.Cmd.BoolVar(&exec.InteractiveTTY, "it", false, "Run container interactively")
}

func (exec *Exec) ParseToArgs(rawArgs []string) []string {
	exec.Cmd.Parse(rawArgs)
	args := []string{"exec"}

	// Always set --user $(id -u):$(id -g) so that when running without user
	// namespaces we at least execute as the current user
	args = appendCurrentUserArgs(args)

	if exec.Detach {
		args = append(args, "-d")
	}

	if exec.InteractiveTTY && !exec.Detach {
		args = append(args, "-it")
	}

	if len(exec.Env) > 0 {
		for _, execEnv := range exec.Env {
			args = append(args, "-e", execEnv)
		}
	}

	argsLeft := exec.Cmd.NArg()
	if argsLeft < 2 {
		fmt.Fprintln(os.Stderr, "Missing [CONTAINER] and/or [CMD] arguments")
		os.Exit(3)
	}
	// add -- to make sure additional arguments are not interpreted as
	// potentially harmful flags. Here these are args for the entrypoint.
	args = append(args, "--")

	// The [CONTAINER] positional argument
	container := exec.Cmd.Args()[0]
	argsLeft--
	if !IsHexOnly(container) {
		container = PrependUsername(container)
	}
	args = append(args, container)

	// [COMMAND] argument
	command := exec.Cmd.Args()[1]
	args = append(args, command)
	argsLeft--

	// optional arguments
	args = append(args, exec.Cmd.Args()[2:]...)
	return args
}
