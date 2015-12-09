package main

import (
	"encoding/json"
	"github.com/rkoesters/xdg/basedir"
	"github.com/rkoesters/xkcd"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

func getComicPath(n int) string {
	return filepath.Join(basedir.CacheHome, "xkcd-gui", strconv.Itoa(n))
}

func getComicInfoPath(n int) string {
	return filepath.Join(getComicPath(n), "info")
}

func getComicImagePath(n int) string {
	return filepath.Join(getComicPath(n), "image")
}

func getComicInfo(n int) (*xkcd.Comic, error) {
	infoPath := getComicInfoPath(n)

	// First, check if we have the file.
	_, err := os.Stat(infoPath)
	if os.IsNotExist(err) {
		err = downloadComicInfo(n)
		if err != nil {
			return nil, err
		}
	} else if err != nil {
		return nil, err
	}

	f, err := os.Open(infoPath)
	if err != nil {
		return nil, err
	}
	c, err := xkcd.New(f)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func downloadComicInfo(n int) error {
	err := os.MkdirAll(getComicPath(n), 0777)
	if err != nil {
		return err
	}

	comic, err := xkcd.Get(n)
	if err != nil {
		return err
	}

	f, err := os.Create(getComicInfoPath(n))
	if err != nil {
		return err
	}

	e := json.NewEncoder(f)
	err = e.Encode(comic)
	if err != nil {
		return err
	}
	return nil
}

func getComicImage(n int) (string, error) {
	imagePath := getComicImagePath(n)

	_, err := os.Stat(imagePath)
	if os.IsNotExist(err) {
		err = downloadComicImage(n)
		if err != nil {
			return "", err
		}
	} else if err != nil {
		return "", err
	}
	return imagePath, nil
}

func downloadComicImage(n int) error {
	c, err := getComicInfo(n)
	if err != nil {
		return err
	}

	f, err := os.Create(getComicImagePath(n))
	if err != nil {
		return err
	}

	resp, err := http.Get(c.Img)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(f, resp.Body)
	if err != nil {
		return err
	}
	return nil
}