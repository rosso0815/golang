package main

// https://godoc.org/gopkg.in/gographics/imagick.v2/imagick

import (
	"log"

	"gopkg.in/gographics/imagick.v2/imagick"
)

func main() {

	imagick.Initialize()
	// Schedule cleanup
	defer imagick.Terminate()
	var err error

	mw := imagick.NewMagickWand()

	err = mw.ReadImage("IMG_3339.JPG")
	if err != nil {
		panic(err)
	}

	//Get original logo size
	log.Println("width      = ", mw.GetImageWidth())
	log.Println("height     = ", mw.GetImageHeight())
	log.Println("properties = ", mw.GetImageProperties("*"))
	log.Println("exif:date  = ", mw.GetImageProperty("exif:DateTimeOriginal"))

	// Calculate half the size
	//hWidth := uint(width / 2)
	//hHeight := uint(height / 2)

	// Resize the image using the Lanczos filter
	// The blur factor is a float, where > 1 is blurry, < 1 is sharp
	// err = mw.ResizeImage(hWidth, hHeight, imagick.FILTER_LANCZOS, 1)
	// if err != nil {
	// 	panic(err)
	// }

	// Set the compression quality to 95 (high quality = low compression)
	// err = mw.SetImageCompressionQuality(95)
	// if err != nil {
	// 	panic(err)
	// }

	// err = mw.DisplayImage(os.Getenv("DISPLAY"))
	// if err != nil {
	// 	panic(err)
	// }

}