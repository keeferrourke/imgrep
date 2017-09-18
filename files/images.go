package files

import (
	/* Standard library packages */
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

func magicLookup(magicb []byte) string {
	magicstr := string(magicb)
	for magic, mimetype := range magicTable {
		if strings.HasPrefix(magicstr, magic) {
			return mimetype
		}
	}

	return ""
}

func IsImage(path string) (bool, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	mime := magicLookup(b)
	if mime == "" {
		return false, err
	}
	return true, err
}
