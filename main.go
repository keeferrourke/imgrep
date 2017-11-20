package main

import (
	/* Standard library packages */
	"fmt"
	"log"
	"os"

	/* Third party */
	// imports as "cli", pinned to v1; cliv2 is going to be drastically
	// different and pinning to v1 avoids issues with unstable API changes
	"gopkg.in/urfave/cli.v1"

	/* Local packages */
	"github.com/keeferrourke/imgrep/files"
	"github.com/keeferrourke/imgrep/storage"
)

var search = cli.Command{
	Name:    "search",
	Aliases: []string{"s"},
	Usage:   "search image database for keywords",
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "no-preindex, n",
			Usage: "run without preindex",
		},
		cli.BoolFlag{
			Name:  "ignore-case, i",
			Usage: "ignore case distinctions",
		},
	},
	Action: func(c *cli.Context) {
		if len(c.Args()) < 1 {
			log.Fatal("args: query required")
		} else {
			for _, arg := range c.Args() {
				files.Query = append(files.Query, arg)
			}
		}
		files.IgnoreCase = c.Bool("ignore-case")
		files.Grep(!c.Bool("no-preindex"))
		for r := range files.Results {
			fmt.Println(files.Results[r])
		}
	},
}

var updateDB = cli.Command{
	Name:    "updatedb",
	Aliases: []string{"init"},
	Usage:   "initialize the database of images",
	Action: func(c *cli.Context) {
		files.InitFromPath(c.Bool("verbose"))
	},
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:        "dir, d",
			Value:       files.WALKPATH,
			Usage:       "specify the base filesystem subtree to scan",
			Destination: &files.WALKPATH,
		},
		cli.BoolFlag{
			Name:  "verbose, v",
			Usage: "enable verbose output",
		},
	},
}

func init() {
	err := storage.InitDB(files.DBFILE)
	if err != nil {
		log.Fatal(err)
	}
}

/* run application */
func main() {
	// customize cli
	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Fprintf(c.App.Writer, "%s v%s\n    %s\n",
			c.App.Name, c.App.Version, c.App.Description)
	}

	// set up the application
	app := cli.NewApp()
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Keefer Rourke",
			Email: "mail@krourke.org",
		},
		cli.Author{
			Name:  "Ivan Zhang",
			Email: "ivan@ivanzhang.ca",
		},
	}
	app.Copyright = "(c) 2017 under the MIT License"
	app.EnableBashCompletion = true
	app.Name = "imgrep"
	app.Description = "image grepper using tesseract OCR to extract words from images"
	app.Usage = "grep images for OCR extracted words"
	app.Version = "0.0.1"
	app.Commands = []cli.Command{
		search,
		updateDB,
	}
	app.CommandNotFound = func(c *cli.Context, command string) {
		fmt.Fprintf(c.App.Writer, "Did you read the manual?\n")
	}

	app.Run(os.Args)
	return
}
