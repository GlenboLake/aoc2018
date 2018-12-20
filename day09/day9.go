package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

type Marble struct {
	id         int
	next, prev *Marble
}

func play(numPlayers, numMarbles int) int {
	players := make([]int, numPlayers)
	currentPlayer := 0
	current := &Marble{}
	current.next = current
	current.prev = current
	for i := 1; i <= numMarbles; i++ {
		if i%23 == 0 {
			remove := current.prev.prev.prev.prev.prev.prev.prev
			players[currentPlayer] += i + remove.id
			current = remove.next
			remove.prev.next = remove.next
			remove.next.prev = remove.prev
		} else {
			before := current.next
			after := before.next
			insert := &Marble{id: i, next: after, prev: before}
			before.next = insert
			after.prev = insert
			current = insert
		}
		currentPlayer = (currentPlayer + 1) % numPlayers
	}
	sort.Ints(players)
	return players[len(players)-1]
}

func main() {
	inputBytes, _ := ioutil.ReadFile("input/day09.txt")
	inputString := strings.TrimSpace(string(inputBytes))
	var numPlayers, numMarbles int
	fmt.Sscanf(inputString, "%d players; last marble is worth %d points", &numPlayers, &numMarbles)

	fmt.Println(play(numPlayers, numMarbles))
	fmt.Println(play(numPlayers, numMarbles*100))
}
