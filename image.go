package xkcd

import (
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"net/http"
)

// Image represents an xkcd comic image.
type Image struct {
	image.Image

	sourceURL    string
	sourceFormat string
}

// NewImage takes an io.Reader and returns an Image struct.
func NewImage(r io.Reader) (*Image, error) {
	var img Image
	var err error

	img.Image, img.sourceFormat, err = image.Decode(r)

	return &img, err
}

// GetImage is a convenience function to download a comic's metadata and
// then download the comic's image. If you already have the comic
// metadata, use Comic's Image() method.
func GetImage(n int) (*Image, error) {
	c, err := Get(n)
	if err != nil {
		return nil, err
	}
	return c.Image()
}

func getImageFromURL(url string) (*Image, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return NewImage(resp.Body)
}

// SourceURL is URL from which the comic image was retrieved. If image
// wasn't retrieved from a URL (e.g. when it is created by calling
// NewImage directly) a blank string will be returned.
func (img *Image) SourceURL() string {
	return img.sourceURL
}

// SourceFormat is the file format that the image was decoded from.
func (img *Image) SourceFormat() string {
	return img.sourceFormat
}
