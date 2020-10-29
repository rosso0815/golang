package main

import (
	"fmt"
	"github.com/xiam/exif"
)

func main() {
	data, _ := exif.Read("20170312_154628.jpg")
	for key, val := range data.Tags {
		fmt.Printf("%s = %s\n", key, val)
	}
}
