package main

import (
	"flag"
	"fmt"
	"log"
	"pfistera/readCSV"
)

func main() {
	wordPtr := flag.String("word", "foo", "a string")
	pPtr := flag.String("p", "foo", "a string")
	numbPtr := flag.Int("numb", 42, "an int")
	flag.Parse()
	fmt.Println("word:", *wordPtr)
	fmt.Println("p:", *pPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("tail:", flag.Args())

	// start
	log.Print("start")
	// read
	readCSV.Read("test.csv")
	readCSV.Clear()
}
