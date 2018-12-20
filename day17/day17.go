package main

import (
	"bufio"
	"fmt"
	"github.com/glenbolake/aoc2018"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	SAND    = '.'
	CLAY    = '#'
	FLOWING = '|'
	STILL   = '~'
)

func Min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func Max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

type Scan [][]rune

func (scan Scan) String() string {
	xMin := len(scan[0])
	for _, s := range scan {
		x := strings.Index(string(s), string(CLAY))
		if x == -1 {
			continue
		}
		if x < xMin {
			xMin = x
		}
	}
	xMin -= 2
	ss := make([]string, 0, len(scan))
	for i := range scan {
		ss = append(ss, fmt.Sprintf("%4d ", i)+string(scan[i][xMin:]))
	}
	return strings.Join(ss, "\n")
}

func isSolid(r rune) bool {
	switch r {
	case CLAY, STILL:
		return true
	case SAND, FLOWING:
		return false
	default:
		fmt.Println("Bad value:", string(r))
		return false
	}
}

func flowWater(clay map[aoc2018.Coord]bool) Scan {
	xMax := 0
	yMin := math.MaxInt32
	yMax := 0
	for c := range clay {
		if c.X > xMax {
			xMax = c.X
		}
		if c.Y < yMin {
			yMin = c.Y
		}
		if c.Y > yMax {
			yMax = c.Y
		}
	}
	scan := make(Scan, yMax+1)
	for row := range scan {
		scan[row] = make([]rune, xMax+2)
		for col := range scan[row] {
			if clay[aoc2018.Coord{Y: row, X: col}] {
				scan[row][col] = CLAY
			} else {
				scan[row][col] = SAND
			}
		}
	}
	sources := []aoc2018.Coord{{X: 500}}
	scan[0][500] = '+'

	for len(sources) > 0 {
		var newSources []aoc2018.Coord
		for _, source := range sources {
			row := source.Y
			col := source.X
			if row == yMax {
				continue
			}
			switch scan[row+1][col] {
			case SAND:
				scan[row+1][col] = FLOWING
				newSources = append(newSources, aoc2018.Coord{Y: row + 1, X: col})
			case CLAY, STILL:
				// See how far the water will flow left and right
				left := col
				right := col
				for isSolid(scan[row+1][left]) && !isSolid(scan[row][left]) {
					left -= 1
				}
				for isSolid(scan[row+1][right]) && !isSolid(scan[row][right]) {
					right += 1
				}
				if isSolid(scan[row][left]) && isSolid(scan[row][right]) {
					// If there are walls on both sides, it's a well and should be filled with still water.
					// The spot flowing from above becomes a new source.
					for i := left + 1; i < right; i++ {
						scan[row][i] = STILL
					}
					newSources = append(newSources, aoc2018.Coord{Y: row - 1, X: col})
				} else {
					// Otherwise, it will be filled with flowing water in both directions. The side(s) with
					// no wall will become sources that will flow down next tick.
					for i := left + 1; i < right; i++ {
						scan[row][i] = FLOWING
					}
					if !isSolid(scan[row][left]) {
						scan[row][left] = FLOWING
						newSources = append(newSources, aoc2018.Coord{Y: row, X: left})
					}
					if !isSolid(scan[row][right]) {
						scan[row][right] = FLOWING
						newSources = append(newSources, aoc2018.Coord{Y: row, X: right})
					}
				}
			}
		}
		sources = newSources
	}

	return scan
}

func part1(clay map[aoc2018.Coord]bool) int {
	yMin := math.MaxInt32
	yMax := 0
	for c := range clay {
		if c.Y < yMin {
			yMin = c.Y
		}
		if c.Y > yMax {
			yMax = c.Y
		}
	}
	scan := flowWater(clay)
	total := 0
	for y := yMin; y <= yMax; y++ {
		total += strings.Count(string(scan[y]), string(FLOWING))
		total += strings.Count(string(scan[y]), string(STILL))
	}
	return total
}

func part2(clay map[aoc2018.Coord]bool) int {
	yMin := math.MaxInt32
	yMax := 0
	for c := range clay {
		if c.Y < yMin {
			yMin = c.Y
		}
		if c.Y > yMax {
			yMax = c.Y
		}
	}
	scan := flowWater(clay)
	total := 0
	for y := yMin; y <= yMax; y++ {
		total += strings.Count(string(scan[y]), string(STILL))
	}
	return total
}

func main() {
	f, _ := os.Open("input/day17.txt")

	scanner := bufio.NewScanner(bufio.NewReader(f))
	clay := map[aoc2018.Coord]bool{}
	var xMin, xMax, yMin, yMax int
	xMin = math.MaxInt32
	yMin = math.MaxInt32
	regex, _ := regexp.Compile(`([xy])=(\d+), [xy]=(\d+)\.\.(\d+)`)
	for scanner.Scan() {
		result := regex.FindStringSubmatch(scanner.Text())
		single, _ := strconv.Atoi(result[2])
		min, _ := strconv.Atoi(result[3])
		max, _ := strconv.Atoi(result[4])
		if result[1] == "x" {
			for y := min; y <= max; y++ {
				clay[aoc2018.Coord{Y: y, X: single}] = true
			}
			xMin = Min(xMin, single)
			xMax = Max(xMax, single)
			yMin = Min(yMin, min)
			yMax = Max(yMax, max)
		} else {
			for x := min; x <= max; x++ {
				clay[aoc2018.Coord{Y: single, X: x}] = true
			}
			xMin = Min(xMin, min)
			xMax = Max(xMax, max)
			yMin = Min(yMin, single)
			yMax = Max(yMax, single)
		}
	}

	fmt.Println(part1(clay))
	fmt.Println(part2(clay))
}
