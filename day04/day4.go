package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
)

type Log struct {
	Moment time.Time
	Event  string
}

type Guard struct {
	Id            int
	AsleepMinutes map[int]int
}

type GuardTeam map[int]Guard

func (g Guard) TimeAsleep() int {
	total := 0
	for _, mins := range g.AsleepMinutes {
		total += mins
	}
	return total
}

func (g Guard) MostAsleep() int {
	var when, howMuch int
	for minute, count := range g.AsleepMinutes {
		if count > howMuch {
			when = minute
			howMuch = count
		}
	}
	return when
}

func checkSleep(input []Log) GuardTeam {
	sort.Slice(input, func(i, j int) bool {
		return input[i].Moment.Before(input[j].Moment)
	})

	var onDuty int
	var asleepTime time.Time
	guards := GuardTeam{}
	for _, event := range input {
		if strings.Contains(event.Event, "begins shift") {
			fmt.Sscanf(event.Event, "Guard #%d begins shift", &onDuty)
			asleepTime = time.Time{}
			if _, ok := guards[onDuty]; !ok {
				guards[onDuty] = Guard{Id: onDuty, AsleepMinutes: map[int]int{}}
			}
		} else if event.Event == "falls asleep" {
			asleepTime = event.Moment
		} else if event.Event == "wakes up" {
			awakeTime := event.Moment
			for t := asleepTime; t != awakeTime; t = t.Add(time.Minute) {
				guards[onDuty].AsleepMinutes[t.Minute()] += 1
			}
		}
	}
	return guards
}

func part1(guards GuardTeam) int {
	worstGuard := 0
	worstSleep := 0
	for id, guard := range guards {
		if guard.TimeAsleep() > worstSleep {
			worstGuard = id
			worstSleep = guard.TimeAsleep()
		}
	}
	return worstGuard * guards[worstGuard].MostAsleep()
}

func part2(guards GuardTeam) int {
	var worstGuard, worstMinute, offenses int
	for id, guard := range guards {
		minute := guard.MostAsleep()
		if guard.AsleepMinutes[minute] > offenses {
			worstGuard = id
			worstMinute = minute
			offenses = guard.AsleepMinutes[minute]
		}
	}

	return worstGuard * worstMinute
}

const format = "2006-01-02 15:04"

func parseLine(line string) Log {
	t, err := time.Parse(format, line[1:17])
	if err != nil {
		fmt.Println(err)
	}

	return Log{
		Moment: t,
		Event:  line[19:],
	}
}

func main() {
	var input []Log

	f, _ := os.Open("input/day04.txt")

	scanner := bufio.NewScanner(bufio.NewReader(f))
	for scanner.Scan() {
		log := parseLine(scanner.Text())
		input = append(input, log)
	}
	guards := checkSleep(input)

	fmt.Println(part1(guards))
	fmt.Println(part2(guards))
}
