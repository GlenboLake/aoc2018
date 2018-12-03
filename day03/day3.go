package main

import (
	"bufio"
	"fmt"
	"os"
)

type Claim struct {
	Id            int
	Left, Top     int
	Width, Height int
}

func part1(claims []Claim) int {
	fabric := make(map[[2]int]int)

	for _, c := range claims {
		for i := c.Left; i < c.Left+c.Width; i++ {
			for j := c.Top; j < c.Top+c.Height; j++ {
				fabric[[2]int{i, j}] += 1
			}
		}
	}
	overlap := 0
	for _, count := range fabric {
		if count >= 2 {
			overlap += 1
		}
	}
	return overlap
}

func part2(claims []Claim) int {
	claimsOk := make(map[int]bool)
	for _, c := range claims {
		claimsOk[c.Id] = true
	}
	fabric := make(map[[2]int][]int)

	for _, c := range claims {
		for i := c.Left; i < c.Left+c.Width; i++ {
			for j := c.Top; j < c.Top+c.Height; j++ {
				fabric[[2]int{i, j}] = append(fabric[[2]int{i, j}], c.Id)
			}
		}
	}

	for _, claimList := range fabric {
		if len(claimList) > 1 {
			for _, claim := range claimList {
				delete(claimsOk, claim)
			}
		}
	}
	for k := range claimsOk {
		return k
	}
	return 0
}

func parseClaim(claim string) Claim {
	var c Claim
	fmt.Sscanf(claim, "#%d @ %d,%d: %dx%d", &c.Id, &c.Left, &c.Top, &c.Width, &c.Height)
	return c
}

func main() {
	f, _ := os.Open("input/day03.txt")

	scanner := bufio.NewScanner(bufio.NewReader(f))
	var claims []Claim
	for scanner.Scan() {
		claims = append(claims, parseClaim(scanner.Text()))
	}

	fmt.Println(part1(claims))
	fmt.Println(part2(claims))
}
