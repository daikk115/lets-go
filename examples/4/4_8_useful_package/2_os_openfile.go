package main

import (
	"fmt"
	"io"
	"os"
)

func openFile() {
	fmt.Println("\n------------------ OpenFile ------------------")
	wd, _ := os.Getwd()
	file, err := os.Open(fmt.Sprintf("%s/go.mod", wd))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))
}

func main() {
	openFile()
}
