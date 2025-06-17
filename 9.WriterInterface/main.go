package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("https://www.google.com")
	if err != nil {
		fmt.Println("Error fetching URL:", err)
		return
	}
	io.Copy(os.Stdout, resp.Body)
}
