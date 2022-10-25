package main

import (
	"fmt"
	"os"
	"regexp"
)

func main() {

	boo, _ := regexp.Compile("^boo.*")

	for i, v := range os.Args {
		fmt.Printf("Index %v : %v\n", i, v)
	}

	fmt.Println(boo.MatchString("richard")) // false
	fmt.Println(boo.MatchString("rebook"))  // fasee
	fmt.Println(boo.MatchString("book"))    // true
}
