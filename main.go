package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// 会場情報
	// https://techbookfest.org/api/event/tbf07
	// サークル情報
	// https://techbookfest.org/api/circle?eventID=tbf07&eventExhibitCourseID=3&visibility=site&limit=100&onlyAdoption=true

	//url := os.Args[1]
	//fmt.Println(url)
	url := "https://techbookfest.org/api/circle?eventID=tbf07&eventExhibitCourseID=3&visibility=site&limit=100&onlyAdoption=true"
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
	fmt.Println(circle)
}
