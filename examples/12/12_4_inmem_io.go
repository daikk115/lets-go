package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	// Create a new bytes.Buffer (empty)
	var buf bytes.Buffer

	// Write data to the buffer (implements io.Writer)
	buf.WriteString("Hello, ")
	buf.Write([]byte("world!"))

	// buf now contains "Hello, world!"

	// Create a byte slice to read into
	data := make([]byte, 2)

	// Read data from the buffer (implements io.Reader)
	n, err := buf.Read(data)
	if err != nil && err != io.EOF {
		panic(err)
	}

	fmt.Printf("Read %d bytes: %s\n", n, string(data[:n]))

	// You can also use io.Copy with bytes.Buffer
	fmt.Println("Remaining data in buffer:")
	io.Copy(os.Stdout, &buf) // prints rest of buffer content
}
