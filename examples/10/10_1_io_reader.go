package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	var r io.Reader = strings.NewReader("Hello, io.Reader!")

	buf := make([]byte, 1) // read 1 bytes at a time, can be changed to 2, 4, 100, etc.

	// Read from the io.Reader until EOF
	for {
		n, err := r.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		fmt.Println(string(buf[:n]))
	}
}
