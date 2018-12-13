package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func nextGen(rules map[string]bool, plants []int) []int {
	var result []int

	plantSet := map[int]bool{}
	for _, p := range plants {
		plantSet[p] = true
	}

	for i := plants[0] - 2; i < plants[len(plants)-1]+3; i++ {
		var state string
		for j := i - 2; j < i+3; j++ {
			if plantSet[j] {
				state += "#"
			} else {
				state += "."
			}
		}
		if rules[state] {
			result = append(result, i)
		}
	}
	return result
}

func part1(rules map[string]bool, plants []int) int {
	for i := 0; i < 20; i++ {
		plants = nextGen(rules, plants)
	}

	total := 0
	for _, plant := range plants {
		total += plant
	}
	return total
}

func part2(rules map[string]bool, plants []int) int {
	diff := 0
	times := 0
	lastTotal := 0
	for _, plant := range plants {
		lastTotal += plant
	}
	for i := 0; i < 50000000000; i++ {
		plants = nextGen(rules, plants)
		total := 0
		for _, plant := range plants {
			total += plant
		}
		newDiff := total - lastTotal
		lastTotal = total
		if newDiff == diff {
			times += 1
			if times > 10 {
				gensRemaining := 50000000000 - i - 1
				return total + diff*gensRemaining
			}
		} else {
			diff = newDiff
			times = 0
		}
	}

	total := 0
	for _, plant := range plants {
		total += plant
	}
	return total
}

func main() {
	f, _ := os.Open("input/day12.txt")

	scanner := bufio.NewScanner(bufio.NewReader(f))
	scanner.Scan()
	var plants []int
	for i, plant := range strings.Split(scanner.Text(), " ")[2] {
		if plant == '#' {
			plants = append(plants, i)
		}
	}
	scanner.Scan()
	rules := map[string]bool{}
	for scanner.Scan() {
		rule := strings.SplitN(scanner.Text(), " => ", 2)
		if rule[1] == "#" {
			rules[rule[0]] = true
		}
	}

	fmt.Println(part1(rules, plants))
	fmt.Println(part2(rules, plants))
}
