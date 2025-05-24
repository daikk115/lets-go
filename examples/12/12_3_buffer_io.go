package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Open file for reading
	file, err := os.Open("doc.go")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Wrap file with a buffered reader
	reader := bufio.NewReader(file)

	for {
		// Read until newline ('\n') or EOF
		line, err := reader.ReadString('\n')
		if err != nil {
			// EOF is expected at end of file
			if err.Error() == "EOF" {
				if len(line) > 0 {
					fmt.Print(line) // print last line if no trailing newline
				}
				break
			}
			panic(err)
		}
		fmt.Println(line) // print each line including newline with \n
	}
}
