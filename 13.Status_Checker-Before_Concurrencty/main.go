package main

import (
	"fmt"
	"net/http"
)

// Takes first link and requests and waits for response before moving to next link
func main() {
	links := []string{
		"https://www.google.com",
		"https://www.example.com",
		"https://www.test.com",
		"https://www.github.com",
		"https://www.stackoverflow.com",
		"https://www.reddit.com",
	}
	for _, link := range links {
		checkLink(link)

	}
}
func checkLink(link string) {
	_, err := http.Get(link)
	if err != nil {
		println(link, "is down")
		return
	}
	fmt.Println(link, "is up")
}
