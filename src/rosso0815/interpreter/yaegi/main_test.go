package main

import (
	"log"
	"testing"
)

func TestSum(t *testing.T) {
	log.Println("@@@ Test Sum")
	var i = Sum(1, 1)
	log.Println("1 + 1 =", i)

}
func TestYaegi(t *testing.T) {
	log.Println("@@@ TestAbs")
	CallYaegi()

}
