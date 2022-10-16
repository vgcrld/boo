package main

import (
	"fmt"
	"time"
)

func main() {

	sec1 := makeTick(1)
	sec2 := makeTick(2)

	go func() {
		i := 0
		for {
			i += 1
			fmt.Println("hello", i)
			sec2()
		}
	}()

	for {
		fmt.Println("here we go")
		sec1()
	}
}

func makeTick(i int64) func() {
	return func() {
		time.Sleep(time.Second * time.Duration(i))
	}
}
