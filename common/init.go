package common

import (
	"flag"
	"os"
)

var (
	PrintVersion = flag.Bool("version", false, "Print version and exit")
	PrintHelp    = flag.Bool("help", false, "Print help and exit")
	ConfigFile   = flag.String("config", "go-relay.yaml", "Config file path")
)

func printUsage() {
	println("Go Public " + Version + " - A simple HTTP relay server.")
	println("Copyright (C) 2023 JustSong. All rights reserved.")
	println("GitHub: https://github.com/songquanpeng/go-relay")
	println("Usage: go-relay [--config <config file path>] [--version] [--help]")
	println("       go-relay init")
}

func init() {
	flag.Parse()
	if *PrintVersion {
		println(Version)
		os.Exit(0)
	}
	if *PrintHelp {
		printUsage()
		os.Exit(0)
	}
	if len(os.Args) > 1 && os.Args[1] == "init" {
		initConfigFile()
		os.Exit(0)
	}
	loadConfigFile()
}
