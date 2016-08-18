// fzb commandline tool
// this commandline tool can be used to validate, format and explore fritzing fzb files.

package main

import (
	"fmt"
	"github.com/paulvollmer/fzb/src/go"
	"os"
)

func main() {
	fmt.Println("fzb")

	fzbData, err := fzb.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fzbData.PrettyPrint()
}
