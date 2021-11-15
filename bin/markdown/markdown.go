package main

import (
	"embed"
	"fmt"
	"path/filepath"

	"github.com/gomarkdown/markdown"
)

//go:embed files/*.md
var files embed.FS

// Our help files
var help = map[string][]byte{
	"index":   getMarkdownInBytes("index.md"),
	"brocade": getMarkdownInBytes("brocade.md"),
}

func main() {

	fmt.Println(string(help["brocade"]))
}

func getMarkdownInBytes(filename string) []byte {
	md, _ := files.ReadFile(filepath.Join("files", filename))
	output := markdown.ToHTML(md, nil, nil)
	return output
}
