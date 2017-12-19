package main

import (
	"fmt"
)

type point struct {
	X int
	Y int
}

func (p point) String() string {
	return fmt.Sprintf("(%v, %v)", p.Y, p.X)
}

func nextPoint(subtract bool, coord string, p point) point {
	if coord == "x" {
		if subtract {
			return point{p.X - 1, p.Y}
		}
		return point{p.X + 1, p.Y}
	}

	// y case
	if subtract {
		return point{p.X, p.Y - 1}
	}
	return point{p.X, p.Y + 1}
}

func nextVal(m map[point]int, loc point) int {
	var neighbours = []point{
		point{
			Y: -1,
			X: -1,
		},
		point{
			Y: -1,
			X: 0,
		},
		point{
			Y: -1,
			X: 1,
		},
		point{
			Y: 0,
			X: -1,
		},
		point{
			Y: 0,
			X: 1,
		},
		point{
			Y: 1,
			X: -1,
		},
		point{
			Y: 1,
			X: 0,
		},
		point{
			Y: 1,
			X: 1,
		},
	}

	sum := 0
	for _, n := range neighbours {
		p := point{
			Y: loc.Y + n.Y,
			X: loc.X + n.X,
		}
		if val, ok := m[p]; ok {
			sum += val
		}
	}
	return sum
}

func main() {
	input := 312051
	m := make(map[point]int)

	loc := point{0, 0}
	c := 1
	m[loc] = c

	actions := 1
	subtract := false

	for c < input {
		for i := 0; i < actions; i++ {
			loc = nextPoint(subtract, "x", loc)

			c = nextVal(m, loc)
			m[loc] = c
			if c > input {
				break
			}
		}

		if c > input {
			break
		}

		subtract = !subtract
		for i := 0; i < actions; i++ {
			loc = nextPoint(subtract, "y", loc)

			c = nextVal(m, loc)
			m[loc] = c
			if c > input {
				break
			}
		}
		actions++
	}

	fmt.Println(c)
}
