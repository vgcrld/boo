package main

import "fmt"

func main() {
	f()
	fmt.Println("Returned normally from f.")
}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println("Calling g.")
	g(-4)
	fmt.Println("Returned normally from g.")
}

// panic goes all the way up the stack when i > 4
// The recursion is ended an the defers are called.
// all the way up to f() in main.
// But the default in f() checks if recover(). That stops the
// panic just before we go back to main.
func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
}
