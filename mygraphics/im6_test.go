package mygraphics

import (
	"fmt"
	"log"
	"os"
	"testing"
	"time"
)

func Test_GetInfoFromSingleFile(t *testing.T) {
	log.Println("@@@ Test_GetInfoFromSingleFile")

	// var testImageHandler ImageHandler

	// mImg1, _ := NewProcessMockImages()
	// testImageHandler = mImg1
	// testImageHandler.ReadFileFromPath("test01.jpg")
	// testImageHandler.ReadFileFromPath("test02.jpg")

	//mImg2, _ := NewProcessImplImages()

	//testImageHandler = mImg2

	//testImageHandler.ReadFileFromPath("test01.jpg")

	//log.Println("GetInfo =", testImageHandler.GetInfo())
	//testImageHandler.SaveFileResized()

	// a image without exif-data
	//testImageHandler.ReadFileFromPath("test02.jpg")
	//log.Println("GetInfo =", testImageHandler.GetInfo())
}

// TestReadmetadata just to verify the actual pictures
// func TestReadmetadata(t *testing.T) {
// 	log.Println("@@@ start")

// Interface Handler
//var img ImageHandler

//r := ImageMagickInternals{}

//img = r

//img.ReadFile("test01.jpg")
//log.Println("r=", r)

// if img1.width != 5472 {
// 	log.Panic("image wrong height")
// }

// if img1.heigth != 3080 {
// 	log.Panic("image wrong heigth")
// }

// if img1.make != "SONY" {
// 	log.Panic("image wrong make")
// }

// if img1.model != "DSC-RX100M3" {
// 	log.Panic("image wrong model")
// }

// if img1.created != "2018:04:20 03:42:56" {
// 	log.Panic("image wrong date")
// }

// 	log.Println("@@@ done")
// }

// func TestResizedImage(t *testing.T) {

// 	log.Println("@@@ TestResizedImage")

//     Info()
//     Info()
//     Info()

// 	//deleteOldFiles()

// 	//img1, _ := ReadMetaInfo("test01.jpg")
// 	//WriteResizedImages(img1)

// 	//img2, _ := ReadMetaInfo("test02.jpg")
// 	//WriteResizedImages(img2)

// 	//deleteOldFiles()
// }

func setup() {
	log.Println("@@@ setup")
	os.RemoveAll("tmp")
	os.Mkdir("tmp", os.ModePerm)
}

func teardown() {
	log.Println("@@@ teardown")
	os.RemoveAll("tmp")
}

func Test_SingleImage(t *testing.T) {
	testCases := []struct {
		file string
		want string
	}{
		{"test01.jpg", "13:31"},
		//{"test02.jpg", "07:31"},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("test %s ", tc.file), func(t *testing.T) {
			log.Println("")
			log.Println("@@@ tc.file=", tc.file)
			myImg, _ := NewFileFromPath(tc.file)
			myImg.GetInfo()
			myImg.SaveFileResized("tmp")

			//			var testImageHandler ImageHandler
			//			testImageHandler = mImg1
			//			testImageHandler.ReadFileFromPath(tc.file)

			//			testImageHandler.ReadFileFromPath("test02.jpg")
			// loc, err := time.LoadLocation(tc.loc)
			// if err != nil {
			// 	t.Fatal("could not load location")
			// }
			// gmt, _ := time.Parse("15:04 2006", tc.gmt)
			// got := gmt.In(loc).Format("15:04")
			// log.Println("gmt=", gmt, " loc=", loc, "got=", got)
			// if got != tc.want {
			// 	t.Errorf("got %s; want %s", got, tc.want)
			// }
		})
	}
}

func TestTime(t *testing.T) {
	testCases := []struct {
		gmt  string
		loc  string
		want string
	}{
		{"12:31 1970", "Europe/Zurich", "13:31"},
		{"12:31 1970", "America/New_York", "07:31"},
		{"12:31 1970", "Australia/Sydney", "22:31"},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s in %s", tc.gmt, tc.loc), func(t *testing.T) {
			loc, err := time.LoadLocation(tc.loc)
			if err != nil {
				t.Fatal("could not load location")
			}
			gmt, _ := time.Parse("15:04 2006", tc.gmt)
			got := gmt.In(loc).Format("15:04")
			log.Println("gmt=", gmt, " loc=", loc, "got=", got)
			if got != tc.want {
				t.Errorf("got %s; want %s", got, tc.want)
			}
		})
	}
}

func TestMain(m *testing.M) {
	log.Println("@@@ TestMain")
	setup()
	m.Run()
	teardown()
}
