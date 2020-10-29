package myconcurrency

import (
	"fmt"
	"strconv"
	"sync"
	"testing"
	"time"
)

// Test_SumBasic : basic test without concurrecncy
func Test_SumBasic(t *testing.T) {
	total := Sum(5, 5)
	if total != 10 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	}
}

// Here's the worker, of which we'll run several
// concurrent instances. These workers will receive
// work on the `jobs` channel and send the corresponding
// results on `results`. We'll sleep a second per job to
// simulate an expensive task.
func worker303(id int, jobs <-chan int, results chan<- int) {

	fmt.Println("@@@ worker", id)

	//j := <-jobs
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}
func TestSum4(t *testing.T) {

	// In order to use our pool of workers we need to send
	// them work and collect their results. We make 2
	// channels for this.
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// This starts up 3 workers, initially blocked
	// because there are no jobs yet.
	for w := 1; w <= 5; w++ {
		go worker303(w, jobs, results)
	}

	// Here we send 5 `jobs` and then `close` that
	// channel to indicate that's all the work we have.
	for j := 1; j <= 10; j++ {
		jobs <- j
	}
	close(jobs)

	// Finally we collect all the results of the work.
	for a := 1; a <= 10; a++ {
		<-results
	}
}

func worker03(wg *sync.WaitGroup, cs chan string, i int) {
	defer wg.Done()
	cs <- "worker" + strconv.Itoa(i)
}

func monitorWorker(wg *sync.WaitGroup, cs chan string) {
	wg.Wait()
	close(cs)
}
func TestSum2(t *testing.T) {
	wg := &sync.WaitGroup{}
	cs := make(chan string)
	go monitorWorker(wg, cs)
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go worker03(wg, cs, i)
	}

	for i := range cs {
		fmt.Println(i)

	}
}
