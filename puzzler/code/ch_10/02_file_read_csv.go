package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	filePath := "./read_csv.csv"

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file.go:", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		return
	}

	for _, row := range records {
		for _, col := range row {
			fmt.Print(col, "\t")
		}
		fmt.Println()
	}
}
