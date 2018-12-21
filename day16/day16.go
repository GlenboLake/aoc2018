package main

import (
	"bufio"
	"fmt"
	"github.com/glenbolake/aoc2018"
	"os"
)

func copyRegs(registers []int) []int {
	result := make([]int, len(registers))
	copy(result, registers)
	return result
}

type Case struct {
	Before, After   []int
	OpCode, A, B, C int
}

func (c Case) Test() map[int]bool {
	options := map[int]bool{}
	for i, op := range aoc2018.OpList {
		if c.testOp(op) {
			options[i] = true
		}
	}
	return options
}

func (c Case) testOp(op aoc2018.OpFunc) bool {
	regCopy := copyRegs(c.Before)
	op(c.A, c.B, c.C, regCopy)
	for i := range regCopy {
		if regCopy[i] != c.After[i] {
			return false
		}
	}
	return true
}

func part1(cases []Case) int {
	count := 0
	for _, c := range cases {
		if len(c.Test()) >= 3 {
			count += 1
		}
	}
	return count
}

func mapOpCodes(cases []Case) map[int]aoc2018.OpFunc {
	opCodes := map[int]aoc2018.OpFunc{}

	possibilities := map[int][]int{}
	list := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	for i := 0; i < 16; i++ {
		possibilities[i] = make([]int, 16)
		copy(possibilities[i], list)
	}

	for _, c := range cases {
		caseList := c.Test()
		var ok []int
		for _, p := range possibilities[c.OpCode] {
			if caseList[p] {
				ok = append(ok, p)
			}
		}
		possibilities[c.OpCode] = ok
	}
	determined := map[int]bool{}
	for len(determined) < 16 {
		for k, v := range possibilities {
			if len(v) == 1 {
				determined[v[0]] = true
			} else {
				var newList []int
				for _, x := range v {
					if !determined[x] {
						newList = append(newList, x)
					}
				}
				possibilities[k] = newList
			}
		}
	}

	for k, v := range possibilities {
		opCodes[k] = aoc2018.OpList[v[0]]
	}

	return opCodes
}

func part2(cases []Case, program []aoc2018.Instruction) int {
	opMap := mapOpCodes(cases)
	for i := 0; i < len(program); i++ {
		program[i].Op = opMap[program[i].OpIndex]
	}

	registers := make([]int, 4)
	for _, inst := range program {
		inst.Execute(registers)
	}

	return registers[0]
}

func main() {
	f, _ := os.Open("input/day16.txt")

	scanner := bufio.NewScanner(bufio.NewReader(f))
	var cases []Case
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		before := make([]int, 4)
		var opCode, a, b, c int
		after := make([]int, 4)
		fmt.Sscanf(scanner.Text(), "Before: [%d, %d, %d, %d]", &before[0], &before[1], &before[2], &before[3])
		scanner.Scan()
		fmt.Sscanf(scanner.Text(), "%d %d %d %d", &opCode, &a, &b, &c)
		scanner.Scan()
		fmt.Sscanf(scanner.Text(), "After: [%d, %d, %d, %d]", &after[0], &after[1], &after[2], &after[3])
		scanner.Scan()
		cases = append(cases, Case{Before: before, After: after, OpCode: opCode, A: a, B: b, C: c})
	}
	var program []aoc2018.Instruction
	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}
		var op, a, b, c int
		fmt.Sscanf(scanner.Text(), "%d %d %d %d", &op, &a, &b, &c)
		program = append(program, aoc2018.Instruction{OpIndex: op, A: a, B: b, C: c})
	}

	fmt.Println(part1(cases))
	fmt.Println(part2(cases, program))
}
