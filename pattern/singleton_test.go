package main

/*

	Author : pfistera
	Purpose : test of Patterns in Golang

*/

import (
	"fmt"
	"testing"

	"pfistera/pattern/singleton"
)

func TestSingleton(t *testing.T) {

	fmt.Println("@@@ TestSingleton")

	s1 := singleton.New()
	s1["this"] = "that"
	fmt.Println("This is ", s1["this"])

	s2 := singleton.New()
	fmt.Println("This is ", s2["this"])
}
