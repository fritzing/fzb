package fzb

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

// the fritzing fzb format go utility
type Fzb struct {
	XMLName   xml.Name   `xml:"module"                 json:"-"             yaml:"-"`
	Title     string     `xml:"title"                  json:"title"         yaml:"title"`
	Instances []Instance `xml:"instances>instance"     json:"instances"     yaml:"instances"`
}

func NewFzb() Fzb {
	f := Fzb{}
	f.Instances = make([]Instance, 0)
	return f
}

func (f *Fzb) TotalInstances() int {
	return len(f.Instances)
}

func (f *Fzb) PrettyPrint() {
	fmt.Printf("Title = %q\n", f.Title)
	totalInstances := f.TotalInstances()
	for k, v := range f.Instances {
		fmt.Printf("\nInstance %v of %v\n", k, totalInstances)
		v.PrettyPrint()
	}
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
	return fzbData, nil
}

func UnmarshalXML(src []byte) (Fzb, error) {
	f := Fzb{}
	err := xml.Unmarshal(src, &f)
	return f, err
}
