package aoc2018

type OpFunc func(A, B, C int, regs []int)

type Instruction struct {
	OpIndex int
	Op      OpFunc
	A, B, C int
}

func (i Instruction) Execute(registers []int) {
	i.Op(i.A, i.B, i.C, registers)
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

var AllOps = map[string]OpFunc{
	"addr": addr, "addi": addi,
	"mulr": mulr, "muli": muli,
	"banr": banr, "bani": bani,
	"borr": borr, "bori": bori,
	"setr": setr, "seti": seti,
	"gtir": gtir, "gtri": gtri, "gtrr": gtrr,
	"eqir": eqir, "eqri": eqri, "eqrr": eqrr,
}

var OpList = []OpFunc{
	addr, addi,
	mulr, muli,
	banr, bani,
	borr, bori,
	setr, seti,
	gtir, gtri, gtrr,
	eqir, eqri, eqrr,
}
