package main

import (
	"fmt"

	"github.com/vgcrld/boo/bin/interface/icom"
)

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
	var a icom.Animator

	// A person that is really cool.
	var p icom.Person

	// A dog that is amazing
	var d icom.Dog

	// Create these
	p, d = icom.Person{Name: "Richard"}, icom.Dog{"fido"}

	// and the interface can be assigned to a person
	a = &p
	a.Talk()
	a.Walk()

	// or a dog and the proper method is called based on the type
	a = &d
	a.Talk()
	a.Walk()

	var aa []icom.Animator
	// aa = append(aa, &p, &d, &p, &p, &p)
	aa = append(
		aa,
		&icom.Dog{Name: "Spot"},
	)
	for _, x := range aa {
		switch v := x.(type) {
		case *icom.Dog:
			fmt.Println("The name is")
			fmt.Println(v.Name)
		}
		fmt.Printf("%T\n", x)
		x.Talk()
		x.Walk()
	}
}
