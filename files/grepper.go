package files

import (
	/* Standard library packages */
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	/* Third party */
	// imports as "cli", pinned to v1; cliv2 is going to be drastically
	// different and pinning to v1 avoids issues with unstable API changes

	/* Local packages */
	"github.com/keeferrourke/imgrep/ocr"
	"github.com/keeferrourke/imgrep/storage"
)

var (
	Query      []string
	Results    []string
	IgnoreCase bool = false
)

/* perform db query */
func Grep(preindex bool) {
	if !preindex {
		filepath.Walk(WALKPATH, GWalker)
	} else {
		for _, arg := range Query {
			res, err := storage.Get(arg, IgnoreCase)
			if err != nil {
				log.Printf("%T %v\n", err, err)
			}
			for i := 0; i < len(res); i++ {
				Results = append(Results, res[i])
			}
		}
	}
}

/* this is an alternative to the Walker function in walker.go:
 * it both walks the filesystem and performs the search;
 * this is slow, but was implemented on request to run imgrep without using the
 * preindexed database */
func GWalker(path string, f os.FileInfo, err error) error {
	if Verbose {
		fmt.Printf("touched: %s\n", path)
	}

	// only try to open existing files
	if _, err := os.Stat(path); !os.IsNotExist(err) && !f.IsDir() {
		isImage, err := IsImage(path)
		if err != nil {
			log.Fatal(err)
		}
		if isImage {
			res, err := ocr.Process(path)
			if err != nil {
				return err
			}
			found := false
			// compare result words to each query word
			for j := 0; j < len(res); j++ {
				r := res[j]
				for i := 0; i < len(Query); i++ {
					if IgnoreCase {
						if strings.Contains(strings.ToLower(r), strings.ToLower(Query[i])) {
							found = true
						}
					} else {
						if strings.Contains(r, Query[i]) {
							found = true
						}
					}
				}
			}
			if found {
				Results = append(Results, path)
			}
		}
	}
	return nil
}
