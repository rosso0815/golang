package main

import (
	"flag"
	"fmt"

	"pfistera/ap_exif"
	//"github.com/gographics/gmagick"
	//"github.com/xiam/exif"
)

/*
func resize(orig string, dest string) {
	mw := gmagick.NewMagickWand()
	defer mw.Destroy()
	mw.ReadImage(orig)
	filter := gmagick.FILTER_LANCZOS
	w := mw.GetImageWidth()
	h := mw.GetImageHeight()

	fmt.Println("w=", w, "h=", h)

	mw.ResizeImage(w/4, h/4, filter, 1)

	mw.WriteImage(dest)
}
*/

func main() {
	f := flag.String("from", "", "original image file ...")
	t := flag.String("to", "", "target file ...")
	flag.Parse()

	fmt.Println("-from=", *f, "-to=", *t)

	ap_exif.SayHi(*f)

	ap_exif.GetDate(*f)

	/*	data, _ := exif.Read(*f)
		for key, val := range data.Tags {
			fmt.Printf("%s = %s\n", key, val)
		}

		gmagick.Initialize()
		defer gmagick.Terminate()

		resize(*f, *t)
	*/

}
