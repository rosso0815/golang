package mygraphics

import (
	"errors"
	"log"
	"path"
	"strings"
	"time"

	"gopkg.in/gographics/imagick.v2/imagick"
)

func init() {
	log.Println("mygraphics_impl -> init")
	imagick.Initialize()
	defer imagick.Terminate()
}

// ReadFileFromPath does the thing
func NewFileFromPath(lPath string) (img *Image, err error) {

	log.Println("@@@ ReadFileFromPath path=", lPath)
	log.Println("path.Ext=", strings.ToLower(path.Ext(lPath)))
	myExt := strings.ToLower(path.Ext(lPath))
	if !strings.Contains(".jpg .heic", myExt) {
		log.Println("NOT .jpg .heic", myExt)
		return nil, errors.New("mygraphics: cannot open not-jpg file")
	}
	img = new(Image)
	mw := imagick.NewMagickWand()
	mw.ReadImage(lPath)
	log.Println("width =", mw.GetImageWidth())
	img.width = mw.GetImageWidth()
	img.heigth = mw.GetImageHeight()
	img.emake = mw.GetImageProperty("exif:Make")
	img.model = mw.GetImageProperty("exif:Model")
	img.model = strings.ReplaceAll(img.model, " ", "")
	img.created = mw.GetImageProperty("exif:DateTimeOriginal")
	img.path = lPath
	log.Println("found", img)
	return img, nil
}

// GetInfo reads the file and analyze it
func (ip *Image) GetInfo() {
	log.Println("@@@ GetInfo", ip)

}

// SaveFileResized reads the file and analyze it
// TODO handle folder of actual file
// TODO handle correct extension
func (img *Image) SaveFileResized(my_path string) (err error) {

	log.Println("@@@ SaveFileResized, heigth:", img.heigth, " width:", img.width)

	//var mw *imagick.MagickWand

	mw := imagick.NewMagickWand()

	mw.ReadImage(img.path)
	log.Println("width =", mw.GetImageWidth())

	// TODO : possibel as cmd args defined ?
	var (
		newWidth  uint = 1980
		newHeigth uint = 960
	)

	// calculate newWidth or newHeigth based on existing values
	if (img.heigth < img.width) && (img.width > newWidth) {
		log.Println("resize width > ", newWidth)
		scale := img.width * 1000 / newWidth
		newHeigth = img.heigth * 1000 / scale
	} else if img.heigth > 960 {
		log.Println("resize width heigth > 960")
		scale := img.heigth * 1000 / newHeigth
		newWidth = img.width * 1000 / scale
	} else {
		log.Println("no resize possible")
		return nil
	}
	log.Println("newWidth=", newWidth, "newHeigth", newHeigth)

	// resize the picture
	err = mw.ResizeImage(newWidth, newHeigth, imagick.FILTER_LANCZOS, 1)
	if err != nil {
		panic(err)
	}

	// convert_date
	layout := "2006:01:02 15:04:05"
	tm, err := time.Parse(layout, img.created)
	if err != nil {
		return errors.New("mygraphics: canno parse date")
	}

	// calculate the new filename
	fixedModel := strings.Replace(img.model, "-", "", 10)

	log.Println("image.model=", img.model, " fixedModel=", fixedModel)

	newFilename := tm.Format("20060102_150405") +
		"_" +
		strings.ToLower(img.emake) +
		"_" +
		strings.ToLower(fixedModel) +
		".jpg"

	newpath := path.Join(path.Dir(img.path), path.Base(newFilename))
	mw.SetImageCompressionQuality(95)
	mw.WriteImage(newpath)
	return nil
}
