package wrap

import (
	"flag"
	"os"
)

type Build struct {
	Cmd       *flag.FlagSet
	Tag       string
	NoCache   bool
	BuildArgs StringSliceFlag
	File      string
	Pull      bool
}

func (build *Build) InitFlags() {
	const tagUsage = "Tag of the image in name:tag format"
	const fileUsage = "Explicit Dockerfile to use"
	build.Cmd = flag.NewFlagSet("build", flag.ExitOnError)
	build.Cmd.StringVar(&build.Tag, "tag", "", tagUsage)
	build.Cmd.StringVar(&build.Tag, "t", "", tagUsage+" (shorthand)")
	build.Cmd.StringVar(&build.File, "file", "", fileUsage)
	build.Cmd.StringVar(&build.File, "f", "", fileUsage+" (shorthand)")
	build.Cmd.Var(&build.BuildArgs, "build-arg", "Set build-time variables")
	build.Cmd.BoolVar(&build.NoCache, "no-cache", false, "Disable caching to rebuild from scratch")
	build.Cmd.BoolVar(&build.NoCache, "pull", false, "Always attempt to pull a newer version of the image")
}

func (build *Build) ParseToArgs(rawArgs []string) []string {
	if err := build.Cmd.Parse(rawArgs); err != nil {
		// Only returns an error if the Usage was shown
		os.Exit(0)
	}
	args := []string{"build"}
	if build.File != "" {
		args = append(args, "--file", build.File)
	}

	if build.Tag != "" {
		args = append(args, "--tag", build.Tag)
	}

	if build.NoCache {
		args = append(args, "--no-cache")
	}

	if build.Pull {
		args = append(args, "--pull")
	}

	if len(build.BuildArgs) > 0 {
		for _, buildArg := range build.BuildArgs {
			args = append(args, "--build-arg", buildArg)
		}
	}

	if build.Cmd.NArg() > 0 {
		// add -- to make sure additional arguments are not interpreted as
		// potentially harmful flags. Here this is the PATH
		args = append(args, "--")
		args = append(args, build.Cmd.Args()...)
	}
	return args
}
