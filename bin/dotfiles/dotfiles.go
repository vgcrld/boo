/*

	This file is used to setup unix .dot files like
	.vimrc and .bashrc.

*/
package main

import (
	"embed"
	"io/ioutil"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
)

//go:embed files/*
var files embed.FS

var home string = os.Getenv("HOME")

func main() {
	writeOutFile(".vimrc", "files/vimrc")
	writeOutFile("README.md", "files/README.md")
	writeOutFile("HELP.md", "files/helpme.md")
}

func writeOutFile(ofile, ifile string) {
	path := filepath.Join(home, ofile)
	log.WithFields(log.Fields{"filepath": path}).Info("writing the .vimrc file")
	vimrc, _ := files.ReadFile(ifile)
	ioutil.WriteFile(path, vimrc, 0644)

}
