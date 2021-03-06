package files

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	/* set up tests */
	// create image data
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
	// png
	pngTest, err := os.Create("test.png")
	defer pngTest.Close()
	if err != nil {
		log.Fatal(err)
	}
	if err := png.Encode(pngTest, img); err != nil {
		log.Fatal(err)
	}
	// jpg
	jpgTest, err := os.Create("test.jpg")
	defer jpgTest.Close()
	if err != nil {
		log.Fatal(err)
	}
	oj := jpeg.Options{
		Quality: 90,
	}
	if err := jpeg.Encode(jpgTest, img, &oj); err != nil {
		log.Fatal(err)
	}
	// gif
	gifTest, err := os.Create("test.gif")
	defer gifTest.Close()
	if err != nil {
		log.Fatal(err)
	}
	og := gif.Options{
		NumColors: 256,
	}
	if err := gif.Encode(gifTest, img, &og); err != nil {
		log.Fatal(err)
	}

	/* run tests */
	m.Run()

	/* clean up tests */
	err = os.RemoveAll("test.png")
	if err != nil {
		log.Fatal(err)
	}
	err = os.RemoveAll("test.jpg")
	if err != nil {
		log.Fatal(err)
	}
	err = os.RemoveAll("test.gif")
	if err != nil {
		log.Fatal(err)
	}
}

func TestIsImage(t *testing.T) {
	// test with empty path
	fmt.Println("Expecting path error...")
	ok := IsImage("")
	if ok {
		t.Error("Test with empty path should have thrown error.")
	}
	ok = IsImage("test.png")
	if !ok {
		t.Error("Could not verify test PNG image.")
	}
	ok = IsImage("test.jpg")
	if !ok {
		t.Error("Could not verify test JPG image.")
	}
	ok = IsImage("test.gif")
	if !ok {
		t.Error("Could not verify test GIF image.")
	}
}
