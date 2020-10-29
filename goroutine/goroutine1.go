package main

import (
	l "log"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func fib(n uint) uint {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fib(n-1) + fib(n-2)
	}
}

func parallel01(number int, goGroup *sync.WaitGroup) {
	l.Println("go routine start", number)
	l.Println("fib=", fib(44))
	l.Printf("go routine %d stop \n", number)
	goGroup.Done()
}

func main() {
	MaxRoutines := 5
	gomaxprocs := 1

	rand.Seed(time.Now().Unix())

	l.Println("### number of cpus = ", runtime.NumCPU())
	l.Println("### defined GOMAXPROCS = ", gomaxprocs)

	runtime.GOMAXPROCS(gomaxprocs)

	goGroup := new(sync.WaitGroup)

	for i := 0; i <= MaxRoutines; i++ {

		//seconds := rand.Intn(MaxSleep)

		l.Println("start goroutine", i)
		go parallel01(i, goGroup)
		goGroup.Add(1)
	}

	goGroup.Wait()

	l.Printf("*** all done ***\n")

}
