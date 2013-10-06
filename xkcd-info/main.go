package main

import (
	"flag"
	"fmt"
	"github.com/rkoesters/xkcd"
	"log"
	"strconv"
)

func main() {
	var comic xkcd.Comic
	var err error

	flag.Parse()

	if flag.NArg() == 0 {
		comic, err = xkcd.GetCurrentComic()
		if err != nil {
			log.Fatal(err)
		}
	} else if flag.NArg() == 1 {
		num, err := strconv.Atoi(flag.Arg(0))
		if err != nil {
			log.Fatal(err)
		}

		comic, err = xkcd.GetComic(num)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		log.Fatal("invalid args")
	}

	printInfo(comic)
}

func printInfo(comic xkcd.Comic) {
	fmt.Printf("Num: %v\n", comic.Num)
	fmt.Printf("Title: %v\n", comic.Title)
	fmt.Printf("Safe-Title: %v\n", comic.SafeTitle)

	fmt.Printf("Img: %v\n", comic.Img)
	fmt.Printf("Alt: %v\n", comic.Alt)

	fmt.Printf("Year: %v\n", comic.Year)
	fmt.Printf("Month: %v\n", comic.Month)
	fmt.Printf("Day: %v\n", comic.Day)

	fmt.Printf("News: %v\n", comic.News)
	fmt.Printf("Link: %v\n", comic.Link)
	fmt.Printf("Transcript: %v\n", comic.Transcript)
}
