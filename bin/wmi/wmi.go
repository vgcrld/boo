package main

import (
	"log"

	"github.com/StackExchange/wmi"
)

type Win32_Process struct {
	Name string
}

func main() {
	var dst []Win32_Process
	q := wmi.CreateQuery(&dst, "")
	err := wmi.Query(q, &dst)
	if err != nil {
		log.Fatal(err)
	}
	for i, v := range dst {
		println(i, v.Name)
	}
}
