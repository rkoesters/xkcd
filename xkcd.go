package xkcd

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	defaultUrl  = "http://xkcd.com/info.0.json"
	templateUrl = "http://xkcd.com/%v/info.0.json"
)

type Comic struct {
	Num       int    "num"
	Title     string "title"
	SafeTitle string "safe_title"

	Img string "img"
	Alt string "alt"

	Year  string "year"
	Month string "month"
	Day   string "day"

	News       string "news"
	Link       string "link"
	Transcript string "transcript"
}

// Get returns the information about the xkcd comic number `num'.
func Get(num int) (Comic, error) {
	url := fmt.Sprintf(templateUrl, num)
	return GetByUrl(url)
}

// GetCurrent returns information for the newest xkcd comic.
func GetCurrent() (Comic, error) {
	return GetByUrl(defaultUrl)
}

// GetByUrl returns infomation downloaded from `url'. Most people
// will use `Get' and `GetCurrent'.
func GetByUrl(url string) (Comic, error) {
	var c Comic

	resp, err := http.Get(url)
	if err != nil {
		return c, err
	}
	defer resp.Body.Close()

	// TODO: handle status codes.

	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&c)
	return c, err
}
