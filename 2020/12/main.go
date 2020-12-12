package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func load() ([]Instr, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, fmt.Errorf("couldn't open file: %v", err)
	}
	defer file.Close()

	a := []Instr{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		n, err := strconv.Atoi(s[1:])
		if err != nil {
			return nil, fmt.Errorf("couldn't convert %s: %v", s[1:], err)
		}

		var action Dir
		switch s[0] {
		case 'N':
			action = N
		case 'E':
			action = E
		case 'S':
			action = S
		case 'W':
			action = W
		case 'L':
			action = L
		case 'R':
			action = R
		case 'F':
			action = F
		default:
			return nil, fmt.Errorf("unknown action %c", s[0])
		}

		a = append(a, Instr{
			Action: action,
			Count:  n,
		})
	}
	return a, err
}

type Point struct {
	Y int
	X int
}

type Boat struct {
	Facing Dir
	Loc    Point
}

type Instr struct {
	Action Dir
	Count  int
}

func mod(x int, y int) int {
	v := x % y
	if v < 0 {
		v += y
	}
	return v
}

const N = 0
const E = 1
const S = 2
const W = 3

const L = 4
const R = 5
const F = 6

type Dir int

const left = -1
const right = 1

func rotate(facing Dir, t int) Dir {
	return Dir(mod((int(facing) + t), 4))
}

func forward(facing Dir, loc Point, count int) Point {
	switch facing {
	case N:
		loc.Y -= count
	case S:
		loc.Y += count
	case E:
		loc.X += count
	case W:
		loc.X -= count
	default:
		log.Printf("invalid facing dir %d", facing)
	}
	return loc
}

func move1(b Boat, instr Instr) Boat {
	nb := b
	switch instr.Action {
	case L:
		for i := 0; i < instr.Count; i += 90 {
			nb.Facing = rotate(nb.Facing, left)
		}
	case R:
		for i := 0; i < instr.Count; i += 90 {
			nb.Facing = rotate(nb.Facing, right)
		}
	case F:
		nb.Loc = forward(nb.Facing, nb.Loc, instr.Count)
	default:
		nb.Loc = forward(instr.Action, nb.Loc, instr.Count)
	}
	return nb
}

func move2(b Boat, w Point, instr Instr) (Boat, Point) {
	switch instr.Action {
	case N:
		w.Y -= instr.Count
	case S:
		w.Y += instr.Count
	case E:
		w.X += instr.Count
	case W:
		w.X -= instr.Count
	case L:
		for i := 0; i < instr.Count; i += 90 {
			w.Y, w.X = 0-w.X, w.Y
		}
	case R:
		for i := 0; i < instr.Count; i += 90 {
			w.Y, w.X = w.X, 0-w.Y
		}
	case F:
		b.Loc.X += (instr.Count * w.X)
		b.Loc.Y += (instr.Count * w.Y)
	default:
		log.Printf("invalid instruction %v", instr)
	}
	return b, w
}

func main() {
	a, err := load()
	if err != nil {
		log.Fatalf("Couldn't load from file: %v", err)
	}

	b := Boat{Facing: E}
	w := Point{X: 10, Y: -1}

	b1 := b
	for _, instr := range a {
		b1 = move1(b1, instr)
	}
	fmt.Println(b1, (b1.Loc.Y + b1.Loc.X))

	b2, w2 := b, w
	for _, instr := range a {
		b2, w2 = move2(b2, w2, instr)
	}
	fmt.Println(b2, (b2.Loc.Y + b2.Loc.X))
}
