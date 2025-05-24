package main

import (
	"io"
	"log"
	"os"
)

func main() {
	// Open file as a source stream (io.Reader)
	file, err := os.Open("doc.go")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// os.Stdout is a target stream (io.Writer)
	// Copy data from file to stdout (stream copy)
	_, err = io.Copy(os.Stdout, file)
	if err != nil {
		log.Fatal(err)
	}
}
