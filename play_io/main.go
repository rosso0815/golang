package main

/*
*	shows the io package to read / write a file
*
*
 */

import (
	//	"bufio"
	"flag"
	"fmt"
	"log"
	//	"io"
	//"io/ioutil"
	"os"
	"pfistera/play_io/io"
)

const APP_VERSION = "0.1"

var versionFlag *bool = flag.Bool("v", false, "Print the version number.")

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	fmt.Println("start")

	//fmt.Println()

	os.Remove("test.txt")

	play_io.WriteFile("test.txt")

	/*
		var file = "test.txt"

		play_io.SayHi()

		log.Println("start file = ", file)

		flag.Parse() // Scan the arguments list

		if *versionFlag {
			log.Println("Version:", APP_VERSION)
			panic("out")
		}

		fmt.Println("")
		dat, err := ioutil.ReadFile(file)
		check(err)
		fmt.Print(string(dat))
		fmt.Println("")
	*/

	//	os.Remove("test.txt")

	log.Println("ende")
}
