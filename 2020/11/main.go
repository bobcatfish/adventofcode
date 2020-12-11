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

func CountAdj1(y, x int, s Seats) int {
	c := 0
	for yi := -1; yi <= 1; yi++ {
		for xi := -1; xi <= 1; xi++ {
			// skip the seat itself
			if yi == 0 && xi == 0 {
				continue
			}
			yy, xx := y+yi, x+xi
			if yy >= 0 && xx >= 0 && yy < len(s) && xx < len(s[0]) {
				if s[yy][xx] == '#' {
					c++
				}
			}
		}
	}
	return c
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

func CountAdj2(y, x int, s Seats) int {
	c := 0
	for yi := -1; yi <= 1; yi++ {
		for xi := -1; xi <= 1; xi++ {
			// skip the seat itself
			if yi == 0 && xi == 0 {
				continue
			}

			sees := false
			yy, xx := y+yi, x+xi
			for {
				if yy < 0 || xx < 0 || yy >= len(s) || xx >= len(s[0]) {
					break
				}
				r := s[yy][xx]

				if r == '#' {
					sees = true
				} else if r == 'L' {
					break
				}
				yy += yi
				xx += xi
			}
			if sees {
				c++
			}
		}
	}
	return c
}

func occupied2(c int) rune {
	if c >= 5 {
		return 'L'
	}
	return '#'
}

type getNewSeat func(c int) rune
type countAdj func(x, y int, s Seats) int

type Rules struct {
	CountAdj countAdj
	Occupied getNewSeat
	Empty    getNewSeat
}

func (s Seats) Tick(r Rules) Seats {
	ss := Seats{}
	for y, row := range s {
		newRow := ""
		for x, seat := range row {
			var n rune
			if seat == '#' {
				adj := r.CountAdj(y, x, s)
				n = r.Occupied(adj)
			} else if seat == 'L' {
				adj := r.CountAdj(y, x, s)
				n = r.Empty(adj)
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
		CountAdj: CountAdj1,
		Occupied: occupied1,
		Empty:    empty,
	})
	fmt.Println(s.Count())

	s = findStable(seats, Rules{
		CountAdj: CountAdj2,
		Occupied: occupied2,
		Empty:    empty,
	})
	fmt.Println(s.Count())
}
