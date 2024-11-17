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

func processUsers(users []struct {
	name string
	age  int
}) {
	for _, user := range users {
		fmt.Printf("Name: %s, Age: %d\n", user.name, user.age)
	}
}

func main() {
	processUsers([]struct {
		name string
		age  int
	}{
		{name: "Alice", age: 30},
		{name: "Bob", age: 25},
		{name: "Bob", age: 25},
		{name: "Bob", age: 25},
		{name: "Bob", age: 25},
		{name: "Bob", age: 25},
	})

}
