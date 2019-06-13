package xkcd_test

import (
	"encoding/json"
	"github.com/rkoesters/xkcd"
	"io"
	"reflect"
	"testing"
	"unicode/utf8"
)

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

	if !comicValidUtf8(comic1) {
		t.Errorf("%+q isn't valid utf-8", comic1)
	}

	if !comicValidUtf8(comic2) {
		t.Errorf("%+q isn't valid utf-8", comic2)
	}

	if !reflect.DeepEqual(comic1, comic2) {
		t.Fatal("comic1 and comic2 don't match")
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

	if !comicValidUtf8(comic1) {
		t.Errorf("%+q isn't valid utf-8", comic1)
	}

	if !comicValidUtf8(comic2) {
		t.Errorf("%+q isn't valid utf-8", comic2)
	}

	if !reflect.DeepEqual(comic1, comic2) {
		t.Fatal("comic1 and comic2 don't match")
	}
}

func TestGet4(t *testing.T) {
	testGet(t, &xkcd.Comic{
		Num:       221,
		Title:     "Random Number",
		SafeTitle: "Random Number",
		Img:       "https://imgs.xkcd.com/comics/random_number.png",
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
	})
}

func TestGet404(t *testing.T) {
	_, err := xkcd.Get(404)
	if err != xkcd.ErrNotFound {
		t.Fatal(err)
	}
}

func TestGet1953(t *testing.T) {
	testGet(t, &xkcd.Comic{
		Num:        1953,
		Title:      "The History of Unicode",
		SafeTitle:  "The History of Unicode",
		Img:        "https://imgs.xkcd.com/comics/the_history_of_unicode.png",
		Alt:        "2048: \"Great news for Maine—we're once again an independent state!!! Thanks, @unicode, for ruling in our favor and sending troops to end New Hampshire's annexation. 🙏🚁🎖️\"",
		Year:       "2018",
		Month:      "2",
		Day:        "9",
		News:       "",
		Link:       "",
		Transcript: "",
	})
}

func TestGet1956(t *testing.T) {
	testGet(t, &xkcd.Comic{
		Num:        1956,
		Title:      "Unification",
		SafeTitle:  "Unification",
		Img:        "https://imgs.xkcd.com/comics/unification.png",
		Alt:        "For a while, some physicists worked on a theory unifying the other forces with both the force of gravity and the film \"Gravity,\" but even after Alfonso Cuarón was held in a deep underground chamber of water for 10^31 years he refused to sell his film to Disney.",
		Year:       "2018",
		Month:      "2",
		Day:        "16",
		News:       "",
		Link:       "",
		Transcript: "",
	})
}

func TestGet2038(t *testing.T) {
	testGet(t, &xkcd.Comic{
		Num:        2038,
		Title:      "Hazard Symbol",
		SafeTitle:  "Hazard Symbol",
		Img:        "https://imgs.xkcd.com/comics/hazard_symbol.png",
		Alt:        "The warning diamond on the Materials Safety Data Sheet for this stuff just has the \"😰\" emoji in all four fields.",
		Year:       "2018",
		Month:      "8",
		Day:        "27",
		News:       "",
		Link:       "",
		Transcript: "",
	})
}

func testGet(t *testing.T, expect *xkcd.Comic) {
	actual, err := xkcd.Get(expect.Num)
	if err != nil {
		t.Error(err)
	}

	if actual == nil {
		t.Fatal("actual == nil")
	}

	if !comicValidUtf8(actual) {
		t.Errorf("%+q isn't valid utf-8", actual)
	}

	if !comicValidUtf8(expect) {
		t.Errorf("%+q isn't valid utf-8", actual)
	}

	if !reflect.DeepEqual(actual, expect) {
		t.Logf("actual=%v", actual)
		t.Logf("expect=%v", expect)
		t.Fail()
	}
}

func comicValidUtf8(c *xkcd.Comic) bool {
	return utf8.ValidString(c.Title) &&
		utf8.ValidString(c.SafeTitle) &&
		utf8.ValidString(c.Img) &&
		utf8.ValidString(c.Alt) &&
		utf8.ValidString(c.Year) &&
		utf8.ValidString(c.Month) &&
		utf8.ValidString(c.Day) &&
		utf8.ValidString(c.News) &&
		utf8.ValidString(c.Link) &&
		utf8.ValidString(c.Transcript)
}
