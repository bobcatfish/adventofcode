package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const days = 100

const NW = 0
const NE = 1
const E = 2
const SE = 3
const SW = 4
const W = 5

func load() ([][]int, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, fmt.Errorf("couldn't open file: %v", err)
	}
	defer file.Close()

	moves := [][]int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		moveLine := []int{}

		for {
			if len(s) == 0 {
				break
			}
			move := -1
			if len(s) >= 2 {
				two := s[:2]
				switch two {
				case "nw":
					move = NW
				case "ne":
					move = NE
				case "se":
					move = SE
				case "sw":
					move = SW
				}
			}
			if move != -1 {
				s = s[2:]
			} else if len(s) >= 1 {
				one := s[:1]
				switch one {
				case "e":
					move = E
				case "w":
					move = W
				}
				if move != -1 {
					s = s[1:]
				}
			}
			if move != -1 {
				moveLine = append(moveLine, move)
			}
		}
		moves = append(moves, moveLine)
	}

	return moves, err
}

type Point struct {
	Y int
	X int
}

func follow(ref Point, moves []int) Point {
	curr := ref
	for _, move := range moves {
		y, x := curr.Y, curr.X

		switch move {
		case NW:
			y--
		case NE:
			x++
		case E:
			y++
			x++
		case SE:
			y++
		case SW:
			x--
		case W:
			y--
			x--
		}
		curr = Point{Y: y, X: x}
	}
	return curr
}

func getNeighbors(p Point) []Point {
	return []Point{
		//NW:
		{Y: p.Y - 1, X: p.X},
		//NE:
		{Y: p.Y, X: p.X + 1},
		//E:
		{Y: p.Y + 1, X: p.X + 1},
		//SE:
		{Y: p.Y + 1, X: p.X},
		//SW:
		{Y: p.Y, X: p.X - 1},
		//W:
		{Y: p.Y - 1, X: p.X - 1},
	}
}

func lives(p Point, neighbors []Point, black map[Point]struct{}) bool {
	count := 0
	for _, n := range neighbors {
		if _, ok := black[n]; ok {
			count++
		}
	}
	_, alive := black[p]

	if alive {
		return count > 0 && count <= 2
	}
	return count == 2
}

func process(black map[Point]struct{}) map[Point]struct{} {
	newBlack := map[Point]struct{}{}
	for p, _ := range black {
		neighbors := getNeighbors(p)

		if lives(p, neighbors, black) {
			newBlack[p] = struct{}{}
		}

		for _, n := range neighbors {
			nn := getNeighbors(n)
			if lives(n, nn, black) {
				newBlack[n] = struct{}{}
			}
		}
	}
	return newBlack
}

func main() {
	moves, err := load()
	if err != nil {
		log.Fatalf("Couldn't load from file: %v", err)
	}
	ref := Point{Y: 0, X: 0}

	black := map[Point]struct{}{}
	for _, move := range moves {
		p := follow(ref, move)

		if _, ok := black[p]; ok {
			delete(black, p)
		} else {
			black[p] = struct{}{}
		}
	}
	fmt.Println(len(black))

	for i := 0; i < days; i++ {
		black = process(black)
		//fmt.Println("Day", i+1, len(black))
	}
	fmt.Println(len(black))

}
