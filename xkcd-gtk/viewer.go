package main

import (
	"errors"
	"fmt"
	"github.com/gotk3/gotk3/gtk"
	"github.com/rkoesters/xkcd"
	"log"
	"math/rand"
)

// Viewer is a struct holding a gtk window for viewing XKCD comics.
type Viewer struct {
	win   *gtk.Window
	hdr   *gtk.HeaderBar
	img   *gtk.Image
	comic *xkcd.Comic
}

// New creates a new XKCD viewer window.
func New() (*Viewer, error) {
	v := new(Viewer)

	// Builder the gtk interface using gtk.Builder.
	builder, err := gtk.BuilderNew()
	if err != nil {
		return nil, err
	}
	err = builder.AddFromFile("viewer.ui")
	if err != nil {
		return nil, err
	}

	// Connect the gtk signals to our functions.
	builder.ConnectSignals(map[string]interface{}{
		"PreviousComic": v.PreviousComic,
		"NextComic":     v.NextComic,
		"RandomComic":   v.RandomComic,
	})

	// We want access to Window, HeaderBar, and Image in the future,
	// so lets get access to them now.
	var ok bool
	obj, err := builder.GetObject("ViewerWindow")
	if err != nil {
		return nil, err
	}
	v.win, ok = obj.(*gtk.Window)
	if !ok {
		return nil, errors.New("window")
	}
	obj, err = builder.GetObject("header")
	if err != nil {
		return nil, err
	}
	v.hdr, ok = obj.(*gtk.HeaderBar)
	if !ok {
		return nil, errors.New("headerbar")
	}
	obj, err = builder.GetObject("ComicImage")
	if err != nil {
		return nil, err
	}
	v.img, ok = obj.(*gtk.Image)
	if !ok {
		return nil, errors.New("image")
	}

	// Closing the window should exit the program.
	v.win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	return v, nil
}

// PreviousComic sets the current comic to the previous comic.
func (v *Viewer) PreviousComic() {
	err := v.SetComic(v.comic.Num - 1)
	if err != nil {
		log.Print(err)
	}
}

// NextComic sets the current comic to the next comic.
func (v *Viewer) NextComic() {
	err := v.SetComic(v.comic.Num + 1)
	if err != nil {
		log.Print(err)
	}
}

// RandomComic sets the current comic to a random comic.
func (v *Viewer) RandomComic() {
	c, err := xkcd.GetCurrent()
	if err != nil {
		log.Print(err)
		return
	}
	err = v.SetComic(rand.Intn(c.Num) + 1)
	if err != nil {
		log.Print(err)
	}
}

// SetComic sets the current comic to the given comic.
func (v *Viewer) SetComic(n int) error {
	var err error
	v.comic, err = getComicInfo(n)
	if err != nil {
		return err
	}

	imgPath, err := getComicImage(n)
	if err != nil {
		log.Printf("error downloading comic: %v", n)
	}
	v.hdr.SetSubtitle(fmt.Sprintf("#%v: %v", v.comic.Num, v.comic.Title))
	v.img.SetFromFile(imgPath)
	v.img.SetTooltipText(v.comic.Alt)

	return nil
}
