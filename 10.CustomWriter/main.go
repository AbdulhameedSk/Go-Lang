package main

import (
	"fmt"
	"io"
	"net/http"
)

type Writer interface {
	Write(p []byte) (n int, err error)
}

type CWriter struct{}

func (CWriter) Write(bs []byte) (int, error) {
	fmt.Println("Data received:", string(bs))
	fmt.Println("Number of bytes received:", len(bs))
	// Here you can process the data as needed, e.g., save to a file, send
	// to another service, etc.
	// For demonstration, we will just print the length of the byte slice.
	// You can also return an error if needed, but for this example, we will return nil
	// to indicate no error occurred.
	return len(bs), nil // Return the number of bytes written
}

func main() {
	resp, err := http.Get("https://www.google.com")
	if err != nil {
		fmt.Println("Error fetching URL:", err)
		return
	}
	cw := CWriter{} // Create an instance of CWriter
	io.Copy(cw, resp.Body)
}
