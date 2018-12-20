package main

import (
	"bufio"
	"fmt"
	"github.com/glenbolake/aoc2018"
	"os"
)

func part1(ip int, program []aoc2018.Instruction) int {
	registers := make([]int, 6)

	for registers[ip] >= 0 && registers[ip] < len(program) {
		instruction := program[registers[ip]]
		aoc2018.AllOps[instruction.OpCode](instruction.A, instruction.B, instruction.C, registers)
		registers[ip] += 1
	}

	return registers[0]
}
func part2(ip int, program []aoc2018.Instruction) int {
	registers := make([]int, 6)
	registers[0] = 1

	for registers[ip] >= 0 && registers[ip] < len(program) {
		if registers[ip] == 1 {
			break
		}
		instruction := program[registers[ip]]
		aoc2018.AllOps[instruction.OpCode](instruction.A, instruction.B, instruction.C, registers)
		registers[ip] += 1
	}
	total := 0
	num := registers[2]
	for i := 1; i <= num; i++ {
		if num%i == 0 {
			total += i
		}
	}

	return total
}

func main() {
	f, _ := os.Open("input/day19.txt")

	scanner := bufio.NewScanner(bufio.NewReader(f))
	scanner.Scan()
	var ip int
	fmt.Sscanf(scanner.Text(), "#ip %d", &ip)
	var program []aoc2018.Instruction
	for scanner.Scan() {
		var inst string
		var a, b, c int
		fmt.Sscanf(scanner.Text(), "%s %d %d %d", &inst, &a, &b, &c)
		program = append(program, aoc2018.Instruction{OpCode: inst, A: a, B: b, C: c})
	}

	fmt.Println(part1(ip, program))
	fmt.Println(part2(ip, program))
}
