package main

import (
	"github.com/nfnt/resize"
	"image/jpeg"
	"log"
	"os"
)

func Resize(path *string) {

	// open "test.jpg"
	file, err := os.Open(*path)
	if err != nil {
		log.Fatal(err)
	}

	// decode jpeg into image.Image
	img, err := jpeg.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	file.Close()

	// resize to width 1000 using Lanczos resampling
	// and preserve aspect ratio
	m := resize.Resize(1000, 0, img, resize.Lanczos3)

	out, err := os.Create("test_resized.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	// write new image to file
	jpeg.Encode(out, m, nil)

}

func main() {
	path := "test.jpg"
	Resize(&path)
	log.Print("done for path ", path)

}

//

//
