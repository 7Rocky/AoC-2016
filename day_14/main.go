package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var regex = getRegexp()

func getRegexp() *regexp.Regexp {
	var r string

	for i := 0; i < 16; i++ {
		r += "|" + strings.Repeat(string(hex.EncodeToString([]byte{byte(i)})[1]), 3)
	}

	return regexp.MustCompile(string(r[1:]))
}

func hash(str string, times int) string {
	hashString := str

	for i := 0; i < times; i++ {
		hash := md5.Sum([]byte(hashString))
		hashString = hex.EncodeToString(hash[:])
	}

	return hashString
}

func findNextHash(index int, char string, times int) {
	stream := strings.Repeat(char, 5)

	for i := index; i < index+1000; i++ {
		if len(hashes) <= i {
			hashString := hash(input+strconv.Itoa(len(hashes)+1), times)
			hashes = append(hashes, hashString)
		}

		if strings.Contains(hashes[i], stream) {
			keys = append(keys, index)
			break
		}
	}
}

var hashes []string
var keys []int

const totalKeys = 64

func getLastIndexKey(times int) int {
	hashes = []string{}
	keys = []int{}

	var index = 1

	for len(keys) < totalKeys {
		var hashString string

		if index > len(hashes) {
			hashString = hash(input+strconv.Itoa(len(hashes)+1), times)
			hashes = append(hashes, hashString)
		} else {
			hashString = hashes[index-1]
		}

		res := regex.FindString(hashString)

		if len(res) == 3 {
			findNextHash(index, string(res[0]), times)
		}

		index++
	}

	return keys[totalKeys-1]
}

var input = "ihaygndm"

func main() {
	fmt.Printf("Index producing the 64th one-time pad key (1): %d\n", getLastIndexKey(1))
	fmt.Printf("Index producing the 64th one-time pad key (2): %d\n", getLastIndexKey(2017))
}
