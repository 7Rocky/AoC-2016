package main

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

const difficulty = "00000"

func main() {
	input := "cxdnnyjw"

	password1 := ""
	password2 := bytes.Repeat([]byte{'_'}, 8)

	index := 0

	for len(password1) < 8 || bytes.Contains(password2, []byte{'_'}) {
		hash := md5.Sum([]byte(input + strconv.Itoa(index)))
		hashString := hex.EncodeToString(hash[:])

		if strings.HasPrefix(hashString, difficulty) {
			if len(password1) < 8 {
				password1 += string(hashString[5])
			}

			position, err := strconv.Atoi(string(hashString[5]))

			if err == nil && position < 8 && password2[position] == '_' {
				password2[position] = hashString[6]
			}
		}

		index++
	}

	fmt.Printf("Password (1): %s\n", password1)
	fmt.Printf("Password (2): %s\n", string(password2))
}
