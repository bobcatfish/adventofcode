package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func load() ([][]rune, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, fmt.Errorf("couldn't open file: %v", err)
	}
	defer file.Close()

	vals := [][]rune{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		val := scanner.Text()
		vals = append(vals, []rune(val))
	}
	return vals, err
}

var chunks = map[rune]rune{
	'(': ')',
	'[': ']',
	'{': '}',
	'<': '>',
}

var scores = map[rune]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

var missingScore = map[rune]int{
	')': 1,
	']': 2,
	'}': 3,
	'>': 4,
}

func getCloses() map[rune]struct{} {
	closes := map[rune]struct{}{}
	for _, close := range chunks {
		closes[close] = struct{}{}
	}
	return closes
}

type stack struct {
	S []rune
}

func (s *stack) pop() (rune, error) {
	if len(s.S) == 0 {
		return 0, fmt.Errorf("empty")
	}
	r := s.S[len(s.S)-1]
	s.S = s.S[:len(s.S)-1]
	return r, nil
}

func (s *stack) push(r rune) {
	s.S = append(s.S, r)
}

func reverse(r []rune) []rune {
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return r
}

func parse(l []rune, closes map[rune]struct{}) (rune, bool, []rune, error) {
	s := stack{}
	for _, r := range l {
		if _, ok := closes[r]; ok {
			expected, err := s.pop()
			if err != nil || expected != r {
				return r, false, nil, nil
			}
		} else {
			close, ok := chunks[r]
			if !ok {
				return 0, false, nil, fmt.Errorf("unexpected char %c", r)
			}
			s.push(close)
		}
	}
	return 0, true, reverse(s.S), nil
}

func main() {
	vals, err := load()
	if err != nil {
		log.Fatalf("Couldn't load nums from file: %v", err)
	}

	closes := getCloses()

	badChar := []rune{}
	incomplete := [][]rune{}
	for _, line := range vals {
		bc, ok, missing, err := parse(line, closes)
		incomplete = append(incomplete, missing)
		if err != nil {
			log.Fatalf("unexpected error for line %s: %v", line, err)
		}
		if !ok {
			badChar = append(badChar, bc)
		}
	}

	sum := 0
	for _, bc := range badChar {
		score, _ := scores[bc]
		sum += score
	}
	fmt.Println(sum)

	scores := []int{}
	for _, m := range incomplete {
		if len(m) > 0 {
			ms := 0
			for _, mm := range m {
				ms *= 5
				ms += missingScore[mm]
			}
			scores = append(scores, ms)
		}
	}
	sort.Ints(scores)
	fmt.Println(scores[len(scores)/2])
}
