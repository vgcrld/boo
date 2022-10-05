package main

import "fmt"

// iota is incremented each time it is executed.
// it only works as a const

const (
	_ = 1 << (10 * iota)
	KiB
	MiB
	GiB
	TiB
	PiB
)

// It gets reset here.
const xx = iota

func main() {
	fmt.Println(KiB, MiB, GiB, TiB, PiB)
	fmt.Println(xx)
}
