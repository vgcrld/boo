package icom

import "fmt"

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
