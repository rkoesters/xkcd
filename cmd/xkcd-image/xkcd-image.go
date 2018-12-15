// xkcd-image downloads an xkcd comic image to the current directory.
//
// Usage:
//	xkcd-image [-n number] [-d output-directory]
package main

import (
	"flag"
	"github.com/rkoesters/xkcd"
	"image/gif"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"path/filepath"
)

var (
	number    = flag.Int("n", 0, "comic number")
	outputDir = flag.String("d", ".", "output directory")
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
		comic, err = xkcd.GetCurrent()
	} else {
		comic, err = xkcd.Get(*number)
	}

	if err != nil {
		log.Fatal(err)
	}

	img, err := comic.Image()
	if err != nil {
		log.Fatal(err)
	}

	outfile, err := os.Create(filepath.Join(*outputDir, img.SourceName()))
	if err != nil {
		log.Fatal(err)
	}
	defer outfile.Close()

	switch img.SourceFormat() {
	case "gif":
		err = gif.Encode(outfile, img, nil)
	case "jpeg":
		err = jpeg.Encode(outfile, img, nil)
	case "png":
		err = png.Encode(outfile, img)
	default:
		log.Fatalf("unknown image format: %v", img.SourceFormat())
	}
	if err != nil {
		log.Fatal(err)
	}
}
