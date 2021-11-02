// Golang program to show how to
// use structs as map keys
package main

// importing required packages
import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
	"github.com/maple-tech/go-hashify"
)

//declaring a struct
type Address struct {
	Name    string
	city    string
	Pincode int
}

func main() {

	// Creating struct instances
	a2 := Address{Name: "Ram", city: "Delhi", Pincode: 2400}
	a1 := Address{"Pam", "Dehradun", 2200}
	a3 := Address{Name: "Sam", city: "Lucknow", Pincode: 1070}

	// Declaring a map
	var mp map[Address]int

	// Checking if the map is empty or not
	if mp == nil {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
	// Declaring and initialising
	// using map literals
	sample := map[Address]int{a1: 1, a2: 2, a3: 3}
	spew.Dump(sample)

	hash, _ := hashify.SHA1String(sample)
	fmt.Println(hash)
}
