package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func load() ([]int, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, fmt.Errorf("couldn't open file: %v", err)
	}
	defer file.Close()

	vals := []int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		val := scanner.Text()
		i, err := strconv.Atoi(val)
		if err != nil {
			return nil, err
		}
		vals = append(vals, i)
	}
	return vals, err
}

func countWindow(window int, vals []int) int {
	n := len(vals) - (window - 1)
	sums := make([]int, n)
	for i := 0; i < len(vals); i++ {
		for j := 0; j < window; j++ {
			si := i - j
			if si >= 0 && si < len(sums) {
				sums[i-j] += vals[i]
			}
		}
	}

	c := 0
	for i, v := range sums {
		if i == 0 {
			continue
		}
		if v > sums[i-1] {
			c++
		}
	}

	return c
}

func main() {
	vals, err := load()
	if err != nil {
		log.Fatalf("Couldn't load nums from file: %v", err)
	}

	fmt.Printf("part 1: %d\n", countWindow(1, vals))
	fmt.Printf("part 2: %d\n", countWindow(3, vals))
}
