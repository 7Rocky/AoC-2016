package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func containsPosition(position complex128, reachedPositions []complex128) bool {
	for _, p := range reachedPositions {
		if p == position {
			return true
		}
	}

	return false
}

func getBlocks(position complex128) int {
	return int(math.Abs(real(position)) + math.Abs(imag(position)))
}

func main() {
	file, _ := os.Open("input.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	instructions := strings.Split(scanner.Text(), ", ")

	position := 0i
	direction := 1i

	twiceReachedPosition := 0i
	var reachedPositions []complex128

	for _, instruction := range instructions {
		if instruction[0] == 'R' {
			direction *= -1i
		} else if instruction[0] == 'L' {
			direction *= 1i
		}

		steps, _ := strconv.Atoi(instruction[1:])

		for i := 0; i < steps; i++ {
			if containsPosition(position, reachedPositions) && twiceReachedPosition == 0 {
				twiceReachedPosition = position
			}

			reachedPositions = append(reachedPositions, position)
			position += direction
		}
	}

	fmt.Printf("Blocks (1): %d\n", getBlocks(position))
	fmt.Printf("Blocks to twice reached position (2): %d\n", getBlocks(twiceReachedPosition))
}
