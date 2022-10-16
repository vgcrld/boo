package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	fmtScanln()
	// bufioReader()
	// bufScanner()
}

// Scanln reads to the new line but you need a list of things to read
func fmtScanln() {

	// Scanln: put data upto LN in d1, d2 and d3, the rest is left on stdin
	var d1, d2, d3 string
	fmt.Scanln(&d1, &d2, &d3)
	fmt.Println(d1, d2, d3)

	// I can't make sense between these two: Scan and Scanln
	var a1, a2, a3 string
	fmt.Scan(&a1, &a2, &a3)
	fmt.Println(a1, a2, a3)

}

// fmt.Scanln does get the whole line but only on word at a time
func bufioReader() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Simple Shell")
	fmt.Println("---------------------")

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)

		if strings.Compare("hi", text) == 0 {
			fmt.Println("hello, Yourself")
		} else {
			fmt.Println("invalid: ", text)
		}

	}

}

func bufScanner() {
	fmt.Println("input text:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("read line: '%s'\n", scanner.Text())
}
