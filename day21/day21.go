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
		if registers[ip] == 28 {
			return registers[1]
		}
		inst := program[registers[ip]]
		aoc2018.AllOps[inst.OpCode](inst.A, inst.B, inst.C, registers)
		registers[ip] += 1
	}
	return 0
}

func part2(ip int, program []aoc2018.Instruction) int {
	registers := make([]int, 6)
	seenStates := map[int]bool{}
	lastSeen := 0
	for {
		if registers[ip] == 28 {
			if seenStates[registers[1]] {
				return lastSeen
			} else {
				seenStates[registers[1]] = true
				lastSeen = registers[1]
			}
		}
		inst := program[registers[ip]]
		aoc2018.AllOps[inst.OpCode](inst.A, inst.B, inst.C, registers)
		registers[ip] += 1
	}
}

func main() {
	f, _ := os.Open("input/day21.txt")

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
