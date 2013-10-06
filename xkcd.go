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
	Num       int
	Title     string
	SafeTitle string `json:"safe_title"`

	Img string
	Alt string

	Year  string
	Month string
	Day   string

	News       string
	Link       string
	Transcript string
}

// Get returns the information about the xkcd comic number `n'.
func Get(n int) (*Comic, error) {
	url := fmt.Sprintf(templateUrl, n)
	return GetByUrl(url)
}

// GetCurrent returns information for the newest xkcd comic.
func GetCurrent() (*Comic, error) {
	return GetByUrl(currentUrl)
}

// GetByUrl returns infomation downloaded from `url'. Most people
// will use `Get' and `GetCurrent'.
func GetByUrl(url string) (*Comic, error) {
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
