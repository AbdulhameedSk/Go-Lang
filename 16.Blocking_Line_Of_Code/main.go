// In Go, a blocking line of code involving channels is any operation that causes the goroutine to pause (block) until the operation can proceed. This typically happens with send (chan <- value) and receive (value := <-chan) operations when conditions aren't ready.
// ch := make(chan int)  // unbuffered channel
// ch <- 10 // BLOCKS until another goroutine receives from the channel
package main

import (
	"fmt"
	"net/http"
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

	fmt.Print("Checking links...\n")
	// fmt.Println(<-c)
	// fmt.Println(<-c)

	// fmt.Println(<-c)

	// fmt.Println(<-c)

	// fmt.Println(<-c)
	//if one more link is down, it will block the main goroutine until it receives a message from the channel
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "is down")
		c <- "Down"
		return
	}
	fmt.Println(link, "is up")
	c <- "Up"
}
