package main

import (
	"log"
	"testing"
)

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func intAdd() func(int) int {
	i := 0
	return func(y int) int {
		i = i + y
		return i
	}
}

func Test_Cars(te *testing.T) {
	log.Println("@@@ Test_Car")

	t := intSeq()
	i := intAdd()

	log.Println("seq=", t())
	log.Println("seq=", t())
	log.Println("seq=", t())
	log.Println("seq=", t())

	log.Println("add=", i(5))
	log.Println("add=", i(5))

	stock := 1234
	p := func(i int) int {
		log.Println("stock=", stock)
		return i * i
	}

	log.Println("stock=", p(2))

	stock = 999

	log.Println("stock=", p(2))
	log.Println("end")
}
