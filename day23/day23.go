package main

import (
	"bufio"
	"fmt"
	"github.com/glenbolake/aoc2018"
	"math"
	"os"
	"strconv"
	"strings"
)

type Nanobot struct {
	id      int
	x, y, z int
	r       int
}

func (n Nanobot) String() string {
	return fmt.Sprintf("(%d,%d,%d)%d", n.x, n.y, n.z, n.r)
}

func (n Nanobot) distTo(other Nanobot) int {
	return aoc2018.Abs(n.x-other.x) + aoc2018.Abs(n.y-other.y) + aoc2018.Abs(n.z-other.z)
}

func part1(bots []Nanobot) int {
	var bestBot Nanobot
	bestRadius := 0
	for _, b := range bots {
		if b.r > bestRadius {
			bestRadius = b.r
			bestBot = b
		}
	}

	return botsInRange(bots, bestBot)
}

func botsInRange(bots []Nanobot, point Nanobot) int {
	count := 0
	for _, bot := range bots {
		if bot.distTo(point) <= bot.r {
			count += 1
		}
	}
	return count
}

type NanobotSet struct {
	data map[Nanobot]struct{}
}

func (ns *NanobotSet) String() string {
	ids := make([]string, 0, len(ns.data))
	for datum := range ns.data {
		ids = append(ids, strconv.Itoa(datum.id))
	}
	return "{" + strings.Join(ids, " ") + "}"
}

func NewSet() *NanobotSet {
	s := &NanobotSet{}
	s.data = make(map[Nanobot]struct{})
	return s
}

func (ns *NanobotSet) Add(nanobot Nanobot) {
	ns.data[nanobot] = struct{}{}
}

func (ns *NanobotSet) Remove(nanobot Nanobot) {
	delete(ns.data, nanobot)
}

func (ns *NanobotSet) Contains(nanobot Nanobot) bool {
	_, ok := ns.data[nanobot]
	return ok
}

func (ns *NanobotSet) Size() int {
	return len(ns.data)
}

func (ns *NanobotSet) Without(other *NanobotSet) *NanobotSet {
	rv := NewSet()
	for bot := range ns.data {
		if other.Contains(bot) {
			continue
		}
		rv.Add(bot)
	}
	return rv
}

func (ns *NanobotSet) Intersect(other *NanobotSet) *NanobotSet {
	rv := NewSet()
	for bot := range ns.data {
		if other.Contains(bot) {
			rv.Add(bot)
		}
	}
	return rv
}

func (ns *NanobotSet) Union(other *NanobotSet) *NanobotSet {
	rv := NewSet()
	for bot := range ns.data {
		rv.Add(bot)
	}
	for bot := range other.data {
		rv.Add(bot)
	}
	return rv
}

func BronKerbosch(graph map[Nanobot]*NanobotSet) *NanobotSet {
	R := NewSet()
	P := NewSet()
	X := NewSet()
	for bot := range graph {
		P.Add(bot)
	}

	return BronKerboschInner(graph, R, P, X)
}

func BronKerboschInner(graph map[Nanobot]*NanobotSet, R, P, X *NanobotSet) *NanobotSet {
	if P.Size() == 0 && X.Size() == 0 {
		return R
	}

	var pivot Nanobot
	var numNeighbors int
	for bot := range P.Union(X).data {
		neighbors := graph[bot]
		if neighbors.Size() > numNeighbors {
			pivot = bot
			numNeighbors = neighbors.Size()
		}
	}
	vs := P.Without(graph[pivot])
	clique := NewSet()
	for v := range vs.data {
		vSet := NewSet()
		vSet.Add(v)

		newClique := BronKerboschInner(graph, R.Union(vSet), P.Intersect(graph[v]), X.Intersect(graph[v]))
		if newClique.Size() > clique.Size() {
			clique = newClique
		}
		P.Remove(v)
		X.Add(v)
	}
	return clique
}

func part2(bots []Nanobot) int {
	graph := map[Nanobot]*NanobotSet{}
	edgeCount := 0
	for _, b := range bots {
		graph[b] = NewSet()
	}
	for i, bot1 := range bots {
		for _, bot2 := range bots[i+1:] {
			dist := bot1.distTo(bot2)
			if dist <= bot1.r+bot2.r {
				edgeCount += 1
				graph[bot1].Add(bot2)
				graph[bot2].Add(bot1)
			}
		}
	}

	clique := BronKerbosch(graph)
	//fmt.Println(clique)

	min, max := 0, math.MaxInt32
	origin := Nanobot{}
	for bot := range clique.data {
		minDist := origin.distTo(bot) - bot.r
		if minDist > min {
			min = minDist
		}
		maxDist := origin.distTo(bot) + bot.r
		if maxDist < max {
			max = maxDist
		}
	}

	return min
}

func main() {
	f, _ := os.Open("input/day23.txt")

	scanner := bufio.NewScanner(bufio.NewReader(f))
	var bots []Nanobot
	i := 0
	for scanner.Scan() {
		b := Nanobot{}
		fmt.Sscanf(scanner.Text(), "pos=<%d,%d,%d>, r=%d", &b.x, &b.y, &b.z, &b.r)
		b.id = i
		i++
		bots = append(bots, b)
	}

	fmt.Println(part1(bots))
	fmt.Println(part2(bots))
}
