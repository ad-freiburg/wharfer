package wrap

import (
	"flag"
	"os"
)

type NetworkList struct {
	Cmd    *flag.FlagSet
	Format string
	Filter string
	Quiet  bool
}

func (c *NetworkList) InitFlags() {
	const filterUsage = "Filter output based on conditions provided"
	c.Cmd = flag.NewFlagSet("network ls", flag.ExitOnError)
	c.Cmd.StringVar(&c.Filter, "filter", "", filterUsage)
	c.Cmd.StringVar(&c.Filter, "f", "", filterUsage+" (shorthand)")
	c.Cmd.StringVar(&c.Format, "format", "", "Pretty-print images using a Go template")
	c.Cmd.BoolVar(&c.Quiet, "q", false, "Only display numeric IDs")
}

func (c *NetworkList) ParseToArgs(rawArgs []string) []string {
	if err := c.Cmd.Parse(rawArgs); err != nil {
		// Only returns an error if the Usage was shown
		os.Exit(0)
	}
	args := []string{"network", "ls"}
	if c.Filter != "" {
		args = append(args, "--filter", c.Filter)
	}
	if c.Format != "" {
		args = append(args, "--format", c.Format)
	}
	if c.Quiet {
		args = append(args, "-q")
	}
	return args
}
