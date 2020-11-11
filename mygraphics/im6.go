package mygraphics

import (
	"errors"
	"log"
	"os"
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

// NewFileFromPath does the thing
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
func (img *Image) GetInfo() {
	log.Println("@@@ GetInfo", img)
}

// SaveFileResized reads the file and analyze it
// TODO handle folder of actual file
// TODO handle correct extension
func (img *Image) SaveFileResized(myPath string) (err error) {

	log.Println("@@@ SaveFileResized, my_path:", myPath, " img.path:", img.path)

	mw := imagick.NewMagickWand()

	mw.ReadImage(img.path)

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
	log.Println("newWidth =", newWidth, ", newHeigth =", newHeigth)

	// resize the picture
	err = mw.ResizeImage(newWidth, newHeigth, imagick.FILTER_LANCZOS, 1)
	if err != nil {
		panic(err)
	}
	log.Println("resize done")

	// convert_date
	tmString := "undefined"
	layout := "2006:01:02 15:04:05"

	if len(img.created) < 1 {
		img.created = "1970:01:01 00:00:00"
	}

	tm, err := time.Parse(layout, img.created)
	if err != nil {
		log.Println("ERROR", err)
		return errors.New("mygraphics: cannot parse date")
	}

	tmString = tm.Format("20060102_150405")

	// calculate the new filename
	fixedModel := strings.Replace(img.model, "-", "", 10)
	log.Println("image.model=", img.model, " fixedModel=", fixedModel, len(fixedModel))
	if len(fixedModel) < 1 {
		fixedModel = "unknown"
	}

	fileVersion := ""
	fileVersionInc := 0
	newFilename := tmString + "_" + strings.ToLower(fixedModel) + ".jpg"
	newpath := path.Join(path.Dir(myPath), path.Base(newFilename))
	log.Println("newpath=", newpath)

	_, err = os.Stat(newpath)
	if err != nil {
		log.Println("file stat error ", err)
	}

	mw.SetImageCompressionQuality(95)
	err = mw.WriteImage(newpath)
	if err != nil {
		log.Println("ERROR mw.WriteImage", err)
	}
	return nil
}
