package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func Abs(val int32) int32 {
	if val < 0 {
		return -val
	} else {
		return val
	}
}

func part1(input string) int {

	var done bool
	for len(input) > 0 && !done {
		done = true
		for i, ch := range input[1:] {
			if Abs(ch-int32(input[i])) == 32 {
				input = input[:i] + input[i+2:]
				done = false
				break
			}
		}
	}
	return len(input)
}

func remove(s string, rem rune) string {
	return strings.Map(func(r rune) rune {
		if r == rem || r == rem-32 {
			return -1
		}
		return r
	}, s)
}

func part2(input string) int {
	options := map[rune]bool{}

	for _, ch := range input {
		if ch >= 'a' && ch <= 'z' {
			options[ch] = true
		}
	}

	best := len(input)
	for o := range options {
		length := part1(remove(input, o))
		if length < best {
			best = length
		}
	}
	return best
}

func main() {
	inputBytes, _ := ioutil.ReadFile("input/day05.txt")
	input := strings.TrimSpace(string(inputBytes))

	fmt.Println("Part 1: ", part1(input))
	fmt.Println("Part 2: ", part2(input))
}
