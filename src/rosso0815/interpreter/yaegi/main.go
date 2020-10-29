package main

import (
	"log"

	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"
)

// CallYaegi = test without args ...
func CallYaegi() {
	log.Println("@@@ Callyaegi")
	i := interp.New(interp.Options{})
	i.Use(stdlib.Symbols)

	_, err := i.Eval(`import "fmt"`)

	if err != nil {
		panic(err)
	}

	_, err = i.Eval(`fmt.Println("Hello Yaegi")`)
	if err != nil {
		panic(err)
	}

}

// Sum = return x + y
func Sum(x int, y int) int {
	log.Println("@@@ Sum ", x, y)
	return x + y
}

func main() {

	log.Println("@@@ start")

	log.Println("1+2 =", Sum(1, 2))

	CallYaegi()

}
