package wrap

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/docker/docker/pkg/namesgenerator"
)

type Run struct {
	Cmd            *flag.FlagSet
	Ports          StringSliceFlag
	Volumes        StringSliceFlag
	Env            StringSliceFlag
	EntryPoint     string
	NoRemove       bool
	Detach         bool
	Init           bool
	InteractiveTTY bool
	Name           string
	RestartPolicy  string
}

func (run *Run) InitFlags() {
	var removeDummy bool
	run.Cmd = flag.NewFlagSet("run", flag.ExitOnError)
	run.Cmd.Var(&run.Volumes, "v", "Volume to bind in container /host:/container format (must be absolute)")
	run.Cmd.Var(&run.Env, "e", "Add an environment variable mapping of the form \"VARIABLE=value\"")
	run.Cmd.Var(&run.Ports, "p", "Port to expose in hostPort:containerPort format")
	run.Cmd.BoolVar(&removeDummy, "rm", true, "Remove container after exit (no-op on wharfer since it's the default)")
	run.Cmd.BoolVar(&run.NoRemove, "no-rm", false, "Do not remove container after exit")
	run.Cmd.BoolVar(&run.Init, "init", true, "Use docker-init enabling the reaping of zombie processes (default unlike in docker)")
	run.Cmd.BoolVar(&run.Detach, "d", false, "Detach container after starting it. Disables interactive mode")
	run.Cmd.BoolVar(&run.InteractiveTTY, "it", false, "Run container interactively")
	run.Cmd.StringVar(&run.Name, "name", "", "Name of the running container instance")
	run.Cmd.StringVar(&run.RestartPolicy, "restart", "", "Restart policy e.g. 'unless-stopped'")
	run.Cmd.StringVar(&run.EntryPoint, "entrypoint", "", "Override the default ENTRYPOINT")
}

func (run *Run) ParseToArgs(rawArgs []string) []string {
	run.Cmd.Parse(rawArgs)
	args := []string{"run"}
	name := namesgenerator.GetRandomName(0)
	if run.Name != "" {
		name = run.Name
	}
	args = append(args, "--name", PrependUsername(name))

	if run.RestartPolicy != "" {
		args = append(args, "--restart", run.RestartPolicy)
		run.NoRemove = true
	}

	if run.EntryPoint != "" {
		args = append(args, "--entrypoint", run.EntryPoint)
	}

	if run.Init {
		args = append(args, "--init")
	}

	if !run.NoRemove {
		args = append(args, "--rm")
	}

	if run.Detach {
		args = append(args, "-d")
	}

	if run.InteractiveTTY && !run.Detach {
		args = append(args, "-it")
	}

	if len(run.Ports) > 0 {
		validPort := regexp.MustCompile(`^[0-9]+:[0-9]+$`)
		for _, runPort := range run.Ports {
			if !validPort.MatchString(runPort) {
				fmt.Fprintln(os.Stderr, "Only numeric <hostPort>:<containerPort> allowed")
				os.Exit(3)
			}
			args = append(args, "-p", runPort)
		}
	}

	if len(run.Volumes) > 0 {
		for _, runVolume := range run.Volumes {
			if !strings.HasPrefix(runVolume, "/") {
				fmt.Fprintln(os.Stderr, "Only absolute paths and host directories are allowed.")
				fmt.Fprintln(os.Stderr, "TIP: use `-v $(pwd)/dir:/dir` to mount a subdirectory of your working directory")
				os.Exit(3)
			}
			args = append(args, "-v", runVolume)
		}
	}

	if len(run.Env) > 0 {
		for _, runEnv := range run.Env {
			args = append(args, "-e", runEnv)
		}
	}

	if run.Cmd.NArg() > 0 {
		// add -- to make sure additional arguments are not interpreted as
		// potentially harmful flags. Here this is the container name
		args = append(args, "--")
		args = append(args, run.Cmd.Args()...)
	}
	return args
}
