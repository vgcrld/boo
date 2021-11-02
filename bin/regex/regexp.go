package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Simple Shell")
	fmt.Println("-------------------------")
	for {
		fmt.Print("text: -> ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1) // convert CRLF to LF
		fmt.Print("regexp: -> ")
		regTxt, _ := reader.ReadString('\n')
		regTxt = strings.Replace(regTxt, "\n", "", -1) // convert CRLF to LF
		re := regexp.MustCompile(regTxt)
		matched := didOrDidNot(re.Match([]byte(text)))
		fmt.Printf("'%v' %v '%v'\n", regTxt, matched, text)
	}
}

func didOrDidNot(s bool) string {
	if s {
		return "matched"
	}
	return "did not match"
}
