// xkcd-xml prints information about an xkcd comic in xml format.
//
// Usage:
//	xkcd-xml [-n number]
package main

import (
	"encoding/xml"
	"flag"
	"github.com/rkoesters/xkcd"
	"log"
	"os"
)

var number = flag.Int("n", 0, "Comic number.")

func main() {
	flag.Parse()

	if flag.NArg() != 0 {
		flag.Usage()
		os.Exit(1)
	}

	var comic *xkcd.Comic
	var err error

	if *number == 0 {
		comic, err = xkcd.GetCurrent()
	} else {
		comic, err = xkcd.Get(*number)
	}

	if err != nil {
		log.Fatal(err)
	}

	e := xml.NewEncoder(os.Stdout)
	err = e.Encode(comic)
	if err != nil {
		log.Fatal(err)
	}
}
