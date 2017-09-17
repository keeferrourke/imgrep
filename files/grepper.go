package files

import (
	/* Standard library packages */
	"fmt"
	"log"

	/* Third party */
	// imports as "cli", pinned to v1; cliv2 is going to be drastically
	// different and pinning to v1 avoids issues with unstable API changes
	"gopkg.in/urfave/cli.v1"

	/* Local packages */
	"github.com/keeferrourke/imgrep/storage"
)

/* perform db query */
func Grep(c *cli.Context) {
	if len(c.Args()) < 1 {
		log.Fatal("args: query required")
	}

	for _, arg := range c.Args() {
		res, _ := storage.Get(arg)
		for i := 0; i < len(res); i++ {
			fmt.Println(res[i])
		}
	}
}
