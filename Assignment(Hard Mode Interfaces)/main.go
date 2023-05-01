package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	var file_name string

	if len(os.Args) > 1 {
		file_name = os.Args[1]
	}

	if file_name == "" {
		fmt.Println("Filename mandatory !")
		os.Exit(1)
	}

	file, err := os.Open(file_name)

	if err != nil {
		fmt.Println("Error in opening file : ", err)
		os.Exit(1)
	}

	_, err = io.Copy(os.Stdout, file)
	if err != nil {
		fmt.Println("Error in printing file : ", err)
		os.Exit(1)
	}

}
