package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	baseBull, targetBull, err := userInput()

	if err != nil {
		fmt.Println(err)
		return
	}

	filepath := filepath.Join("data", "inbreeding.csv")

	file, err := os.Open(filepath)

	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()

	if err != nil {
		fmt.Println("Failed to read csv: ", err)
		return
	}

	column := getColumnIndex(records, baseBull)
	row := getRowIndex(records, targetBull)
	inbreedingValue, err := getCellValue(records, column, row)

	if err != nil {
		fmt.Println("Failed to get inbreeding value")
	}

	fmt.Println(inbreedingValue)
}

func userInput() (string, string, error) {
	if len(os.Args) != 3 {
		return "", "", fmt.Errorf("usage: ./your_program <input1> <input2>")
	}

	input1 := os.Args[1]
	input2 := os.Args[2]

	return input1, input2, nil
}

func getColumnIndex(data [][]string, value string) int {
	for columnIndex, column := range data {
		if column[0] == value {
			fmt.Println("Column number: ", columnIndex)
			return columnIndex
		}
	}

	return -1
}

func getRowIndex(data [][]string, value string) int {
	for rowIndex, row := range data {
		if row[0] == value {
			fmt.Println("Row number: ", rowIndex)
			return rowIndex
		}
	}

	return -1
}

func getCellValue(data [][]string, column int, row int) (string, error) {
	if row >= len(data) || column >= len(data[0]) {
		return "", fmt.Errorf("index out of bounds")
	}

	return data[row][column], nil
}
