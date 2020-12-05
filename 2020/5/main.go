package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func convert(s string) (string, error) {
	// oh dear i have go 1.9 and arch linux
	// var b strings.Builder
	ss := ""
	for _, r := range s {
		switch r {
		case 'F':
			ss += "0"
		case 'B':
			ss += "1"
		case 'L':
			ss += "0"
		case 'R':
			ss += "1"
		default:
			return "", fmt.Errorf("unexpected rune %c", r)
		}
	}
	return ss, nil
}

func load() ([]string, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, fmt.Errorf("couldn't open file: %v", err)
	}
	defer file.Close()

	vals := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		val, err := convert(scanner.Text())
		if err != nil {
			return nil, fmt.Errorf("couldn't convert %s: %v", val, err)
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

func lazyInt(r rune) int {
	if r == '0' {
		return 0
	}
	return 1
}

func getSeats(vals []string) []Seat {
	seats := []Seat{}
	for _, v := range vals {
		row, column := 0, 0
		for i, r := range v {
			ri := lazyInt(r)
			if i > 6 {
				shift := (3 - (i - 6))
				column += ri << uint(shift)
			} else {
				shift := 6 - i
				row += ri << uint(shift)
			}
		}
		seats = append(seats, Seat{
			Pass:   v,
			Row:    row,
			Column: column,
			Id:     (row * 8) + column,
		})
	}
	return seats
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
	vals, err := load()
	if err != nil {
		log.Fatalf("Couldn't load nums from file: %v", err)
	}
	seats := getSeats(vals)
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
