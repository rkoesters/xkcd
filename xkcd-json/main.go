package main

import (
	"encoding/json"
	"flag"
	"github.com/rkoesters/xkcd"
	"log"
	"os"
)

var number = flag.Int("n", 0, "Comic number.")

func main() {
	flag.Parse()

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

	e := json.NewEncoder(os.Stdout)
	err = e.Encode(comic)
	if err != nil {
		log.Fatal(err)
	}
}
