package main

import "fmt"

func main() {

	// here you must use [5]int (arrays are too structured)
	fmt.Println("An Array")
	x := [5]int{1, 2, 3, 4, 5}
	updateArray(2, 100, &x)
	updateArray(0, 1001, &x)
	fmt.Println(x)

	fmt.Println("A Slice")
	u := make([]int, 1000)
	updateSlice(2, 100, &u)
	updateSlice(0, 11, &u)
	updateSlice(100, 1024, &u)
	updateSlice(200, 1024, &u)
	fmt.Println(u)

}

// here you must recieve [5]int
func updateArray(idx, val int, a *[5]int) {
	(*a)[idx] = val
}

func updateSlice(idx, val int, a *[]int) {
	(*a)[idx] = val
}
