package main

import (
	"fmt"
	"strconv"
)

func part1(add int) string {
	scoreboard := []int{3, 7}
	elf1, elf2 := 0, 1

	for len(scoreboard) < 10+add {
		newScores := scoreboard[elf1] + scoreboard[elf2]
		if newScores > 9 {
			scoreboard = append(scoreboard, 1)
		}
		scoreboard = append(scoreboard, newScores%10)
		elf1 = (elf1 + scoreboard[elf1] + 1) % len(scoreboard)
		elf2 = (elf2 + scoreboard[elf2] + 1) % len(scoreboard)
	}
	rv := ""
	for i := 0; i < 10; i++ {
		rv += strconv.Itoa(scoreboard[i+add])
	}
	return rv
}

func slicesEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func part2(input int) int {
	scoreboard := []int{3, 7}
	elf1, elf2 := 0, 1
	var seq []int
	for input > 0 {
		seq = append([]int{input % 10}, seq...)
		input /= 10
	}

	for {
		newScores := scoreboard[elf1] + scoreboard[elf2]
		if newScores > 9 {
			scoreboard = append(scoreboard, 1)
		}
		scoreboard = append(scoreboard, newScores%10)
		elf1 = (elf1 + scoreboard[elf1] + 1) % len(scoreboard)
		elf2 = (elf2 + scoreboard[elf2] + 1) % len(scoreboard)
		if len(scoreboard) > len(seq) {
			if slicesEqual(scoreboard[len(scoreboard)-len(seq):], seq) {
				return len(scoreboard) - len(seq)
			}
			if newScores > 9 {
				//fmt.Println(len(scoreboard), len(seq))
				if slicesEqual(scoreboard[len(scoreboard)-len(seq)-1:len(scoreboard)-1], seq) {
					return len(scoreboard) - len(seq) - 1
				}
			}
		}
	}

	return 0
}

const INPUT = 793061

func main() {
	fmt.Println(part1(INPUT))
	fmt.Println(part2(INPUT))
}
