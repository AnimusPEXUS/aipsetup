package main

import (
	"log"

	"github.com/gotk3/gotk3/gtk"
)

//go:generate go-bindata -pkg ui -o ./ui/bindata.go ./ui/main.glade

func main() {
	gtk.Init(nil)

	win, err := UIMainWindowNew()
	if err != nil {
		log.Fatal(err)
	}
	win.Show()

	gtk.Main()
}
