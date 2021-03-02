package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func runProgram(instructions []string, registers map[byte]int) int {
	pc := 0
	clk := -1
	count := 0

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
					var num int

					if strings.Contains("abcd", args[2]) {
						num = registers[args[2][0]]
					} else {
						num, _ = strconv.Atoi(args[2])
					}

					pc += num - 1
				}
			}
		case "out":
			var nextClk int

			if strings.Contains("abcd", reg) {
				nextClk = registers[reg[0]]
			} else {
				nextClk, _ = strconv.Atoi(reg)
			}

			if nextClk != 0 && nextClk != 1 {
				return -1
			}

			if clk == -1 {
				clk = 1 - nextClk
			}

			if clk+nextClk == 1 {
				clk = nextClk
				count++

				if count == 100 {
					return 0
				}
			} else {
				return -1
			}
		}

		pc++
	}

	return -1
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}

	return n * factorial(n-1)

}

func main() {
	file, _ := os.Open("input.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var instructions []string

	for scanner.Scan() {
		instructions = append(instructions, scanner.Text())
	}

	a := 0

	for runProgram(instructions, map[byte]int{'a': a, 'b': 0, 'c': 0, 'd': 0}) != 0 {
		a++
	}

	fmt.Printf("Minimum value of register 'a' to generate CLK (1): %d\n", a)
}
