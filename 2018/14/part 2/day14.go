package main

import (
	"fmt"
	"log"
	"strconv"
)

const (
	after = "430971"
	num   = 10
)

func next(rec []int, elf1, elf2 int) ([]int, error) {
	n := []int{}
	new := rec[elf1] + rec[elf2]
	s := strconv.Itoa(new)
	for _, r := range s {
		i, err := strconv.Atoi(string(r))
		if err != nil {
			return nil, fmt.Errorf("couldn't convert %c to int: %v", r, err)
		}
		n = append(n, i)
	}
	return n, nil
}

func move(rec []int, i int) int {
	steps := rec[i] + 1
	return (i + steps) % len(rec)
}

func main() {
	elf1, elf2 := 0, 1
	rec := []int{3, 7}
	i := 0
	for {
		n, err := next(rec, elf1, elf2)
		if err != nil {
			log.Fatalf("Couldn't determine next value: %v", err)
		}
		for _, nn := range n {
			rec = append(rec, nn)
			if nn == int(after[i]-'0') {
				i++
				if i >= len(after) {
					// who even am i
					goto Done
				}
			} else {
				i = 0
			}
		}

		elf1 = move(rec, elf1)
		elf2 = move(rec, elf2)
	}
Done:
	fmt.Println(len(rec) - len(after))
}
