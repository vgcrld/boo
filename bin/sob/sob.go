package main

import "fmt"

func main() {

	var bitch uint
	p(
		"welcome",
		"you sick",
		200,
		"year old",
		"son of a",
		&bitch,
	)
}

// Helper 'cause fmt.Println("") every 5 seconds is annoying
func p(i ...interface{}) {
	fmt.Println(i...)
}
