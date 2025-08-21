package main

import (
	"fmt"
)

func main() {

	//create a 2d array of names. mix man and womens names. may it 10x10
	names := [10][10]string{
		{"Alice", "Bob", "Charlie", "Diana", "Ethan", "Fiona", "George", "Hannah", "Ian", "Jasmine"},
		{"Kieran", "Lily", "Mike", "Nina", "Oscar", "Paula", "Quentin", "Rachel", "Sam", "Tina"},
	}
	//iterate over the array and print each name
	for i := 0; i < len(names); i++ {
		for j := 0; j < len(names[i]); j++ {
			if names[i][j] != "" { // Check if the name is not empty
				fmt.Printf("Name at %d%d: %s\n", i, j, names[i][j])
			}
		}
	}

}

func sayIt(names ...string) {
	for _, name := range names {
		fmt.Printf("Hello, %s!\n", name)
	}

}
