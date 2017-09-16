package srv

import (
    /* Standard library packages */
    //"html"
    "path/filepath"
    "fmt"
    "log"
    "net/http"

    /* Third party */
    // imports as "cli", pinned to v1; cliv2 is going to be drastically
    // different and pinning to v1 avoids issues with unstable API changes
    "gopkg.in/urfave/cli.v1"

    /* Local packages */
    "github.com/keeferrourke/htn17/files"
)

var (
    PORT string = "1337"
)

func StartServer(c *cli.Context) (err error) {
    fmt.Println("In server")

    err = filepath.Walk(files.WALKPATH, files.Walker)

    log.Fatal(http.ListenAndServe(":" + PORT, nil))

    return nil

}
