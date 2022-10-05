package main

import (
	"fmt"
	"os"
)

const (
	_ = iota
	a
	b
	c
)

func main() {

	xx := []struct {
		name string
		age  int
	}{
		{name: "richard", age: 52},
		{name: "richard", age: 52},
		{name: "richard", age: 52},
		{name: "richard", age: 52},
		{name: "richard", age: 52},
		{name: "richard", age: 52},
		{name: "richard", age: 52},
		{name: "richard", age: 52},
		{name: "richard", age: 52},
	}

	printStruct(xx)
	printTemp(os.TempDir)
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
