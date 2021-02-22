package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func generateTraps(rows []string) []string {
	lastRow, newRow := rows[len(rows)-1], []byte{}

	if lastRow[:2] == "^^" || lastRow[:2] == ".^" {
		newRow = append(newRow, '^')
	} else {
		newRow = append(newRow, '.')
	}

	for i := 1; i < len(lastRow)-1; i++ {
		tiles := lastRow[i-1 : i+2]

		if tiles == "^^." || tiles == ".^^" || tiles == "^.." || tiles == "..^" {
			newRow = append(newRow, '^')
		} else {
			newRow = append(newRow, '.')
		}
	}

	if lastRow[len(lastRow)-2:] == "^^" || lastRow[len(lastRow)-2:] == "^." {
		newRow = append(newRow, '^')
	} else {
		newRow = append(newRow, '.')
	}

	return append(rows, string(newRow))
}

func main() {
	file, _ := os.Open("input.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var rows []string

	for scanner.Scan() {
		rows = append(rows, scanner.Text())
	}

	for len(rows) < 40 {
		rows = generateTraps(rows)
	}

	safeTraps := 0

	for _, row := range rows {
		safeTraps += strings.Count(row, ".")
	}

	fmt.Printf("Number of safe tiles (1): %d\n", safeTraps)

	for len(rows) < 400000 {
		rows = generateTraps(rows)
	}

	for _, row := range rows[40:] {
		safeTraps += strings.Count(row, ".")
	}

	fmt.Printf("Number of safe tiles (2): %d\n", safeTraps)
}
