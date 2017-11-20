package ocr

import (
	/* Standard library packages */
	"errors"
	"os"
	"strings"

	/* Third party */
	"github.com/otiai10/gosseract"
)

// Process attempts to run OCR on the file located at the provided path
func Process(path string) ([]string, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, errors.New("path: cannot stat file")
	}

	client := gosseract.NewClient()
	client.SetImage(path)
	out, err := client.Text()
	if err != nil {
		return nil, err
	}

	if out == "" {
		return nil, nil
	}
	return strings.Split(out, " "), nil
}
