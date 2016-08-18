package main

import (
	"flag"
	"fmt"
	"github.com/paulvollmer/fzb/src/go"
	"github.com/paulvollmer/go-verbose"
	"os"
)

const (
	version = "0.1.0"
)

// fzb commandline tool
// this commandline tool can be used to validate, format and explore fritzing fzb files.
func main() {
	flagFile := flag.String("f", "", "validate a file")
	flagDir := flag.String("d", "", "validate a directory")
	flagVerbose := flag.Bool("v", false, "verbose mode")
	flagVersion := flag.Bool("version", false, "print version and exit")
	flag.Parse()
	if *flagVersion {
		fmt.Println(version)
		os.Exit(0)
	}

	debug := verbose.New(os.Stdout, *flagVerbose)

	report := ""

	fmt.Println("fzb validate", version)
	if *flagFile != "" {
		debug.Println("Read Fzb", *flagFile)
		report = fzb.ValidateFile(os.Args[0], *flagFile)
		fmt.Println(report)
	}

	if *flagDir != "" {
		debug.Println("Read Folder", *flagDir)
		report = fzb.ValidateDir(*flagDir)
		fmt.Println(report)
	}

	if report != "" {
		os.Exit(1)
	}

	fmt.Println("Fzb Data valid")
}
