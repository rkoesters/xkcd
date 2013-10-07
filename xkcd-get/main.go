package main

import (
	"flag"
	"github.com/rkoesters/xkcd"
	"io"
	"log"
	"net/http"
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

	resp, err := http.Get(comic.Img)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	_, err = io.Copy(os.Stdout, resp.Body)
	if err != nil {
		log.Fatal(err)
	}
}
