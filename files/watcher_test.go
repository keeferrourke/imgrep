package files

import (
	"io/ioutil"
	"testing"
)

func TestInitWatcher(t *testing.T) {
	d1 := []byte("hello\ngo\n")
	err := ioutil.WriteFile("dat1", d1, 0644)
	if err != nil {
		t.Error(err)
	}
	InitWatcher(".")
	err = ioutil.WriteFile("dat1", []byte("yolo"), 0644)
	if err != nil {
		t.Error(err)
	}
}
