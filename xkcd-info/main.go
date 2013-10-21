// xkcd-info prints information about an xkcd comic.
//
// Usage:
//	xkcd-info [-n number]
package main

import (
	"flag"
	"fmt"
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

	printInfo(comic)
}

func printInfo(comic *xkcd.Comic) {
	fmt.Printf("Num:        %v\n", comic.Num)
	fmt.Printf("Title:      %v\n", comic.Title)
	fmt.Printf("Safe-Title: %v\n", comic.SafeTitle)
	fmt.Printf("Img:        %v\n", comic.Img)
	fmt.Printf("Alt:        %v\n", comic.Alt)
	fmt.Printf("Year:       %v\n", comic.Year)
	fmt.Printf("Month:      %v\n", comic.Month)
	fmt.Printf("Day:        %v\n", comic.Day)
	fmt.Printf("News:       %v\n", comic.News)
	fmt.Printf("Link:       %v\n", comic.Link)
	fmt.Printf("Transcript: %v\n", comic.Transcript)
}
