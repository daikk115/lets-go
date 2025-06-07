package main

import (
	"bufio"
	"fmt"
	"os"
)

func bufferingIO() {
	fmt.Println("\n------------------ bufferingIO ------------------")
	wd, _ := os.Getwd()
	file, err := os.Open(fmt.Sprintf("%s/go.mod", wd))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}

func main() {
	bufferingIO()
}
