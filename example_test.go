package xkcd_test

import (
	"fmt"
	"github.com/rkoesters/xkcd"
	"log"
)

func ExampleGet() {
	comic, err := xkcd.Get(140)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Number: %v\n", comic.Num)
	fmt.Printf("Image: %v\n", comic.Img)
	fmt.Printf("Alt Text: %v\n", comic.Alt)

	// Output:
	// Number: 140
	// Image: http://imgs.xkcd.com/comics/delicious.png
	// Alt Text: I'm currently in the I Have Cheese phase of this cycle.
}
