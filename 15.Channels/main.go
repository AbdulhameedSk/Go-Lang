// Channels are like pipes that connect concurrent go routines.
// Channels also have types
package main

//Giving one output every time a link is checked
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
	fmt.Println(<-c) // receive from channel c
}

// sending data to channel c syntax: c <- value
// receiving data from channel c syntax: value := <-c
// print ing data from channel c syntax: fmt.Println(<-c)
func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		println(link, "is down")
		c <- "Down" // send the link to channel c
		return
	}
	fmt.Println(link, "is up")
	c <- "Up" // send the link to channel c
}
