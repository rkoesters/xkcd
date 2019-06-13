package xkcd_test

import (
	"github.com/rkoesters/xkcd"
	"testing"
)

func TestGetImage(t *testing.T) {
	img, err := xkcd.GetImage(1)
	if err != nil {
		t.Error(err)
	}

	if img == nil {
		t.Error("img == nil")
	}
}
