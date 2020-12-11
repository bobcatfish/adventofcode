package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	//"time"
)

type Seats []string

func (s Seats) Count() int {
	c := 0
	for _, row := range s {
		for _, seat := range row {
			if seat == '#' {
				c++
			}
		}
	}
	return c
}

func (s Seats) Print() {
	for _, row := range s {
		for _, seat := range row {
			fmt.Printf("%c", seat)
		}
		fmt.Println()
	}
	fmt.Println()
}

func (s Seats) Equals(ss Seats) bool {
	for y, row := range s {
		if ss[y] != row {
			return false
		}
	}
	return true
}

func CountAdj(y, x int, s Seats, check checkXY) int {
	c := 0
	for yi := -1; yi <= 1; yi++ {
		for xi := -1; xi <= 1; xi++ {
			// skip the seat itself
			if yi == 0 && xi == 0 {
				continue
			}
			if check(yi, y, xi, x, s) {
				c++
			}
		}
	}
	return c
}

func checkXYNeighbors(yi, y, xi, x int, s Seats) bool {
	yy, xx := y+yi, x+xi
	if yy >= 0 && xx >= 0 && yy < len(s) && xx < len(s[0]) {
		if s[yy][xx] == '#' {
			return true
		}
	}
	return false
}

func occupied1(c int) rune {
	if c >= 4 {
		return 'L'
	}
	return '#'
}

func empty(c int) rune {
	if c == 0 {
		return '#'
	}
	return 'L'
}

func checkXYVisible(yi, y, xi, x int, s Seats) bool {
	yy, xx := y+yi, x+xi
	for {
		if yy < 0 || xx < 0 || yy >= len(s) || xx >= len(s[0]) {
			break
		}
		r := s[yy][xx]

		if r == '#' {
			return true
		} else if r == 'L' {
			break
		}
		yy += yi
		xx += xi
	}
	return false
}

func occupied2(c int) rune {
	if c >= 5 {
		return 'L'
	}
	return '#'
}

type getNewSeat func(c int) rune
type checkXY func(yi, y, xi, x int, s Seats) bool

type Rules struct {
	CheckXY  checkXY
	Occupied getNewSeat
}

func (s Seats) Tick(r Rules) Seats {
	ss := Seats{}
	for y, row := range s {
		newRow := ""
		for x, seat := range row {
			var n rune
			if seat == '#' {
				adj := CountAdj(y, x, s, r.CheckXY)
				n = r.Occupied(adj)
			} else if seat == 'L' {
				adj := CountAdj(y, x, s, r.CheckXY)
				n = empty(adj)
			} else {
				n = seat
			}
			newRow += string(n)
		}
		ss = append(ss, newRow)
	}
	return ss
}

func load() (Seats, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, fmt.Errorf("couldn't open file: %v", err)
	}
	defer file.Close()

	seats := Seats{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		seats = append(seats, s)
	}
	return seats, err
}

func findStable(curr Seats, r Rules) Seats {
	for {
		/*
			curr.Print()
			time.Sleep(1000 * time.Millisecond)
		*/

		next := curr.Tick(r)
		if curr.Equals(next) {
			return next
		}
		curr = next
	}
}

func main() {
	seats, err := load()
	if err != nil {
		log.Fatalf("Couldn't load from file: %v", err)
	}

	s := findStable(seats, Rules{
		CheckXY:  checkXYNeighbors,
		Occupied: occupied1,
	})
	fmt.Println(s.Count())

	s = findStable(seats, Rules{
		CheckXY:  checkXYVisible,
		Occupied: occupied2,
	})
	fmt.Println(s.Count())
}
