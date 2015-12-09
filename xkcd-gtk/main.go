package main

import (
	"flag"
	"github.com/gotk3/gotk3/gtk"
	"github.com/rkoesters/xkcd"
	"log"
	"os"
)

var number = flag.Int("n", 0, "Comic number.")

func main() {
	flag.Parse()
	gtk.Init(nil)

	if flag.NArg() != 0 {
		flag.Usage()
		os.Exit(1)
	}

	if *number == 0 {
		c, err := xkcd.GetCurrent()
		if err != nil {
			log.Fatal(err)
		}
		*number = c.Num
	}

	viewer, err := New()
	if err != nil {
		log.Fatal(err)
	}

	viewer.SetComic(*number)
	viewer.win.ShowAll()

	gtk.Main()
}
