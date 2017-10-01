package files

import (
	/* Standard library packages */
	"errors"
	"io/ioutil"
	"log"
	"strings"
	/* Third party */ /* Local packages */)

var magicTable = map[string]string{
	"\xff\xd8\xff":      "image/jpeg",
	"\x89PNG\r\n\x1a\n": "image/png",
	"GIF87a":            "image/gif",
	"GIF89a":            "image/gif",
}

func magicLookup(b []byte) (string, error) {
	imgdata := string(b)
	for magic, mimetype := range magicTable {
		if strings.HasPrefix(imgdata, magic) {
			return mimetype, nil
		}
	}

	return "", errors.New("file: image format unrecognized")
}

func IsImage(path string) (bool, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("%T %v\n", err, err)
		return false, err
	}

	_, err = magicLookup(b)
	if err != nil {
		return false, err
	}
	return true, nil
}
