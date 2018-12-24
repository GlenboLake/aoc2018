package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

const (
	ImmuneSystem = iota
	Infection    = iota
)

type UnitGroup struct {
	team       int
	count      int
	hp         int
	damage     int
	attackType string
	initiative int
	weak       string
	immune     string
}

func (ug UnitGroup) EffectivePower(boost int) int {
	if ug.team == ImmuneSystem {
		return ug.count * (ug.damage + boost)
	} else {
		return ug.count * ug.damage
	}
}

func (ug *UnitGroup) WeakTo(attack string) bool {
	for _, w := range strings.Split(ug.weak, ", ") {
		if w == attack {
			return true
		}
	}
	return false
}

func (ug *UnitGroup) ImmuneTo(attack string) bool {
	for _, i := range strings.Split(ug.immune, ", ") {
		if i == attack {
			return true
		}
	}
	return false
}

func parseGroup(team int, line string) UnitGroup {
	regex, _ := regexp.Compile(`(\d+) units each with (\d+) hit points(?:.*) with an attack that does (\d+) (\w+) damage at initiative (\d+)`)
	weakRegex, _ := regexp.Compile(`weak to ([\w, ]+)`)
	immuneRegex, _ := regexp.Compile(`immune to ([\w, ]+)`)

	result := regex.FindStringSubmatch(line)
	count, _ := strconv.Atoi(result[1])
	hp, _ := strconv.Atoi(result[2])
	damage, _ := strconv.Atoi(result[3])
	attackType := result[4]
	init, _ := strconv.Atoi(result[5])
	weak := weakRegex.FindStringSubmatch(line)
	var weaknesses string
	if len(weak) > 1 {
		weaknesses = weak[1]
	}
	immune := immuneRegex.FindStringSubmatch(line)
	var immunities string
	if len(immune) > 1 {
		immunities = immune[1]
	}

	group := UnitGroup{
		team:       team,
		count:      count,
		hp:         hp,
		damage:     damage,
		attackType: attackType,
		initiative: init,
		weak:       weaknesses,
		immune:     immunities,
	}

	return group
}

func simulateBattle(team1, team2 []*UnitGroup, boost int) (int, int) {
	//fmt.Println("Initial:")
	//fmt.Println("Team 1:")
	//for _, ug := range team1 {
	//	fmt.Println(ug.count, "units", ug.EffectivePower(boost), "power")
	//}
	//fmt.Println("Team 2:")
	//for _, ug := range team2 {
	//	fmt.Println(ug.count, "units", ug.EffectivePower(boost), "power")
	//}

	team1copy := make([]*UnitGroup, len(team1))
	for i, ug := range team1 {
		ugc := *ug
		team1copy[i] = &ugc
	}
	team2copy := make([]*UnitGroup, len(team2))
	for i, ug := range team2 {
		ugc := *ug
		team2copy[i] = &ugc
	}

	for len(team1copy) > 0 && len(team2copy) > 0 {
		// Order by effective power and initiative
		//fmt.Print("New round: ")
		var deathToll int
		units := make([]*UnitGroup, 0, len(team1copy)+len(team2copy))
		units = append(units, team1copy...)
		units = append(units, team2copy...)
		sort.Slice(units, func(i, j int) bool {
			if units[i].EffectivePower(boost) == units[j].EffectivePower(boost) {
				return units[i].initiative > units[j].initiative
			}
			return units[i].EffectivePower(boost) > units[j].EffectivePower(boost)
		})

		// Make list of valid targets
		targetsRemaining := map[*UnitGroup]struct{}{}
		for _, u := range units {
			targetsRemaining[u] = struct{}{}
		}
		targets := map[*UnitGroup]*UnitGroup{}

		for _, attacker := range units {
			var target *UnitGroup
			bestDamage := 0
			for t := range targetsRemaining {
				if t.team == attacker.team {
					continue
				}
				if t.ImmuneTo(attacker.attackType) {
					continue
				}
				damage := attacker.EffectivePower(boost)
				if t.WeakTo(attacker.attackType) {
					damage *= 2
				}
				if damage > bestDamage || target == nil {
					target = t
					bestDamage = damage
				} else if damage == bestDamage {
					if t.EffectivePower(boost) > target.EffectivePower(boost) {
						target = t
					} else if t.EffectivePower(boost) == target.EffectivePower(boost) {
						if t.initiative > target.initiative {
							target = t
						}
					}
				}
			}
			targets[attacker] = target
			delete(targetsRemaining, target)
		}

		// Attack phase
		sort.Slice(units, func(i, j int) bool {
			return units[i].initiative > units[j].initiative
		})

		for _, attacker := range units {
			if attacker.count <= 0 {
				continue
			}
			defender := targets[attacker]
			if defender == nil {
				continue
			}
			attackDamage := attacker.EffectivePower(boost)
			if defender.WeakTo(attacker.attackType) {
				attackDamage *= 2
			}
			defendersKilled := attackDamage / defender.hp
			if defendersKilled > defender.count {
				defendersKilled = defender.count
			}
			deathToll += defendersKilled
			defender.count -= defendersKilled
		}
		if deathToll == 0 {
			return Infection, 0
		}

		// Remove dead units
		var alive1 []*UnitGroup
		for _, ug := range team1copy {
			if ug.count > 0 {
				alive1 = append(alive1, ug)
			}
		}
		team1copy = alive1
		var alive2 []*UnitGroup
		for _, ug := range team2copy {
			if ug.count > 0 {
				alive2 = append(alive2, ug)
			}
		}
		team2copy = alive2
	}

	var winner, alive int
	if len(team1copy) == 0 {
		winner = Infection
		for _, ug := range team2copy {
			alive += ug.count
		}
	} else {
		winner = ImmuneSystem
		for _, ug := range team1copy {
			alive += ug.count
		}
	}

	return winner, alive
}

func part1(team1, team2 []*UnitGroup) int {
	_, alive := simulateBattle(team1, team2, 0)
	return alive
}

func part2(team1, team2 []*UnitGroup) int {
	high, low := math.MaxInt32, 0
	var result int

	for high-low > 1 {
		testBoost := (high + low) / 2

		winner, alive := simulateBattle(team1, team2, testBoost)
		if winner == ImmuneSystem {
			high = testBoost
			result = alive
		} else {
			low = testBoost
		}
	}
	return result
}

func main() {
	f, _ := os.Open("input/day24.txt")

	scanner := bufio.NewScanner(bufio.NewReader(f))
	var immuneSystem, infection []*UnitGroup
	var team int
	for scanner.Scan() {
		switch scanner.Text() {
		case "":
			break
		case "Immune System:":
			team = ImmuneSystem
		case "Infection:":
			team = Infection
		default:
			group := parseGroup(team, scanner.Text())
			switch team {
			case 0:
				immuneSystem = append(immuneSystem, &group)
			case 1:
				infection = append(infection, &group)
			}
		}
	}

	fmt.Println(part1(immuneSystem, infection))
	fmt.Println(part2(immuneSystem, infection))
}
