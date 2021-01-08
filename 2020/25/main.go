package main

import (
	"fmt"
	"math"
)

func loop(subj, loop int) int {
	curr := 1
	for i := 0; i < loop; i++ {
		curr = curr * subj
		curr %= 20201227
	}
	return curr
}

func wangJangle(pub int) int {
	i := 8
	val := int(math.Pow(float64(7), float64(8)))
	if val == pub {
		return i
	}
	for {
		i++
		val = (val * 7) % 20201227

		if val == pub {
			return i
		}
	}
	return -1

}

func main() {
	doorPub, cardPub := 1327981, 2822615
	//doorPub, cardPub := 17807724, 5764801

	cardLoop := wangJangle(cardPub)
	fmt.Println("Card loop", cardLoop)
	doorLoop := wangJangle(doorPub)
	fmt.Println("Door loop", doorLoop)

	doorKey := loop(doorPub, cardLoop)
	fmt.Println("Door key", doorKey)

	cardKey := loop(cardPub, doorLoop)
	fmt.Println("Card key", cardKey)

}
