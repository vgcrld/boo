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

	addCentOsVim()
	writeOutFile(".vimrc", "files/vimrc")
	addPathogen()
	addNerdtree()

}

func writeOutFile(ofile, ifile string) {
	path := filepath.Join(home, ofile)
	log.WithFields(log.Fields{"filepath": path}).Info("writing the .vimrc file")
	vimrc, _ := files.ReadFile(ifile)
	ioutil.WriteFile(path, vimrc, 0644)

}

func runOsCommand(cmd string, args ...string) bool {

	log.WithFields(log.Fields{"command": cmd}).Info("running command")
	command := exec.Command(cmd, args...)
	stdout, err := command.Output()
	if err != nil {
		return false
	}
	fmt.Println(string(stdout))
	return true

}

func addCentOsVim() {

	if ok := runOsCommand("apt-get", "update"); ok {
		runOsCommand("apt-get", "install", "-y", "vim")
	} else {
		panic("unable to add vim with apt-get")
	}

}

func addPathogen() {

	if ok := runOsCommand("mkdir", "-p", filepath.Join(home, ".vim/autoload"), filepath.Join(home, ".vim/bundle")); ok {
		runOsCommand("curl", "-LSso", filepath.Join(home, ".vim/autoload/pathogen.vim"), "https://tpo.pe/pathogen.vim")
	} else {
		panic("unable to add pathogen")
	}

}

func addNerdtree() {

	nerdtreeCmd := []string{
		"clone",
		"https://github.com/preservim/nerdtree.git",
		filepath.Join(home, ".vim/bundle/nerdtree"),
	}
	ok := runOsCommand("git", nerdtreeCmd...)
	if !ok {
		panic("unable to add pathogen")
	}

}
