// fzb commandline tool
// this commandline tool can be used to validate, format and explore fritzing fzb files.

package main

import (
	"flag"
	"fmt"
	"github.com/paulvollmer/fzb/src/go"
	"os"
)

func main() {
	flagFile := flag.String("f", "", "the file")
	flag.Parse()

	// fmt.Println("Read Fzb", *flagFile)
	fzbData, err := fzb.ReadFile(*flagFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// fmt.Println("Check Data...")
	err, warn := fzbData.Check()
	if warn == "" && err == nil {
		fmt.Println(*flagFile, "is valid")
	}

	if warn != "" {
		fmt.Print(warn)
	}
	if err != nil {
		fmt.Print(err)
	}

	// fzbData.PrettyPrint()
}
