package main

import "fmt"

type Coordinate struct {
	X, Y int
}

func (c Coordinate) String() string {
	return fmt.Sprintf("%d,%d", c.X, c.Y)
}

const SERIAL = 9005

func power(serial int) map[Coordinate]int {
	powers := map[Coordinate]int{}
	for x := 1; x <= 300; x++ {
		rack := x + 10
		for y := 1; y <= 300; y++ {
			power := (rack*y + serial) * rack
			power = power / 100 % 10
			powers[Coordinate{x, y}] = power - 5
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

func calcGridPower(grid map[Coordinate]int, square Square) int {
	power, ok := powerCache[square]
	if !ok {
		if square.Size == 1 {
			power = grid[Coordinate{square.X, square.Y}]
		} else if square.Size < 1 {
			return 0
		} else {
			power = calcGridPower(grid, Square{square.X, square.Y, square.Size - 1}) +
				calcGridPower(grid, Square{square.X + 1, square.Y + 1, square.Size - 1}) +
				calcGridPower(grid, Square{square.X + square.Size - 1, square.Y, 1}) +
				calcGridPower(grid, Square{square.X, square.Y + square.Size - 1, 1}) -
				calcGridPower(grid, Square{square.X + 1, square.Y + 1, square.Size - 2})
		}
		powerCache[square] = power
	}
	return power
}

func part1(serial int) Coordinate {
	grid := power(serial)

	var bestPos Coordinate
	bestValue := 0

	for x := 1; x < 299; x++ {
		for y := 1; y < 299; y++ {
			value := calcGridPower(grid, Square{x, y, 3})
			if value > bestValue {
				bestPos = Coordinate{x, y}
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
				square := Square{x, y, size}
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

func main() {
	resetCache()
	fmt.Println(part1(SERIAL))
	fmt.Println(part2(SERIAL))
}
