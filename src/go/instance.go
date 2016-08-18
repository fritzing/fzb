package fzb

import (
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

type View struct {
	Layer string `xml:"layer,attr"`
	Geometry
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
