package main

import "fmt"

const (
	/*
		numPlayers = 9
		lastMarble = 25
	*/
	numPlayers = 486
	lastMarble = 70833
	//lastMarble = 70833 * 100
	magic = 23
)

func mod(a, b int) int {
	v := a % b
	if v < 0 {
		return b + v
	}
	return v
}

func placeMarble(marbles []int, currIndex, currMarble, currLen int) (int, []int, int, int) {
	var score int

	if currMarble%magic == 0 {
		score += currMarble
		moveIndex := mod((currIndex - 7), currLen)
		score += marbles[moveIndex]
		copy(marbles[moveIndex:currLen], marbles[moveIndex+1:currLen])
		currIndex = moveIndex
		currLen--
	} else {
		nextMarble := mod(currIndex+1, currLen)

		copy(marbles[nextMarble+2:], marbles[nextMarble+1:])
		marbles[nextMarble+1] = currMarble

		currIndex = nextMarble + 1
		currLen++
	}
	return score, marbles, currIndex, currLen
}

func main() {
	marbles := make([]int, lastMarble)
	marbles[0] = 0
	marbles[1] = 2
	marbles[2] = 1

	currLen := 3
	currIndex := 1
	currPlayer := 3
	scores := map[int]int{}

	for i := 0; i < numPlayers; i++ {
		scores[i] = 0
	}

	for currMarble := 3; currMarble <= lastMarble; currMarble++ {
		var score int
		score, marbles, currIndex, currLen = placeMarble(marbles, currIndex, currMarble, currLen)
		scores[currPlayer] += score

		currPlayer = (currPlayer + 1) % numPlayers

		/*
			for i, m := range marbles {
				if i == currIndex {
					fmt.Printf("(%d) ", m)
				} else {
					fmt.Printf("%d ", m)
				}
			}
			fmt.Println()
			fmt.Println(scores)
		*/
	}

	max := 0
	for _, score := range scores {
		if score > max {
			max = score
		}
	}

	fmt.Println(max)
}
