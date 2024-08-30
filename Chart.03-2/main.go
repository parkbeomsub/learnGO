package main

// request/ response
import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFaild = errors.New("Request failed")

type requestResult struct {
	url    string
	status string
}

func main() {
	c := make(chan requestResult)
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

	for _, url := range urls {
		go hitURL(url, c)

	}
	for i := 0; i < len(urls); i++ {
		fmt.Println(<-c)
	}
}

func hitURL(url string, c chan<- requestResult) {
	fmt.Println("Checking ", url)
	resp, err := http.Get(url)
	status := "OK"
	if err != nil || resp.StatusCode >= 400 {
		status = "FAILED"

	}
	c <- requestResult{url: url, status: status}

}
