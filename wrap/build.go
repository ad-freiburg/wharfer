package wrap

import (
	"flag"
)

type Build struct {
	Cmd     *flag.FlagSet
	Tag     string
	NoCache bool
}

func (build *Build) InitFlags() {
	const tagUsage = "Tag of the image in name:tag format"
	build.Cmd = flag.NewFlagSet("build", flag.ExitOnError)
	build.Cmd.StringVar(&build.Tag, "tag", "", tagUsage)
	build.Cmd.StringVar(&build.Tag, "t", "", tagUsage+" (shorthand)")
	build.Cmd.BoolVar(&build.NoCache, "no-cache", false, "Disable caching to rebuild from scratch")
}

func (build *Build) ParseToArgs(rawArgs []string) []string {
	build.Cmd.Parse(rawArgs)
	args := []string{"build"}
	if build.Tag != "" {
		args = append(args, "--tag", build.Tag)
	}

	if build.NoCache {
		args = append(args, "--no-cache")
	}

	if build.Cmd.NArg() > 0 {
		// add -- to make sure additional arguments are not interpreted as
		// potentially harmful flags. Here this is the PATH
		args = append(args, "--")
		args = append(args, build.Cmd.Args()...)
	}
	return args
}
