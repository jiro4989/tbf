package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func fetchCircleOverviews(tbfID, baseURL string) []CircleOverview {
	var cursor string
	var cos []CircleOverview
	for {
		url := baseURL + "&cursor=" + cursor
		circle := fetchCircle(url)
		for _, c := range circle.List {
			url := fmt.Sprintf("https://techbookfest.org/event/%s/circle/%s", tbfID, c.ID)
			co := CircleOverview{
				Name:            c.Name,
				NameRuby:        c.NameRuby,
				URL:             url,
				GenreFreeFormat: c.GenreFreeFormat,
			}
			cos = append(cos, co)
		}
		if circle.Cursor == "" {
			break
		}
		cursor = circle.Cursor
		time.Sleep(1 * time.Second)
	}
	return cos
}

func fetchCircle(url string) Circle {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	var circle Circle
	if err := json.Unmarshal(b, &circle); err != nil {
		panic(err)
	}

	return circle
}
