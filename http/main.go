package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	// bs := make([]byte, 99999)

	// // Read function here is an interface that is used by GO to read any type of input
	// // Read method always takes byte type of variable as an object and updates the same
	// resp.Body.Read(bs)

	// fmt.Println(string(bs))
	// io.Copy(something that implements writer_interface, something that implements reader_interface)
	// io.Copy(os.Stdout, resp.Body)

	lw := logWriter{}

	io.Copy(lw, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("\n Just wrote this many bytes : ", len(bs))
	return len(bs), nil
}
