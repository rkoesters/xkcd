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
	if len(os.Args) != 2 {
		log.Fatal("args")
	}

	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	info, err := xkcd.GetComicInfo(num)
	if err != nil {
		log.Fatal(err)
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
