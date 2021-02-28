package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func getMarker(s string) (int, int) {
	regex := regexp.MustCompile(`(\d+)x(\d+)`)
	res := regex.FindStringSubmatch(s)

	x, _ := strconv.Atoi(res[1])
	y, _ := strconv.Atoi(res[2])

	return x, y
}

func decompressV1(compressed string) int {
	decompressed := ""

	for i := 0; i < len(compressed); i++ {
		r := compressed[i]

		if !strings.Contains(compressed[i:], ")") {
			return len(decompressed) + len(compressed[i:])
		}

		if r == '(' {
			closingParen := strings.Index(compressed[i:], ")")

			x, y := getMarker(compressed[i : i+closingParen])

			decompressed += strings.Repeat(compressed[closingParen+i+1:closingParen+i+x+1], y)
			i += closingParen + x

			continue
		}

		decompressed += string(r)
	}

	return len(decompressed)
}

func decompressV2(compressed string) int {
	length := 0

	for i := 0; i < len(compressed); i++ {
		r := compressed[i]

		if r == '(' {
			closingParen := strings.Index(compressed[i:], ")")

			x, y := getMarker(compressed[i : i+closingParen])

			length += y * decompressV2(compressed[closingParen+i+1:closingParen+i+x+1])
			i += closingParen + x

			continue
		}

		length++
	}

	return length
}

func main() {
	file, _ := os.Open("input.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	compressed := scanner.Text()

	fmt.Printf("Length of decompressed file V1 (1): %d\n", decompressV1(compressed))
	fmt.Printf("Length of decompressed file V2 (2): %d\n", decompressV2(compressed))
}
