// Package xkcd allows access to metadata for xkcd comics.
package xkcd

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

// Comic is a struct that contains information about an xkcd comic.
type Comic struct {
	Num        int    `json:"num"`
	Title      string `json:"title"`
	SafeTitle  string `json:"safe_title"`
	Img        string `json:"img"`
	Alt        string `json:"alt"`
	Year       string `json:"year"`
	Month      string `json:"month"`
	Day        string `json:"day"`
	News       string `json:"news"`
	Link       string `json:"link"`
	Transcript string `json:"transcript"`
}

// New reads from an io.Reader and returns a *Comic struct. Assumes text
// is UTF-8. WARNING: this function will improperly decode comics
// retrieved directly from xkcd.com as xkcd.com provides ISO8859-1
// encoded JSON instead of UTF-8 encoded JSON. Use Get or GetCurrent for
// retrieving comics directly from xkcd.com.
func New(r io.Reader) (*Comic, error) {
	d := json.NewDecoder(r)
	c := new(Comic)
	err := d.Decode(c)
	return c, err
}

const (
	currentURL  = "http://xkcd.com/info.0.json"
	templateURL = "http://xkcd.com/%v/info.0.json"
)

// Get fetches information about the xkcd comic number `n'.
func Get(n int) (*Comic, error) {
	url := fmt.Sprintf(templateURL, n)
	return getByURL(url)
}

// GetCurrent fetches information for the newest xkcd comic.
func GetCurrent() (*Comic, error) {
	return getByURL(currentURL)
}

// ErrNotFound is returned when the xkcd server returns a 4xx response
// code (like 404).
var ErrNotFound = errors.New("Error retrieving comic")

// getByURL returns information downloaded from `url'.
func getByURL(url string) (*Comic, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return nil, ErrNotFound
	}
	return New(resp.Body)
}

// Image retrieves the comic image from the xkcd server and returns it
// as an image.Image. Comics in GIF format won't contain the entire
// image, see image/gif.Decode for more information.
func (c *Comic) Image() (*Image, error) {
	return getImageFromURL(c.Img)
}
