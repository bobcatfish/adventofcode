package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type square struct {
	Size int
	Rows []string
	Grid string
}

func (s square) printGrid() {
	for _, r := range s.Rows {
		fmt.Println(r)
	}
}

func (s square) split() [][]square {
	var d int
	if s.Size%2 == 0 {
		d = 2
	} else if s.Size%3 == 0 {
		d = 3
	} else {
		log.Fatalf("Not divis %v\n", s)
	}

	n := s.Size / d

	ss := make([][]square, n)
	for y := 0; y < n; y++ {
		ss[y] = make([]square, n)
		for x := 0; x < n; x++ {
			var sq square
			sq.Size = d

			for j := 0; j < d; j++ {
				sr := (y * d) + j
				so := (x * d)
				sq.Rows = append(sq.Rows, s.Rows[sr][so:so+d])
			}

			sq.Grid = strings.Join(sq.Rows, "/")
			ss[y][x] = sq
		}
	}
	return ss
}

func getRotations(s square) []square {
	var ss []square

	for i := 0; i < 4; i++ {
		var r string
		for x := 0; x < s.Size; x++ {
			for y := s.Size - 1; y >= 0; y-- {
				r += string(s.Rows[y][x])
			}
			if x < s.Size-1 {
				r += "/"
			}
		}
		//fmt.Println("rotation", r, "of", s)
		s = newSquare(r)
		ss = append(ss, s)
	}
	return ss
}

func getFlips(s square) []square {
	var ss []square

	var r string
	for y := 0; y < s.Size; y++ {
		for x := s.Size - 1; x >= 0; x-- {
			r += string(s.Rows[y][x])
		}
		if y < s.Size-1 {
			r += "/"
		}
	}
	//fmt.Println("flip", r, "of", s)
	ss = append(ss, newSquare(r))

	r = ""
	for y := s.Size - 1; y >= 0; y-- {
		for x := 0; x < s.Size; x++ {
			r += string(s.Rows[y][x])
		}
		if y > 0 {
			r += "/"
		}
	}
	//fmt.Println("flip", r, "of", s)
	return append(ss, newSquare(r))
}

func addRule(m map[string]square, line string) {
	p := strings.Split(line, " => ")
	input := newSquare(p[0])
	output := newSquare(p[1])

	m[input.Grid] = output

	r := getRotations(input)
	f := getFlips(input)

	for _, rot := range r {
		m[rot.Grid] = output
		flips := getFlips(rot)
		for _, flip := range flips {
			m[flip.Grid] = output
		}
	}
	for _, flip := range f {
		m[flip.Grid] = output
		rots := getRotations(flip)
		for _, rot := range rots {
			m[rot.Grid] = output
		}
	}
}

func newSquare(line string) square {
	var s square
	s.Rows = strings.Split(line, "/")
	s.Size = len(s.Rows)
	s.Grid = line
	return s
}

func combine(ss [][]square) square {
	var s string

	if len(ss[0]) == 1 {
		return ss[0][0]
	}

	for y := 0; y < len(ss); y++ {
		for j := 0; j < ss[y][0].Size; j++ {
			for x := 0; x < len(ss[y]); x++ {
				s += ss[y][x].Rows[j]
			}
			if j < ss[y][0].Size-1 {
				s += "/"
			}
		}
		if y < len(ss)-1 {
			s += "/"
		}
	}
	//fmt.Println("old", ss)
	//fmt.Println("new", s)
	return newSquare(s)
}

func applyRule(s square, m map[string]square) square {
	ss := s.split()

	for y := 0; y < len(ss); y++ {
		for x := 0; x < len(ss[y]); x++ {
			v, ok := m[ss[y][x].Grid]
			if !ok {
				log.Fatalf("No rule for %v\n", ss[y][x].Grid)
			}
			ss[y][x] = v
		}
	}
	return combine(ss)
}

func count(s square) int {
	var count int
	for _, r := range s.Grid {
		if r == '#' {
			count++
		}
	}
	return count
}

const iterations = 18

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Couldn't open file: %v\n", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	m := make(map[string]square)
	for scanner.Scan() {
		line := scanner.Text()
		addRule(m, line)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	/*
		for k, v := range m {
			fmt.Println(k, v)
		}
		os.Exit(0)
	*/

	s := newSquare(".#./..#/###")

	for i := 0; i < iterations; i++ {
		fmt.Println("start", i)
		s = applyRule(s, m)
		fmt.Println("done", i)
	}

	s.printGrid()

	fmt.Println(count(s))
}
