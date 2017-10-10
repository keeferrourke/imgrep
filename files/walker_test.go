package files

import (
	"fmt"
	"testing"

	"github.com/keeferrourke/imgrep/storage"
)

func TestProcessImage(t *testing.T) {
	storage.InitDB(DBFILE)
	testStaticDir := "../.res/test/static/"
	testTable := []struct {
		path      string
		assertion func(error)
	}{
		{
			testStaticDir + "test.png",
			func(err error) {
				if err != nil {
					t.Errorf("Expected err to be %v but got %v", nil, err)
				}
			}},
		{
			testStaticDir + "test.txt",
			func(err error) {
				if err != ErrImageNotRecognized {
					t.Errorf("Expected err to be %v but got %v", nil, err)
				}
			},
		},
		{
			testStaticDir + "doesnotexist",
			func(err error) {
				errString := fmt.Sprintf("open %s: no such file or directory", testStaticDir+"doesnotexist")
				if err.Error() != errString {
					t.Errorf("Expected err to be %v but got %v", errString, err.Error())
				}
			},
		},
	}

	for _, testCase := range testTable {
		err := processImage(testCase.path)
		testCase.assertion(err)

		// clean up
		storage.Delete(testCase.path)
	}

}
