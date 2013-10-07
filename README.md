xkcd
====

`xkcd` is a go library for accessing xkcd comics.
[Click here for documentation.](http://godoc.org/github.com/rkoesters/xkcd)

Example
-------

	package main

	import (
		"fmt"
		"github.com/rkoesters/xkcd"
		"log"
	)

	func main() {
		comic, err := xkcd.Get(140)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Number: ", comic.Num)
		fmt.Println("Image: ", comic.Img)
		fmt.Println("Alt Text: ", comic.Alt)
	}


Utilities
=========

This comes with a couple example utilities.

xkcd-info
---------

`xkcd-info` prints information about a xkcd comic.

### Usage

	xkcd-info [-n number]

xkcd-get
--------

`xkcd-get` downloads a xkcd comic and prints it to stdout.

### Usage

	xkcd-get [-n number] [> filename]

xkcd-json
---------

`xkcd-json` prints information about xkcd comic in json format.

### Usage

	xkcd-json [-n number]
