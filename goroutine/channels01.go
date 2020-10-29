package main

import (
	"fmt"
	"sync"
	"time"
)

func giveAnswer(wg *sync.WaitGroup, in chan int, out chan int) {
	for {
		fmt.Println("@@@ give_answer")
		select {
		case <-in:
			out <- 1
			wg.Done()
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func main() {
	x := 0
	y := 0
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x = <-c
	y = <-c
	go sum(s[len(s)/2:], c)
	y = <-c
	fmt.Println(x, y, x+y)

	/*
		var wg sync.WaitGroup
		in := make(chan int)
		out := make(chan int)
		go give_answer(&wg, in, out)
		for {
			out <- 9
			in <- 9
			for {
				select {
				case <-out:
					fmt.Println("out")
				}
			}
			wg.Wait()
		}*/

}
