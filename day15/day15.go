package main

import (
	"bufio"
	"fmt"
	"github.com/glenbolake/aoc2018"
	"math"
	"os"
	"sort"
	"strings"
)

var NOWHERE = aoc2018.Coord{}

type Unit struct {
	Type   string
	Loc    aoc2018.Coord
	Attack int
	HP     int
}

func (u Unit) String() string {
	return fmt.Sprintf("%s%s[%d]", u.Type, u.Loc, u.HP)
}

func NewElf(loc aoc2018.Coord, attackPower int) Unit {
	return Unit{"E", loc, attackPower, 200}
}

func NewGoblin(loc aoc2018.Coord) Unit {
	return Unit{"G", loc, 3, 200}
}

func (u Unit) ChooseTarget(m Map) *Unit {
	var targets []*Unit
	for _, unit := range m.SortedUnits() {
		if u.Type != unit.Type && u.Loc.DistTo(unit.Loc) == 1 {
			targets = append(targets, unit)
		}
	}
	if len(targets) == 0 {
		return nil
	}
	sort.Slice(targets, func(i, j int) bool {
		if targets[i].HP == targets[j].HP {
			if targets[i].Loc.Y == targets[j].Loc.Y {
				return targets[i].Loc.X < targets[j].Loc.X
			} else {
				return targets[i].Loc.Y < targets[j].Loc.Y
			}
		} else {
			return targets[i].HP < targets[j].HP
		}
	})
	return targets[0]
}

type Map struct {
	Units  []*Unit
	Spaces map[aoc2018.Coord]bool
}

func (m Map) String() string {
	var maxX, maxY int
	for s := range m.Spaces {
		if s.X > maxX {
			maxX = s.X
		}
		if s.Y > maxY {
			maxY = s.Y
		}
	}
	width := maxX + 2
	height := maxY + 2
	units := map[aoc2018.Coord]Unit{}
	for _, u := range m.SortedUnits() {
		units[u.Loc] = *u
	}
	s := ""
	for row := 0; row < height; row++ {
		var rowSummary []string
		for col := 0; col < width; col++ {
			coord := aoc2018.Coord{col, row}
			if m.Spaces[coord] {
				if u, ok := units[coord]; ok {
					s += u.Type
					rowSummary = append(rowSummary, fmt.Sprintf("%s(%d)", u.Type, u.HP))
				} else {
					s += "."
				}
			} else {
				s += "#"
			}
		}
		s += fmt.Sprintf("   %s\n", strings.Join(rowSummary, ", "))
	}
	return s
}

func (m Map) CountLivingElves() int {
	count := 0
	for _, u := range m.Units {
		if u.Type == "E" {
			count += 1
		}
	}
	return count
}

func (m *Map) Remove(unit *Unit) {
	for i := 0; i < len(m.Units); i++ {
		if m.Units[i] == unit {
			m.Units = append(m.Units[:i], m.Units[i+1:]...)
		}
	}
}

func (m Map) FreeSpaces() map[aoc2018.Coord]bool {
	freeSpaces := map[aoc2018.Coord]bool{}
	unitSpaces := map[aoc2018.Coord]bool{}
	for _, u := range m.SortedUnits() {
		unitSpaces[u.Loc] = true
	}
	for k, v := range m.Spaces {
		if v && !unitSpaces[k] {
			freeSpaces[k] = true
		}
	}

	return freeSpaces
}

func (m Map) SortedUnits() []*Unit {
	rv := append([]*Unit{}, m.Units...)
	sort.Slice(rv, func(i, j int) bool {
		if rv[i].Loc.Y == rv[j].Loc.Y {
			return rv[i].Loc.X < rv[j].Loc.X
		} else {
			return rv[i].Loc.Y < rv[j].Loc.Y
		}
	})
	return rv
}

func (m Map) TeamsRemaining() map[string]bool {
	teams := map[string]bool{}
	for _, u := range m.Units {
		teams[u.Type] = true
	}
	return teams
}

func (m Map) TotalHP() int {
	total := 0
	for _, u := range m.Units {
		total += u.HP
	}
	return total
}

// If a unit is supposed to move, analyze the map to determine which space it should move to.
func (m Map) FindMovement(unit Unit) aoc2018.Coord {
	// Check for adjacent enemies
	for _, u := range m.SortedUnits() {
		if unit.Type != u.Type && unit.Loc.DistTo(u.Loc) == 1 {
			return NOWHERE
		}
	}
	// Build an A* map
	distanceFromUnit := m.Travel(unit.Loc)
	var adjacent []aoc2018.Coord
	for coord, dist := range distanceFromUnit {
		if dist == 1 {
			adjacent = append(adjacent, coord)
		}
	}
	sort.Slice(adjacent, func(i, j int) bool {
		if adjacent[i].Y == adjacent[j].Y {
			return adjacent[i].X < adjacent[j].X
		}
		return adjacent[i].Y < adjacent[j].Y
	})
	// Find close enemies
	nearEnemies := m.SpacesNearEnemies(unit)

	// Figure out the target square to walk towards
	var targets []aoc2018.Coord
	distance := math.MaxInt32
	for space := range nearEnemies {
		dist, ok := distanceFromUnit[space]
		if !ok {
			// Space must be reachable to count
			continue
		}
		if dist < distance {
			distance = dist
			targets = []aoc2018.Coord{space}
		} else if dist == distance {
			targets = append(targets, space)
		}
	}
	if len(targets) == 0 {
		return aoc2018.Coord{}
	}
	sort.Slice(targets, func(i, j int) bool {
		if targets[i].Y == targets[j].Y {
			return targets[i].X < targets[j].X
		}
		return targets[i].Y < targets[j].Y
	})
	target := targets[0]

	// Find out which adjacent square is closest to the target
	aStar := m.Travel(target)
	distToTarget := math.MaxInt32
	rv := NOWHERE
	for _, a := range adjacent {
		dist, ok := aStar[a]
		if ok && dist < distToTarget {
			distToTarget = dist
			rv = a
		}
	}

	return rv
}

func (m Map) Travel(coord aoc2018.Coord) map[aoc2018.Coord]int {
	freeSpaces := m.FreeSpaces()
	rv := map[aoc2018.Coord]int{coord: 0}
	lastGen := map[aoc2018.Coord]bool{coord: true}
	value := 0
	for len(lastGen) > 0 {
		value += 1
		nextGen := map[aoc2018.Coord]bool{}
		for coord := range lastGen {
			spaces := []aoc2018.Coord{
				{coord.X, coord.Y - 1},
				{coord.X, coord.Y + 1},
				{coord.X - 1, coord.Y},
				{coord.X + 1, coord.Y}}
			for _, s := range spaces {
				if _, seen := rv[s]; seen {
					continue
				}
				if !freeSpaces[s] {
					continue
				}
				nextGen[s] = true
				rv[s] = value
			}
		}
		lastGen = nextGen
	}
	return rv
}

func (m Map) SpacesNearEnemies(unit Unit) map[aoc2018.Coord]bool {
	var enemyTeam string
	if unit.Type == "E" {
		enemyTeam = "G"
	} else if unit.Type == "G" {
		enemyTeam = "E"
	} else {
		panic("BAD UNIT TYPE")
	}

	freeSpaces := m.FreeSpaces()
	units := m.SortedUnits()
	rv := map[aoc2018.Coord]bool{}
	for _, enemy := range units {
		if enemy.Type != enemyTeam {
			continue
		}
		if freeSpaces[aoc2018.Coord{enemy.Loc.X, enemy.Loc.Y - 1}] {
			rv[aoc2018.Coord{enemy.Loc.X, enemy.Loc.Y - 1}] = true
		}
		if freeSpaces[aoc2018.Coord{enemy.Loc.X, enemy.Loc.Y + 1}] {
			rv[aoc2018.Coord{enemy.Loc.X, enemy.Loc.Y + 1}] = true
		}
		if freeSpaces[aoc2018.Coord{enemy.Loc.X - 1, enemy.Loc.Y}] {
			rv[aoc2018.Coord{enemy.Loc.X - 1, enemy.Loc.Y}] = true
		}
		if freeSpaces[aoc2018.Coord{enemy.Loc.X + 1, enemy.Loc.Y}] {
			rv[aoc2018.Coord{enemy.Loc.X + 1, enemy.Loc.Y}] = true
		}
	}
	return rv
}

func simulateBattle(battlefield *Map) int {
	round := 0
	for {
		round += 1
		for _, u := range battlefield.SortedUnits() {
			if u.HP <= 0 {
				continue
			}
			if len(battlefield.TeamsRemaining()) == 1 {
				return round - 1
			}
			moveTo := battlefield.FindMovement(*u)
			if moveTo != NOWHERE {
				u.Loc = moveTo
			}
			attackTarget := u.ChooseTarget(*battlefield)
			if attackTarget != nil {
				attackTarget.HP -= u.Attack
				if attackTarget.HP <= 0 {
					battlefield.Remove(attackTarget)
				}
			}
		}
	}
}

func part1(input []string) int {
	battlefield := parseMap(input, 3)
	rounds := simulateBattle(&battlefield)
	return rounds * battlefield.TotalHP()
}

func part2(input []string) int {
	min := 1
	max := 200
	elfCount := parseMap(input, 0).CountLivingElves()

	for max-min > 1 {
		test := (min + max) / 2
		battlefield := parseMap(input, test)
		simulateBattle(&battlefield)
		if battlefield.CountLivingElves() < elfCount {
			min = test
		} else {
			max = test
		}
	}
	battlefield := parseMap(input, max)
	rounds := simulateBattle(&battlefield)
	return rounds * battlefield.TotalHP()
}

func parseMap(input []string, elfAttack int) Map {
	dungeon := Map{}
	dungeon.Spaces = map[aoc2018.Coord]bool{}
	for y, line := range input {
		for x, ch := range line {
			if ch != '#' {
				dungeon.Spaces[aoc2018.Coord{x, y}] = true
			}
			if ch == 'E' {
				elf := NewElf(aoc2018.Coord{x, y}, elfAttack)
				dungeon.Units = append(dungeon.Units, &elf)
			} else if ch == 'G' {
				goblin := NewGoblin(aoc2018.Coord{x, y})
				dungeon.Units = append(dungeon.Units, &goblin)
			}
		}
	}
	return dungeon
}

func main() {
	f, _ := os.Open("input/day15.txt")

	scanner := bufio.NewScanner(bufio.NewReader(f))
	var input []string
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	fmt.Println(part1(input))
	fmt.Println(part2(input))
}
