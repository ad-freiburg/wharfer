package wrap

import (
	"flag"
)

type Ps struct {
	Cmd    *flag.FlagSet
	Quiet  bool
	Filter string
}

func (ps *Ps) InitFlags() {
	const filterUsage = "Filter output based on conditions provided"
	ps.Cmd = flag.NewFlagSet("ps", flag.ExitOnError)
	ps.Cmd.StringVar(&ps.Filter, "filter", "", filterUsage)
	ps.Cmd.StringVar(&ps.Filter, "f", "", filterUsage+" (shorthand)")
	ps.Cmd.BoolVar(&ps.Quiet, "q", false, "Only display numeric IDs")
}

func (ps *Ps) ParseToArgs(rawArgs []string) []string {
	ps.Cmd.Parse(rawArgs)
	args := []string{"ps"}
	if ps.Quiet {
		args = append(args, "-q")
	}
	if ps.Filter != "" {
		args = append(args, "--filter", ps.Filter)
	}
	return args
}
