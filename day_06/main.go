package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func findCommonCharacters(letterCount map[rune]int) (string, string) {
	var mostCommon, leastCommon rune
	max, min := 0, math.MaxUint32

	for c, n := range letterCount {
		if n > max {
			max = n
			mostCommon = c
		}

		if n < min {
			min = n
			leastCommon = c
		}
	}

	return string(mostCommon), string(leastCommon)
}

func main() {
	file, _ := os.Open("input.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var letterCounts []map[rune]int

	for scanner.Scan() {
		for i, c := range scanner.Text() {
			if len(letterCounts) == i {
				letterCounts = append(letterCounts, make(map[rune]int))
			}

			letterCounts[i][c]++
		}
	}

	mostCommonCorrection, leastCommonCorrection := "", ""

	for _, letterCounts := range letterCounts {
		mostCommon, leastCommon := findCommonCharacters(letterCounts)

		mostCommonCorrection += mostCommon
		leastCommonCorrection += leastCommon
	}

	fmt.Printf("Error-corrected message (1): %s\n", mostCommonCorrection)
	fmt.Printf("Error-corrected message (2): %s\n", leastCommonCorrection)
}
