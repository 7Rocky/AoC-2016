package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type ipv7 struct {
	supernet []string
	hypernet []string
}

func isABBA(s string) bool {
	return s[0] == s[3] && s[1] == s[2] && s[0] != s[1]
}

func isABA(s string) bool {
	return s[0] == s[2] && s[0] != s[1]
}

func matchABAWithBAB(a, b string) bool {
	return a[0] == b[1] && a[1] == b[0]
}

func hasABBA(s string) bool {
	supports := false

	for i := 0; i < len(s)-3; i++ {
		if isABBA(s[i : i+4]) {
			supports = true
		}
	}

	return supports
}

func getABAStrings(ss []string) []string {
	var abaStrings []string

	for _, s := range ss {
		for i := 0; i < len(s)-2; i++ {
			if isABA(s[i : i+3]) {
				abaStrings = append(abaStrings, s[i:i+3])
			}
		}
	}

	return abaStrings
}

var supportsTLS = func(ip ipv7) bool {
	supports := false

	for _, f := range ip.supernet {
		if hasABBA(f) {
			supports = true
			break
		}
	}

	for _, s := range ip.hypernet {
		if hasABBA(s) {
			supports = false
			break
		}
	}

	return supports
}

var supportsSSL = func(ip ipv7) bool {
	abaStrings := getABAStrings(ip.supernet)
	babStrings := getABAStrings(ip.hypernet)

	for _, a := range abaStrings {
		for _, b := range babStrings {
			if matchABAWithBAB(a, b) {
				return true
			}
		}
	}

	return false
}

func countIPSupporting(ips []ipv7, supports func(ip ipv7) bool) int {
	count := 0

	for _, ip := range ips {
		if supports(ip) {
			count++
		}
	}

	return count
}

func main() {
	file, _ := os.Open("input.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var ips []ipv7

	for scanner.Scan() {
		count := strings.Count(scanner.Text(), "[")
		regex := regexp.MustCompile(strings.Repeat(`(.+)\[(.+)\]`, count) + "(.+)")

		res := regex.FindStringSubmatch(scanner.Text())

		var supernet, hypernet []string

		for j := 0; j < count; j++ {
			supernet = append(supernet, res[2*j+1])
			hypernet = append(hypernet, res[2*j+2])
		}

		supernet = append(supernet, res[2*count+1])

		ips = append(ips, ipv7{supernet, hypernet})
	}

	fmt.Printf("Number of IPv7 supporting TLS (1): %d\n", countIPSupporting(ips, supportsTLS))
	fmt.Printf("Number of IPv7 supporting SSL (2): %d\n", countIPSupporting(ips, supportsSSL))
}
