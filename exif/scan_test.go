package exif

import (
	"fmt"
	"testing"
)

func TestXxx(*testing.T) {
	fmt.Println("read 20170312_154628.jpg")

	file := "20170312_154628.jpg"

	e := Exif{}

	_ = e.ReadFile(file)

}
