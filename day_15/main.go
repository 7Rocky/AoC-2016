package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func extendedGcd(a, b int) (int, int, int) {
	if a%b > 0 {
		u, v, d := extendedGcd(b, a%b)
		u = v
		v = (d - a*u) / b

		return u, v, d
	}

	return 0, 1, b
}

func prod(arr []int) int {
	prod := 1

	for _, a := range arr {
		prod *= a
	}

	return prod
}

func inv(n, m int) int {
	a, _, _ := extendedGcd(n, m)

	for a < 0 {
		a += m
	}

	return a
}

func next(rems, mods, partials []int) int {
	last := len(rems) - 1
	result := rems[last]

	for i, p := range partials {
		result -= prod(mods[:i]) * p
	}

	next := (result * inv(prod(mods[:last]), mods[last])) % mods[last]

	for next < 0 {
		next += mods[last]
	}

	return next
}

func solveCrt(rems, mods []int) int {
	var partials []int
	result := 0

	for i := 0; i < len(mods); i++ {
		partial := next(rems[:i+1], mods[:i+1], partials)
		partials = append(partials, partial)
		result += prod(mods[:i]) * partial
	}

	return result
}

func main() {
	file, _ := os.Open("input.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var rems, mods []int

	regex := regexp.MustCompile(`Disc #(\d+) has (\d+) positions; at time=0, it is at position (\d+).`)

	for scanner.Scan() {
		res := regex.FindAllStringSubmatch(scanner.Text(), -1)

		for i := range res {
			id, _ := strconv.Atoi(res[i][1])
			mod, _ := strconv.Atoi(res[i][2])
			init, _ := strconv.Atoi(res[i][3])

			rems = append(rems, (-init-id)%mod)
			mods = append(mods, mod)
		}
	}

	fmt.Printf("Time to press the button (1): %d\n", solveCrt(rems, mods))

	newDiscID, newDiscMod, newDiscInit := len(rems)+1, 11, 0

	fmt.Printf("Time to press the button with an extra disc (2): %d\n", solveCrt(append(rems, (-newDiscInit-newDiscID)%newDiscMod), append(mods, newDiscMod)))
}
