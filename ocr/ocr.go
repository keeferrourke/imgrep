package ocr

import (
	/* Standard library packages */
	"fmt"
	"strings"

	/* Third party */
	// imports as "cli", pinned to v1; cliv2 is going to be drastically
	// different and pinning to v1 avoids issues with unstable API changes
	"github.com/otiai10/gosseract"
	/* Local packages */)

func Process(path string) []string {
	client, _ := gosseract.NewClient()
	out, _ := client.Src(path).Out()

	s := fmt.Sprintf(out)
	return strings.Split(s, " ")
}
