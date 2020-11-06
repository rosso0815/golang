package mygraphics

// Image structur to handle images
type Image struct {
	width   uint
	heigth  uint
	emake   string
	model   string
	created string
	path    string
}

// ImageHandler used for internal handling of a internal array of images
//type ImageHandler interface {
//	AddFileFromPath(path string) (index int, err error)
//	//AllGetInfo(index int) (img Image, err error)
//	//AllSaveFileResized(index int, targetDir string) (err error)
//}
