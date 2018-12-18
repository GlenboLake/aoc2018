package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	OPEN       rune = '.'
	TREES      rune = '|'
	LUMBERYARD rune = '#'
)

type CollectionArea []string

func (c CollectionArea) String() string {
	return strings.Join(c, "\n")
}

func (c CollectionArea) GetRow(x int) string {
	if x < 0 || x >= len(c) {
		s := ""
		for range c[0] {
			s += "."
		}
		return s
	} else {
		return c[x]
	}
}

func (c CollectionArea) ResourceValue() int {
	trees := 0
	lumberyards := 0
	for _, row := range c {
		trees += strings.Count(row, string(TREES))
		lumberyards += strings.Count(row, string(LUMBERYARD))
	}
	return trees * lumberyards
}

type caHelper struct {
	rowNum int
	result string
}

func (c *CollectionArea) Tick() {
	newData := make(chan caHelper)
	for i, row := range *c {
		go tickRow(row, c.GetRow(i-1), c.GetRow(i+1), i, newData)
	}
	for i := 0; i < len(*c); i++ {
		d := <-newData
		(*c)[d.rowNum] = d.result
	}
}

func tickRow(row, above, below string, rowNum int, results chan<- caHelper) {
	var result []rune
	for i := 0; i < len(row); i++ {
		value := rune(row[i])
		near := map[rune]int{}
		if i > 0 {
			near[rune(row[i-1])] += 1
			near[rune(above[i-1])] += 1
			near[rune(below[i-1])] += 1
		}
		near[rune(above[i])] += 1
		near[rune(below[i])] += 1
		if i < len(row)-1 {
			near[rune(row[i+1])] += 1
			near[rune(above[i+1])] += 1
			near[rune(below[i+1])] += 1
		}
		switch value {
		case OPEN:
			if near[TREES] >= 3 {
				result = append(result, TREES)
			} else {
				result = append(result, OPEN)
			}
		case TREES:
			if near[LUMBERYARD] >= 3 {
				result = append(result, LUMBERYARD)
			} else {
				result = append(result, TREES)
			}
		case LUMBERYARD:
			if near[LUMBERYARD] > 0 && near[TREES] > 0 {
				result = append(result, LUMBERYARD)
			} else {
				result = append(result, OPEN)
			}
		}
	}
	results <- caHelper{rowNum, string(result)}
}

func part1(area CollectionArea) int {
	myCopy := make(CollectionArea, len(area))
	copy(myCopy, area)
	for i := 0; i < 10; i++ {
		myCopy.Tick()
	}
	return myCopy.ResourceValue()
}

func stringsIndexOf(s []string, item string) int {
	for i, str := range s {
		if str == item {
			return i
		}
	}
	return -1
}

func part2(area CollectionArea) int {
	history := []string{""}
	limit := 1000000000
	for i := 1; i <= limit; i++ {
		area.Tick()
		rep := strings.Join(area, "")
		index := stringsIndexOf(history, rep)
		if index == -1 {
			history = append(history, rep)
		} else {
			loopLength := i - index
			answerIndex := limit % loopLength
			for answerIndex < index {
				answerIndex += loopLength
			}
			answer := strings.Count(history[answerIndex], string(TREES)) * strings.Count(history[answerIndex], string(LUMBERYARD))
			return answer
		}
	}
	return area.ResourceValue()
}

func main() {
	f, _ := os.Open("input/day18.txt")

	var input CollectionArea
	scanner := bufio.NewScanner(bufio.NewReader(f))
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	fmt.Println(part1(input))
	fmt.Println(part2(input))
}
