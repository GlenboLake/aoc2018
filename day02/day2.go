package main

import (
	"bufio"
	"fmt"
	"os"
)

func letterCounts(s string) (bool, bool) {
	counts := make(map[rune]int)
	for _, l := range s {
		counts[l] += 1
	}

	var twos, threes bool
	for _, b := range counts {
		if b == 2 {
			twos = true
		}
		if b == 3 {
			threes = true
		}
	}
	return twos, threes
}

func part1(input []string) int {
	var twos, threes int
	for _, line := range input {
		two, three := letterCounts(line)
		if two {
			twos += 1
		}
		if three {
			threes += 1
		}
	}
	return twos * threes
}

func part2(input []string) string {
	for i, s1 := range input {
		for _, s2 := range input[i+1:] {
			matches := ""
			for j := 0; j < len(s1); j++ {
				if s1[j] == s2[j] {
					matches += string(rune(s1[j]))
				}
			}
			if len(matches) == len(s1)-1 {
				return matches
			}
		}
	}
	return ""
}

func main() {
	f, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(bufio.NewReader(f))
	input := make([]string, 0, 250)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	fmt.Println(12, part1([]string{"abcdef", "bababc", "abbcde", "abcccd", "aabcdd", "abcdee", "ababab"}))
	fmt.Println(part1(input))

	fmt.Println("fgij", part2([]string{"abcde", "fghij", "klmno", "pqrst", "fguij", "axcye", "wvxyz"}))
	fmt.Println(part2(input))
}
