package fzb

import (
	"testing"
)

func Test_Fzb(t *testing.T) {
	f := NewFzb()
	f.Title = "test"

	fxml, err := f.MarshalXML()
	if err != nil {
		t.Error(err)
	}
	t.Log(string(fxml))
}
