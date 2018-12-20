package main

import (
	"fmt"
	"github.com/glenbolake/aoc2018"
	"time"
)

const SERIAL = 9005

func sumAreaTable(serial int) map[aoc2018.Coord]int {
	table := map[aoc2018.Coord]int{}
	for x := 1; x <= 300; x++ {
		rack := x + 10
		for y := 1; y <= 300; y++ {
			table[aoc2018.Coord{X: x, Y: y}] = (rack*y+serial)*rack/100%10 - 5 +
				table[aoc2018.Coord{X: x - 1, Y: y}] + table[aoc2018.Coord{X: x, Y: y - 1}] -
				table[aoc2018.Coord{X: x - 1, Y: y - 1}]
		}
	}
	return table
}

func getSquare(areaTable map[aoc2018.Coord]int, x, y, size int) int {
	return areaTable[aoc2018.Coord{X: x - 1, Y: y - 1}] + areaTable[aoc2018.Coord{X: x + size - 1, Y: y + size - 1}] -
		areaTable[aoc2018.Coord{X: x - 1, Y: y + size - 1}] - areaTable[aoc2018.Coord{X: x + size - 1, Y: y - 1}]
}

func power(serial int) map[aoc2018.Coord]int {
	powers := map[aoc2018.Coord]int{}
	for x := 1; x <= 300; x++ {
		rack := x + 10
		for y := 1; y <= 300; y++ {
			power := (rack*y + serial) * rack
			power = power / 100 % 10
			powers[aoc2018.Coord{X: x, Y: y}] = power - 5
		}
	}
	return powers
}

type Square struct {
	X, Y, Size int
}

func (s Square) String() string {
	return fmt.Sprintf("%d,%d,%d", s.X, s.Y, s.Size)
}

var powerCache map[Square]int

func resetCache() {
	powerCache = map[Square]int{}
}

func calcGridPower(grid map[aoc2018.Coord]int, square Square) int {
	power, ok := powerCache[square]
	if !ok {
		if square.Size == 1 {
			power = grid[aoc2018.Coord{X: square.X, Y: square.Y}]
		} else if square.Size < 1 {
			return 0
		} else {
			power = calcGridPower(grid, Square{X: square.X, Y: square.Y, Size: square.Size - 1}) +
				calcGridPower(grid, Square{X: square.X + 1, Y: square.Y + 1, Size: square.Size - 1}) +
				calcGridPower(grid, Square{X: square.X + square.Size - 1, Y: square.Y, Size: 1}) +
				calcGridPower(grid, Square{X: square.X, Y: square.Y + square.Size - 1, Size: 1}) -
				calcGridPower(grid, Square{X: square.X + 1, Y: square.Y + 1, Size: square.Size - 2})
		}
		powerCache[square] = power
	}
	return power
}

func part1(serial int) aoc2018.Coord {
	grid := power(serial)

	var bestPos aoc2018.Coord
	bestValue := 0

	for x := 1; x < 299; x++ {
		for y := 1; y < 299; y++ {
			value := calcGridPower(grid, Square{X: x, Y: y, Size: 3})
			if value > bestValue {
				bestPos = aoc2018.Coord{X: x, Y: y}
				bestValue = value
			}
		}
	}

	return bestPos
}

func part2(serial int) Square {
	grid := power(serial)

	var bestSquare Square
	bestValue := 0

	for size := 1; size <= 300; size++ {
		for x := 1; x < 301-size; x++ {
			for y := 1; y < 301-size; y++ {
				square := Square{X: x, Y: y, Size: size}
				value := calcGridPower(grid, square)
				if value > bestValue {
					bestSquare = square
					bestValue = value
				}
			}
		}
	}
	return bestSquare
}

func faster(serial int) (string, string) {
	table := sumAreaTable(serial)

	var bestX1, bestY1, bestValue1 int
	for x := 1; x <= 298; x++ {
		for y := 1; y <= 298; y++ {
			value := getSquare(table, x, y, 3)
			if value > bestValue1 {
				bestX1, bestY1, bestValue1 = x, y, value
			}
		}
	}

	var bestX2, bestY2, bestSize, bestValue2 int
	for size := 1; size <= 300; size++ {
		for x := 1; x <= 301-size; x++ {
			for y := 1; y <= 301-size; y++ {
				value := getSquare(table, x, y, size)
				if value > bestValue2 {
					bestX2, bestY2, bestSize, bestValue2 = x, y, size, value
				}
			}
		}
	}
	return fmt.Sprintf("%d,%d", bestX1, bestY1), fmt.Sprintf("%d,%d,%d", bestX2, bestY2, bestSize)
}

func main() {
	start := time.Now()
	resetCache()
	fmt.Println(part1(SERIAL))
	fmt.Println(part2(SERIAL))
	fmt.Println("Method 1:", time.Now().Sub(start))

	// method 2
	start = time.Now()
	p1, p2 := faster(SERIAL)
	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println("Method 2:", time.Now().Sub(start))
}
