package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func getRow(s string) (int, error) {
	vv := 0
	for i, r := range s {
		v := 0
		switch r {
		case 'F':
		case 'B':
			v = 1
		default:
			return -1, fmt.Errorf("unexpected rune %c for row in %s", r, s)
		}
		shift := len(s) - 1 - i
		vv += v << uint(shift)
	}
	return vv, nil
}

func getCol(s string) (int, error) {
	vv := 0
	for i, r := range s {
		v := 0
		switch r {
		case 'L':
		case 'R':
			v = 1
		default:
			return -1, fmt.Errorf("unexpected rune %c for column in %s", r, s)
		}
		shift := len(s) - 1 - i
		vv += v << uint(shift)
	}
	return vv, nil
}

func convert(s string) (Seat, error) {
	ss := Seat{Pass: s}

	row, col := s[:7], s[7:]
	var err error
	ss.Row, err = getRow(row)
	if err != nil {
		return ss, fmt.Errorf("couldn't convert row: %s", err)
	}
	ss.Column, err = getCol(col)
	if err != nil {
		return ss, fmt.Errorf("couldn't convert column: %s", err)
	}
	ss.Id = (ss.Row * 8) + ss.Column
	return ss, nil
}

func load() ([]Seat, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, fmt.Errorf("couldn't open file: %v", err)
	}
	defer file.Close()

	vals := []Seat{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		val, err := convert(scanner.Text())
		if err != nil {
			return nil, fmt.Errorf("couldn't convert %v: %v", val, err)
		}
		vals = append(vals, val)
	}
	return vals, err
}

type Seat struct {
	Pass   string
	Row    int
	Column int
	Id     int
}

func findGap(seats []Seat) (int, error) {
	for i := 1; i < len(seats); i++ {
		seat := seats[i]
		if i == len(seats)-1 {
			return -1, fmt.Errorf("didn't find gap!!")
		}

		expected := seat.Id - 1
		if seats[i-1].Id != expected {
			return expected, nil
		}
	}
	return -1, fmt.Errorf("inconceivable!")
}

func main() {
	seats, err := load()
	if err != nil {
		log.Fatalf("Couldn't load nums from file: %v", err)
	}

	sort.Slice(seats, func(i, j int) bool {
		return seats[i].Id < seats[j].Id
	})
	fmt.Println(seats[len(seats)-1])

	gap, err := findGap(seats)
	if err != nil {
		log.Fatalf("error looking for gap: %v", err)
	}
	fmt.Println(gap)
}
