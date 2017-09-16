package files

import (
    /* Standard library packages */
    "fmt"
    "log"
    "os"
    "os/user"

    /* Third party */

    /* Local packages */
    "github.com/keeferrourke/htn17/ocr"
)

var (
    WALKPATH string
)

func init () {
    u, err := user.Current()
    if err != nil {
        log.Fatal(err)
    }
    WALKPATH = u.HomeDir + string(os.PathSeparator) + "Pictures/Screenshots"
}

func Walker(path string, f os.FileInfo, err error) error {
    fmt.Printf("touched: %s\n", path)

    // only try to open existing files
    if _, err := os.Stat(path); !os.IsNotExist(err) && !f.IsDir() {
        isImage, err := IsImage(path)
        if err != nil {
            log.Fatal(err)
        }
        if isImage {
            fmt.Printf(ocr.Process(path))
        }
    }

    return nil
}
