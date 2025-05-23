package main

import (
	"encoding/csv"
	"os"
)

func main() {
	records := [][]string{
		{"Name", "Age", "Country"},
		{"Alice", "30", "USA"},
		{"Bob", "25", "UK"},
	}

	f, _ := os.Create("data.csv")
	writer := csv.NewWriter(f)
	defer f.Close()

	for _, record := range records {
		writer.Write(record)
	}
	writer.Flush()
}
