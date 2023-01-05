package main

import (
	"embed"
	"fmt"
	"net/http"
)

//go:embed public/*
var files embed.FS

func main() {

	index, err := files.ReadFile("public/index.html")
	if err != nil {
		fmt.Println("this jawn didn't work")
	}
	// fmt.Print(string(index))
	// os.Exit(1)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, string(index))
	})
	http.ListenAndServe(":8000", nil)
}
