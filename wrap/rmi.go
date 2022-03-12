package wrap

import (
	"flag"
	"os"
)

type Rmi struct {
	Cmd   *flag.FlagSet
	Force bool
}

func (rmi *Rmi) InitFlags() {
	rmi.Cmd = flag.NewFlagSet("rmi", flag.ExitOnError)
	rmi.Cmd.BoolVar(&rmi.Force, "f", false, "Force removal of the image")
}

func (rmi *Rmi) ParseToArgs(rawArgs []string) []string {
	if err := rmi.Cmd.Parse(rawArgs); err != nil {
		// Only returns an error if the Usage was shown
		os.Exit(0)
	}
	args := []string{"rmi"}

	if rmi.Force {
		args = append(args, "-f")
	}

	if rmi.Cmd.NArg() > 0 {
		// add -- to make sure additional arguments are not interpreted as
		// potentially harmful flags. Here this is the image to rm
		args = append(args, "--")
		for _, arg := range rmi.Cmd.Args() {
			if !IsHexOnly(arg) {
				arg = PrependUsername(arg)
			}
			args = append(args, arg)
		}
	}
	return args
}
