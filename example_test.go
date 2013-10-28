package xkcd_test

import (
	"fmt"
	"github.com/rkoesters/xkcd"
	"log"
	"strings"
)

func ExampleNew() {
	r := strings.NewReader(`{"month": "3", "num": 1190, "link": "http:\/\/geekwagon.net\/projects\/xkcd1190\/", "year": "2013", "news": "", "safe_title": "Time", "transcript": "", "alt": "The end.", "img": "http:\/\/imgs.xkcd.com\/comics\/time.png", "title": "Time", "day": "25"}`)

	comic, err := xkcd.New(r)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Number: %v\n", comic.Num)
	fmt.Printf("Image: %v\n", comic.Img)
	fmt.Printf("Alt Text: %v\n", comic.Alt)
	// Output:
	// Number: 1190
	// Image: http://imgs.xkcd.com/comics/time.png
	// Alt Text: The end.
}

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
