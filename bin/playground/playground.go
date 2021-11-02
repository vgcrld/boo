package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/vgcrld/boo/commons"
	"github.com/vgcrld/boo/commons/ui"
)

type Apple struct {
	Type string
}

func (a *Apple) peel() string {
	return fmt.Sprintf("I peeled my %v apple", a.Type)
}

type Orange struct {
	Type string
}

func (o *Orange) peel() string {
	return fmt.Sprintf("I peeled my %v orange", o.Type)
}

type Fruit interface {
	peel() string
}

func main() {

	// whe we are done we are going to say done.
	defer fmt.Println("done!")

	// Let's trap SIGINT and SIGTERM, make the user enter 'quit'
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		for {
			<-sigs
			fmt.Println("Enter 'quit' to end.")
		}
	}()

	// Some basic playground stuff
	showVersions()
	showFruits()

	fmt.Println("Maps Code")
	fmap := fruitMaps()
	fmt.Println(fmap)
	deleteFromMap(100, fmap)
	printTheMap(fmap)

	showBunches(
		(&Apple{Type: "McIntosh"}).peel(),
		10.6,
		time.Now(),
		time.Millisecond,
		"test"[2:],
		(&Orange{Type: "bobba"}).peel(),
		(&Orange{Type: "bobba"}).peel(),
	)

	// Loop this
	go looper("Main (Enter 'quit' to end.)", 5000)
	for {
		val := readConsole()
		if val == "quit" {
			break
		}
	}

}

func printTheMap(fmap *map[int]interface{}) {
	for i, v := range *fmap {
		fmt.Println(i, v)
	}
}

func deleteFromMap(i int, m *map[int]interface{}) {
	fmt.Println("delete key from map:", i)
	delete(*m, i)
}

func fruitMaps() *map[int]interface{} {
	fruits := make(map[int]interface{})
	fruits[200] = Orange{Type: "Florida"}
	fruits[201] = Orange{Type: "California"}
	fruits[100] = Apple{Type: "Red Delicious"}
	fruits[101] = Apple{Type: "McIntosh"}
	return &fruits
}

func showVersions() {
	var version int = commons.VERSION
	var versionui int = ui.VERSION
	fmt.Println("version:", version, "version-ui", versionui)
}

func showFruits() {
	var fruits []Fruit
	fruits = append(fruits, &Orange{Type: "Florida"})
	fruits = append(fruits, &Apple{Type: "Red Delicious"})
	fruits = append(fruits, &Orange{Type: "California"})
	for i, v := range fruits {
		fmt.Printf("%v %T %s\n", i, v, v.peel())
	}
}

func looper(name string, i time.Duration) {
	for {
		fmt.Println("Loooper", name, time.Now())
		time.Sleep(i * time.Millisecond)
	}
}

func readConsole() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1) // convert CRLF to LF
	return text
}

func showBunches(stuffs ...interface{}) {
	for i, v := range stuffs {
		fmt.Println(i, v)
	}
}
