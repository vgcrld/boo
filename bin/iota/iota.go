package main

import (
	"fmt"
	"math"
)

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
const newIota = iota

func main() {
	fmt.Printf("KiB=%v, MiB=%v, GiB=%v, TiB=%v, PiB=%v\n", KiB, MiB, GiB, TiB, PiB)
	fmt.Println(newIota, "ioto is reset outside of the const")
	xx := 127.0 / 5.0
	i, f := math.Modf(xx)
	fmt.Printf("int=%v, fractional=%v\n", i, f)

}
