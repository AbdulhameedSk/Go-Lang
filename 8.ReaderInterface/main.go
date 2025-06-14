// Every value has a type.
// Every function should have a return type.
// Before interfaces, every function should specify the type of its arguments.

// No matter what the source of input to our program is, if it implements Reader, it will output a byte slice.

// type Reader interface {
//     Read(p []byte) (n int, err error)
// }

// So the Reader takes raw data and a byte slice as input (which is empty and is the data we want to read),
// returns an int (the number of bytes read) and an error, if any.

package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error fetching URL:", err)
		return
		os.Exit(1)
	}
	// bs:=[]byte{} // Create an empty byte slice to hold the response body
	bs := make([]byte, 99999) // Create a byte slice to hold the response body 99999 bytes	
	resp.Body.Read(bs)
	fmt.Println("Response:", string(bs))
}
