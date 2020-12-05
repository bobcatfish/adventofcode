package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const NEEDED = 2020

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

func find2(vals []int) ([]int, error) {
	needed := map[int]int{NEEDED - vals[0]: vals[0]}
	for _, v := range vals[1:] {
		if vv, ok := needed[v]; ok {
			return []int{v, vv}, nil
		}
		needed[NEEDED-v] = v
	}
	return []int{}, fmt.Errorf("didn't find them!")
}

func find3(vals []int) ([]int, error) {
	needed := map[int][]int{}
	for i, a := range vals {
		for _, b := range vals[i:] {
			needed[NEEDED-(a+b)] = []int{a, b}
		}
	}
	for _, v := range vals {
		if vv, ok := needed[v]; ok {
			return append(vv, v), nil
		}
	}

	return []int{}, fmt.Errorf("didn't find them!")
}

func main() {
	vals, err := load()
	if err != nil {
		log.Fatalf("Couldn't load nums from file: %v", err)
	}

	a, err := find2(vals)
	if err != nil {
		log.Fatalf("couldn't find 2 values: %v", err)
	}
	fmt.Println("2 vals:", a, a[0]*a[1])

	b, err := find3(vals)
	if err != nil {
		log.Fatalf("couldn't find 2 values: %v", err)
	}
	fmt.Println("3 vals:", b, b[0]*b[1]*b[2])
}
