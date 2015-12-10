package main

import (
	"errors"
	"fmt"
	"github.com/gotk3/gotk3/gtk"
	"github.com/rkoesters/xkcd"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// Viewer is a struct holding a gtk window for viewing XKCD comics.
type Viewer struct {
	comic *xkcd.Comic
	win   *gtk.Window
	hdr   *gtk.HeaderBar
	img   *gtk.Image
	// The labels in the details popover.
	detailsNumber *gtk.Label
	detailsTitle  *gtk.Label
	detailsDate   *gtk.Label
	detailsLink   *gtk.Label
}

// New creates a new XKCD viewer window.
func New() (*Viewer, error) {
	v := new(Viewer)

	// Builder the gtk interface using gtk.Builder.
	builder, err := gtk.BuilderNew()
	if err != nil {
		return nil, err
	}
	err = builder.AddFromFile("xkcd-gtk.ui")
	if err != nil {
		return nil, err
	}

	// Connect the gtk signals to our functions.
	builder.ConnectSignals(map[string]interface{}{
		"PreviousComic": v.PreviousComic,
		"NextComic":     v.NextComic,
		"RandomComic":   v.RandomComic})

	// We want access to Window, HeaderBar, and Image in the future,
	// so lets get access to them now.
	var ok bool
	obj, err := builder.GetObject("viewer-window")
	if err != nil {
		return nil, err
	}
	v.win, ok = obj.(*gtk.Window)
	if !ok {
		return nil, errors.New("error getting viewer-window")
	}
	obj, err = builder.GetObject("header")
	if err != nil {
		return nil, err
	}
	v.hdr, ok = obj.(*gtk.HeaderBar)
	if !ok {
		return nil, errors.New("error getting header")
	}
	obj, err = builder.GetObject("comic-image")
	if err != nil {
		return nil, err
	}
	v.img, ok = obj.(*gtk.Image)
	if !ok {
		return nil, errors.New("error getting comic-image")
	}

	// Get details labels.
	v.detailsNumber, err = getLabel(builder, "details-number")
	if err != nil {
		return nil, err
	}
	v.detailsTitle, err = getLabel(builder, "details-title")
	if err != nil {
		return nil, err
	}
	v.detailsDate, err = getLabel(builder, "details-date")
	if err != nil {
		return nil, err
	}
	v.detailsLink, err = getLabel(builder, "details-link")
	if err != nil {
		return nil, err
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
	c, err := getNewestComicInfo()
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
	c, err := getComicInfo(n)
	if err != nil {
		return err
	}
	v.comic = c

	imgPath, err := getComicImage(n)
	if err != nil {
		log.Printf("error downloading comic: %v", n)
	}
	v.hdr.SetSubtitle(fmt.Sprintf("#%v: %v", v.comic.Num, v.comic.Title))
	v.img.SetFromFile(imgPath)
	v.img.SetTooltipText(v.comic.Alt)

	// Update the details popover.
	v.detailsNumber.SetText(strconv.Itoa(v.comic.Num))
	v.detailsTitle.SetText(v.comic.Title)
	v.detailsDate.SetText(formatDate(v.comic.Year, v.comic.Month, v.comic.Day))
	if v.comic.Link != "" {
		fmtLink := fmt.Sprintf("<a href=\"%v\">%[1]v</a>", v.comic.Link)
		v.detailsLink.SetMarkup(fmtLink)
	} else {
		v.detailsLink.SetText("")
	}

	return nil
}

// formatDate takes a year, month, and date as strings and turns them
// into a pretty date.
func formatDate(year, month, day string) string {
	date := strings.Join([]string{year, month, day}, "-")
	t, err := time.Parse("2006-1-2", date)
	if err != nil {
		return ""
	}
	return t.Format("Jan _2, 2006")
}

func getLabel(b *gtk.Builder, id string) (*gtk.Label, error) {
	obj, err := b.GetObject(id)
	if err != nil {
		return nil, err
	}
	label, ok := obj.(*gtk.Label)
	if !ok {
		return nil, fmt.Errorf("error getting label: %v", id)
	}
	return label, nil
}
