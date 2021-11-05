/*

	This file is used to setup unix .dot files like
	.vimrc and .bashrc.

*/
package main

import (
	"embed"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"

	log "github.com/sirupsen/logrus"
)

//go:embed files/*
var files embed.FS

var home string = os.Getenv("HOME")

func main() {

	runOsCommand("mkdir", "-p", "/tmp/bob")
	runOsCommand("touch", "/tmp/bob/file1")
	runOsCommand("touch", "/tmp/bob/file2")
	runOsCommand("touch", "/tmp/bob/file3")
	runOsCommand("ls", "-l", "/tmp/bob")

	// Write out some files
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

func runOsCommand(cmd string, args ...string) bool {
	command := exec.Command(cmd, args...)
	stdout, err := command.Output()
	if err != nil {
		return false
	}
	fmt.Println(string(stdout))
	return true
}
