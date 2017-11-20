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
	// Query is a slice of keywords used in a database query
	Query []string
	// Results is a slice of filenames of images that match the query
	Results []string
	// IgnoreCase is a boolean flag to ignore case specifiers in keyword strings
	IgnoreCase = false
)

// Grep performs the actual search by matching keywords in the database or by
// walking the filesystem (no-preindex) and checking each file
func Grep(preindex bool) {
	if !preindex {
		filepath.Walk(WALKPATH, GWalker)
	} else {
		for _, arg := range Query {
			res, err := storage.Get(arg, IgnoreCase)
			if err != nil {
				log.Printf("%T %v\n", err, err)
			}
			for _, r := range res {
				if _, err := os.Stat(r); !os.IsNotExist(err) {
					Results = append(Results, r)
				} else if storage.Lookup(r) { // if file !exists and in database
					storage.Delete(r)
				}
			}
		}
	}
}

// GWalker is an alternative to the Walker function in walker.go:
// it both walks the filesystem and performs the search;
// this is slow, but was implemented on request to run imgrep without using the
// preindexed database
func GWalker(path string, f os.FileInfo, err error) error {
	if Verbose {
		fmt.Printf("touched: %s\n", path)
	}

	// only try to open existing files
	if _, err := os.Stat(path); !os.IsNotExist(err) && !f.IsDir() {
		isImage := IsImage(path)
		if isImage {
			res, err := ocr.Process(path)
			if err != nil {
				return err
			}
			found := false
			// compare result words to each query word
			for _, r := range res {
				for _, q := range Query {
					if IgnoreCase {
						if strings.Contains(strings.ToLower(r), strings.ToLower(q)) {
							found = true
						}
					} else {
						if strings.Contains(r, q) {
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
