package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	dimX = 50
	dimY = 6
)

type rule struct {
	kind string
	x    int
	y    int
}

func rect(x, y int) {
	for j := 0; j < y; j++ {
		screen[j] = strings.Repeat("#", x) + screen[j][x:]
	}
}

func row(x, y int) {
	screen[y] = screen[y][dimX-x:] + screen[y][:dimX-x]
}

func column(x, y int) {
	var column string

	for j := 0; j < dimY; j++ {
		column += string(screen[j][x])
	}

	for j := 0; j < dimY; j++ {
		screen[j] = screen[j][:x] + string(column[(j-y+dimY)%dimY]) + screen[j][x+1:]
	}
}

func countPixelsLit() int {
	count := 0

	for _, row := range screen {
		count += strings.Count(row, "#")
	}

	return count
}

func printScreen() {
	for _, row := range screen {
		fmt.Println(row)
	}
}

var screen []string

func main() {
	file, _ := os.Open("input.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for i := 0; i < dimY; i++ {
		screen = append(screen, strings.Repeat(".", dimX))
	}

	var rules []rule

	for scanner.Scan() {
		var x, y int
		var kind string

		if strings.HasPrefix(scanner.Text(), "rect") {
			kind = "rect"
			regex := regexp.MustCompile(`rect (\d+)x(\d+)`)
			res := regex.FindAllStringSubmatch(scanner.Text(), -1)

			for i := range res {
				x, _ = strconv.Atoi(res[i][1])
				y, _ = strconv.Atoi(res[i][2])
			}
		}

		if strings.HasPrefix(scanner.Text(), "rotate row") {
			kind = "row"
			regex := regexp.MustCompile(`rotate row y=(\d+) by (\d+)`)
			res := regex.FindAllStringSubmatch(scanner.Text(), -1)

			for i := range res {
				y, _ = strconv.Atoi(res[i][1])
				x, _ = strconv.Atoi(res[i][2])
			}
		}

		if strings.HasPrefix(scanner.Text(), "rotate column") {
			kind = "column"
			regex := regexp.MustCompile(`rotate column x=(\d+) by (\d+)`)
			res := regex.FindAllStringSubmatch(scanner.Text(), -1)

			for i := range res {
				x, _ = strconv.Atoi(res[i][1])
				y, _ = strconv.Atoi(res[i][2])
			}
		}

		rules = append(rules, rule{kind, x, y})
	}

	for _, r := range rules {
		switch r.kind {
		case "rect":
			rect(r.x, r.y)
		case "row":
			row(r.x, r.y)
		case "column":
			column(r.x, r.y)
		}
	}

	fmt.Printf("Number of pixels lit (1): %d\n", countPixelsLit())
	fmt.Printf("Message(2):\n\n")
	printScreen()
}
