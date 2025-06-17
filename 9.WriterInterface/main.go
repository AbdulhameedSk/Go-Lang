package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// Writer is interface
type Writer interface {
	Write(p []byte) (n int, err error)
}

// Writer interface is used to write data to a destination.
//byte[] is working as argument here
//it sends data to some output channal like text file, console, etc.
// The Write method takes a byte slice as input and returns the number of bytes written and an error, if any.

func main() {
	resp, err := http.Get("https://www.google.com")
	if err != nil {
		fmt.Println("Error fetching URL:", err)
		return
	}
	// Copy the response body to the standard output
	//stdout is type file
	//first argument is Writer interface
	//second argument is Reader interface
	//os.stout is file type and it has funcito called Write
	io.Copy(os.Stdout, resp.Body)
}
