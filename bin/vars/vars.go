package main

import "fmt"

var someVar string = "outside"

func main() {

	// someVar here is scoped to the block, result inside
	{
		var someVar string = "inside"
		fmt.Println(someVar)
	}

	// result here is outside
	fmt.Println(someVar)

	// someVar here is scoped to the global and can be changed
	{
		someVar = "bob"
		fmt.Println(someVar)
	}

	// result is bob
	fmt.Println(someVar)

}
