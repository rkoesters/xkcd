package xkcd

import (
	"golang.org/x/text/encoding/charmap"
)

// XKCD's JSON is ISO8859-1 and we want it in UTF-8.
func fixComic(c *Comic) {
	c.Title = fixEncoding(c.Title)
	c.SafeTitle = fixEncoding(c.SafeTitle)
	c.Img = fixEncoding(c.Img)
	c.Alt = fixEncoding(c.Alt)
	c.Year = fixEncoding(c.Year)
	c.Month = fixEncoding(c.Month)
	c.Day = fixEncoding(c.Day)
	c.News = fixEncoding(c.News)
	c.Link = fixEncoding(c.Link)
	c.Transcript = fixEncoding(c.Transcript)
}

// Convert ISO8859-1 text erroneously decoded as UTF-8 into actual
// UTF-8.
func fixEncoding(s string) string {
	d := charmap.ISO8859_1.NewEncoder()
	ret, err := d.String(s)
	if err != nil {
		// input wasn't ISO8859-1, just return it as is.
		return s
	}
	return ret
}
