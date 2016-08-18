package fzb

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
)

// the fritzing fzb format go utility
type Fzb struct {
	XMLName         xml.Name   `xml:"module"                 json:"-"                 yaml:"-"`
	Title           string     `xml:"title"                  json:"title"             yaml:"title"`
	Icon            string     `xml:"icon,attr"              json:"icon"              yaml:"icon"`
	FritzingVersion string     `xml:"fritzingVersion,attr"   json:"fritzingVersion"   yaml:"fritzingVersion"`
	Instances       []Instance `xml:"instances>instance"     json:"instances"         yaml:"instances"`
}

func NewFzb() Fzb {
	f := Fzb{}
	f.Instances = make([]Instance, 0)
	return f
}

func ReadFile(src string) (Fzb, error) {
	fzbBytes, err := ioutil.ReadFile(src)
	if err != nil {
		return Fzb{}, err
	}
	fzbData, err := UnmarshalXML(fzbBytes)
	if err != nil {
		return fzbData, err
	}
	err = fzbData.Check()
	return fzbData, err
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

func (f *Fzb) Check() error {
	if f.TotalInstances() < 16 {
		return errors.New("Minimum number of Instances must be 16")
	}
	for _, v := range f.Instances {
		v.Check()
	}
	return nil
}

func (f *Fzb) PrettyPrint() {
	fmt.Printf("Title = %q\n", f.Title)
	fmt.Printf("Icon  = %q\n", f.Icon)
	totalInstances := f.TotalInstances()
	for k, v := range f.Instances {
		fmt.Printf("\nInstance %v of %v\n", k, totalInstances)
		v.PrettyPrint()
	}
}
