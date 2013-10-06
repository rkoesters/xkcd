xkcd
====

`xkcd` is a go library for accessing xkcd comics.

Example
-------

	package main

	import (
		"fmt"
		"github.com/rkoesters/xkcd"
		"log"
	)

	func main() {
		info, err := xkcd.GetComicInfo(100)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Number: ", info.Num)
		fmt.Println("Image: ", info.Img)
		fmt.Println("Alt Text: ", info.Alt)
	}


Utilities
=========

This comes with a couple example utilities.

xkcd-info
---------

`xkcd-info` prints information about a xkcd comic.

### Usage

	xkcd-info [number]

xkcd-get
--------

`xkcd-get` downloads a xkcd comic and prints it to stdout.

### Usage

	xkcd-get [number] [> filename]
