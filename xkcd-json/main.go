package main

import (
	"encoding/json"
	"github.com/rkoesters/xkcd"
	"log"
	"os"
	"strconv"
)

func main() {
	var comic *xkcd.Comic
	var err error

	switch len(os.Args) {
	case 1:
		comic, err = xkcd.GetCurrent()
		if err != nil {
			log.Fatal(err)
		}
	case 2:
		num, err := strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}

		comic, err = xkcd.Get(num)
		if err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatal("args")
	}

	e := json.NewEncoder(os.Stdout)
	err = e.Encode(comic)
	if err != nil {
		log.Fatal(err)
	}
}
