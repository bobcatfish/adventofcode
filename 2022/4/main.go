package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func parseNums(ss []string) ([]int, error) {
	nums := []int{}
	for _, sss := range ss {
		i, err := strconv.Atoi(sss)
		if err != nil {
			return nil, fmt.Errorf("couldn't convert %s to int: %v", sss, err)
		}
		nums = append(nums, i)
	}
	return nums, nil
}

func load() ([][]int, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, fmt.Errorf("couldn't open file: %v", err)
	}
	defer file.Close()

	vals := [][]int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		val := scanner.Text()
		re := regexp.MustCompile(`^(\d*)-(\d*),(\d*)-(\d*)$`)
		vv := re.FindStringSubmatch(val)
		if len(vv) != 5 {
			return nil, fmt.Errorf("couldn't parse %s", val)
		}
		nums, err := parseNums(vv[1:])
		if err != nil {
			return nil, err
		}
		vals = append(vals, nums)
	}
	return vals, err
}

func sortPairs(pairs [][]int) {
	for _, pair := range pairs {
		p1, p2 := pair[:2], pair[2:]
		if p1[0] > p2[0] || p1[0] == p2[0] && p1[1] < p2[1] {
			p1[0], p1[1], p2[0], p2[1] = p2[0], p2[1], p1[0], p1[1]
		}
	}
}

func contains(pair []int) bool {
	p1, p2 := pair[:2], pair[2:]
	return p2[1] <= p1[1]
}

func overlap(pair []int) bool {
	p1, p2 := pair[:2], pair[2:]
	return p2[0] <= p1[1]
}

func main() {
	vals, err := load()
	if err != nil {
		log.Fatalf("Couldn't load from file: %v", err)
	}

	sortPairs(vals)

	count := 0
	for _, pair := range vals {
		if contains(pair) {
			count++
		}
	}
	fmt.Println(count)

	count2 := 0
	for _, pair := range vals {
		if overlap(pair) {
			count2++
		}
	}
	fmt.Println(count2)
}
