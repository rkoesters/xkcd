package main

import (
	"flag"
	"fmt"
	"github.com/rkoesters/xkcd"
	"log"
	"strconv"
)

func main() {
	var info xkcd.ComicInfo
	var err error

	flag.Parse()

	if flag.NArg() == 0 {
		info, err = xkcd.GetCurrentComicInfo()
		if err != nil {
			log.Fatal(err)
		}
	} else if flag.NArg() == 1 {
		num, err := strconv.Atoi(flag.Arg(0))
		if err != nil {
			log.Fatal(err)
		}

		info, err = xkcd.GetComicInfo(num)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		log.Fatal("invalid args")
	}

	printInfo(info)
}

func printInfo(info xkcd.ComicInfo) {
	fmt.Printf("Num: %v\n", info.Num)
	fmt.Printf("Title: %v\n", info.Title)
	fmt.Printf("Safe-Title: %v\n", info.SafeTitle)

	fmt.Printf("Img: %v\n", info.Img)
	fmt.Printf("Alt: %v\n", info.Alt)

	fmt.Printf("Year: %v\n", info.Year)
	fmt.Printf("Month: %v\n", info.Month)
	fmt.Printf("Day: %v\n", info.Day)

	fmt.Printf("News: %v\n", info.News)
	fmt.Printf("Link: %v\n", info.Link)
	fmt.Printf("Transcript: %v\n", info.Transcript)
}
