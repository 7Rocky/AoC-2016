package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func runProgram(instructions []string, registers map[byte]int) {
	pc := 0

	for pc < len(instructions) {
		instr := instructions[pc]
		args := strings.Split(instr, " ")
		reg := args[1]

		switch string(instr[:3]) {
		case "cpy":
			dest := args[2]

			if strings.Contains("abcd", reg) {
				registers[dest[0]] = registers[reg[0]]
			} else {
				registers[dest[0]], _ = strconv.Atoi(reg)
			}
		case "inc":
			registers[reg[0]]++
		case "dec":
			registers[reg[0]]--
		case "jnz":
			if strings.Contains("abcd", reg) {
				if registers[reg[0]] != 0 {
					num, _ := strconv.Atoi(args[2])
					pc += num - 1
				}
			} else {
				cond, _ := strconv.Atoi(reg)

				if cond != 0 {
					num, _ := strconv.Atoi(args[2])
					pc += num - 1
				}
			}
		}

		pc++
	}
}

func main() {
	file, _ := os.Open("input.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var instructions []string

	for scanner.Scan() {
		instructions = append(instructions, scanner.Text())
	}

	var registers = map[byte]int{
		'a': 0,
		'b': 0,
		'c': 0,
		'd': 0,
	}

	runProgram(instructions, registers)

	fmt.Printf("Value of register 'a' (1): %d\n", registers['a'])

	registers = map[byte]int{
		'a': 0,
		'b': 0,
		'c': 1,
		'd': 0,
	}

	runProgram(instructions, registers)

	fmt.Printf("Value of register 'a' initiallizing register c to be 1 (2): %d\n", registers['a'])
}
