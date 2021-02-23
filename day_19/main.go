package main

import "fmt"

type elf struct {
	number     int
	next, prev *elf
}

func round(elf, opElf *elf, level int) int {
	length := input

	for length > 1 {
		elf = elf.next
		length--

		if level == 1 {
			elf.next = elf.next.next
		} else {
			opElf.prev.next = opElf.next
			opElf.next.prev = opElf.prev

			opElf = opElf.next

			if length%2 == 0 {
				opElf = opElf.next
			}
		}
	}

	return elf.number
}

func generateElves() (*elf, *elf) {
	lastElf := &elf{number: input}
	pElf := lastElf
	var opElf *elf

	for i := 1; i < input; i++ {
		pElf.next = &elf{number: i, prev: pElf}
		pElf.next.prev = pElf
		pElf = pElf.next

		if i == input/2+1 {
			opElf = pElf
		}
	}

	pElf.next = lastElf
	lastElf.prev = pElf

	return lastElf, opElf
}

const input = 3012210

func main() {
	elf1, _ := generateElves()
	elf2, opElf := generateElves()

	fmt.Printf("Elf with all presents (1): %d\n", round(elf1, nil, 1))
	fmt.Printf("Elf with all presents (2): %d\n", round(elf2, opElf, 2))
}
