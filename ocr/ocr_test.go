package ocr

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	/* set up tests */
	// make empty directory
	err := os.MkdirAll("testing", 755)
	if err != nil {
		log.Fatalf("%T, %v\n", err, err)
	}
	// touch empty file
	empty, err := os.Create("empty.txt")
	defer empty.Close()
	if err != nil {
		log.Fatalf("%T %v\n", err, err)
	}
	// create png image
	const w, h = 24, 24
	img := image.NewNRGBA(image.Rect(0, 0, w, h))
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			img.Set(x, y, color.NRGBA{
				R: uint8(x + y),
				G: uint8(x + y),
				B: uint8(x + y),
				A: 255,
			})
		}
	}
	pngTest, err := os.Create("test.png")
	defer pngTest.Close()
	if err != nil {
		log.Fatalf("%T, %v\n", err, err)
	}
	if err := png.Encode(pngTest, img); err != nil {
		log.Fatalf("%T, %v\n", err, err)
	}

	/* run tests */
	m.Run()

	/* clean up tests */
	err = os.RemoveAll("testing")
	if err != nil {
		log.Fatalf("%T, %v\n", err, err)
	}
	err = os.RemoveAll("empty.txt")
	if err != nil {
		log.Fatalf("%T, %v\n", err, err)
	}
	err = os.RemoveAll("test.png")
	if err != nil {
		log.Fatalf("%T, %v\n", err, err)
	}
}

func TestProcess(t *testing.T) {
	// test with empty path
	_, err := Process("")
	if err == nil {
		t.Error("Test with empty path should have thrown error.")
	}

	// test on empty directory
	_, err = Process("testing")
	if err == nil {
		t.Error("Test on directory inode should have thrown error.")
	}

	// test on empty file
	_, err = Process("empty.txt")
	if err == nil {
		t.Error("Test on empty file should have thrown error.")
	}

	// test image without any text
	kw, err := Process("test.png")
	if err != nil {
		t.Errorf("%T %v\n", err, err)
	} else if kw != nil {
		t.Error("Did not expect any results.")
	}

	// test image with text
}
