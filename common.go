package aoc2018

import "fmt"

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type Coord struct {
	X, Y int
}

func (c Coord) DistTo(other Coord) int {
	return Abs(c.X-other.X) + Abs(c.Y-other.Y)
}

func (c Coord) String() string {
	return fmt.Sprintf("%d,%d", c.X, c.Y)
}
