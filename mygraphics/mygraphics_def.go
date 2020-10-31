package mygraphics

// Image structur to handle images
type Image struct {
	width   uint
	heigth  uint
	make    string
	model   string
	created string
	path    string
}

// ImageHandler used for internal handling of a internal array of images
type ImageHandler interface {
	AddFileFromPath(path string) (index int, err error)
	GetInfo(index int) (img Image, err error)
	SaveFileResized(index int, targetDir string) (err error)
}
