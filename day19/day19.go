package main

import (
	"bufio"
	"fmt"
	"os"
)

type Op func(A, B, C int, regs []int)

type Instruction struct {
	OpCode  string
	A, B, C int
}

func addr(A, B, C int, regs []int) {
	regs[C] = regs[A] + regs[B]
}

func addi(A, B, C int, regs []int) {
	regs[C] = regs[A] + B
}

func mulr(A, B, C int, regs []int) {
	regs[C] = regs[A] * regs[B]
}

func muli(A, B, C int, regs []int) {
	regs[C] = regs[A] * B
}

func banr(A, B, C int, regs []int) {
	regs[C] = regs[A] & regs[B]
}

func bani(A, B, C int, regs []int) {
	regs[C] = regs[A] & B
}

func borr(A, B, C int, regs []int) {
	regs[C] = regs[A] | regs[B]
}

func bori(A, B, C int, regs []int) {
	regs[C] = regs[A] | B
}

func setr(A, _, C int, regs []int) {
	regs[C] = regs[A]
}

func seti(A, _, C int, regs []int) {
	regs[C] = A
}

func gtir(A, B, C int, regs []int) {
	if A > regs[B] {
		regs[C] = 1
	} else {
		regs[C] = 0
	}
}

func gtri(A, B, C int, regs []int) {
	if regs[A] > B {
		regs[C] = 1
	} else {
		regs[C] = 0
	}
}

func gtrr(A, B, C int, regs []int) {
	if regs[A] > regs[B] {
		regs[C] = 1
	} else {
		regs[C] = 0
	}
}

func eqir(A, B, C int, regs []int) {
	if A == regs[B] {
		regs[C] = 1
	} else {
		regs[C] = 0
	}
}

func eqri(A, B, C int, regs []int) {
	if regs[A] == B {
		regs[C] = 1
	} else {
		regs[C] = 0
	}
}

func eqrr(A, B, C int, regs []int) {
	if regs[A] == regs[B] {
		regs[C] = 1
	} else {
		regs[C] = 0
	}
}

var AllOps = map[string]Op{
	"addr": addr, "addi": addi,
	"mulr": mulr, "muli": muli,
	"banr": banr, "bani": bani,
	"borr": borr, "bori": bori,
	"setr": setr, "seti": seti,
	"gtir": gtir, "gtri": gtri, "gtrr": gtrr,
	"eqir": eqir, "eqri": eqri, "eqrr": eqrr,
}

func part1(ip int, program []Instruction) int {
	registers := make([]int, 6)

	for registers[ip] >= 0 && registers[ip] < len(program) {
		instruction := program[registers[ip]]
		AllOps[instruction.OpCode](instruction.A, instruction.B, instruction.C, registers)
		registers[ip] += 1
	}

	return registers[0]
}
func part2(ip int, program []Instruction) int {
	registers := make([]int, 6)
	registers[0] = 1

	for registers[ip] >= 0 && registers[ip] < len(program) {
		if registers[ip] == 1 {
			break
		}
		instruction := program[registers[ip]]
		AllOps[instruction.OpCode](instruction.A, instruction.B, instruction.C, registers)
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
	var program []Instruction
	for scanner.Scan() {
		var inst string
		var a, b, c int
		fmt.Sscanf(scanner.Text(), "%s %d %d %d", &inst, &a, &b, &c)
		program = append(program, Instruction{OpCode: inst, A: a, B: b, C: c})
	}

	fmt.Println(part1(ip, program))
	fmt.Println(part2(ip, program))
}
