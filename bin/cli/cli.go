package main

/*

	The docu for this jawn can be found: https://github.com/urfave/cli/blob/v2.3.0/docs/v2/manual.md

	It's kind of confusing. But go is not the easiest to flow-follow...

*/

import (
	"embed" // love this
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

//go:embed helps/*
var files embed.FS

var read cli.Command = cli.Command{
	Name:        "read",
	Description: "read some stuff",
}

func main() {
	desc, _ := files.ReadFile("helps/description.txt")
	app := &cli.App{
		Action:      actionFunc,
		Name:        "momifier",
		Description: string(desc),
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func actionFunc(c *cli.Context) error {
	readme, err := files.ReadFile("helps/somedata.txt")
	if err != nil {
		fmt.Println(err)
		panic("You messed up, bitch.")
	}
	fmt.Println(string(readme), err)
	return nil
}
