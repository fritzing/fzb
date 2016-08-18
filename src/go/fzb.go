package fzb

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	// "os"
	"path/filepath"
	"strings"
)

const (
	MinimumInstances = 4
)

// the fritzing fzb format go utility
type Fzb struct {
	XMLName         xml.Name   `xml:"module"                 json:"-"                 yaml:"-"`
	Title           string     `xml:"title"                  json:"title"             yaml:"title"`
	Icon            string     `xml:"icon,attr"              json:"icon"              yaml:"icon"`
	FritzingVersion string     `xml:"fritzingVersion,attr"   json:"fritzingVersion"   yaml:"fritzingVersion"`
	Instances       []Instance `xml:"instances>instance"     json:"instances"         yaml:"instances"`
}

// NewFzb return a new Fzb object
func NewFzb() Fzb {
	f := Fzb{}
	f.Instances = make([]Instance, 0)
	return f
}

// PrettyPrint the data to stdout
func (f *Fzb) PrettyPrint() {
	fmt.Printf("Title           = %q\n", f.Title)
	fmt.Printf("Icon            = %q\n", f.Icon)
	fmt.Printf("FritzingVersion = %q\n", f.FritzingVersion)
	totalInstances := f.TotalInstances()
	for k, v := range f.Instances {
		fmt.Printf("\nInstance %v of %v\n", k, totalInstances)
		v.PrettyPrint()
	}
}

// ReadFile read the given file and return a Fzb object
func ReadFile(src string) (Fzb, error) {
	fzbBytes, err := ioutil.ReadFile(src)
	if err != nil {
		return Fzb{}, err
	}
	fzbData, err := UnmarshalXML(fzbBytes)
	return fzbData, err
}

// type FzbDir map[string]Fzb
//
// func ReadDir() FzbDir {
// 	store := FzbDir{}
// 	return store
// }

func UnmarshalXML(src []byte) (Fzb, error) {
	f := Fzb{}
	err := xml.Unmarshal(src, &f)
	return f, err
}

// MarshalXML return the marshaled data as byte array
func (f *Fzb) ParseXML() ([]byte, error) {
	b, err := xml.MarshalIndent(f, "", "  ")
	return b, err
}

func (f *Fzb) TotalInstances() int {
	return len(f.Instances)
}

func (f *Fzb) Validate(basepath string) (error, string) {
	errMsg := ""
	warnMsg := ""
	if f.Title == "" {
		errMsg += "ERROR > Missing Title\n"
	}
	if f.FritzingVersion == "" {
		warnMsg += "WARN  > Missing FritzingVersion\n"
	}

	// check if icon exist...
	if f.Icon == "" {
		warnMsg += "WARN  > Missing Icon Path\n"
	}

	// check if icon is lowercase
	if f.Icon != strings.ToLower(f.Icon) {
		warnMsg += fmt.Sprintf("WARN  > Icon Filename is not lowercase. %s\n", f.Icon)
	}

	iconExt := filepath.Ext(f.Icon)
	if iconExt != ".png" {
		warnMsg += fmt.Sprintf("WARN  > Icon Extension not '.png' - %s\n", iconExt)
	}

	// check fi file exist / is file readable...
	tmpPath := filepath.Join(basepath, f.Icon)
	_, err := ioutil.ReadFile(tmpPath)
	if err != nil {
		errMsg += fmt.Sprintf("ERROR > Icon File %q - %s\n", f.Icon, err)
	}
	// fmt.Println("icon exist", tmpPath)

	tmptotal := f.TotalInstances()
	if tmptotal < MinimumInstances {
		errMsg += fmt.Sprintf("ERROR > Minimum number of Instances must be 4! current %v\n", tmptotal)
	}

	// check if a instance file exist
	// for _, v := range f.Instances {
	// fmt.Println("PATH", v.Path)
	// if v.Path != "" {
	// 	instancePathExt := filepath.Ext(v.Path)
	// 	// fmt.Println("instancePathExt", instancePathExt)
	// 	if instancePathExt == "" {
	// 		errMsg += "WARN  > Instance Path has no .fzp extension - " + v.Path + "\n"
	// 	} else if instancePathExt != ".fzp" {
	// 		errMsg += "WARN  > Instance Path is not a .fzp file. " + v.Path + "\n"
	// 	} else if instancePathExt == ".fzp" {
	// 		_, err := ioutil.ReadFile(v.Path)
	// 		if err != nil {
	// 			errMsg += "ERROR> file not found " + v.Path + "\n"
	// 		}
	// 	}
	// }
	// }

	if errMsg != "" {
		return errors.New(errMsg), warnMsg
	}
	return nil, warnMsg
}

// ValidateFile validate a .fzb file and print result to stdout
func ValidateFile(basepath, src string) string {
	tmpReport := ""

	// check if file is a fzb
	if filepath.Ext(src) == ".fzb" {

		fzbData, err := ReadFile(src)
		if err != nil {
			return fmt.Sprintf("ERROR> %q Read File: %s\n", src, err)
		}

		// fmt.Println("Validate Data...")
		err, warn := fzbData.Validate(basepath)
		if warn == "" && err == nil {
			return ""
		}

		if warn != "" || err != nil {
			tmpReport = fmt.Sprintf("\n===== @ %q\n", src)
		}
		if warn != "" {
			tmpReport += warn
		}
		if err != nil {
			tmpReport += err.Error()
		}

	}
	return tmpReport
}

// ValidateDir validate all .fzb files at the given directory.
func ValidateDir(src string) string {
	d, err := ioutil.ReadDir(src)
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}

	// totalFiles := len(d)
	// fmt.Println("Total Files", totalFiles)
	// fmt.Println("Start Validating files...")

	tmpReport := ""
	for _, v := range d {
		tmpfilepath := src + "/" + v.Name()
		// fmt.Println("tmpfilepath", tmpfilepath)
		tmpReport += ValidateFile(src, tmpfilepath)
		// err, _ =
		// if err != nil {
		// 	fmt.Println("\n", v.Name())
		// 	fmt.Println(err)
		// }
	}

	return tmpReport
}
