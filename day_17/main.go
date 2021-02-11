package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
)

var directions = map[complex128]string{
	-1i: "U",
	1i:  "D",
	-1:  "L",
	1:   "R",
}

func adjacentPositions(pos position) []position {
	var positions []position

	hash := md5.Sum([]byte(input + pos.steps))
	hashString := hex.EncodeToString(hash[:])

	for i, d := range []complex128{-1i, 1i, -1, 1} {
		coord := pos.coordinates + d

		if real(coord) >= 0 && imag(coord) >= 0 && real(coord) <= 3 && imag(coord) <= 3 && strings.Contains("bcdef", string(hashString[i])) {
			positions = append(positions, position{coordinates: coord, steps: directions[d]})
		}
	}

	return positions
}

func breadthFirstSearch(root position, target complex128) (string, int) {
	queue := []position{root}

	var pos position
	var paths []string

	for len(queue) > 0 {
		pos = queue[0]
		queue = queue[1:]

		if pos.coordinates == target {
			paths = append(paths, pos.steps)
			continue
		}

		adjacent := adjacentPositions(pos)

		for _, next := range adjacent {
			next.steps = pos.steps + next.steps
			queue = append(queue, next)
		}
	}

	return paths[0], len(paths[len(paths)-1])
}

type position struct {
	coordinates complex128
	steps       string
}

const input = "mmsxrhfx"

func main() {
	shortestPath, maxLength := breadthFirstSearch(position{coordinates: 0, steps: ""}, 3+3i)

	fmt.Printf("Shortest path to reach the vault (1): %s\n", shortestPath)
	fmt.Printf("Longest path length to reach the vault (2): %d\n", maxLength)
}
