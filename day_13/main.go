package main

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
)

func decToBin(number, length int) []int {
	var remainder int
	remainders := make([]int, length)

	for i := 0; i < length; i++ {
		number, remainder = number/2, number%2
		remainders[length-i-1] = remainder
	}

	return remainders
}

func sum(arr []int) int {
	sum := 0

	for _, a := range arr {
		sum += a
	}

	return sum
}

func getPointRepresentation(x, y int) string {
	number := x*x + 3*x + 2*x*y + y + y*y + input
	bin := decToBin(number, int(math.Floor(math.Log2(float64(number))+1)))

	if sum(bin)%2 == 1 {
		return "#"
	}

	return "."
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

		if coord[0] >= 0 && coord[1] >= 0 && getPointRepresentation(coord[0], coord[1]) != "#" {
			positions = append(positions, position{coordinates: coord})
		}
	}

	for _, y := range []int{-1, 1} {
		coord := []int{pos.coordinates[0], pos.coordinates[1] + y}

		if coord[0] >= 0 && coord[1] >= 0 && getPointRepresentation(coord[0], coord[1]) != "#" {
			positions = append(positions, position{coordinates: coord})
		}
	}

	return positions
}

func getCoordinatesString(coordinates []int) string {
	return strconv.Itoa(coordinates[0]) + "," + strconv.Itoa(coordinates[1])
}

func breadthFirstSearch(root position, target []int) int {
	queue := []position{root}
	visitedPositions[getCoordinatesString(root.coordinates)] = 0

	var pos position

	for len(queue) > 0 {
		pos = queue[0]
		queue = queue[1:]

		if reflect.DeepEqual(pos.coordinates, target) {
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

	return pos.steps
}

const input = 1358

func main() {
	minSteps := breadthFirstSearch(position{coordinates: []int{1, 1}, steps: 0}, []int{31, 39})

	fmt.Printf("Minimum steps to reach target (1): %d\n", minSteps)

	numPositions := 0

	for _, v := range visitedPositions {
		if v <= 50 {
			numPositions++
		}
	}

	fmt.Printf("Number of positions reached in 50 steps (2): %d\n", numPositions)
}
