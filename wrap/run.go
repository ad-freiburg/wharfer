package wrap

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Run struct {
	Cmd            *flag.FlagSet
	Ports          StringSliceFlag
	Volumes        StringSliceFlag
	Remove         bool
	Detach         bool
	Init           bool
	InteractiveTTY bool
	Name           string
	RestartPolicy  string
}

func (run *Run) InitFlags() {
	run.Cmd = flag.NewFlagSet("run", flag.ExitOnError)
	run.Cmd.Var(&run.Volumes, "v", "Volume to bind in container /host:/container format (must be absolute)")
	run.Cmd.Var(&run.Ports, "p", "Port to expose in hostPort:containerPort format")
	run.Cmd.BoolVar(&run.Remove, "rm", false, "Remove container after running it")
	run.Cmd.BoolVar(&run.Init, "init", false, "Use docker-init enabling the reaping of zombie processes")
	run.Cmd.BoolVar(&run.Detach, "d", false, "Detach container after starting it. Disables interactive mode")
	run.Cmd.BoolVar(&run.InteractiveTTY, "it", false, "Run container interactively (default unlike in docker)")
	run.Cmd.StringVar(&run.Name, "name", "", "Name of the running container instance")
	run.Cmd.StringVar(&run.RestartPolicy, "restart", "", "Restart policy e.g. 'unless-stopped'")
}

func (run *Run) ParseToArgs(rawArgs []string) []string {
	run.Cmd.Parse(rawArgs)
	args := []string{"run"}
	if run.Name != "" {
		args = append(args, "--name", run.Name)
	}

	if run.RestartPolicy != "" {
		args = append(args, "--restart", run.RestartPolicy)
	}

	if run.Init {
		args = append(args, "--init")
	}

	if run.Remove {
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

	if run.Cmd.NArg() > 0 {
		// add -- to make sure additional arguments are not interpreted as
		// potentially harmful flags. Here this is the container name
		args = append(args, "--")
		args = append(args, run.Cmd.Args()...)
	}
	return args
}
