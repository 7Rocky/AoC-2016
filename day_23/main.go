package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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
					var num int

					if strings.Contains("abcd", args[2]) {
						num = registers[args[2][0]]
					} else {
						num, _ = strconv.Atoi(args[2])
					}

					pc += num - 1
				}
			}
		case "tgl":
			num := registers[reg[0]] + pc

			if num >= 0 && num < len(instructions) {
				count := strings.Count(instructions[num], " ")

				if count == 1 {
					if strings.Contains(instructions[num], "inc") {
						instructions[num] = "dec" + instructions[num][3:]
					} else {
						instructions[num] = "inc" + instructions[num][3:]
					}
				} else if count == 2 {
					if strings.Contains(instructions[num], "jnz") {
						instructions[num] = "cpy" + instructions[num][3:]
					} else {
						instructions[num] = "jnz" + instructions[num][3:]
					}
				}
			}
		}

		pc++
	}
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

	eggs := 7

	var registers = map[byte]int{
		'a': eggs,
		'b': 0,
		'c': 0,
		'd': 0,
	}

	runProgram(instructions, registers)

	fmt.Printf("Value of register 'a' (1): %d\n", registers['a'])

	eggs = 12
	x := 1

	regex := regexp.MustCompile(`cpy (\d+) [cd]`)

	for _, instr := range instructions {
		if regex.MatchString(instr) {
			num, _ := strconv.Atoi(regex.FindStringSubmatch(instr)[1])
			x *= num
		}
	}

	fmt.Printf("Value of register 'a' (2): %d\n", factorial(eggs)+x)
}
