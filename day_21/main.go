package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	regexSwapPostion    = regexp.MustCompile(`swap position (\d) with position (\d)`)
	regexSwapLetter     = regexp.MustCompile(`swap letter (\D) with letter (\D)`)
	regexRotate         = regexp.MustCompile(`rotate (left|right) (\d) steps?`)
	regexRotatePosition = regexp.MustCompile(`rotate based on position of letter (\D)`)
	regexReverse        = regexp.MustCompile(`reverse positions (\d) through (\d)`)
	regexMove           = regexp.MustCompile(`move position (\d) to position (\d)`)
)

func reverse(arr []byte) []byte {
	var reversed []byte

	for i := len(arr) - 1; i >= 0; i-- {
		reversed = append(reversed, arr[i])
	}

	return reversed
}

func scramble(operation string, password []byte) []byte {

	if regexSwapPostion.MatchString(operation) {
		res := regexSwapPostion.FindAllStringSubmatch(operation, -1)

		x, _ := strconv.Atoi(res[0][1])
		y, _ := strconv.Atoi(res[0][2])

		password[x], password[y] = password[y], password[x]
	} else if regexSwapLetter.MatchString(operation) {
		res := regexSwapLetter.FindAllStringSubmatch(operation, -1)

		x := bytes.IndexByte(password, res[0][1][0])
		y := bytes.IndexByte(password, res[0][2][0])

		password[x], password[y] = password[y], password[x]
	} else if regexRotate.MatchString(operation) {
		res := regexRotate.FindAllStringSubmatch(operation, -1)

		direction := res[0][1]
		steps, _ := strconv.Atoi(res[0][2])

		if direction == "right" {
			for i := 0; i < steps; i++ {
				password = append(password[len(password)-1:], password[:len(password)-1]...)
			}
		} else if direction == "left" {
			for i := 0; i < steps; i++ {
				password = append(password[1:], password[0])
			}
		}
	} else if regexRotatePosition.MatchString(operation) {
		res := regexRotatePosition.FindAllStringSubmatch(operation, -1)

		steps := bytes.IndexByte(password, res[0][1][0])

		if steps >= 4 {
			steps++
		}

		for i := 0; i < steps+1; i++ {
			password = append(password[len(password)-1:], password[:len(password)-1]...)
		}
	} else if regexReverse.MatchString(operation) {
		res := regexReverse.FindAllStringSubmatch(operation, -1)

		x, _ := strconv.Atoi(res[0][1])
		y, _ := strconv.Atoi(res[0][2])

		reversed := reverse(password[x : y+1])

		if y+1 < len(password) {
			reversed = append(reversed, password[y+1:]...)
		}

		password = append(password[:x], reversed...)
	} else if regexMove.MatchString(operation) {
		res := regexMove.FindAllStringSubmatch(operation, -1)

		x, _ := strconv.Atoi(res[0][1])
		y, _ := strconv.Atoi(res[0][2])

		var partialPassword []byte

		if x > y {
			partialPassword = append(partialPassword, password[0:y]...)
			partialPassword = append(partialPassword, password[x])
			partialPassword = append(partialPassword, password[y:x]...)
			password = append(partialPassword, password[x+1:]...)
		} else {
			partialPassword = append(partialPassword, password[0:x]...)
			partialPassword = append(partialPassword, password[x+1:y+1]...)
			partialPassword = append(partialPassword, password[x])
			password = append(partialPassword, password[y+1:]...)
		}
	}

	return password
}

func unscramble(operation string, password []byte) []byte {
	if regexRotate.MatchString(operation) {
		if strings.Contains(operation, "left") {
			password = scramble(strings.Replace(operation, "left", "right", 1), password)
		} else {
			password = scramble(strings.Replace(operation, "right", "left", 1), password)
		}
	} else if regexRotatePosition.MatchString(operation) {
		wanted := string(password)

		for wanted != string(scramble(operation, []byte(string(password)))) {
			password = append(password[1:], password[0])
		}
	} else if regexMove.MatchString(operation) {
		res := regexMove.FindAllStringSubmatch(operation, -1)

		x, _ := strconv.Atoi(res[0][1])
		y, _ := strconv.Atoi(res[0][2])

		password = scramble(fmt.Sprintf("move position %d to position %d", y, x), password)
	} else {
		password = scramble(operation, password)
	}

	return password
}

func main() {
	file, _ := os.Open("input.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)

	password := []byte("abcdefgh")

	var operations []string

	for scanner.Scan() {
		operations = append(operations, scanner.Text())
	}

	for _, operation := range operations {
		password = scramble(operation, password)
	}

	fmt.Printf("Scrambled result (1): %s\n", password)

	password = []byte("fbgdceah")

	for i := len(operations) - 1; i >= 0; i-- {
		password = unscramble(operations[i], password)
	}

	fmt.Printf("Unscrambled password (2): %s\n", password)
}
