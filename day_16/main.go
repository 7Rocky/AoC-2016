package main

import (
	"fmt"
)

func reverseNegate(s []byte) []byte {
	var r []byte

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '1' {
			r = append(r, '0')
		} else {
			r = append(r, '1')
		}
	}

	return r
}

func modifiedDragonCurve(input string, length int) []byte {
	a := []byte(input)
	inputLength := len(input)
	separators := []byte{'0'}

	b := reverseNegate(a)

	for (len(separators)+1)*inputLength+len(separators) < length {
		reversedSeparators := reverseNegate(separators)
		separators = append(separators, '0')
		separators = append(separators, reversedSeparators...)
	}

	result := a

	for i, c := range separators {
		result = append(result, c)

		if i%2 == 0 {
			result = append(result, b...)
		} else {
			result = append(result, a...)
		}
	}

	return result[:length]
}

func getChecksum(input string, length int) string {
	result := modifiedDragonCurve(input, length)
	var checksum []byte

	for len(checksum)%2 == 0 {
		checksum = []byte{}

		for i := 0; i < len(result); i += 2 {
			if result[i] == result[i+1] {
				checksum = append(checksum, '1')
			} else {
				checksum = append(checksum, '0')
			}
		}

		result = checksum
	}

	return string(checksum)
}

func main() {
	input := "10111011111001111"

	length := 272
	fmt.Printf("Correct checksum (1): %s\n", getChecksum(input, length))

	length = 35651584
	fmt.Printf("Correct checksum (2): %s\n", getChecksum(input, length))
}
