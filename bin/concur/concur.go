package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/vgcrld/boo/bin/concur/common"
)

var (
	wg         sync.WaitGroup
	workerWait sync.WaitGroup
)

type lambda struct {
	val      int
	name     string
	function func(i int) int
}

func main() {

	// runTickers()
	// channelExampleWithStruct()
	// channelExampleBasic()
	// channelExampleWithGoRoutine()
	// common.PrintSomeRandomeSlices()

	channelWorker()

	// r := common.RandomeSizeIntSlice(10)
	// fmt.Println(r)
}

func goWorker(c chan lambda) {
	fmt.Println("Starting Working Go routine...")
	for {
		val := <-c
		val.function(val.val)
		// fmt.Printf("Do the math: %v%v%v=%v\n", val.val, val.name, val.val, result)
	}
}

func sum(i int) int {
	r := i + i
	return r
}

func times(i int) int {
	return i * i
}

func channelWorker() {

	// Our Queue is a type lambda
	var queue chan lambda
	queue = make(chan lambda)

	for range 3 {
		workerWait.Add(1)
		go goWorker(queue)

	}

	st := time.Now()
	for i := 0; i < 1e7; i++ {
		queue <- lambda{val: i, name: "+", function: sum}
		queue <- lambda{val: i, name: "*", function: times}
	}

	fmt.Println("We are done here:", time.Now().Sub(st))

}

func channelExampleWithStruct() {

	// Remember the flow, follow the main go thread and if not buffered the add to chan will block
	// and a single pull from chan will block.
	queue := make(chan common.Event, 100)
	var w sync.WaitGroup
	w.Add(1)
	go func() {

		student := []common.User{
			{
				FirstName: "Richard",
				LastName:  "Davis",
				Email:     "vgcrld@gmail.com",
				Tickets:   5,
			},
			{
				FirstName: "Robert",
				LastName:  "Davis",
				Email:     "vgcrwd@gmail.com",
				Tickets:   1,
			},
		}
		event := common.NewEvent("Math 101", "Malvern, PA")
		event.Students = student
		queue <- event
		queue <- common.NewEvent("Math 102", "Malvern, PA")
		queue <- common.NewEvent("Math 201", "Malvern, PA")
		queue <- common.NewEvent("Science 201", "Malvern, PA")
		fmt.Println("done waiting")
		w.Done()
	}()
	fmt.Println("About to wait")
	w.Wait()
	for len(queue) > 0 {
		val := <-queue
		fmt.Println(val.Class)
		fmt.Println(val.Location)
		fmt.Println(val.Students)

	}
}

// use channel in a go routine. channels block so you have
// to be careful to open then use them and make sure nothing
// is hanging around waiting for stuff.
// Especially in the main go routine. You can't block the main! that what that means.
// extra go routines won't block then will just die with the main.
func channelExampleWithGoRoutine() {

	// This routine takes the messages channel, sleeps and then puts true on it.
	// this go routine will block until messages is ready to receive. if that
	// never happens and the main go routine ends it will just go away.
	// It may even finish before it can print Working because the main go routine is moving on.
	messages := make(chan bool)
	go func() {
		fmt.Println("Working")
		time.Sleep(time.Second)
		fmt.Println("Done sleeping")
		messages <- true // If there is mothing to pull this off it will stop here.
		messages <- true // If there is mothing to pull this off it will stop here.
		messages <- true // If there is mothing to pull this off it will stop here.
		messages <- true // If there is mothing to pull this off it will stop here.
		fmt.Println("finished, len:", len(messages))
		close(messages) // If you don't close it will block and cause deadlock. because range in the main will keep waiting.
	}()

	// time.Sleep(time.Second * 4)
	// If you don't close the channel in the qo routine then
	// this range will deadlock - saying the routines are asleep.
	for i := range messages {
		fmt.Println("pulled off value", i)
	}

}

// Play with channels a bit
func channelExampleBasic() {
	// Super simple. Make, add, pull off end.
	// a channel can send and receive. <-channel (send), channel<- (receive)
	// this only works in the main go routine because it is buffered (10)
	// the minute the buffer is full (default 1) to main go will wait until
	// something pulls it off the queue.
	queue := make(chan string, 10)
	defer close(queue)
	queue <- "My name is rich"
	fmt.Println(<-queue)

	// load it up and pull it off
	queue <- "this"
	queue <- "is"
	queue <- "a"
	queue <- "test"
	for len(queue) > 0 {
		fmt.Println(<-queue)
	}
}

// Run some tickers
func runTickers() {
	short := makeTick(100)
	medium := makeTick(1000)
	long := makeTick(1500)

	fmt.Println("Staring Go Routines...")

	wg.Add(3)
	go runner("long message", 2, long) // this wont block the short ones
	go runner("medium queue", 5, medium)
	go runner("short stuff", 25, short)
	wg.Wait()

}

// Return a function that is sleep of seconds
func makeTick(seconds int64) func() {
	return func() {
		time.Sleep(time.Millisecond * time.Duration(seconds))
	}
}

// Just do some work
func runner(name string, count int, fn func()) {
	for i := 0; i < count; i++ {
		fmt.Println(name, i)
		fn()
	}
	wg.Done()
}
