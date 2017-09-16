package files

import (
    /* Standard library packages */
    "fmt"
    "os"

    /* Third party */
    // imports as "cli", pinned to v1; cliv2 is going to be drastically
    // different and pinning to v1 avoids issues with unstable API changes
    "gopkg.in/urfave/cli.v1"

    /* Local packages */
    "github.com/keeferrourke/htn17/ocr"
    "github.com/keeferrourke/htn17/storage"
)

/* error handling */
type errorString struct {
    s string
}

func (e *errorString) Error() string {
    return e.s
}

func New(err string) error {
    return &errorString{err}
}

/* perform db query */
func Grep(c *cli.Context) error {
    if len(c.Args()) < 1 {
        err := errors.New("args: query required")
        log.Fatal(err)
        return err
    }


}
