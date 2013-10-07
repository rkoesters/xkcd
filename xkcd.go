/*
Package xkcd allows access to metadata for xkcd comics.
*/
package xkcd

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	currentUrl  = "http://xkcd.com/info.0.json"
	templateUrl = "http://xkcd.com/%v/info.0.json"
)

type Comic struct {
	Num       int    `json:"num"`
	Title     string `json:"title"`
	SafeTitle string `json:"safe_title"`

	Img string `json:"img"`
	Alt string `json:"alt"`

	Year  string `json:"year"`
	Month string `json:"month"`
	Day   string `json:"day"`

	News       string `json:"news"`
	Link       string `json:"link"`
	Transcript string `json:"transcript"`
}

// Get returns the information about the xkcd comic number `n'.
func Get(n int) (*Comic, error) {
	url := fmt.Sprintf(templateUrl, n)
	return getByUrl(url)
}

// GetCurrent returns information for the newest xkcd comic.
func GetCurrent() (*Comic, error) {
	return getByUrl(currentUrl)
}

// getByUrl returns infomation downloaded from `url'.
func getByUrl(url string) (*Comic, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// TODO: handle status codes.

	dec := json.NewDecoder(resp.Body)

	c := new(Comic)
	err = dec.Decode(c)
	return c, err
}
