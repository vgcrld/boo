package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {

	// copy();
	// pipe()
	readAll()

}

func copy() {
	r := strings.NewReader("some io.Reader stream to be read\n")

	// Copy r to Stdout
	// if z, err := io.Copy(os.Stdout, r); err != nil {
	if z, err := io.Copy(os.Stdout, r); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("data", z)
	}
}

func pipe() {
	r, w := io.Pipe()

	go func() {
		fmt.Fprint(w, "some io.Reader stream to be read\n")
		w.Close()
	}()

	if _, err := io.Copy(os.Stdout, r); err != nil {
		log.Fatal(err)
	}

}

func readAll() {
	r := strings.NewReader("Go is a general-purpose language designed with systems programming in mind.")

	b, err := io.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", b)
}
