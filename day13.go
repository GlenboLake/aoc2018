package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

const (
	UP       = iota
	RIGHT    = iota
	DOWN     = iota
	LEFT     = iota
	STRAIGHT = iota
)

type Cart struct {
	x, y     int
	dir      int
	nextTurn int
}

func (c Cart) String() string {
	return fmt.Sprintf("%s@%d,%d", dirName(c.dir), c.x, c.y)
}

func (c *Cart) TurnLeft() {
	c.dir = (c.dir + 3) % 4
}

func (c *Cart) TurnRight() {
	c.dir = (c.dir + 1) % 4
}

func (c *Cart) Turn() {
	switch c.nextTurn {
	case LEFT:
		c.dir = (c.dir + 3) % 4
		c.nextTurn = STRAIGHT
	case STRAIGHT:
		c.nextTurn = RIGHT
	case RIGHT:
		c.dir = (c.dir + 1) % 4
		c.nextTurn = LEFT
	}
}

func (c *Cart) Move() {
	switch c.dir {
	case UP:
		c.y -= 1
	case DOWN:
		c.y += 1
	case LEFT:
		c.x -= 1
	case RIGHT:
		c.x += 1
	}
}

func dirName(dir int) string {
	switch dir {
	case UP:
		return "^"
	case DOWN:
		return "v"
	case LEFT:
		return "<"
	case RIGHT:
		return ">"
	default:
		return "?"
	}
}

func dirFromChar(char rune) int {
	switch char {
	case '^':
		return UP
	case 'v':
		return DOWN
	case '<':
		return LEFT
	case '>':
		return RIGHT
	default:
		return -1
	}
}

func tick(cart *Cart, track []string) {
	cart.Move()
	switch track[cart.y][cart.x] {
	case '+':
		cart.Turn()
	case '/':
		switch cart.dir {
		case UP, DOWN:
			cart.TurnRight()
		case LEFT, RIGHT:
			cart.TurnLeft()
		}
	case '\\':
		switch cart.dir {
		case UP, DOWN:
			cart.TurnLeft()
		case LEFT, RIGHT:
			cart.TurnRight()
		}
	}
}

func parseCarts(track []string) ([]string, []Cart) {
	var carts []Cart
	for i, row := range track {
		for j, char := range row {
			d := dirFromChar(char)
			if d == UP || d == DOWN {
				track[i] = track[i][:j] + "|" + track[i][j+1:]
			} else if d == LEFT || d == RIGHT {
				track[i] = track[i][:j] + "-" + track[i][j+1:]
			} else {
				continue
			}
			carts = append(carts, Cart{j, i, d, LEFT})
		}
	}
	return track, carts
}

func sortCarts(carts []Cart) []Cart {
	sort.Slice(carts, func(i, j int) bool {
		if carts[i].y == carts[j].y {
			return carts[i].x < carts[j].x
		} else {
			return carts[i].y < carts[j].y
		}
	})
	return carts
}

func collision(carts *[]Cart, ignore map[*Cart]bool) bool {
	type Pos struct {
		x, y int
	}
	positions := map[Pos]bool{}
	for i, c := range *carts {
		if ignore[&(*carts)[i]] {
			continue
		}
		p := Pos{c.x, c.y}
		if positions[p] {
			return true
		} else {
			positions[p] = true
		}
	}
	return false
}

func trackWithCarts(track []string, carts []Cart) string {
	t := make([]string, len(track))
	copy(t, track)
	for _, c := range carts {
		t[c.y] = t[c.y][:c.x] + dirName(c.dir) + t[c.y][c.x+1:]
	}
	return strings.Join(t, "\n")
}

func part1(track []string, carts []Cart) string {

	for {
		for i := range sortCarts(carts) {
			tick(&carts[i], track)
			if collision(&carts, nil) {
				return fmt.Sprintf("%d,%d", carts[i].x, carts[i].y)
			}
		}
	}
}

func part2(track []string, carts []Cart) string {

	//for _, r := range track {
	//	fmt.Println(r)
	//}
	//fmt.Println(carts)

	for len(carts) > 1 {
		fmt.Printf("Looking at %d carts\n", len(carts))
		removed := map[*Cart]bool{}
		for i := range sortCarts(carts) {
			if removed[&carts[i]] {
				continue
			}
			tick(&carts[i], track)
			if collision(&carts, removed) {
				x, y := carts[i].x, carts[i].y
				fmt.Println("Collision at", x, y)
				found := 0
				newCarts := make([]Cart, len(carts)-2)
				for j := 0; j < len(carts); j++ {
					if carts[j].x == x && carts[j].y == y {
						found += 1
						removed[&carts[j]] = true
					} else {
						newCarts = append(newCarts, carts[j])
					}
				}
				fmt.Println("Found", found)
			}
		}
		if len(removed) > 0 {
			fmt.Printf("Removing %d carts\n", len(removed))
			remaining := make([]Cart, 0, len(carts)-len(removed))
			for i, c := range carts {
				if !removed[&carts[i]] {
					remaining = append(remaining, c)
				}
			}
			carts = remaining
		}
	}
	return fmt.Sprintf("%d,%d", carts[0].x, carts[0].y)
}

func main() {
	f, _ := os.Open("input/day13.txt")

	scanner := bufio.NewScanner(bufio.NewReader(f))
	var input []string
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	track, carts := parseCarts(input)
	p1Carts := make([]Cart, len(carts))
	copy(p1Carts, carts)
	fmt.Println(part1(track, p1Carts))

	//input = []string{
	//	`/>-<\  `,
	//	`|   |  `,
	//	`| /<+-\`,
	//	`| | | v`,
	//	`\>+</ |`,
	//	`  |   ^`,
	//	`  \<->/`,
	//}

	fmt.Println(part2(track, carts))
}
