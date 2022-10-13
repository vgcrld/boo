package shr

import "fmt"

func RunPointers() {
	fmt.Printf("v%v.%v.%v.%v\n", Version, Release, Modification, Patch)
	i := pointers()
	fmt.Println("this will be a pointer to the struct:", i)
	// to access i you need to dereference it *i
	for _, x := range *i {
		fmt.Println(x.Name, x.Age)
	}
}

// send back this struct pointer
func pointers() *[]struct {
	Name string
	Age  int
} {
	x := []struct {
		Name string
		Age  int
	}{
		{"richard", 52},
		{"robert", 56},
	}
	return &x
}
