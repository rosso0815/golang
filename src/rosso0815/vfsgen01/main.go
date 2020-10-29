//go:generate go run generate/generate_static.go

package main

import (
	"bufio"
	"fmt"
	"log"

	"pfistera/vfsgen01/assets"
)

func main() {
	log.Println("@@@ start")

	file, err := assets.Assets.Open("index.html")

	if err != nil {
		log.Panicln("error")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	log.Println("--- scan file start")
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
	log.Println("--- scan file finished")
}
