package resize

import (
	"log"
	"os"
	"path/filepath"
	"regexp"
	"time"

	"gopkg.in/gographics/imagick.v3/imagick"
)

/*
# Unterordner mit Bild-Dateiname
    if [ ! -d @eaDir/"$f" ]
    then
        mkdir @eaDir/"$f"
    fi

    # Bild -> XL
    if [ ! -f @eaDir/"$f"/SYNOPHOTO_THUMB_XL.jpg ]
    then
        #convert $f -resize 1280x1280\> -quality 90 -unsharp 0.5x0.5+1.25+0.0 @eaDir/$f/SYNOPHOTO_THUMB_XL.jpg
        vipsthumbnail -s 1280x1280 "$f" -o @eaDir/"$f"/SYNOPHOTO_THUMB_XL.jpg
    fi

    # XL -> L
    if [ ! -f @eaDir/"$f"/SYNOPHOTO_THUMB_L.jpg ]
    then
        #convert @eaDir/$f/SYNOPHOTO_THUMB_XL.jpg -resize 800x800\> -quality 90 -unsharp 0.5x0.5+1.25+0.0 @eaDir/$f/SYNOPHOTO_THUMB_L.jpg
        vipsthumbnail -s 800x800 "$f" -o @eaDir/"$f"/SYNOPHOTO_THUMB_L.jpg
    fi

    # L -> M
    if [ ! -f @eaDir/"$f"/SYNOPHOTO_THUMB_M.jpg ]
    then
        #convert @eaDir/$f/SYNOPHOTO_THUMB_L.jpg  -resize 320x320\> -quality 90 -unsharp 0.5x0.5+1.25+0.0 @eaDir/$f/SYNOPHOTO_THUMB_M.jpg
        vipsthumbnail -s 320x320 "$f" -o @eaDir/"$f"/SYNOPHOTO_THUMB_M.jpg
    fi

    # M -> S
    if [ ! -f @eaDir/"$f"/SYNOPHOTO_THUMB_S.jpg ]
    then
        #convert @eaDir/$f/SYNOPHOTO_THUMB_M.jpg  -resize 120x120\> -quality 90 -unsharp 0.5x0.5+1.25+0.0 @eaDir/$f/SYNOPHOTO_THUMB_S.jpg
        vipsthumbnail -s 120x120 "$f" -o @eaDir/"$f"/SYNOPHOTO_THUMB_S.jpg
    fi

    # L -> B
    if [ ! -f @eaDir/"$f"/SYNOPHOTO_THUMB_B.jpg ]
    then
        #convert @eaDir/$f/SYNOPHOTO_THUMB_L.jpg  -resize 640x640\> -quality 90 -unsharp 0.5x0.5+1.25+0.0 @eaDir/$f/SYNOPHOTO_THUMB_B.jpg
        vipsthumbnail -s 640x640 "$f" -o @eaDir/"$f"/SYNOPHOTO_THUMB_B.jpg
    fi
*/

//-----------------------------------------------------------------------------
//
// TODO add thumbs
//-----------------------------------------------------------------------------
func Resize(path string, filename string, doScale bool, doThumbs bool, start time.Time) {

	var n_width uint
	var n_height uint
	var scale float64
	var dateCreated string
	var model string
	var new_filename string

	log.Println("start resize", time.Since(start))

	imagick.Initialize()
	defer imagick.Terminate()
	mw := imagick.NewMagickWand()
	log.Println("after init ", time.Since(start))

	// TODO add path + filename
	abs, _ := filepath.Abs(path)
	log.Println("abs - path ", abs+string(os.PathSeparator)+filename, time.Since(start))
	log.Println("after abs ", time.Since(start))
	err := mw.ReadImage(abs + string(os.PathSeparator) + filename)
	if err != nil {
		panic(err)
	}
	log.Println("after read", time.Since(start))

	// Get original logo size
	width := mw.GetImageWidth()
	height := mw.GetImageHeight()

	if width > height {
		scale = float64(width) / 1920
		n_width = uint(float64(width) / scale)
		n_height = uint(float64(height) / scale)
		log.Printf("scale=%v old_width=%v old_height=%v new_width=%v new_height=%v\n", scale, width, height, n_width, n_height)
		log.Println(time.Since(start))

	} else {
		log.Println("scale undefined")
		panic("exit")
	}

	//log.Println(mw.GetImageProperties(""))
	dateCreated = mw.GetImageProperty("exif:DateTimeDigitized")
	model = mw.GetImageProperty("exif:Model")

	log.Println("exif:DateTimeDigitized", dateCreated)
	log.Println("exif:Model", model)
	//log.Println("exif:ExifImageLength", mw.GetImageProperty("exif:ExifImageLength"))
	//log.Println("exif:ExifImageWidth", mw.GetImageProperty("exif:ExifImageWidth"))

	//ident := mw.IdentifyImage()
	// log.Println("ident", ident)

	re := regexp.MustCompile("[^A-Za-z0-9]+")
	model = re.ReplaceAllString(model, "")
	log.Println("model=", model)

	//"2017:03:12 15:46:29
	//re = regexp.MustCompile("exif:DateTimeOriginal:\\s*(.*)")
	//dateCreated = re.FindStringSubmatch(ident)[1]
	layout := "2006:01:02 15:04:05"
	tm, err := time.Parse(layout, dateCreated)
	if err != nil {
		panic(err)
	}
	log.Printf("dateTimeOriginal=%q t=%v\n", dateCreated, tm)

	new_filename = tm.Format("20060102_150405") + "_" + model + ".jpg"
	log.Println("new_filename", new_filename)

	if doScale {

		// Resize the image using the Lanczos filter
		// The blur factor is a float, where > 1 is blurry, < 1 is sharp
		err = mw.ResizeImage(n_width, n_height, imagick.FILTER_LANCZOS, 1)
		if err != nil {
			panic(err)
		}
		mw.WriteImage(new_filename)

	} else {

		log.Println("will not scale")
		return

	}

	if doThumbs {

		log.Println("doThumbs =" + abs + string(os.PathSeparator) + "@eaDir")

		if stat, err := os.Stat(path + string(os.PathSeparator) + "@eaDir"); err == nil && stat.IsDir() {
			log.Println("path does exists")
		} else {
			log.Println("path does not exists")
		}

		//os.Mkdir("pfistera/play/test", 0777)
		//os.Mkdir("pfistera", 0777)

		/*
			   		  if [ ! -d @eaDir/"$f" ]
				       then
				           mkdir @eaDir/"$f"
				       fi

				       # Bild -> XL
				       if [ ! -f @eaDir/"$f"/SYNOPHOTO_THUMB_XL.jpg ]
				       then
				           #convert $f -resize 1280x1280\> -quality 90 -unsharp 0.5x0.5+1.25+0.0 @eaDir/$f/SYNOPHOTO_THUMB_XL.jpg
				           vipsthumbnail -s 1280x1280 "$f" -o @eaDir/"$f"/SYNOPHOTO_THUMB_XL.jpg
				       fi
		*/

	}

}
