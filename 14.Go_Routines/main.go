package main
//Does not give output coz first main completes and child routines are not completed so main exits befire child routines complete 
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
	for _, link := range links {
		go checkLink(link)

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
