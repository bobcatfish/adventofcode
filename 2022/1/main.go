package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

type Elf struct {
	Food []int
}

func newElf() *Elf {
	return &Elf{
		Food: []int{},
	}
}

func (e *Elf) Calories() int {
	return sum(e.Food)
}

func sum(vals []int) int {
	sum := 0
	for _, c := range vals {
		sum += c
	}
	return sum
}

func load() ([]*Elf, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, fmt.Errorf("couldn't open file: %v", err)
	}
	defer file.Close()

	elf := newElf()
	elves := []*Elf{elf}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		val := scanner.Text()
		if val == "" {
			elf = newElf()
			elves = append(elves, elf)
		} else {
			i, err := strconv.Atoi(val)
			if err != nil {
				return nil, err
			}
			elf.Food = append(elf.Food, i)
		}
	}
	return elves, err
}

func main() {
	elves, err := load()
	if err != nil {
		log.Fatalf("couldn't get elf vals: %v", err)
	}

	sort.SliceStable(elves, func(i, j int) bool {
		return elves[i].Calories() > elves[j].Calories()
	})

	fmt.Println(elves[0].Calories())
	food := []int{elves[0].Calories(), elves[1].Calories(), elves[2].Calories()}
	fmt.Println(sum(food))
}
