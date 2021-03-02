package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var permutations [][]byte

func perm(a []byte, f func([]byte), i int) {
	if i > len(a) {
		f(a)
		return
	}

	perm(a, f, i+1)

	for j := i + 1; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		perm(a, f, i+1)
		a[i], a[j] = a[j], a[i]
	}
}

type position struct {
	coordinates []int
	steps       int
}

var visitedPositions = map[string]int{}

func adjacentPositions(pos position) []position {
	var positions []position

	for _, x := range []int{-1, 1} {
		coord := []int{pos.coordinates[0] + x, pos.coordinates[1]}

		if coord[0] >= 0 && coord[1] >= 0 && layout[coord[1]][coord[0]] != '#' {
			positions = append(positions, position{coordinates: coord})
		}
	}

	for _, y := range []int{-1, 1} {
		coord := []int{pos.coordinates[0], pos.coordinates[1] + y}

		if coord[0] >= 0 && coord[1] >= 0 && layout[coord[1]][coord[0]] != '#' {
			positions = append(positions, position{coordinates: coord})
		}
	}

	return positions
}

func getCoordinatesString(coordinates []int) string {
	return strconv.Itoa(coordinates[0]) + "," + strconv.Itoa(coordinates[1])
}

func breadthFirstSearch(root position, target byte) position {
	queue := []position{root}
	visitedPositions = map[string]int{}
	visitedPositions[getCoordinatesString(root.coordinates)] = 0

	var pos position

	for len(queue) > 0 {
		pos = queue[0]
		queue = queue[1:]

		if layout[pos.coordinates[1]][pos.coordinates[0]] == target {
			break
		}

		adjacent := adjacentPositions(pos)

		for _, next := range adjacent {
			if _, ok := visitedPositions[getCoordinatesString(next.coordinates)]; !ok {
				next.steps = pos.steps + 1
				queue = append(queue, next)
				visitedPositions[getCoordinatesString(next.coordinates)] = next.steps
			}
		}
	}

	return pos
}

func getPathSteps(path []byte, initialPosition []int) int {
	pos := breadthFirstSearch(position{coordinates: initialPosition}, path[0])

	for i := 1; i < len(path); i++ {
		pos = breadthFirstSearch(pos, path[i])
	}

	return pos.steps
}

func getShortestPathLength(initialPosition []int) int {
	steps := math.MaxFloat64

	for _, path := range permutations {
		steps = math.Min(float64(steps), float64(getPathSteps(path, initialPosition)))
	}

	return int(steps)
}

var layout []string

func main() {
	file, _ := os.Open("input.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var initialPosition []int

	for scanner.Scan() {
		layout = append(layout, scanner.Text())

		if strings.Contains(scanner.Text(), "0") {
			initialPosition = []int{strings.Index(scanner.Text(), "0"), len(layout) - 1}
		}
	}

	perm([]byte{'1', '2', '3', '4', '5', '6', '7'}, func(a []byte) {
		x := make([]byte, len(a))
		copy(x, a)
		permutations = append(permutations, x)
	}, 0)

	fmt.Printf("Shortest path passing through all numbers (1): %d\n",
		getShortestPathLength(initialPosition))

	for i := range permutations {
		permutations[i] = append(permutations[i], '0')
	}

	fmt.Printf("Shortest path passing through all numbers and returning to 0 (2): %d\n",
		getShortestPathLength(initialPosition))
}
