package wrap

import (
	"flag"
)

type Images struct {
	Cmd    *flag.FlagSet
	Format string
}

func (images *Images) InitFlags() {
	images.Cmd = flag.NewFlagSet("images", flag.ExitOnError)
	images.Cmd.StringVar(&images.Format, "format", "", "Pretty-print images using a Go template")
}

func (images *Images) ParseToArgs(rawArgs []string) []string {
	images.Cmd.Parse(rawArgs)
	args := []string{"images"}
	if images.Format != "" {
		args = append(args, "--format", images.Format)
	}
	return args
}
