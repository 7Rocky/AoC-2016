package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var blacklist [][]int

	for scanner.Scan() {
		limits := strings.Split(scanner.Text(), "-")
		min, _ := strconv.Atoi(limits[0])
		max, _ := strconv.Atoi(limits[1])

		blacklist = append(blacklist, []int{min, max})
	}

	sort.Slice(blacklist, func(i, j int) bool {
		return blacklist[i][0] < blacklist[j][0]
	})

	max := blacklist[0][1]
	allowedIPs := 0
	lowestAllowedIP := 0

	for i := 1; i < len(blacklist); i++ {
		if max+1 >= blacklist[i][0] {
			if max < blacklist[i][1] {
				max = blacklist[i][1]
			}
		} else {
			allowedIPs += blacklist[i][0] - max - 1

			if lowestAllowedIP == 0 {
				lowestAllowedIP = max + 1
			}

			max = blacklist[i][1]
		}
	}

	fmt.Printf("Lowest IP allowed (1): %d\n", lowestAllowedIP)
	fmt.Printf("Number of allowed IPs (2): %d\n", allowedIPs)
}
