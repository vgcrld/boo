package common

import (
	"fmt"
	"math/rand"
	"time"
)

func PrintSomeRandomeSlices() {
	for i := 0; i < 10; i++ {
		fmt.Println(RandomeSizeIntSlice(30))
	}
}

func RandomeSizeIntSlice(s int) []int {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	top := r1.Intn(s)
	ret := []int{}
	for i := 0; i <= top; i++ {
		ret = append(ret, i)
	}
	return ret

}
