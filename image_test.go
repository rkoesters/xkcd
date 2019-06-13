package xkcd_test

import (
	"github.com/rkoesters/xkcd"
	"testing"
)

func TestGetImage1(t *testing.T) {
	testGetImage(t, 1,
		"https://imgs.xkcd.com/comics/barrel_cropped_(1).jpg",
		"barrel_cropped_(1).jpg",
		"jpeg",
	)
}

func TestGetImage1000(t *testing.T) {
	testGetImage(t, 1000,
		"https://imgs.xkcd.com/comics/1000_comics.png",
		"1000_comics.png",
		"png",
	)
}

func TestGetImage1116(t *testing.T) {
	testGetImage(t, 1116,
		"https://imgs.xkcd.com/comics/traffic_lights.gif",
		"traffic_lights.gif",
		"gif",
	)
}

func testGetImage(t *testing.T, num int, url, name, format string) {
	img, err := xkcd.GetImage(num)
	if err != nil {
		t.Error(err)
	}

	if img == nil {
		t.Fatal("img == nil")
	}

	if url != img.SourceURL() {
		t.Errorf("url: expected=%v actual=%v", url, img.SourceURL())
	}

	if name != img.SourceName() {
		t.Errorf("name: expected=%v actual=%v", name, img.SourceName())
	}

	if format != img.SourceFormat() {
		t.Errorf("format: expected=%v actual=%v", format, img.SourceFormat())
	}
}
