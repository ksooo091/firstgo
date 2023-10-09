package main

import (
	"fmt"
	"net/http"
)

type urlGetResult struct {
	url    string
	status string
}

func main() {
	//var results = map[string]string{}
	// 위와 같음
	var results = make(map[string]string)
	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.google.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://academy.nomadcoders.co/",
		"https://app.sesac.site/",
	}
	c := make(chan urlGetResult)
	for _, url := range urls {

		go hitURL(url, c)

	}
	for range urls {
		result := <-c
		results[result.url] = result.status

	}
	for url, status := range results {
		fmt.Println(url, status)
	}

}

func hitURL(url string, c chan<- urlGetResult) {

	resp, err := http.Get(url)

	if err != nil || resp.StatusCode >= 400 {
		fmt.Println(err)
		c <- urlGetResult{url: url, status: "Fail"}
	} else {
		c <- urlGetResult{url: url, status: resp.Status}
	}
}
