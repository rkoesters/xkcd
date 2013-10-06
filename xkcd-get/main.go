package main

import (
	"github.com/rkoesters/xkcd"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	var comic xkcd.Comic
	var err error

	switch len(os.Args) {
	case 1:
		comic, err = xkcd.GetCurrentComic()
		if err != nil {
			log.Fatal(err)
		}
	case 2:
		num, err := strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}

		comic, err = xkcd.GetComic(num)
		if err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatal("args")
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
