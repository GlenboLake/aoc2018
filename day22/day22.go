package main

import (
	"fmt"
	"github.com/glenbolake/aoc2018"
	"math"
)

const (
	depth   = 4845
	targetX = 6
	targetY = 770
)

const (
	rocky  = 0
	wet    = 1
	narrow = 2
)

var erosionMemo = map[aoc2018.Coord]int{}

func calcErosion(x, y int) int {
	input := aoc2018.Coord{X: x, Y: y}
	value, seen := erosionMemo[input]
	if seen {
		return value
	}
	var geoIndex int
	if (x == 0 && y == 0) || (x == targetX && y == targetY) {
		geoIndex = 0
	} else if x != 0 && y != 0 {
		geoIndex = calcErosion(x-1, y) * calcErosion(x, y-1)
	} else if x == 0 {
		geoIndex = 48271 * y
	} else { // y==0
		geoIndex = 16807 * x
	}
	erosion := (geoIndex + depth) % 20183
	erosionMemo[input] = erosion
	return erosion
}

func terrain(x, y int) int {
	return calcErosion(x, y) % 3
}

func part1() int {
	risk := 0
	for y := 0; y <= targetY; y++ {
		for x := 0; x <= targetX; x++ {
			terrain := terrain(x, y)
			risk += terrain
		}
	}
	return risk - terrain(targetX, targetY)
}

const (
	nothing   = iota
	torch     = iota
	climbGear = iota
)

var tools = []int{nothing, torch, climbGear}

func toolValid(x, y, tool int) bool {
	switch terrain(x, y) {
	case rocky:
		return tool == torch || tool == climbGear
	case wet:
		return tool == climbGear || tool == nothing
	case narrow:
		return tool == torch || tool == nothing
	}
	return false
}

type State struct {
	tool int
	x, y int
}

func validNext(state State) map[State]int {
	tool, x, y := state.tool, state.x, state.y
	rv := map[State]int{}
	for _, t := range tools {
		if t != tool && toolValid(x, y, t) {
			rv[State{tool: t, x: x, y: y}] = 7
		}
	}
	if x > 0 && toolValid(x-1, y, tool) {
		rv[State{tool: tool, x: x - 1, y: y}] = 1
	}
	if toolValid(x+1, y, tool) {
		rv[State{tool: tool, x: x + 1, y: y}] = 1
	}
	if y > 0 && toolValid(x, y-1, tool) {
		rv[State{tool: tool, x: x, y: y - 1}] = 1
	}
	if toolValid(x, y+1, tool) {
		rv[State{tool: tool, x: x, y: y + 1}] = 1
	}
	return rv
}

func checkDone(newTimes map[State]int, bestTime int) bool {
	for _, t := range newTimes {
		if t < bestTime {
			return false
		}
	}
	return true
}

func part2() int {
	goal := State{tool: torch, x: targetX, y: targetY}
	times := map[State]int{{tool: torch}: 0, goal: math.MaxInt32}
	newTimes := map[State]int{{tool: torch}: 0}

	for !checkDone(newTimes, times[goal]) {
		nextNewTimes := map[State]int{}
		for state, time := range newTimes {
			for newState, addedTime := range validNext(state) {
				if t, ok := times[newState]; !ok || t > time+addedTime {
					times[newState] = time + addedTime
					nextNewTimes[newState] = time + addedTime
				}
			}
		}
		newTimes = nextNewTimes
	}

	return times[goal]
}

func main() {
	fmt.Println(part1())
	fmt.Println(part2())
}
