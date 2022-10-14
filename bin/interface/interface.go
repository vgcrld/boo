package main

import (
	"fmt"

	"github.com/vgcrld/boo/bin/interface/icom"
)

// Interface into the characters
type Animator interface {
	Talk()
	Walk()
}

// A Human Being
type Person struct {
	Name string
}

// Make the Human talk
func (p *Person) Talk() {
	fmt.Println("Hello, World!")
}

// Make the Human talk
func (p *Person) Walk() {
	fmt.Println("Here I am on 2 feet")
}

// A Dog animal
type Dog struct {
	Name string
}

// Make the dog talk
func (d *Dog) Talk() {
	fmt.Println("Bow Wow")
}

// Make the dog walk
func (d *Dog) Walk() {
	fmt.Println("I'm down on all fours.")
}

// Start of the main PGM.
func main() {

	icom.HELLO()

	hh := icom.NewHolder()
	hh.Name = "Richard"
	hh.Age = 52
	hh.Address = "13 Lovalee Lane"
	hh.Code = 14
	fmt.Printf("%T: Name=%v\n", hh, hh.Name)

	// We create the interface
	var a Animator

	// A person that is really cool.
	var p Person

	// A dog that is amazing
	var d Dog

	// Create these
	p, d = Person{"Richard"}, Dog{"fido"}

	// and the interface can be assigned to a person
	a = &p
	a.Talk()
	a.Walk()

	// or a dog and the proper method is called based on the type
	a = &d
	a.Talk()
	a.Walk()

	var aa []Animator
	// aa = append(aa, &p, &d, &p, &p, &p)
	aa = append(
		aa,
		&Dog{Name: "Spot"},
	)
	for _, x := range aa {
		switch v := x.(type) {
		case *Dog:
			fmt.Println("The name is")
			fmt.Println(v.Name)
		}
		fmt.Printf("%T\n", x)
		x.Talk()
		x.Walk()
	}
}
