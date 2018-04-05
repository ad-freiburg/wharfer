package main

import (
	"fmt"
	"github.com/ad-freiburg/wharfer/wrap"
	"log"
	"os"
	"os/exec"
	"strings"
)

func execDocker(args ...string) {
	const dockerbin = "/usr/bin/docker"
	cmd := exec.Command(dockerbin, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

var build wrap.Build
var run wrap.Run
var ps wrap.Ps
var kill wrap.Kill
var rm wrap.Rm
var logs wrap.Logs

func init() {
	build.InitFlags()
	run.InitFlags()
	ps.InitFlags()
	kill.InitFlags()
	logs.InitFlags()
	rm.InitFlags()
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage:", os.Args[0], "COMMAND")
		fmt.Fprintln(os.Stderr, "Commands:")
		fmt.Fprintln(os.Stderr, "\tbuild")
		fmt.Fprintln(os.Stderr, "\trun")
		os.Exit(1)
	}

	var args []string
	switch os.Args[1] {
	case "build":
		args = build.ParseToArgs(os.Args[2:])
	case "run":
		args = run.ParseToArgs(os.Args[2:])
	case "kill":
		args = kill.ParseToArgs(os.Args[2:])
	case "rm":
		args = rm.ParseToArgs(os.Args[2:])
	case "logs":
		args = logs.ParseToArgs(os.Args[2:])
	case "ps":
		args = ps.ParseToArgs(os.Args[2:])
	default:
		fmt.Fprintln(os.Stderr, "Unknown subcommand", os.Args[1])
		os.Exit(2)
	}

	execDocker(args...)
}
