package main

import (
	"log"
	"testing"
	"time"
)

func routine123(c chan int) {

	var sum int

	for {
		select {
		case msg := <-c:
			sum = sum + msg
			log.Println("answer again", sum, msg)
			//case <-quit:
			//	log.Println("quit")
			//	return

			//default:
			//	log.Println(".")
			//			time.Sleep(50 * time.Millisecond)
		}
	}
}

func TestChannel02(t *testing.T) {

	in := make(chan int, 0)

	go routine123(in)

	in <- 1
	in <- 2
	in <- 3

	//	time.Sleep(200 * time.Millisecond)

	in <- 5

	//go answer(9999)
	//quit <- 0

	//	in <- 3
	//quit <- 0
	//x := <-in
	//l.Println("in=", x)

	log.Println("done")
	time.Sleep(1 * time.Second)
}
