package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/ad-freiburg/wharfer/wrap"
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

var buildCmd wrap.Build
var runCmd wrap.Run
var psCmd wrap.Ps
var killCmd wrap.Kill
var rmCmd wrap.Rm
var loadCmd wrap.Load
var saveCmd wrap.Save
var logsCmd wrap.Logs
var pullCmd wrap.Pull
var imagesCmd wrap.Images
var networkCreateCmd wrap.NetworkCreate
var networkListCmd wrap.NetworkList
var networkRemoveCmd wrap.NetworkRemove
var attachCmd wrap.Attach
var execCmd wrap.Exec

func init() {
	buildCmd.InitFlags()
	runCmd.InitFlags()
	psCmd.InitFlags()
	killCmd.InitFlags()
	loadCmd.InitFlags()
	saveCmd.InitFlags()
	logsCmd.InitFlags()
	rmCmd.InitFlags()
	pullCmd.InitFlags()
	imagesCmd.InitFlags()
	networkCreateCmd.InitFlags()
	networkListCmd.InitFlags()
	networkRemoveCmd.InitFlags()
	attachCmd.InitFlags()
	execCmd.InitFlags()
}

var version = "no-release"

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage:", os.Args[0], "<COMMAND>|--version")
		fmt.Fprintln(os.Stderr, "Commands:")
		fmt.Fprintln(os.Stderr, "\tbuild")
		fmt.Fprintln(os.Stderr, "\trun")
		fmt.Fprintln(os.Stderr, "\tps")
		fmt.Fprintln(os.Stderr, "\tkill")
		fmt.Fprintln(os.Stderr, "\trm")
		fmt.Fprintln(os.Stderr, "\tload")
		fmt.Fprintln(os.Stderr, "\tsave")
		fmt.Fprintln(os.Stderr, "\tlogs")
		fmt.Fprintln(os.Stderr, "\tpull")
		fmt.Fprintln(os.Stderr, "\timages")
		fmt.Fprintln(os.Stderr, "\tnetwork")
		fmt.Fprintln(os.Stderr, "\tattach")
		fmt.Fprintln(os.Stderr, "\texec")
		os.Exit(1)
	}

	var args []string
	switch os.Args[1] {
	case "build":
		args = buildCmd.ParseToArgs(os.Args[2:])
	case "run":
		args = runCmd.ParseToArgs(os.Args[2:])
	case "kill":
		args = killCmd.ParseToArgs(os.Args[2:])
	case "rm":
		args = rmCmd.ParseToArgs(os.Args[2:])
	case "save":
		args = saveCmd.ParseToArgs(os.Args[2:])
	case "load":
		args = loadCmd.ParseToArgs(os.Args[2:])
	case "logs":
		args = logsCmd.ParseToArgs(os.Args[2:])
	case "ps":
		args = psCmd.ParseToArgs(os.Args[2:])
	case "pull":
		args = pullCmd.ParseToArgs(os.Args[2:])
	case "images":
		args = imagesCmd.ParseToArgs(os.Args[2:])
	case "network":
		if len(os.Args) < 3 {
			fmt.Fprintln(os.Stderr, "Missing subcommand")
			os.Exit(1)
		} else {
			switch os.Args[2] {
			case "create":
				args = networkCreateCmd.ParseToArgs(os.Args[3:])
			case "ls":
				args = networkListCmd.ParseToArgs(os.Args[3:])
			case "rm":
				args = networkRemoveCmd.ParseToArgs(os.Args[3:])
			case "--help":
				fmt.Fprintln(os.Stderr, "Commands:")
				fmt.Fprintln(os.Stderr, "\tcreate")
				fmt.Fprintln(os.Stderr, "\tls")
				fmt.Fprintln(os.Stderr, "\trm")
				os.Exit(1)
			default:
				fmt.Fprintln(os.Stderr, "Unknown subcommand", os.Args[2])
				os.Exit(1)
			}
		}
	case "attach":
		args = attachCmd.ParseToArgs(os.Args[2:])
	case "exec":
		args = execCmd.ParseToArgs(os.Args[2:])
	case "--version":
		fmt.Fprintln(os.Stderr, os.Args[0], "version", version)
		os.Exit(0)
	default:
		fmt.Fprintln(os.Stderr, "Unknown subcommand", os.Args[1])
		os.Exit(2)
	}

	execDocker(args...)
}
