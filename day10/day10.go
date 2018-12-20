package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Vector struct {
	x, y int
}

type Particle struct {
	position, velocity Vector
}

func (p *Particle) Tick() {
	p.position.x += p.velocity.x
	p.position.y += p.velocity.y
}

func (p *Particle) TickBack() {
	p.position.x -= p.velocity.x
	p.position.y -= p.velocity.y
}

func skySize(stars []*Particle) int {
	xmin := stars[0].position.x
	xmax := stars[0].position.x
	ymin := stars[0].position.y
	ymax := stars[0].position.y

	for _, s := range stars {
		if s.position.x < xmin {
			xmin = s.position.x
		}
		if s.position.x > xmax {
			xmax = s.position.x
		}
		if s.position.y < ymin {
			ymin = s.position.y
		}
		if s.position.y > ymax {
			ymax = s.position.y
		}
	}
	return (xmax - xmin) * (ymax - ymin)
}

func skyText(stars []*Particle) string {
	xmin := stars[0].position.x
	xmax := stars[0].position.x
	ymin := stars[0].position.y
	ymax := stars[0].position.y

	sky := map[Vector]bool{}

	for _, s := range stars {
		sky[s.position] = true
		if s.position.x < xmin {
			xmin = s.position.x
		}
		if s.position.x > xmax {
			xmax = s.position.x
		}
		if s.position.y < ymin {
			ymin = s.position.y
		}
		if s.position.y > ymax {
			ymax = s.position.y
		}
	}

	skyText := ""
	for y := ymin; y <= ymax; y++ {
		for x := xmin; x <= xmax; x++ {
			if sky[Vector{x: x, y: y}] {
				skyText += "#"
			} else {
				skyText += " "
			}
		}
		skyText += "\n"
	}
	return skyText
}

func solve(stars []*Particle) {
	size := skySize(stars)
	t := 0
	for {

		for _, s := range stars {
			s.Tick()
		}
		newSize := skySize(stars)
		if newSize > size {
			for _, s := range stars {
				s.TickBack()
			}
			fmt.Println(skyText(stars))
			fmt.Println(t)
			return
		}
		t += 1
		size = newSize
	}
}

func main() {
	f, _ := os.Open("input/day10.txt")

	var input []*Particle
	scanner := bufio.NewScanner(bufio.NewReader(f))
	for scanner.Scan() {
		px, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()[10:16]))
		py, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()[18:24]))
		vx, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()[36:38]))
		vy, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()[40:42]))
		input = append(input, &Particle{
			position: Vector{x: px, y: py},
			velocity: Vector{x: vx, y: vy},
		})
	}

	solve(input)
}
