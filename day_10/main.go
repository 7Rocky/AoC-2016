package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type bot struct {
	id         int
	microchips []int
	giveBots   []int
	output     []int
}

const (
	targetMicrochipA = 61
	targetMicrochipB = 17
)

var bots map[int]*bot = make(map[int]*bot)
var outputs []int

func giveBots(b *bot) {
	sort.Slice(b.microchips, func(i, j int) bool {
		return b.microchips[i] < b.microchips[j]
	})

	if b.giveBots[0] != -1 {
		botLow := bots[b.giveBots[0]]
		botLow.microchips = append(botLow.microchips, b.microchips[0])
	} else {
		outputs[b.output[0]] = b.microchips[0]
	}

	if b.giveBots[1] != -1 {
		botHigh := bots[b.giveBots[1]]
		botHigh.microchips = append(botHigh.microchips, b.microchips[1])
	} else {
		outputs[b.output[1]] = b.microchips[1]
	}

	b.microchips = []int{}
}

func indexOf(arr []int, a int) int {
	for i := range arr {
		if arr[i] == a {
			return i
		}
	}

	return -1
}

func prod(arr []int) int {
	p := 1

	for _, a := range arr {
		p *= a
	}

	return p
}

func getMatches(line, reString string) (int, int, int) {
	regex := regexp.MustCompile(reString)
	res := regex.FindAllStringSubmatch(line, -1)

	var id, low, high int

	for i := range res {
		id, _ = strconv.Atoi(res[i][1])
		low, _ = strconv.Atoi(res[i][2])
		high, _ = strconv.Atoi(res[i][3])
	}

	return id, low, high
}

func main() {
	file, _ := os.Open("input.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)

	maxOutputIndex := 0

	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "value") {
			regex := regexp.MustCompile(`value (\d+) goes to bot (\d+)`)
			res := regex.FindAllStringSubmatch(line, -1)

			var microchip, id int

			for i := range res {
				microchip, _ = strconv.Atoi(res[i][1])
				id, _ = strconv.Atoi(res[i][2])
			}

			if b, ok := bots[id]; !ok {
				bots[id] = &bot{id, []int{microchip}, []int{-1, -1}, []int{-1, -1}}
			} else {
				b.microchips = append(b.microchips, microchip)
			}
		} else if strings.Count(line, "output") == 1 {
			if strings.Contains(line, "high to output") {
				id, low, high := getMatches(line, `bot (\d+) gives low to bot (\d+) and high to output (\d+)`)

				if high > maxOutputIndex {
					maxOutputIndex = high
				}

				if b, ok := bots[id]; !ok {
					bots[id] = &bot{id, []int{}, []int{low, -1}, []int{-1, high}}
				} else {
					b.giveBots = []int{low, -1}
					b.output = []int{-1, high}
				}
			} else {
				id, low, high := getMatches(line, `bot (\d+) gives low to output (\d+) and high to bot (\d+)`)

				if low > maxOutputIndex {
					maxOutputIndex = low
				}

				if b, ok := bots[id]; !ok {
					bots[id] = &bot{id, []int{}, []int{-1, high}, []int{low, -1}}
				} else {
					b.giveBots = []int{-1, high}
					b.output = []int{low, -1}
				}
			}
		} else if strings.Count(line, "output") == 2 {
			id, low, high := getMatches(line, `bot (\d+) gives low to output (\d+) and high to output (\d+)`)

			if low > maxOutputIndex {
				maxOutputIndex = low
			}

			if high > maxOutputIndex {
				maxOutputIndex = high
			}

			if b, ok := bots[id]; !ok {
				bots[id] = &bot{id, []int{}, []int{-1, -1}, []int{low, high}}
			} else {
				b.giveBots = []int{-1, -1}
				b.output = []int{low, high}
			}
		} else {
			id, low, high := getMatches(line, `bot (\d+) gives low to bot (\d+) and high to bot (\d+)`)

			if b, ok := bots[id]; !ok {
				bots[id] = &bot{id, []int{}, []int{low, high}, []int{-1, -1}}
			} else {
				b.giveBots = []int{low, high}
			}
		}
	}

	outputs = make([]int, maxOutputIndex+1)

	var targetBotID int

	for indexOf(outputs[:3], 0) != -1 {
		for id, b := range bots {
			if indexOf(b.microchips, targetMicrochipA) != -1 && indexOf(b.microchips, targetMicrochipB) != -1 {
				targetBotID = id
			}

			if len(b.microchips) > 1 {
				giveBots(b)
			}
		}
	}

	fmt.Printf("Bot comparing value-%d and value-%d (1): %d\n",
		targetMicrochipA, targetMicrochipB, targetBotID)
	fmt.Printf("Product of first three outputs (2): %d\n", prod(outputs[:3]))
}
