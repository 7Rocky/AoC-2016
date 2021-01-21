package main

import (
	"bufio"
	"fmt"
	"os"
)

var steps = map[rune]complex128{
	'U': 1i,
	'R': 1,
	'D': -1i,
	'L': -1,
}

func getBathroomCode(buttons map[complex128]string, instructions []string, position complex128) string {
	var code string

	for _, instruction := range instructions {
		for _, d := range instruction {
			step := steps[d]

			if _, ok := buttons[position+step]; ok {
				position += step
			}
		}

		code += buttons[position]
	}

	return code
}

func main() {
	file, _ := os.Open("input.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var instructions []string

	for scanner.Scan() {
		instructions = append(instructions, scanner.Text())
	}

	buttons := map[complex128]string{
		-1 + 1i: "1",
		1i:      "2",
		1 + 1i:  "3",
		-1:      "4",
		0:       "5",
		1:       "6",
		-1 - 1i: "7",
		-1i:     "8",
		1 - 1i:  "9",
	}

	fmt.Printf("Bathroom code (1): %s\n", getBathroomCode(buttons, instructions, 0))

	buttons = map[complex128]string{
		2i:      "1",
		-1 + 1i: "2",
		1i:      "3",
		1 + 1i:  "4",
		-2:      "5",
		-1:      "6",
		0:       "7",
		1:       "8",
		2:       "9",
		-1 - 1i: "A",
		-1i:     "B",
		1 - 1i:  "C",
		-2i:     "D",
	}

	fmt.Printf("Bathroom code (2): %s\n", getBathroomCode(buttons, instructions, -2))
}
