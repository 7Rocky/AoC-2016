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

	for char, count := range letterCount {
		if count > max {
			max = count
			mostCommon = char
		}

		if count < min {
			min = count
			leastCommon = char
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
		for i, char := range scanner.Text() {
			if len(letterCounts) == i {
				letterCounts = append(letterCounts, make(map[rune]int))
			}

			letterCounts[i][char]++
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
