package files

import (
	"errors"
	"io/ioutil"
	"log"
	"strings"
)

var magicTable = map[string]string{
	"\xff\xd8\xff":      "image/jpeg",
	"\x89PNG\r\n\x1a\n": "image/png",
	"GIF87a":            "image/gif",
	//"GIF89a":            "image/gif", // animated gif
}

var ErrImageNotRecognized = errors.New("file: image format unrecognized")

func magicLookup(b []byte) (string, error) {
	imgdata := string(b)
	for magic, mimetype := range magicTable {
		if strings.HasPrefix(imgdata, magic) {
			return mimetype, nil
		}
	}

	return "", ErrImageNotRecognized
}

// IsImage determines if a file located at the provided path is one of a
// JPEG, PNG, or GIF image
func IsImage(path string) bool {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("%T %v\n", err, err)
		return false
	}

	_, err = magicLookup(b)
	if err != nil {
		return false
	}
	return true
}
