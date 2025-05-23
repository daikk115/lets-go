package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("data.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	for i, row := range records {
		if i == 0 {
			fmt.Println("Header:", row)
			continue
		}
		fmt.Printf("Name: %s, Age: %s, Country: %s\n", row[0], row[1], row[2])
	}
}
