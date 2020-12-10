package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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
		s := scanner.Text()
		i, err := strconv.Atoi(s)
		if err != nil {
			return nil, fmt.Errorf("couldn't convert %s to a num: %v", s, err)
		}
		vals = append(vals, i)
	}
	return vals, err
}

func countDiffs(ii []int) map[int]int {
	diffs := map[int]int{}

	for i, v := range ii {
		if i > 0 {
			diff := v - ii[i-1]
			diffs[diff]++
		}
	}

	return diffs
}

var memos map[string]int

func memo(s string, val int) {
	if memos == nil {
		memos = map[string]int{}
	}

	memos[s] = val
}

func countPoss(prev int, ii []int, curr int) int {
	poss := 1

	for {
		// assuming can't remove built in
		if curr == len(ii)-1 {
			break
		}
		next := ii[curr+1]
		if next-prev <= 3 {
			end := ii[curr+1:]
			s := fmt.Sprintf("%d, %v", prev, end)
			if v, ok := memos[s]; ok {
				poss += v
			} else {
				p := countPoss(prev, end, 0)
				memo(s, p)
				poss += p
			}

		}
		curr++
		prev = ii[curr-1]
	}

	return poss
}

func main() {
	ii, err := load()
	if err != nil {
		log.Fatalf("Couldn't load from file: %v", err)
	}
	sort.Ints(ii)

	// add charging outlet
	ii = append([]int{0}, ii...)

	// add highest adapter
	ii = append(ii, ii[len(ii)-1]+3)

	diffs := countDiffs(ii)
	fmt.Println(diffs)
	fmt.Println(diffs[1] * diffs[3])

	// assuming can't remove charging adapter
	poss := countPoss(ii[0], ii, 1)
	fmt.Println(poss)
}
