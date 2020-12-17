package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type getNeighbors func(c Coord) []Coord

func getNeighbors3d(c Coord) []Coord {
	cc := []Coord{}
	for zi := -1; zi <= 1; zi++ {
		for yi := -1; yi <= 1; yi++ {
			for xi := -1; xi <= 1; xi++ {
				if yi == 0 && xi == 0 && zi == 0 {
					continue
				}
				cc = append(cc, Coord{Z: c.Z + zi, Y: c.Y + yi, X: c.X + xi})
			}
		}
	}
	return cc
}

func getNeighbors4d(c Coord) []Coord {
	cc := []Coord{}
	for wi := -1; wi <= 1; wi++ {
		for zi := -1; zi <= 1; zi++ {
			for yi := -1; yi <= 1; yi++ {
				for xi := -1; xi <= 1; xi++ {
					if yi == 0 && xi == 0 && zi == 0 && wi == 0 {
						continue
					}
					cc = append(cc, Coord{W: c.W + wi, Z: c.Z + zi, Y: c.Y + yi, X: c.X + xi})
				}
			}
		}
	}
	return cc
}

func countActive(n []Coord, s Space) int {
	c := 0
	for _, nn := range n {
		if _, ok := s[nn]; ok {
			c++
		}
	}
	return c
}

func becomesActive(c Coord, s Space, f getNeighbors) (bool, []Coord) {
	neighbors := f(c)

	count := countActive(neighbors, s)

	if _, ok := s[c]; ok {
		// if active
		if count == 2 || count == 3 {
			return true, neighbors
		}
	} else {
		// if inactive
		if count == 3 {
			return true, neighbors
		}
	}
	return false, neighbors
}

func Tick(space Space, f getNeighbors) Space {
	ss := Space{}

	for c, _ := range space {
		a, neighbors := becomesActive(c, space, f)
		if a {
			ss[c] = struct{}{}
		}
		for _, n := range neighbors {
			if b, _ := becomesActive(n, space, f); b {
				ss[n] = struct{}{}
			}
		}
	}
	return ss
}

type Coord struct {
	W int
	Z int
	Y int
	X int
}

type Space map[Coord]struct{}

func load() (Space, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, fmt.Errorf("couldn't open file: %v", err)
	}
	defer file.Close()

	space := Space{}
	scanner := bufio.NewScanner(file)
	y := 0
	for scanner.Scan() {
		s := scanner.Text()
		for x, r := range s {
			if r == '#' {
				coord := Coord{
					Y: y,
					X: x,
				}
				space[coord] = struct{}{}
			}
		}
		y++
	}
	return space, err
}

const cycles = 6

func main() {
	startingSpace, err := load()
	if err != nil {
		log.Fatalf("Couldn't load from file: %v", err)
	}

	space := startingSpace
	for i := 0; i < cycles; i++ {
		space = Tick(space, getNeighbors3d)
	}
	fmt.Println(len(space))

	space2 := startingSpace
	for i := 0; i < cycles; i++ {
		space2 = Tick(space2, getNeighbors4d)
	}
	fmt.Println(len(space2))
}
