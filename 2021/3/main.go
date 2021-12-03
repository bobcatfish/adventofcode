package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func load() ([]string, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, fmt.Errorf("couldn't open file: %v", err)
	}
	defer file.Close()

	v := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		val := scanner.Text()
		v = append(v, val)

	}
	return v, err
}

func common(i int, vals []string, o int, most bool) int {
	c0, c1 := 0, 0
	for j := 0; j < len(vals); j++ {
		v := vals[j][i]
		if v == 48 {
			c0++
		} else {
			c1++
		}
	}
	if c0 == c1 {
		return o
	}
	if most && c0 > c1 {
		return 0
	} else if most {
		return 1
	} else if !most && c0 < c1 {
		return 0
	}
	return 1
}

func elim(i, v int, vals []string) []string {
	n := []string{}

	for j := 0; j < len(vals); j++ {
		match := vals[j][i] == (byte(v) + byte('0'))
		if match || (len(n) == 0 && j == len(vals)-1) {
			n = append(n, vals[j])
		}
	}

	return n
}

func bin2ints(digits []int) (int, int) {
	gamma, epsilon := 0, 0
	for i, d := range digits {
		if d == 1 {
			gamma += (1 << (uint(len(digits)) - 1 - uint(i)))
		} else {
			epsilon += (1 << (uint(len(digits)) - 1 - uint(i)))
		}

	}
	return gamma, epsilon
}

func main() {
	vals, err := load()
	if err != nil {
		log.Fatalf("Couldn't load nums from file: %v", err)
	}

	digits := make([]int, len(vals[0]))
	for i := 0; i < len(vals[0]); i++ {
		digits[i] = common(i, vals, -1, true)
	}

	gamma, epsilon := bin2ints(digits)

	fmt.Println(gamma, epsilon, gamma*epsilon)

	os := make([]string, len(vals))
	copy(os, vals)

	for i := 0; i < len(vals[0]); i++ {
		mc := common(i, os, 1, true)
		os = elim(i, mc, os)
	}

	cs := make([]string, len(vals))
	copy(cs, vals)

	for i := 0; i < len(vals[0]); i++ {
		lc := common(i, cs, 0, false)
		cs = elim(i, lc, cs)
	}

	osb, _ := strconv.ParseInt(os[0], 2, 64)
	csb, _ := strconv.ParseInt(cs[0], 2, 64)

	fmt.Println(osb, csb, osb*csb)
}
