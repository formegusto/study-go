package main

import (
	"fmt"

	"github.com/formegusto/study-go/urlutils"
)

func main() {
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
	}
	for _,url := range urls {
		err := urlutils.HitURL(url)
		if err != nil {
			results[url] = "failed"
		} else {
			results[url] = "success"
		}
	}

	for url, result := range results {
		fmt.Println(url, ":" , result)
	}
}