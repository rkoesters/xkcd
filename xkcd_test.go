package xkcd_test

import (
	"github.com/rkoesters/xkcd"
	"testing"
)

func TestGet(t *testing.T) {
	comic, err := xkcd.Get(221)
	if err != nil {
		t.Fatal(err)
	}

	expected := &xkcd.Comic{
		Num:       221,
		Title:     "Random Number",
		SafeTitle: "Random Number",
		Img:       "http://imgs.xkcd.com/comics/random_number.png",
		Alt:       "RFC 1149.5 specifies 4 as the standard IEEE-vetted random number.",
		Year:      "2007",
		Month:     "2",
		Day:       "9",
		News:      "",
		Link:      "",
		Transcript: `int getRandomNumber()
{
  return 4; 
 chosen by fair dice roll.
               
 guarenteed to be random.
}
{{title text: RFC 1149.5 specifies 4 as the standard IEEE-vetted random number.}}`,
	}

	if *comic != *expected {
		t.Fail()
	}
}

func TestGetCurrent(t *testing.T) {
	comic1, err := xkcd.GetCurrent()
	if err != nil {
		t.Fatal(err)
	}

	comic2, err := xkcd.Get(comic1.Num)
	if err != nil {
		t.Fatal(err)
	}

	if *comic1 != *comic2 {
		t.Fail()
	}
}

func Test404(t *testing.T) {
	_, err := xkcd.Get(404)
	if err.Error() != "404 Not Found" {
		t.Fatal(err)
	}
}
