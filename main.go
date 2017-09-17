package main

import (
    /* Standard library packages */
    "fmt"
    "os"

    /* Third party */
    // imports as "cli", pinned to v1; cliv2 is going to be drastically
    // different and pinning to v1 avoids issues with unstable API changes
    "gopkg.in/urfave/cli.v1"

    /* Local packages */
    "github.com/keeferrourke/imgrep/files"
    "github.com/keeferrourke/imgrep/storage"
)

/* cli commands */
// search db
var Search = cli.Command {
    Name: "search",
    Aliases: []string{"s", "find"},
    Usage: "search image database for keywords",
    Action: files.Grep,
}

// update/initialize sql db
var UpdateDB = cli.Command {
    Name: "updatedb",
    Aliases: []string{"init"},
    Usage: "initialize the database of images",
    Action: files.InitFromPath,
    Flags: []cli.Flag {
        cli.StringFlag {
            Name: "dir, d",
            Value: files.WALKPATH,
            Usage: "specify the base filesystem subtree to scan",
            Destination: &files.WALKPATH,
        },
        cli.BoolFlag {
            Name: "verbose, v",
            Usage: "enable verbose output",
        },
    },
}

func init() {
    storage.InitDB(files.DBFILE)
}

/* run application */
func main() {
    // customize cli
    cli.VersionPrinter = func(c *cli.Context) {
        fmt.Fprintf(c.App.Writer, "%s %s - %s\n",
                    c.App.Name, c.App.Version, c.App.Description)
    }

    // set up the application
    app := cli.NewApp()
    app.Authors = []cli.Author {
        cli.Author {
            Name: "Keefer Rourke",
        },
        cli.Author {
            Name: "Ivan Zhang",
        },
        cli.Author {
            Name: "Thomas Dedinsky",
        },
    }
    app.Copyright = "(c) 2017 under the MIT License"
    app.EnableBashCompletion = true
    app.Name = "imgrep"
    app.Description = "go-cli image grepper using tesseract"
    app.Usage = "grep image files for words"
    app.Version = "v0"
    app.Commands = []cli.Command{
        Search,
        UpdateDB,
    }
    app.CommandNotFound = func(c *cli.Context, command string) {
        fmt.Fprintf(c.App.Writer, "Did you read the manual?\n");
    }

    app.Run(os.Args)
    return
}
