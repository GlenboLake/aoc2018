package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func part1(input []string) string {
	requirements := map[string][]string{}
	for _, line := range input {
		var a, b string
		fmt.Sscanf(line, "Step %s must be finished before step %s can begin.", &a, &b)
		requirements[b] = append(requirements[b], a)
		if _, present := requirements[a]; !present {
			requirements[a] = []string{}
		}
	}

	order := ""
	for {
		var options []string
		for node, reqs := range requirements {
			if strings.Contains(order, node) {
				continue
			}
			met := 0
			for _, r := range reqs {
				if strings.Contains(order, r) {
					met += 1
				}
			}
			if met == len(reqs) {
				options = append(options, node)
			}
		}
		if len(options) == 0 {
			break
		} else {
			sort.Strings(options)
			order += options[0]
		}
	}

	return order
}

type Worker struct {
	task     string
	timeLeft int
}

func (w *Worker) assign(task string) {
	w.task = task
	w.timeLeft = taskTime(task)
}

func idleWorkers(workers []Worker) []int {
	var idle []int
	for i, w := range workers {
		if w.task == "" {
			idle = append(idle, i)
		}
	}
	return idle
}

func taskTime(task string) int {
	return int(task[0]) - 4
}

func part2(input []string) int {
	requirements := map[string][]string{}
	for _, line := range input {
		var a, b string
		fmt.Sscanf(line, "Step %s must be finished before step %s can begin.", &a, &b)
		requirements[b] = append(requirements[b], a)
		if _, present := requirements[a]; !present {
			requirements[a] = []string{}
		}
	}

	order := ""
	workers := make([]Worker, 5, 5)

	t := 0
	for len(order) < len(requirements) {
		// Check for completed work
		var completed []string
		for i := 0; i < len(workers); i++ {
			if workers[i].timeLeft == 0 && workers[i].task != "" {
				completed = append(completed, workers[i].task)
				workers[i].task = ""
			}
		}
		sort.Strings(completed)
		for _, task := range completed {
			order += task
		}

		// Assign new work
		idle := idleWorkers(workers)
		currentWork := map[string]bool{}
		for _, w := range workers {
			if w.task != "" {
				currentWork[w.task] = true
			}
		}
		if len(idle) > 0 {
			var options []string
			for node, reqs := range requirements {
				if strings.Contains(order, node) || currentWork[node] {
					continue
				}
				met := 0
				for _, r := range reqs {
					if strings.Contains(order, r) {
						met += 1
					}
				}
				if met == len(reqs) {
					options = append(options, node)
				}
			}
			sort.Strings(options)
			numAssignments := 0
			if len(options) < len(idle) {
				numAssignments = len(options)
			} else {
				numAssignments = len(idle)
			}
			for i := 0; i < numAssignments; i++ {
				workers[idle[i]].assign(options[i])
			}
		}

		// Tick down the clock
		for i := 0; i < len(workers); i++ {
			if workers[i].timeLeft > 0 {
				workers[i].timeLeft -= 1
			}
		}
		t += 1
	}

	return t - 1
}

func main() {
	f, _ := os.Open("input/day07.txt")

	var input []string
	scanner := bufio.NewScanner(bufio.NewReader(f))
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	fmt.Println(part1(input))
	fmt.Println(part2(input))
}
