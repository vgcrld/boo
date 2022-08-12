package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
)

func main() {
	baseURL := "https://www.galileosuite.com/"
	if len(os.Args) == 2 {
		baseURL = os.Args[1]
	}
	fmt.Println(baseURL)
	openBrowser(baseURL)
}
func openBrowser(targetURL string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", targetURL).Start()
		// TODO: "Windows Subsytem for Linux" is also recognized as "linux", but then we need
		// err = exec.Command("rundll32.exe", "url.dll,FileProtocolHandler", targetURL).Start()
	case "windows":
		err = exec.Command("rundll32.exe", "url.dll,FileProtocolHandler", targetURL).Start()
	case "darwin":
		err = exec.Command("open", targetURL).Start()
	default:
		err = fmt.Errorf("unsupported platform %v", runtime.GOOS)
	}
	if err != nil {
		log.Fatal(err)
	}
}
