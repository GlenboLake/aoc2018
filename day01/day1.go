package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func part1(input []int) int {
	sum := 0
	for _, val := range input {
		sum += val
	}

	return sum
}

func part2(input []int) int {
	history := map[int]bool{0: true}
	freq := 0

	for {
		for _, val := range input {
			freq += val
			_, seen := history[freq]
			if seen {
				return freq
			} else {
				history[freq] = true
			}
		}
	}
}

func main() {
	f, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(bufio.NewReader(f))
	input := make([]int, 0, 1000)
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		input = append(input, num)
	}
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}
