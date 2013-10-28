package xkcd_test

import (
	"encoding/json"
	"github.com/rkoesters/xkcd"
	"io"
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

	t.Log("comic: ", comic)
	t.Log("expected: ", expected)

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

	t.Log("comic1: ", comic1)
	t.Log("comic2: ", comic2)

	if *comic1 != *comic2 {
		t.Fatal("comic1 and comic2 don't match")
	}
}

func Test404(t *testing.T) {
	_, err := xkcd.Get(404)
	if err.Error() != "404 Not Found" {
		t.Fatal(err)
	}
}

func TestNew(t *testing.T) {
	comic1, err := xkcd.GetCurrent()
	if err != nil {
		t.Fatal(err)
	}

	r, w := io.Pipe()

	go func() {
		e := json.NewEncoder(w)
		err = e.Encode(comic1)
		if err != nil {
			t.Fatal(err)
		}
	}()

	comic2, err := xkcd.New(r)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("comic1: ", comic1)
	t.Log("comic2: ", comic2)

	if *comic1 != *comic2 {
		t.Fatal("comic1 and comic2 don't match")
	}
}
