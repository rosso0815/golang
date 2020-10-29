package exif

import (
	"errors"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/rwcarlsen/goexif/exif"
	"github.com/rwcarlsen/goexif/tiff"
)

// Exif tbd
type Exif struct {
	filename string
	model    string
	created  time.Time
	width    string
}

/*
func myWalk(name exif.FieldName, tag exif.*tiff.Tag) error {
	log.Println("walk name=" + name)
	return nil
}
*/

// Walker tbd
type Walker struct{}

// Walk tbd
func (_ Walker) Walk(name exif.FieldName, tag *tiff.Tag) error {
	data, _ := tag.MarshalJSON()
	log.Printf("    %v: %v\n", name, string(data))
	return nil
}
func (e Exif) ReadFile(filename string) error {
	log.Println("exif.ReadFile=" + filename)
	e.filename = filename

	f, err := os.Open(e.filename)
	if err != nil {
		return errors.New("exit , cannot open file" + filename)
	} else {
		log.Print("open file", filename)
	}

	x, err := exif.Decode(f)
	if err != nil {
		log.Print("error ", err.Error())
		return errors.New("decode " + err.Error())
	}

	x.Walk(Walker{})

	// lets analyze the file ...
	tm, err := x.DateTime() // normally, don't ignore errors!
	if err != nil {
		log.Printf("error get DateTime")
		return nil
	} else {
		log.Print("dateTime=", tm)
		e.created = tm
	}

	camModel, err := x.Get(exif.Model) // normally, don't ignore errors!
	if err != nil {
		log.Printf("error get exif.Model")
		return nil
	} else {
		log.Printf("model=%s\n", camModel.String())
		e.model = camModel.String()
	}

	width, err := x.Get(exif.ImageWidth)
	if err != nil {
		log.Println("error get exif.ImageWidth", err)
		return nil
	} else {
		log.Printf("exif.ImageWidth=" + width.String())
		e.width = width.String()
	}

	//log.Printf("DateTime=%s\n", value.String())
	//focal, _ := x.Get(exif.FocalLength)
	//numer, denom, _ := focal.Rat2(0) // retrieve first (only) rat. value
	//fmt.Printf("%v/%v", numer, denom)

	// Two convenience functions exist for date/time taken and GPS coords:
	//tm, _ := x.DateTime()
	//fmt.Println("Taken: ", tm)

	//lat, long, _ := x.LatLong()
	//fmt.Println("lat, long: ", lat, ", ", long)

	log.Print("---")
	return nil
}

//-----------------------------------------------------------------------------
func printFile1(fp string, fi os.FileInfo, err error) error {
	if err != nil {
		log.Println(err)
		return nil
	}
	if !!fi.IsDir() {
		log.Println("directory:" + fp)
		return nil
	}
	log.Println("file:" + fp)
	return nil
}

//-----------------------------------------------------------------------------
func addFiles(files *[]string, types []string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Println(err)
			return nil
		}
		if !!info.IsDir() {
			log.Println("directory:" + path)
			return nil
		}

		//log.Println("file:" + path)
		*files = append(*files, path)
		return nil
	}
}

//-----------------------------------------------------------------------------
/*
func main() {

	dirPtr := flag.String("dir", "", "define input directory")
	flag.Parse()
	//log.Println("length=", len(*dirPtr))
	log.Println("scan directory:", *dirPtr)

	// build filelist
	log.Println("build filelist")
	fileList := []string{}
	filepath.Walk(*dirPtr, addFiles(&fileList, []string{"JPG", "jpg"}))

	// work through file list
	log.Println("@@@ show files")

	for _, file := range fileList {
		log.Println("main range file=", file)
		err := fileGetMeta(file)
		if err != nil {
			log.Print("found error ", err.Error())
		}
	}

}
*/
//-----------------------------------------------------------------------------
