package urlutils

import (
	"errors"
	"fmt"
	"net/http"
)

type RequestResult struct {
	Url string
	Status string
}

var errRequestFailed = errors.New("Request Failed")
func HitURL(url string) error {
	fmt.Println("Checking: ", url)
	res, err := http.Get(url)
	if err != nil || res.StatusCode >= 400 {
		return errRequestFailed
	}
	return nil
}

// chan<- 이 채널은 보낼 수만 있고 받을 수는 없다.
func HitURLByGo(url string, c chan<- RequestResult) {
	fmt.Println("Checking: ", url)
	res, err := http.Get(url)
	if err != nil || res.StatusCode >= 400 {
		c <- RequestResult { Url: url, Status: "FAILED" }
	} else {
		c <- RequestResult { Url: url, Status: "OK" }
	}
}

func HitURLTest() {
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
		err := HitURL(url)
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