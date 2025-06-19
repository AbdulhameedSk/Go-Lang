package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"https://www.google.com",
		"https://www.example.com",
		"https://www.test.com",
		"https://www.github.com",
		"https://www.stackoverflow.com",
		"https://www.reddit.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	// for  {
	// 	go checkLink(<-c,c)
	// }
	for l := range c {
		// go checkLink(l, c)
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(l, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "is down")
		c <- link
		return
	}
	fmt.Println(link, "is up")
	c <- link
}
