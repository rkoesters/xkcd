package main

import (
	"fmt"
	"github.com/gotk3/gotk3/gtk"
	"github.com/rkoesters/xkcd"
	"log"
	"errors"
)

type XKCDViewer struct {
	win   *gtk.Window
	hdr   *gtk.HeaderBar
	img   *gtk.Image
	comic *xkcd.Comic
}

func New() (*XKCDViewer, error) {
	v := new(XKCDViewer)

	builder, err := gtk.BuilderNew()
	if err != nil {
		return nil, err
	}
	err = builder.AddFromFile("window.ui")
	if err != nil {
		return nil, err
	}

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

	v.win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	return v, nil
}

func (v *XKCDViewer) SetComic(n int) error {
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
