package ocr

import (
	"testing"
)

func TestProcess(t *testing.T) {
	// test with empty path
	_, err := Process("")
	if err == nil {
		t.Error("Test with empty path should have thrown error.")
	}
	// add more test cases as they are relevant; try with image and non-image
	// files
}
