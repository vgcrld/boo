package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	wg sync.WaitGroup
	ch chan int
)

const (
	threads  = 10
	values   = 500
	max_time = 5000 // Maximum sleep time in milliseconds
)

func main() {

	ch = make(chan int)

	for i := range threads {
		wg.Add(1)
		go fk(i)
	}

	for range values {
		ch <- 10 + rand.Intn(41)
	}
	close(ch)
	wg.Wait()

	fmt.Println("All goroutines have finished")
	fmt.Println("Exiting main function")
}

func fk(tname int) {
	fmt.Printf("Starting goroutine %d\n", tname)
	for n := range ch {
		st := rand.Intn(max_time)
		result := n * 2 // Example calculation
		fmt.Printf("go%d Received: %d, sleeptime: %dms, calc: %d * 2 = %d\n", tname, n, st, n, result)
		time.Sleep(time.Duration(st) * time.Millisecond)
	}
	fmt.Printf("Goroutine %d finished\n", tname)
	wg.Done()
}
