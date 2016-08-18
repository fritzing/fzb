package fzb

import (
	"errors"
	"fmt"
)

type Instance struct {
	IconView   View   `xml:"views>iconView"`
	ModuleRef  string `xml:"moduleIdRef,attr"`
	ModelIndex string `xml:"modelIndex,attr"`
	Path       string `xml:"path,attr"`
}

func (i *Instance) PrettyPrint() {
	fmt.Printf("ModuleRef  = %q\n", i.ModuleRef)
	fmt.Printf("ModelIndex = %q\n", i.ModelIndex)
	fmt.Printf("Path       = %q\n", i.Path)
	i.IconView.PrettyPrint()
}

func (i *Instance) Validate() error {
	fmt.Println("check instance", i)
	if i.IconView.Layer != "icon" {
		return errors.New("IconView not valid. must be 'icon'")
	}

	if i.ModuleRef == "" {
		return errors.New("IconView ModuleRef not set")
	}

	return nil
}

type View struct {
	Layer    string `xml:"layer,attr"`
	Geometry `xml:"geometry"`
}

func (v *View) PrettyPrint() {
	fmt.Printf("Layer     = %q\n", v.Layer)
	fmt.Println("Geometry =", v.Geometry)
}

// FIXME: the xml decoder does not parse the attributes of the geometry tag
type Geometry struct {
	X int `xml:"x,attr"`
	Y int `xml:"y,attr"`
	Z int `xml:"z,attr"`
}
