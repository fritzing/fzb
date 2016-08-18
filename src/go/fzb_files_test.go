package fzb

import (
	"testing"
)

var TestDataPassed = []struct {
	src string
}{
	{"../fixture/test16.fzb"},
}

func Test_Fzb_ReadFile_Passed(t *testing.T) {
	for ti, tt := range TestDataPassed {
		t.Log("TEST", ti, tt.src)

		_, err := ReadFile(tt.src)
		if err != nil {
			t.Error("ReadFile Error:", err)
		}
		// t.Log(f)

		// if f.Title != "Test Parts" {
		// 	t.Error("fzb Title not equal")
		// }

		// if f.TotalInstances() > 1 {
		// 	t.Error("Fzb Instances total not equal")
		// }
		//
		// if f.Instances[0].Path != "Basic" {
		// 	t.Error("Fzb Instances[0] Path not equal")
		// }

		// f.PrettyPrint()

	}
}

// the following test files produce errors, so we can check if all errors are fine
var TestDataFailed = []struct {
	src string
}{
	// {"../fixture/test-empty.fzb"},
	{"../fixture/test-module-empty.fzb"},
	{"../fixture/test-title.fzb"},
	{"../fixture/test1.fzb"},
}

func Test_Fzb_ReadFile_Failed(t *testing.T) {
	for ti, tt := range TestDataFailed {
		t.Log("TEST ", ti, tt.src)

		tmpfzb, err := ReadFile(tt.src)
		if err != nil {
			t.Error("Missing fixture", err)
		}
		err, warn := tmpfzb.Validate("../fixture")
		if err == nil {
			t.Error("ReadFile missing Error")
		}
		t.Log(warn)
	}
}
