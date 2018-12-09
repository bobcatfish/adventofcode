package main

import "fmt"

const (
	numPlayers = 486
	lastMarble = 70833 * 100
	magic      = 23
)

type marble struct {
	id   int
	next *marble
	prev *marble
}

func (m marble) String() string {
	return fmt.Sprintf("[%d] next: %d, prev: %d", m.id, m.next.id, m.prev.id)
}

func placeMarble(currM *marble, m *marble, length int) (*marble, int, int) {
	var score int
	var curr *marble

	if m.id%magic == 0 {
		score += m.id

		marbleToReplace := currM
		for i := 0; i < 7; i++ {
			marbleToReplace = marbleToReplace.prev
		}

		before := marbleToReplace.prev
		after := marbleToReplace.next
		score += marbleToReplace.id

		before.next = after
		after.prev = before

		curr = after

		length--
	} else {
		after := currM.next

		m.next = after.next
		m.prev = after

		after.next.prev = m
		after.next = m

		curr = m
		length++
	}

	return curr, score, length
}

func main() {

	next := &marble{
		id: 1,
	}
	prev := &marble{
		id: 0,
	}
	currMarble := &marble{
		id:   2,
		next: next,
		prev: prev,
	}
	next.next = prev
	next.prev = currMarble

	prev.next = currMarble
	prev.prev = next

	currPlayer := 3

	scores := map[int]int{}

	for i := 0; i < numPlayers; i++ {
		scores[i] = 0
	}

	length := 3
	for currMarbleVal := 3; currMarbleVal <= lastMarble; currMarbleVal++ {
		var score int
		marble := &marble{
			id: currMarbleVal,
		}
		currMarble, score, length = placeMarble(currMarble, marble, length)
		scores[currPlayer] += score

		currPlayer = (currPlayer + 1) % numPlayers
	}

	max := 0
	for _, score := range scores {
		if score > max {
			max = score
		}
	}

	fmt.Println(max)
}
