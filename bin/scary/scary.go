package main

import (
	"fmt"

	"github.com/vgcrld/boo/commons"
	"github.com/vgcrld/boo/commons/ui"
)

func main() {
	fmt.Printf("Version=%v UI_Version=%v\n", commons.VERSION, ui.VERSION)

}
