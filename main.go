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
    "github.com/keeferrourke/htn17/srv"
    "github.com/keeferrourke/htn17/files"
)

/* cli commands */
// server start
var Start = cli.Command {
    Name: "start",
    Aliases: []string{"run"},
    Usage: "start the threat-detection server",
    Action: srv.StartServer,
    Flags: []cli.Flag {
        cli.StringFlag {
            Name: "port, p",
            Value: "1337",
            Usage: "set `PORT` for the server at run-time",
            Destination: &srv.PORT,
        },
        cli.StringFlag {
            Name: "dir, d",
            Value: files.WALKPATH,
            Usage: "specify the base filesystem subtree to scan",
            Destination: &files.WALKPATH,
        },
    },
}


// bootstrap the application
func init() {
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
    app.Name = "gocymon"
    app.Description = "Cross platform threat-detection daemon"
    app.Usage = "run threat-detection server"
    app.Version = "v0"
    app.Commands = []cli.Command{
        Start,
    }
    app.CommandNotFound = func(c *cli.Context, command string) {
        fmt.Fprintf(c.App.Writer, "Did you read the manual?\n");
    }

    app.Run(os.Args)
    return
}
