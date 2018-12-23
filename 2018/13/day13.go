package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type dir int

const (
	upCart    = '^'
	downCart  = 'v'
	leftCart  = '<'
	rightCart = '>'

	dirUp       dir = 0
	dirDown     dir = 1
	dirLeft     dir = 2
	dirRight    dir = 3
	dirStraight dir = 4
)

var pattern = []dir{
	dirLeft,
	dirStraight,
	dirRight,
}

var allCarts = map[rune]bool{
	upCart:    true,
	downCart:  true,
	leftCart:  true,
	rightCart: true,
}

type point struct {
	y int
	x int
}

type cart struct {
	d       dir
	pattern int
}

type square struct {
	p point
	c *cart
	r rune
}

func down(loc point) point {
	return point{
		x: loc.x,
		y: loc.y + 1,
	}
}

func left(loc point) point {
	return point{
		x: loc.x - 1,
		y: loc.y,
	}
}

func right(loc point) point {
	return point{
		x: loc.x + 1,
		y: loc.y,
	}
}

func up(loc point) point {
	return point{
		x: loc.x,
		y: loc.y - 1,
	}
}

func getNext(s *square, c *cart) point {
	var p point
	if c.d == dirLeft {
		p = left(s.p)
	} else if c.d == dirRight {
		p = right(s.p)
	} else if c.d == dirUp {
		p = up(s.p)
	} else if c.d == dirDown {
		p = down(s.p)
	}

	return p
}

func getNextDir(s *square, c *cart) dir {
	var d dir
	if s.r == '\\' {
		if c.d == dirLeft {
			d = dirUp
		} else if c.d == dirRight {
			d = dirDown
		} else if c.d == dirUp {
			d = dirLeft
		} else if c.d == dirDown {
			d = dirRight
		}
	} else if s.r == '/' {
		if c.d == dirLeft {
			d = dirDown
		} else if c.d == dirRight {
			d = dirUp
		} else if c.d == dirUp {
			d = dirRight
		} else if c.d == dirDown {
			d = dirLeft
		}
	} else if s.r == '|' || s.r == '-' {
		d = c.d
	} else if s.r == '+' {
		if s.r == '+' {
			p := c.pattern
			c.pattern = (p + 1) % len(pattern)
			dir := pattern[p]
			if dir == dirLeft {
				if c.d == dirLeft {
					d = dirDown
				} else if c.d == dirRight {
					d = dirUp
				} else if c.d == dirUp {
					d = dirLeft
				} else if c.d == dirDown {
					d = dirRight
				}
			} else if dir == dirRight {
				if c.d == dirLeft {
					d = dirUp
				} else if c.d == dirRight {
					d = dirDown
				} else if c.d == dirUp {
					d = dirRight
				} else if c.d == dirDown {
					d = dirLeft
				}
			} else {
				d = c.d
			}
		}
	} else {
		log.Fatalf("wut %c", s.r)
	}
	return d
}

type newDest struct {
	prevS *square
	s     *square
	c     *cart
}

func tick(m [][]*square) map[point]bool {
	todo := []newDest{}

	for y := 0; y < len(m); y++ {
		for x := 0; x < len(m[y]); x++ {
			s := m[y][x]
			if s.c != nil {
				next := getNext(s, s.c)

				todo = append(todo, newDest{
					s:     m[next.y][next.x],
					c:     s.c,
					prevS: s,
				})
			}
		}
	}

	// move the carts
	collocs := map[point]bool{}
	wasCollision := map[point]bool{}
	for _, new := range todo {
		// collision
		if new.s.c != nil || wasCollision[new.prevS.p] {
			new.s.c = nil
			new.prevS.c = nil
			wasCollision[new.s.p] = true
			collocs[new.s.p] = true
		} else {
			new.s.c = new.c
			new.c.d = getNextDir(new.s, new.c)
			new.prevS.c = nil
		}
	}

	return collocs
}

func findCart(m [][]*square) []point {
	p := []point{}
	for y := 0; y < len(m); y++ {
		for x := 0; x < len(m[y]); x++ {
			if m[y][x].c != nil {
				p = append(p, point{x: x, y: y})
			}
		}
	}
	return p
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Couldn't open file: %v", err)
	}
	defer file.Close()

	vals := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		vals = append(vals, scanner.Text())
	}

	m := [][]*square{}
	for y, v := range vals {
		row := []*square{}
		for x, vv := range v {
			s := &square{
				p: point{
					y: y,
					x: x,
				},
			}
			if _, ok := allCarts[vv]; ok {
				var d dir
				var r rune
				if vv == upCart {
					d = dirUp
					r = '|'
				} else if vv == downCart {
					d = dirDown
					r = '|'
				} else if vv == leftCart {
					d = dirLeft
					r = '-'
				} else {
					d = dirRight
					r = '-'
				}
				c := cart{
					d:       d,
					pattern: 0,
				}
				s.c = &c
				s.r = r
			} else {
				s.r = vv
			}
			row = append(row, s)
		}
		m = append(m, row)
	}

	for {
		cols := tick(m)
		if len(cols) > 0 {
			for col := range cols {
				fmt.Printf("collision at %d,%d\n", col.x, col.y)
			}
			cartCount := len(findCart(m))
			if cartCount <= 1 {
				break
			}
		}
	}

	loc := findCart(m)
	fmt.Printf("Last cart at %d,%d\n", loc[0].x, loc[0].y)
}
