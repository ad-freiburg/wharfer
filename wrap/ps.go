package wrap

import (
	"flag"
)

type Ps struct {
	Cmd *flag.FlagSet
}

func (ps *Ps) InitFlags() {
	ps.Cmd = flag.NewFlagSet("ps", flag.ExitOnError)
}

func (ps *Ps) ParseToArgs(rawArgs []string) []string {
	ps.Cmd.Parse(rawArgs)
	args := []string{"ps"}
	return args
}
