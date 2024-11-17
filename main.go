package main

import (
	"fmt"
)

const (
	_ = iota
	a
	b
	c
)

func main() {

	// This is a slice of anonymous structs
	xx := []struct {
		name string
		age  int
	}{
		{name: "richard", age: 52},
		{name: "susan", age: 52},
		{name: "robert", age: 52},
		{name: "mark", age: 52},
		{name: "joe", age: 52},
		{name: "ben", age: 52},
		{name: "lizzi", age: 52},
		{name: "harry", age: 52},
		{name: "jesse", age: 52},
	}

	for _, x := range xx {
		fmt.Println(x.name)
	}

}

func printStruct(xx []struct {
	name string
	age  int
}) {
	fmt.Println(xx)
}

func printTemp(f func() string) {
	fmt.Println(f())
}
