package files

import (
	/* Standard library packages */
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"

	/* Third party */
	// imports as "cli", pinned to v1; cliv2 is going to be drastically
	// different and pinning to v1 avoids issues with unstable API changes
	"gopkg.in/urfave/cli.v1"

	/* Local packages */
	"github.com/keeferrourke/imgrep/ocr"
	"github.com/keeferrourke/imgrep/storage"
)

var Args []string

/* perform db query */
func Grep(c *cli.Context) {
	if len(c.Args()) < 1 {
		log.Fatal("args: query required")
	}

	if c.Bool("no-preindex") {
		for _, arg := range c.Args() {
			Args = append(Args, arg)
		}

		wg := sync.WaitGroup{}
		filepath.Walk(WALKPATH, GWrapper(&wg))
		wg.Wait()
	} else {
		for _, arg := range c.Args() {
			res, _ := storage.Get(arg)
			for i := 0; i < len(res); i++ {
				fmt.Println(res[i])
			}
		}
	}
}

func GWrapper(wg *sync.WaitGroup) func(path string, f os.FileInfo, err error) error {
	return func(path string, f os.FileInfo, err error) error {
		go func() {
			wg.Add(1)
			defer wg.Done()
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
					// rather than indexing in sqlite db, compare results from OCR
					// scan with Args string slice
					res := ocr.Process(path)
					found := false
					for j := 0; j < len(res); j++ {
						r := res[j]

						for i := 0; i < len(Args); i++ {
							if strings.Contains(strings.ToLower(r), strings.ToLower(Args[i])) {
								found = true
							}
						}
					}

					if found {
						fmt.Println(path)
					}
				}
			}
		}()
		return nil
	}
}
