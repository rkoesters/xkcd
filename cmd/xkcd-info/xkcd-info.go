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
	"net/http"
	"os"
)

var (
	number = flag.Int("n", 0, "Comic number.")
	useNew = flag.Bool("new", false, "Use xkcd.New instead of xkcd.Get.")
)

func main() {
	flag.Parse()

	if flag.NArg() != 0 {
		flag.Usage()
		os.Exit(1)
	}

	var comic *xkcd.Comic
	var err error

	if *number == 0 {
		if !*useNew {
			comic, err = xkcd.GetCurrent()
		} else {
			comic, err = getComicWithNew("http://xkcd.com/info.0.json")
		}
	} else {
		if !*useNew {
			comic, err = xkcd.Get(*number)
		} else {
			url := fmt.Sprintf("http://xkcd.com/%v/info.0.json", *number)
			comic, err = getComicWithNew(url)
		}
	}

	if err != nil {
		log.Fatal(err)
	}

	printInfo(comic)
}

func getComicWithNew(url string) (*xkcd.Comic, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return nil, xkcd.ErrNotFound
	}
	return xkcd.New(resp.Body)
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
