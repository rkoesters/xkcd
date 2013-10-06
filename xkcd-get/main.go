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
	var info xkcd.ComicInfo
	var err error

	switch len(os.Args) {
	case 1:
		info, err = xkcd.GetCurrentComicInfo()
		if err != nil {
			log.Fatal(err)
		}
	case 2:
		num, err := strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}

		info, err = xkcd.GetComicInfo(num)
		if err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatal("args")
	}

	resp, err := http.Get(info.Img)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	_, err = io.Copy(os.Stdout, resp.Body)
	if err != nil {
		log.Fatal(err)
	}
}
