package files

import (
	/* Standard library packages */
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"runtime"

	/* Local packages */
	"github.com/keeferrourke/imgrep/ocr"
	"github.com/keeferrourke/imgrep/storage"
)

var (
	// WALKPATH is the path to be walked during database initialization
	WALKPATH string
	// CONFDIR is the directory where program files will be stored
	CONFDIR string
	// DBFILE is the path to the file where the keyword database is stored
	DBFILE string

	// Verbose is a boolean flag toggling verbosity of *Walker functions
	Verbose = false
)

func init() {
	u, err := user.Current()
	if err != nil {
		panic(err)
	}
	WALKPATH, err = os.Getwd()
	if err != nil {
		panic(err)
	}
	CONFDIR = u.HomeDir + string(os.PathSeparator)
	if runtime.GOOS == "windows" {
		CONFDIR += "AppData" + string(os.PathSeparator) + "Local"
		CONFDIR += string(os.PathSeparator) + "imgrep"
	} else {
		CONFDIR += ".imgrep"
		if _, err := os.Stat(CONFDIR); os.IsNotExist(err) {
			err = os.Mkdir(CONFDIR, os.ModePerm)
			if err != nil {
				panic(err)
			}
		}
	}
	DBFILE = CONFDIR + string(os.PathSeparator) + "imgrep.db"
}

// Walker is executed when walking the file system.
// It checks if files are images and runs Tesseract OCR to index keywords
// accordingly.
func Walker(path string, f os.FileInfo, err error) error {
	if Verbose {
		fmt.Printf("touched: %s\n", path)
	}

	// only try to open existing files
	if _, err := os.Stat(path); !os.IsNotExist(err) && !f.IsDir() {
		isImage := IsImage(path)
		if isImage { // only process images
			var keys []string
			keys, err := ocr.Process(path)
			if err != nil {
				return err
			}
			storage.Insert(path, keys...)
		}
	}
	return nil
}

// InitFromPath initializes the keyword database by walking the directory tree
// and subsequently processing images.
func InitFromPath(verbose bool) error {
	if verbose {
		Verbose = true
	}

	err := filepath.Walk(WALKPATH, Walker)
	return err
}
