package main

import (
	"embed"
	"fmt"
	"net/http"
)

//go:embed public/*
var files embed.FS

func main() {

	http.HandleFunc("/grid", makeHandleFunc("index.html"))
	http.HandleFunc("/file", makeHandleFunc("file.json"))
	http.ListenAndServe(":8000", nil)
}

func makeHandleFunc(f string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		file := getFromEmbed(f)
		fmt.Fprintf(w, string(file))
	}
}

func getFromEmbed(f string) string {
	file, err := files.ReadFile("public/" + f)
	if err != nil {
		file = []byte("file not found in public")
	}
	return string(file)
}
