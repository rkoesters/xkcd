package main

import (
	"fmt"
	"github.com/gotk3/gotk3/gtk"
	"github.com/rkoesters/xkcd"
	"log"
	"math/rand"
	"time"
)

type XKCDViewer struct {
	win   *gtk.Window
	hdr   *gtk.HeaderBar
	img   *gtk.Image
	comic *xkcd.Comic
}

func New() (*XKCDViewer, error) {
	var err error
	v := new(XKCDViewer)
	v.win, err = gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		return nil, err
	}
	v.win.Connect("destroy", func() {
		gtk.MainQuit()
	})
	v.win.SetDefaultSize(800, 600)

	// Create HeaderBar
	err = v.setupHeaderBar()
	if err != nil {
		return nil, err
	}

	// Create image widget.
	content, err := v.setupImage()
	if err != nil {
		return nil, err
	}

	v.win.SetTitlebar(v.hdr)
	v.win.Add(content)

	return v, nil
}

func (v *XKCDViewer) setupHeaderBar() error {
	var err error

	v.hdr, err = gtk.HeaderBarNew()
	if err != nil {
		return err
	}
	v.hdr.SetShowCloseButton(true)
	v.hdr.SetTitle("XKCD")

	box, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
	if err != nil {
		return err
	}

	style, err := box.GetStyleContext()
	if err != nil {
		return err
	}
	style.AddClass("linked")

	btnPrev, err := gtk.ButtonNewFromIconName("go-previous-symbolic", gtk.ICON_SIZE_MENU)
	if err != nil {
		return err
	}
	btnPrev.Connect("clicked", func() {
		err := v.SetComic(v.comic.Num - 1)
		if err != nil {
			log.Print(err)
		}
	})

	btnNext, err := gtk.ButtonNewFromIconName("go-next-symbolic", gtk.ICON_SIZE_MENU)
	if err != nil {
		return err
	}
	btnNext.Connect("clicked", func() {
		err := v.SetComic(v.comic.Num + 1)
		if err != nil {
			log.Print(err)
		}
	})

	box.Add(btnPrev)
	box.Add(btnNext)
	v.hdr.PackStart(box)

	btnRand, err := gtk.ButtonNewFromIconName("media-playlist-shuffle-symbolic", gtk.ICON_SIZE_MENU)
	if err != nil {
		return err
	}
	rand.Seed(time.Now().Unix())
	btnRand.Connect("clicked", func() {
		c, err := xkcd.GetCurrent()
		if err != nil {
			log.Print(err)
		}
		err = v.SetComic(rand.Intn(c.Num) + 1)
		if err != nil {
			log.Print(err)
		}
	})

	btnMenu, err := gtk.ButtonNewFromIconName("open-menu-symbolic", gtk.ICON_SIZE_MENU)
	if err != nil {
		return err
	}

	btnAbout, err := gtk.ButtonNewWithLabel("About")
	if err != nil {
		return nil
	}
	btnAbout.Connect("clicked", func() {
		abdi, err := gtk.AboutDialogNew()
		if err != nil {
			log.Print(err)
			return
		}
		abdi.SetProgramName("XKCD GUI")
		abdi.SetLicense("TODO")
		abdi.SetVersion("v0.1")
		abdi.SetWebsite("http://github.com/rkoesters/xkcd")
		abdi.ShowAll()
	})

	v.hdr.PackEnd(btnMenu)
	v.hdr.PackEnd(btnRand)

	return nil
}

func (v *XKCDViewer) setupImage() (*gtk.ScrolledWindow, error) {
	swin, err := gtk.ScrolledWindowNew(nil, nil)
	if err != nil {
		return nil, err
	}

	v.img, err = gtk.ImageNew()
	if err != nil {
		return nil, err
	}

	swin.Add(v.img)

	return swin, nil
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
