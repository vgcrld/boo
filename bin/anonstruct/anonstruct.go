package main

import (
	"fmt"

	"github.com/vgcrld/boo/commons"
	"github.com/vgcrld/boo/commons/ui"
)

func main() {

	defer func() {
		xx := []struct {
			name string
			age  int
		}{
			{name: "bob", age: 56},
			{name: "sue", age: 54},
			{name: "richard", age: 52},
		}
		fmt.Println("bye!", xx)
		for _, e := range xx {
			fmt.Println("Struct Name", e.name)
			fmt.Println("Struct Age", e.age)
		}
	}()

	fmt.Printf("Welcome to version= %v, ui=%v\n", commons.VERSION, ui.VERSION)

}
