package main

/*
	This code must run on Windows.
	It calls the WMI locally.

*/
import (
	"fmt"
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/yusufpapurcu/wmi"
)

// In this wmi package you create the type that matches the class you want to
// collect and then you add the members that you need.
// see as an example:
// https://docs.microsoft.com/en-us/windows/win32/cimwin32prov/win32-process
type Win32_Process struct {
	Name            string
	ParentProcessId uint32
}

type Win32_DiskPartition struct {
	Name string
}

func main() {
	disks()
}

func processes() {
	var dst []Win32_Process
	q := wmi.CreateQuery(&dst, "")
	err := wmi.Query(q, &dst)
	if err != nil {
		log.Fatal(err)
	}
	spew.Dump(dst)
}

func disks() {
	var dst []Win32_DiskPartition
	q := wmi.CreateQuery(&dst, "")
	err := wmi.Query(q, &dst)
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range dst {
		fmt.Println(v.Name)
	}

}
