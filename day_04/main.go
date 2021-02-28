package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func getLetterCount(name string) map[byte]int {
	letterCounts := make(map[byte]int)

	for _, c := range name {
		if c != '-' {
			letterCounts[byte(c)] = strings.Count(name, string(c))
		}
	}

	return letterCounts
}

func getCorrectChecksum(r room) string {
	letterCount := getLetterCount(r.name)

	checksum := ""

	for len(checksum) < 5 {
		max := 0
		var letter byte = 255

		for c, n := range letterCount {
			if n >= max && !strings.Contains(checksum, string(c)) {
				if n > max {
					max = n
					letter = c
				} else if n == max && c < letter {
					max = n
					letter = c
				}
			}
		}

		checksum += string(letter)
	}

	return checksum
}

func getRealRooms(rooms []room) []room {
	var realRooms []room

	for _, r := range rooms {
		if r.checksum == getCorrectChecksum(r) {
			realRooms = append(realRooms, r)
		}
	}

	return realRooms
}

func sumIds(rooms []room) int {
	sum := 0

	for _, r := range rooms {
		sum += r.id
	}

	return sum
}

func caesarDecypher(str string, shift int) string {
	decypher := ""

	for _, c := range str {
		if c == '-' {
			decypher += " "
		}

		if c >= 'a' && c <= 'z' {
			d := ((int(c)-'a'+1)+shift)%('z'-'a'+1) + 'a' - 1
			decypher += string(byte(d))
		}
	}

	return decypher
}

func getDecryptedNames(rooms []room) []room {
	var decypherNames []room

	for _, r := range rooms {
		decypherNames = append(decypherNames, room{
			name:     caesarDecypher(r.name, r.id),
			id:       r.id,
			checksum: r.checksum,
		})
	}

	return decypherNames
}

type room struct {
	name     string
	id       int
	checksum string
}

func main() {
	file, _ := os.Open("input.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var rooms []room

	regex := regexp.MustCompile(`(\D+)-(\d+)\[(\D+)\]`)

	for scanner.Scan() {
		res := regex.FindStringSubmatch(scanner.Text())

		name := res[1]
		id, _ := strconv.Atoi(res[2])
		checksum := res[3]

		rooms = append(rooms, room{name: name, id: id, checksum: checksum})
	}

	realRooms := getRealRooms(rooms)

	decryptedNames := getDecryptedNames(rooms)
	northPoleID := 0

	for _, r := range decryptedNames {
		if strings.Contains(r.name, "pole") {
			northPoleID = r.id
		}
	}

	fmt.Printf("Sum of real rooms' IDs (1): %d\n", sumIds(realRooms))
	fmt.Printf("North Pole object storage ID (2): %d\n", northPoleID)
}
