package main

import (
	"bufio"
	"fmt"
	"github.com/glenbolake/aoc2018"
	"os"
	"strings"
)

type Point4D struct {
	a, b, c, d int
}

func (p *Point4D) DistTo(other Point4D) int {
	return aoc2018.Abs(p.a-other.a) + aoc2018.Abs(p.b-other.b) + aoc2018.Abs(p.c-other.c) + aoc2018.Abs(p.d-other.d)
}

type Constellation struct {
	stars map[Point4D]struct{}
}

func NewConstellation() Constellation {
	c := Constellation{}
	c.Clear()
	return c
}

func (c *Constellation) Add(star Point4D) {
	c.stars[star] = struct{}{}
}

func (c *Constellation) Touches(p Point4D) bool {
	for star := range c.stars {
		if star.DistTo(p) <= 3 {
			return true
		}
	}
	return false
}
func (c *Constellation) Clear() {
	c.stars = map[Point4D]struct{}{}
}

func part1(stars []Point4D) int {
	var currentConstellation := NewConstellation()
	remainingStars := map[Point4D]struct{}{}
	for _, s := range stars {
		remainingStars[s] = struct{}{}
	}

	numConstellations := 0
	for len(remainingStars) > 0 {
		if len(currentConstellation.stars) == 0 {
			numConstellations ++
			for s := range remainingStars {
				currentConstellation.Add(s)
				delete(remainingStars, s)
				break
			}
		}
		var added []Point4D
		for s := range remainingStars {
			if currentConstellation.Touches(s) {
				currentConstellation.Add(s)
				added = append(added, s)
			}
		}
		for _, s := range added {
			delete(remainingStars, s)
		}
		if len(added) == 0 {
			currentConstellation.Clear()
		}
	}
	return numConstellations
}

func main() {
	f, _ := os.Open("input/day25.txt")

	scanner := bufio.NewScanner(bufio.NewReader(f))
	var stars []Point4D
	for scanner.Scan() {
		star := Point4D{}
		fmt.Sscanf(strings.TrimSpace(scanner.Text()), "%d,%d,%d,%d", &star.a, &star.b, &star.c, &star.d)
		stars = append(stars, star)
	}

	fmt.Println(part1(stars))
}
