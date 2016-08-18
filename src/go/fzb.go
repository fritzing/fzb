package fzb

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"
	// "os"
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
	if err != nil {
		return fzbData, err
	}
	return fzbData, err
}

type FzbDir map[string]Fzb

func ReadDir() FzbDir {
	store := FzbDir{}
	return store
}

func UnmarshalXML(src []byte) (Fzb, error) {
	f := Fzb{}
	err := xml.Unmarshal(src, &f)
	return f, err
}

func (f *Fzb) MarshalXML() ([]byte, error) {
	b, err := xml.MarshalIndent(f, "", "  ")
	return b, err
}

func (f *Fzb) TotalInstances() int {
	return len(f.Instances)
}

func (f *Fzb) Validate() (error, string) {
	errMsg := ""
	warnMsg := ""
	if f.Title == "" {
		errMsg = "ERROR > Missing Title\n"
	}
	if f.FritzingVersion == "" {
		warnMsg = "WARN  >  Missing FritzingVersion\n"
	}

	tmptotal := f.TotalInstances()
	if tmptotal < 16 {
		errMsg += fmt.Sprintf("ERROR > Minimum number of Instances must be 16! current %v\n", tmptotal)
	}

	// for _, v := range f.Instances {
	// 	v.Validate()
	// }
	if errMsg != "" {
		return errors.New(errMsg), warnMsg
	}
	return nil, warnMsg
}

// ValidateFile validate a .fzb file and print result to stdout
func ValidateFile(src string) string {
	tmpReport := ""

	// check if file is a fzb
	if filepath.Ext(src) == ".fzb" {

		fzbData, err := ReadFile(src)
		if err != nil {
			return fmt.Sprintf("ERROR @ %q Read File: %s\n", src, err)
		}

		// fmt.Println("Validate Data...")
		err, warn := fzbData.Validate()
		if warn == "" && err == nil {
			return fmt.Sprintf("OK    @ %q - file is valid\n\n", src)
		}

		if warn != "" || err != nil {
			tmpReport = fmt.Sprintf("ERROR @ %q\n", src)
		}
		if warn != "" {
			tmpReport += warn
		}
		if err != nil {
			tmpReport += err.Error() + "\n"
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

	totalFiles := len(d)
	fmt.Println("Total Files", totalFiles)
	fmt.Println("Start Validating files...")

	tmpReport := ""
	for _, v := range d {
		tmpfilepath := src + "/" + v.Name()
		// fmt.Println("tmpfilepath", tmpfilepath)
		tmpReport += ValidateFile(tmpfilepath)
		// err, _ =
		// if err != nil {
		// 	fmt.Println("\n", v.Name())
		// 	fmt.Println(err)
		// }
	}
	return tmpReport
}
