package wrap

import (
	"flag"
	"os"
)

type Ps struct {
	Cmd    *flag.FlagSet
	All    bool
	Filter string
	Quiet  bool
}

func (ps *Ps) InitFlags() {
	const allUsage = "Show all containers (default shows just running)"
	const filterUsage = "Filter output based on conditions provided"
	ps.Cmd = flag.NewFlagSet("ps", flag.ExitOnError)
	ps.Cmd.BoolVar(&ps.All, "all", false, allUsage)
	ps.Cmd.BoolVar(&ps.All, "a", false, allUsage+" (shorthand)")
	ps.Cmd.StringVar(&ps.Filter, "filter", "", filterUsage)
	ps.Cmd.StringVar(&ps.Filter, "f", "", filterUsage+" (shorthand)")
	ps.Cmd.BoolVar(&ps.Quiet, "q", false, "Only display numeric IDs")
}

func (ps *Ps) ParseToArgs(rawArgs []string) []string {
	if err := ps.Cmd.Parse(rawArgs); err != nil {
		// Only returns an error if the Usage was shown
		os.Exit(0)
	}
	args := []string{"ps"}
	if ps.All {
		args = append(args, "-a")
	}
	if ps.Filter != "" {
		args = append(args, "--filter", ps.Filter)
	}
	if ps.Quiet {
		args = append(args, "-q")
	}
	return args
}
