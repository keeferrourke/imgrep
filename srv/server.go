package srv

import (
    /* Standard library packages */
    "fmt"
    "log"
    "net/http"

    /* Third party */
    // imports as "cli", pinned to v1; cliv2 is going to be drastically
    // different and pinning to v1 avoids issues with unstable API changes
    "gopkg.in/urfave/cli.v1"

    /* Local packages */
)

var (
    PORT string = "1337"
)

func StartServer(c *cli.Context) (err error) {
    fmt.Println("In server")

    log.Fatal(http.ListenAndServe(":" + PORT, nil))

    return nil

}
