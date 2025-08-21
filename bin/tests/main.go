package main

import (
	"fmt"
)

type Number interface {
	int | float32 | float64
}

func main() {

	a := Add(1, 2)
	fmt.Println(a)

}

func Add[T Number](a T, b T) T {
	return a + b
}
