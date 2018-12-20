package main

import (
	"bufio"
	"fmt"
	"github.com/glenbolake/aoc2018"
	"os"
)

const (
	NORTH = 'N'
	SOUTH = 'S'
	EAST  = 'E'
	WEST  = 'W'
)

func mapDistances(input string) map[aoc2018.Coord]int {
	var branchSources []aoc2018.Coord
	currentPosition := aoc2018.Coord{}
	doors := map[aoc2018.Coord]map[rune]bool{currentPosition: {}}
	distances := map[aoc2018.Coord]int{currentPosition: 0}

	for _, ch := range input {
		switch ch {
		case NORTH:
			doors[currentPosition][NORTH] = true
			dist := distances[currentPosition]
			currentPosition = aoc2018.Coord{X: currentPosition.X, Y: currentPosition.Y - 1}
			if doors[currentPosition] == nil {
				doors[currentPosition] = map[rune]bool{}
			}
			doors[currentPosition][SOUTH] = true
			if val, ok := distances[currentPosition]; val > dist+1 || !ok {
				distances[currentPosition] = dist + 1
			}
		case SOUTH:
			doors[currentPosition][SOUTH] = true
			dist := distances[currentPosition]
			currentPosition = aoc2018.Coord{X: currentPosition.X, Y: currentPosition.Y + 1}
			if doors[currentPosition] == nil {
				doors[currentPosition] = map[rune]bool{}
			}
			doors[currentPosition][NORTH] = true
			if val, ok := distances[currentPosition]; val > dist+1 || !ok {
				distances[currentPosition] = dist + 1
			}
		case EAST:
			doors[currentPosition][EAST] = true
			dist := distances[currentPosition]
			currentPosition = aoc2018.Coord{X: currentPosition.X + 1, Y: currentPosition.Y}
			if doors[currentPosition] == nil {
				doors[currentPosition] = map[rune]bool{}
			}
			doors[currentPosition][WEST] = true
			if val, ok := distances[currentPosition]; val > dist+1 || !ok {
				distances[currentPosition] = dist + 1
			}
		case WEST:
			doors[currentPosition][WEST] = true
			dist := distances[currentPosition]
			currentPosition = aoc2018.Coord{X: currentPosition.X - 1, Y: currentPosition.Y}
			if doors[currentPosition] == nil {
				doors[currentPosition] = map[rune]bool{}
			}
			doors[currentPosition][EAST] = true
			if val, ok := distances[currentPosition]; val > dist+1 || !ok {
				distances[currentPosition] = dist + 1
			}
		case '(':
			branchSources = append(branchSources, currentPosition)
		case '|':
			currentPosition = branchSources[len(branchSources)-1]
		case ')':
			currentPosition = branchSources[len(branchSources)-1]
			branchSources = branchSources[:len(branchSources)-1]
		}
	}
	return distances
}

func part1(input string) int {
	distances := mapDistances(input)
	maxDistance := 0

	for _, d := range distances {
		if d > maxDistance {
			maxDistance = d
		}
	}

	return maxDistance
}

func part2(input string) int {
	distances := mapDistances(input)
	count := 0

	for _, d := range distances {
		if d >= 1000 {
			count += 1
		}
	}

	return count
}

func main() {
	f, _ := os.Open("input/day20.txt")
	defer f.Close()

	r := bufio.NewReader(f)
	input, _ := r.ReadString('\n')

	fmt.Println(part1(input))
	fmt.Println(part2(input))
}
