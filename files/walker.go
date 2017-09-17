package files

import (
	/* Standard library packages */
	"fmt"
	"log"
	"os"
	"os/user"
	"path/filepath"

	/* Third party */
	// imports as "cli", pinned to v1; cliv2 is going to be drastically
	// different and pinning to v1 avoids issues with unstable API changes
	"gopkg.in/urfave/cli.v1"

	/* Local packages */
	"github.com/keeferrourke/imgrep/ocr"
	"github.com/keeferrourke/imgrep/storage"
)

var (
	WALKPATH string
	CONFDIR  string
	DBFILE   string

	verb bool = false
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
	CONFDIR = u.HomeDir + string(os.PathSeparator) + ".imgrep"
	if _, err := os.Stat(CONFDIR); os.IsNotExist(err) {
		err = os.Mkdir(CONFDIR, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
	DBFILE = CONFDIR + string(os.PathSeparator) + "imgrep.db"
}

func Walker(path string, f os.FileInfo, err error) error {
	if verb {
		fmt.Printf("touched: %s\n", path)
	}

	// only try to open existing files
	if _, err := os.Stat(path); !os.IsNotExist(err) && !f.IsDir() {
		isImage, err := IsImage(path)
		if err != nil {
			log.Fatal(err)
		}
		if isImage {
			storage.Insert(path, ocr.Process(path)...)
		}
	}
	return nil
}

func InitFromPath(c *cli.Context) error {
	if c.Bool("verbose") {
		verb = true
	}

	err := filepath.Walk(WALKPATH, Walker)
	return err
}
