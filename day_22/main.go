package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

type fileNode struct {
	name  complex128
	size  int
	used  int
	avail int
	use   int
}

func getViablePairs() int {
	var viablePairs = map[complex128]complex128{}

	for _, f1 := range files {
		for _, f2 := range files {
			if f1.name != f2.name && f1.used > 0 && f1.used <= f2.avail {
				viablePairs[f1.name] = f2.name
			}
		}
	}

	return len(viablePairs)
}

var files []fileNode

func main() {
	file, _ := os.Open("input.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var regex = regexp.MustCompile(`/dev/grid/node-x(\d+)-y(\d+)\s+(\d+)T\s+(\d+)T\s+(\d+)T\s+(\d+)%`)

	scanner.Scan()
	scanner.Scan()

	for scanner.Scan() {
		res := regex.FindAllStringSubmatch(scanner.Text(), -1)

		x, _ := strconv.Atoi(res[0][1])
		y, _ := strconv.Atoi(res[0][2])
		name := complex(float64(x), float64(y))
		size, _ := strconv.Atoi(res[0][3])
		used, _ := strconv.Atoi(res[0][4])
		avail, _ := strconv.Atoi(res[0][5])
		use, _ := strconv.Atoi(res[0][6])

		files = append(files, fileNode{name: name, size: size, used: used, avail: avail, use: use})
	}

	fmt.Printf("Viable file pairs (1): %d\n", getViablePairs())

	var empty complex128
	var maxX int

	for _, f := range files {
		if f.used == 0 {
			empty = f.name
		}

		maxX = int(math.Max(float64(maxX), real(f.name)))
	}

	result := int(real(empty)+imag(empty)) + maxX - 1 + (maxX-1)*5 + 1

	fmt.Printf("Minimum steps to move goal data (2): %d\n", result)
}
