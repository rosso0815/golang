package main

import (
	"fmt"
	"pfistera/my_imgfix/resize"
	"time"
	//	"github.com/pkg/profile"
)

func main() {

	start := time.Now()

	// TODO sssss
	//defer profile.Start().Stop()
	//	defer profile.Start(profile.CPUProfile).Stop()

	fmt.Println("start", time.Since(start))

	resize.Resize(".", "input.jpg", false, false, start)

	fmt.Println("done", time.Since(start))
}
