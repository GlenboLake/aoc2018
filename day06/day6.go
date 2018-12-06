package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type Point struct {
	X, Y int
}

var NoPoint Point

func abs(val int) int {
	if val < 0 {
		return -val
	} else {
		return val
	}
}

func manhattan(a, b Point) int {
	return abs(a.X-b.X) + abs(a.Y-b.Y)
}

func findClosest(p Point, list []Point) Point {
	var closest []Point
	closestDist := math.MaxInt32
	closestCount := 0
	for _, point := range list {
		dist := manhattan(p, point)
		if dist < closestDist {
			closest = []Point{point}
			closestDist = dist
			closestCount = 1
		} else if dist == closestDist {
			closestCount += 1
			closest = append(closest, point)
		}
	}
	if len(closest) == 1 {
		return closest[0]
	} else {
		return NoPoint
	}
}

func part1(input []Point) int {
	minX := input[0].X
	maxX := input[0].X
	minY := input[0].Y
	maxY := input[0].Y

	for _, p := range input[1:] {
		if p.X < minX {
			minX = p.X
		} else if p.X > maxX {
			maxX = p.X
		}
		if p.Y < minY {
			minY = p.Y
		} else if p.Y > maxY {
			maxY = p.Y
		}
	}
	left := minX - 100
	right := maxX + 100
	top := minY - 100
	bottom := maxY + 100

	infinite := map[Point]bool{}

	var c Point
	for x := left; x <= right; x++ {
		c = findClosest(Point{x, top}, input)
		if c != NoPoint {
			infinite[c] = true
		}
		c = findClosest(Point{x, bottom}, input)
		if c != NoPoint {
			infinite[c] = true
		}
	}
	for y := top; y <= bottom; y++ {
		c = findClosest(Point{left, y}, input)
		if c != NoPoint {
			infinite[c] = true
		}
		c = findClosest(Point{right, y}, input)
		if c != NoPoint {
			infinite[c] = true
		}
	}
	areas := map[Point]int{}
	for _, p := range input {
		if !infinite[p] {
			areas[p] = 0
		}
	}
	for x := left; x < right; x++ {
		for y := top; y < bottom; y++ {
			c = findClosest(Point{x, y}, input)
			if c != NoPoint && infinite[c] == false {
				areas[c] += 1
			}
		}
	}

	biggestArea := 0
	for _, area := range areas {
		if area > biggestArea {
			biggestArea = area
		}
	}

	return biggestArea
}

func part2(input []Point, threshold int) int {
	minX := input[0].X
	maxX := input[0].X
	minY := input[0].Y
	maxY := input[0].Y

	for _, p := range input[1:] {
		if p.X < minX {
			minX = p.X
		} else if p.X > maxX {
			maxX = p.X
		}
		if p.Y < minY {
			minY = p.Y
		} else if p.Y > maxY {
			maxY = p.Y
		}
	}

	area := 0
	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			p := Point{x, y}
			total := 0
			for _, i := range input {
				total += manhattan(p, i)
			}
			if total < threshold {
				area += 1
			}
		}
	}
	return area
}

func main() {
	f, _ := os.Open("input/day06.txt")

	var input []Point
	scanner := bufio.NewScanner(bufio.NewReader(f))
	for scanner.Scan() {
		var p Point
		fmt.Sscanf(scanner.Text(), "%d, %d", &p.X, &p.Y)
		input = append(input, p)
	}

	sample := []Point{
		{1, 1},
		{1, 6},
		{8, 3},
		{3, 4},
		{5, 5},
		{8, 9},
	}

	fmt.Println(part1(input))
	fmt.Println(16, part2(sample, 32))
	fmt.Println(part2(input, 10000))
}
