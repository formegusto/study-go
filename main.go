package main

import (
	"fmt"

	"github.com/formegusto/study-go/urlutils"
)

func main() {
	results := make(map[string]string)
	c := make(chan urlutils.RequestResult)

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
		go urlutils.HitURLByGo(url, c)
	}

	for i:=0;i<len(urls);i++ {
		result := <-c
		results[result.Url] = result.Status
	}

	for url, status := range results {
		fmt.Println(url, status)
	}
}