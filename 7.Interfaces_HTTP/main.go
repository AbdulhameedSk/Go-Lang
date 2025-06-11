package main

import (
	"fmt"
	"net/http"
	"os"
)
//Response is a struct type response has (status code-String, body-io.ReadCloser Interface, StatusCode-int)
//when it comes to body it is interface which further has interfaces named as Reader and Closer
// io.Reader is an interface that wraps the Read method, which reads data from a source.
// io.Closer is an interface that wraps the Close method, which closes the resource.
func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error fetching URL:", err)
		return
		os.Exit(1)
	}
	fmt.Println("Response:", resp)
}
