package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
)

func countPossible(triangles [][]int) int {
	count := 0

	for _, triangle := range triangles {
		if triangle[0]+triangle[1] > triangle[2] {
			count++
		}
	}

	return count
}

func sortSides(triangle []int) {
	sort.Slice(triangle, func(i, j int) bool {
		return triangle[i] < triangle[j]
	})
}

func main() {
	file, _ := os.Open("input.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var trianglesRows, trianglesColumns [][]int

	var triangle1, triangle2, triangle3 []int

	regex := regexp.MustCompile(`(\d+)\s+(\d+)\s+(\d+)`)

	i := 0

	for scanner.Scan() {
		res := regex.FindAllStringSubmatch(scanner.Text(), -1)

		for i := range res {
			a, _ := strconv.Atoi(res[i][1])
			b, _ := strconv.Atoi(res[i][2])
			c, _ := strconv.Atoi(res[i][3])

			triangle1 = append(triangle1, a)
			triangle2 = append(triangle2, b)
			triangle3 = append(triangle3, c)

			triangle := []int{a, b, c}

			sortSides(triangle)

			trianglesRows = append(trianglesRows, triangle)
		}

		i++

		if i == 3 {
			sortSides(triangle1)
			sortSides(triangle2)
			sortSides(triangle3)

			trianglesColumns = append(trianglesColumns, triangle1, triangle2, triangle3)
			triangle1, triangle2, triangle3 = []int{}, []int{}, []int{}
			i = 0
		}
	}

	fmt.Printf("Possible triangles by rows (1): %d\n", countPossible(trianglesRows))
	fmt.Printf("Possible triangles by columns (2): %d\n", countPossible(trianglesColumns))
}
