package main

import (
	"bytes"
	"embed"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strconv"
)

//go:embed frontend/dist
var frontend embed.FS

func main() {
	var port int
	flag.IntVar(&port, "port", 8080, "The port to listen on")
	flag.Parse()

	http.Handle("/api/v1/law", http.HandlerFunc(getRandomLaw))

	stripped, err := fs.Sub(frontend, "frontend/dist")
	if err != nil {
		log.Fatalln(err)
	}

	frontendFS := http.FileServer(http.FS(stripped))
	http.Handle("/", frontendFS)

	strPort := strconv.FormatInt(int64(port), 10)
	baseURL := "http://localhost:" + strPort
	if len(os.Args) == 2 {
		baseURL = os.Args[1]
	}
	openBrowser(baseURL)

	log.Fatalln(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

type Law struct {
	Name       string `json:"name,omitempty"`
	Definition string `json:"definition,omitempty"`
}

var HackerLaws = []Law{
	{
		Name:       "Amdahl's Law",
		Definition: "Amdahl's Law is a formula which shows the potential speedup of a computational task which can be achieved by increasing the resources of a system.",
	},
	{
		Name:       "Conway's Law",
		Definition: "This law suggests that the technical boundaries of a system will reflect the structure of the organisation.",
	},
	{
		Name:       "Gall's Law",
		Definition: "A complex system that works is invariably found to have evolved from a simple system that worked.",
	},
}

func getRandomLaw(w http.ResponseWriter, r *http.Request) {
	randomLaw := HackerLaws[rand.Intn(len(HackerLaws))]
	j, err := json.Marshal(randomLaw)
	if err != nil {
		http.Error(w, "couldn't retrieve random hacker law", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	io.Copy(w, bytes.NewReader(j))
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
