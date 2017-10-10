package files

import (
	"testing"

	"github.com/keeferrourke/imgrep/storage"
)

func TestProcessImage(t *testing.T) {
	storage.InitDB(DBFILE)
	testStaticDir := "../.res/test/static/"
	testTable := []struct {
		path string
		err  error
	}{
		{testStaticDir + "test.png", nil},
	}

	for _, testCase := range testTable {
		err := processImage(testCase.path)
		if err != testCase.err {
			t.Errorf("Expected err to be %v but got %v", testCase.err, err)
		}
		storage.Delete(testCase.path)
	}

}
