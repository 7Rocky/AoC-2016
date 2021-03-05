package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"regexp"
	"sort"
)

type situation struct {
	state []complex128
	steps int
}

func indexOfString(arr []string, s interface{}) int {
	for i, a := range arr {
		if a == s {
			return i
		}
	}

	return -1
}

func indexOfInt(arr []int, s interface{}) int {
	for i, a := range arr {
		if a == s {
			return i
		}
	}

	return -1
}

func ordinalComplex(c complex128) float64 {
	return real(c) + 4*imag(c)
}

func sortState(state []complex128) {
	sort.Slice(state, func(i, j int) bool {
		return ordinalComplex(state[i]) < ordinalComplex(state[j])
	})
}

func possibleStates(state []complex128) [][]complex128 {
	var possibleStates [][]complex128

	for _, d := range []complex128{1, 1i, -1, -1i} {
		for i := 1; i < len(state); i++ {
			if real(state[i]) == real(state[0]) {
				stateCopy := make([]complex128, len(state))
				copy(stateCopy, state)

				stateCopy[0] += complex(real(d)+imag(d), 0)
				stateCopy[i] += d

				possibleStates = append(possibleStates, stateCopy)
			}

			for j := 1; j < len(state); j++ {
				if j != i && real(state[i]) == real(state[0]) && real(state[j]) == real(state[0]) {
					stateCopy := make([]complex128, len(state))
					copy(stateCopy, state)

					stateCopy[0] += complex(real(d)+imag(d), 0)
					stateCopy[i] += d
					stateCopy[j] += d

					possibleStates = append(possibleStates, stateCopy)
				}
			}
		}
	}

	for _, d := range []complex128{1 + 1i, -1 - 1i} {
		for i := 1; i < len(state); i++ {
			if real(state[i]) == real(state[0]) && imag(state[i]) == real(state[0]) {
				stateCopy := make([]complex128, len(state))
				copy(stateCopy, state)

				stateCopy[0] += complex(real(d), 0)
				stateCopy[i] += d

				possibleStates = append(possibleStates, stateCopy)
			}
		}
	}

	return possibleStates
}

func getMembers(state []complex128) ([]int, []int) {
	var generators, microchips []int

	for _, s := range state[1:] {
		generators = append(generators, int(real(s)))
		microchips = append(microchips, int(imag(s)))
	}

	return generators, microchips
}

func count(state []int, floor int) int {
	count := 0

	for i := 0; i < len(state); i++ {
		if state[i] == floor {
			count++
		}
	}

	return count
}

func containsFloorLessThan(state []complex128, floor int) bool {
	for i, s := range state {
		if i == 0 {
			if real(s) < float64(floor) {
				return true
			}
		} else if real(s) < float64(floor) || imag(s) < float64(floor) {
			return true
		}
	}

	return false
}

func containsFloorGreaterThan(state []complex128, floor int) bool {
	for _, s := range state {
		if real(s) > float64(floor) || imag(s) > float64(floor) {
			return true
		}
	}

	return false
}

func isPossible(state []complex128) bool {
	if isValid, ok := validStates[getStateString(state)]; ok {
		return isValid
	}

	if containsFloorLessThan(state, 1) || containsFloorGreaterThan(state, 4) {
		validStates[getStateString(state)] = false
		return false
	}

	generators, microchips := getMembers(state)

	for i, m := range microchips {
		if m != generators[i] {
			indexGenerator := indexOfInt(generators, m)

			if indexGenerator != -1 && count(microchips, m) >= count(generators, m) {
				validStates[getStateString(state)] = false
				return false
			}
		}
	}

	validStates[getStateString(state)] = true
	return true
}

func containsState(situations []situation, state []complex128) bool {
	for _, s := range situations {
		if reflect.DeepEqual(s.state, state) {
			return true
		}
	}

	return false
}

func adjacentSituations(sit situation) []situation {
	var situations []situation

	for _, s := range possibleStates(sit.state) {
		sortState(s)

		if isPossible(s) {
			if !containsState(situations, s) && !reflect.DeepEqual(s, sit.state) {
				situations = append(situations, situation{state: s})
			}
		}
	}

	graph[getStateString(sit.state)] = situations

	return situations
}

func getStateString(state []complex128) string {
	return fmt.Sprintf("%v", state)
}

var graph map[string][]situation = map[string][]situation{}
var validStates map[string]bool = map[string]bool{}
var visitedStates map[string]bool = map[string]bool{}

func breadthFirstSearch(root situation, target []complex128) int {
	queue := []situation{root}
	visitedStates[getStateString(root.state)] = true

	var sit situation

	for len(queue) > 0 {
		sit = queue[0]
		queue = queue[1:]

		if reflect.DeepEqual(sit.state, target) {
			break
		}

		adjacent, ok := graph[getStateString(sit.state)]

		if !ok {
			adjacent = adjacentSituations(sit)
		}

		for _, next := range adjacent {
			if _, ok := visitedStates[getStateString(next.state)]; !ok {
				next.steps = sit.steps + 1
				queue = append(queue, next)
				visitedStates[getStateString(next.state)] = true
			}
		}
	}

	return sit.steps
}

func main() {
	file, _ := os.Open("input.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var microchips []string

	floor := 1.0
	state := []complex128{1}
	regex := regexp.MustCompile(`(\w+)(-compatible| generator)`)

	for scanner.Scan() {
		res := regex.FindAllStringSubmatch(scanner.Text(), -1)

		for i := range res {
			index := indexOfString(microchips, res[i][1])

			if index == -1 {
				microchips = append(microchips, res[i][1])
				state = append(state, 0)
				index = len(state) - 2
			}

			if res[i][2] == "-compatible" {
				state[index+1] += complex(0, floor)
			} else {
				state[index+1] += complex(floor, 0)
			}
		}

		floor++
	}

	target := make([]complex128, len(state))

	for i := range target {
		target[i] = 4 + 4i
	}

	target[0] = 4

	fmt.Printf("Minimum number of steps with %d pairs generator-microchip (1): ", len(microchips))
	fmt.Println(breadthFirstSearch(situation{state: state}, target))

	microchips = append(microchips, "elerium", "dilithium")
	state = append(state, 1+1i, 1+1i)

	target = append(target, 4+4i, 4+4i)

	graph = map[string][]situation{}
	validStates = map[string]bool{}
	visitedStates = map[string]bool{}

	fmt.Printf("Minimum number of steps with %d pairs generator-microchip (2): ", len(microchips))
	fmt.Println(breadthFirstSearch(situation{state: state}, target))
}
