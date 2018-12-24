package main

import (
	"fmt"
	"log"
	"strconv"
)

const (
	after = 430971
	num   = 10
)

func addNext(rec []int, elf1, elf2 int) ([]int, error) {
	new := rec[elf1] + rec[elf2]
	s := strconv.Itoa(new)
	for _, r := range s {
		i, err := strconv.Atoi(string(r))
		if err != nil {
			return nil, fmt.Errorf("couldn't convert %c to int: %v", r, err)
		}
		rec = append(rec, i)
	}
	return rec, nil
}

func move(rec []int, i int) int {
	steps := rec[i] + 1
	return (i + steps) % len(rec)
}

func main() {
	elf1, elf2 := 0, 1
	rec := []int{3, 7}
	for {
		var err error
		rec, err = addNext(rec, elf1, elf2)
		if err != nil {
			log.Fatalf("Couldn't determine next value: %v", err)
		}
		elf1 = move(rec, elf1)
		elf2 = move(rec, elf2)
		if len(rec) >= (1 + after + num) {
			break
		}
	}
	for i := after; i < after+num; i++ {
		fmt.Printf("%d", rec[i])
	}
	fmt.Println()
}
