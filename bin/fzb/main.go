package main

import (
	"flag"
	"fmt"
	"github.com/paulvollmer/fzb/src/go"
	"github.com/paulvollmer/go-verbose"
	"io/ioutil"
	"os"
)

// fzb commandline tool
// this commandline tool can be used to validate, format and explore fritzing fzb files.
func main() {
	flagFile := flag.String("f", "", "validate a file")
	flagDir := flag.String("d", "", "validate a directory")
	flagVerbose := flag.Bool("v", false, "verbose mode")
	flag.Parse()

	debug := verbose.New(os.Stdout, *flagVerbose)

	if *flagFile != "" {
		debug.Println("Read Fzb", *flagFile)
		processFile(*flagFile)
	}

	if *flagDir != "" {
		debug.Println("Read Folder", *flagDir)
		processDir(*flagDir)
	}

}
func processDir(src string) error {

	d, err := ioutil.ReadDir(src)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	totalFiles := len(d)
	fmt.Println("total files", totalFiles)
	for _, v := range d {
		tmpfilepath := src + "/" + v.Name()
		// fmt.Println("tmpfilepath", tmpfilepath)
		err := processFile(tmpfilepath)
		if err != nil {
			fmt.Println("\n", v.Name())
			fmt.Println(err)
		}
	}
	return err
}

func processFile(src string) error {
	fzbData, err := fzb.ReadFile(src)
	if err != nil {
		fmt.Println(err)
		return err
	}

	// fmt.Println("Check Data...")
	err, warn := fzbData.Check()
	if warn == "" && err == nil {
		fmt.Println(src, "is valid")
		return nil
	}

	if warn != "" {
		fmt.Print(warn)
	}
	if err != nil {
		fmt.Print(err)
	}
	return err
}
