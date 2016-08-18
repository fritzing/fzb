package fzb

import (
	"testing"
)

func Test_Fzb(t *testing.T) {
	f := NewFzb()
	t.Log(f)
}

func Test_Fzb_ReadFile(t *testing.T) {
	f, err := ReadFile("../fixture/test1.fzb")
	if err != nil {
		t.Error(err)
	}
	t.Log(f)

	if f.Title != "Test Parts" {
		t.Error("fzb Title not equal")
	}

	if f.TotalInstances() != 1 {
		t.Error("Fzb Instances total not equal")
	}

	if f.Instances[0].Path != "Basic" {
		t.Error("Fzb Instances[0] Path not equal")
	}

	f.PrettyPrint()
}
