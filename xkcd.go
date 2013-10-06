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

type ComicInfo struct {
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

func GetComicInfo(num int) (ComicInfo, error) {
	url := fmt.Sprintf(templateUrl, num)
	return GetComicInfoByUrl(url)
}

func GetCurrentComicInfo() (ComicInfo, error) {
	return GetComicInfoByUrl(defaultUrl)
}

func GetComicInfoByUrl(url string) (ComicInfo, error) {
	var info ComicInfo

	resp, err := http.Get(url)
	if err != nil {
		return info, err
	}
	defer resp.Body.Close()

	// TODO: handle status codes.

	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&info)
	return info, err
}
