package ocr

import (
	/* Standard library packages */
	"errors"
	"fmt"
	"os"
	"strings"

	/* Third party */
	// imports as "cli", pinned to v1; cliv2 is going to be drastically
	// different and pinning to v1 avoids issues with unstable API changes
	"github.com/otiai10/gosseract"
	/* Local packages */)

func Process(path string) ([]string, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, errors.New("path: cannot stat file")
	}

	client, err := gosseract.NewClient()
	if err != nil {
		return nil, err
	}
	out, err := client.Src(path).Out()
	if err != nil {
		return nil, err
	}

	s := fmt.Sprintf(out)
	return strings.Split(s, " "), nil
}
