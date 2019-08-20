package wrap

import (
	"flag"
	"os"
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
	if err := images.Cmd.Parse(rawArgs); err != nil {
		// Only returns an error if the Usage was shown
		os.Exit(0)
	}
	args := []string{"images"}
	if images.Format != "" {
		args = append(args, "--format", images.Format)
	}

	if images.Cmd.NArg() > 0 {
		// add -- to make sure additional arguments are not interpreted as
		// potentially harmful flags. Here this is the name of an image.
		args = append(args, "--")
		args = append(args, images.Cmd.Args()...)
	}
	return args
}
