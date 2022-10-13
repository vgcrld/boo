package main

import (
	"fmt"

	"github.com/fatih/color"
)

func main() {

	// Make the closures vars. fn1, fn2 are functions (ready to go)
	// i+=i so 1, 2, 4, etc.. Start with 1, 12
	fn1 := closureTest(1)
	fn2 := closureTest(12)
	/*
		The output of the following is this.
		The cool part of closures is they wrap
		the local variable and they keep that value with them.

		1 12
		2 24
		4 48
		8 96
		16 192
		32 384
		64 768
		128 1536
		256 3072
		512 6144
	*/
	for i := 0; i < 10; i++ {
		p(fn1(), fn2())
	}

	// The fn1, fn2 functions hold their current value, and can be
	// continually incremented.
	p("let's print fn2 again. It holds the last value")
	pf("fn2=%v\n", fn2())

	p("and again")
	pf("This is fn2=%v\n", fn2())

	// Added in some color this is kind of cool.
	color.Cyan("This is in color, baby!")
	c := color.New(color.FgRed).Add(color.Underline)
	c.Println("You go now!")
}

// This will print anything passed to it
func p(i ...interface{}) {
	fmt.Println(i...)
}

// This will print anything passed to it using fmt.Printf
func pf(f string, i ...interface{}) {
	fmt.Printf(f, i...)
}

// Return a function that closes "sum" that is local.
// This function then holds it's own sum
func closureTest(startVal int) func() int {
	sum := startVal
	return func() int {
		defer func() { sum += sum }()
		return sum
	}
}
