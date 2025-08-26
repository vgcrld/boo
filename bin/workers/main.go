package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	wg sync.WaitGroup
	ch chan int
)

// Default values for the application.
const (
	threads  = 2
	values   = 50
	max_time = 5000 // Maximum sleep time in milliseconds
)

func main() {

	t := flag.Int("threads", threads, "Number of goroutines to start")
	v := flag.Int("values", values, "Number of values to process")
	m := flag.Int("max-time", max_time, "Maximum sleep time in milliseconds")

	flag.Parse()

	fmt.Printf("Starting %d goroutines to process %d values with max sleep time %dms\n", *t, *v, *m)

	ch = make(chan int)

	for i := range *t {
		wg.Add(1)
		go fk(i, *m)
	}

	for range *v {
		ch <- 10 + rand.Intn(41)
	}
	close(ch)
	wg.Wait()

	fmt.Println("All goroutines have finished")
	fmt.Println("Exiting main function")
}

func fk(tname, msecs int) {
	fmt.Printf("Starting goroutine %d\n", tname)
	for n := range ch {
		st := rand.Intn(msecs)
		fmt.Printf("go%d Received: %d, sleeptime: %dms\n", tname, n, st)
		time.Sleep(time.Duration(st) * time.Millisecond)
	}
	fmt.Printf("Goroutine %d finished\n", tname)
	wg.Done()
}
