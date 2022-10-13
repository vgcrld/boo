package icom

import "fmt"

type Holder struct {
	Name    string
	Age     int
	Address string
	Code    rune
}

func NewHolder() Holder {
	return Holder{
		Name:    "",
		Age:     0,
		Address: "",
		Code:    0,
	}
}

func HELLO() {
	fmt.Println("hello, people")
}
