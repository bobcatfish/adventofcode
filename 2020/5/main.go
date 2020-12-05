package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func getVal(s string, one rune) int {
	vv := 0
	for i, r := range s {
		if r == one {
			shift := len(s) - 1 - i
			vv += 1 << uint(shift)
		}
	}
	return vv
}

func convert(s string) (Seat, error) {
	ss := Seat{Pass: s}
	row, col := s[:7], s[7:]
	ss.Row = getVal(row, 'B')
	ss.Column = getVal(col, 'R')
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
