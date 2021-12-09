package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strings"
)

var num1 = []int{2, 5}
var num7 = []int{0, 2, 5}
var num4 = []int{1, 2, 3, 5}
var num2 = []int{0, 2, 3, 4, 6}
var num3 = []int{0, 2, 3, 5, 6}
var num5 = []int{0, 1, 3, 5, 6}
var num0 = []int{0, 1, 2, 4, 5, 6}
var num6 = []int{0, 1, 3, 4, 5, 6}
var num9 = []int{0, 1, 2, 3, 5, 6}
var num8 = []int{0, 1, 2, 3, 4, 5, 6}
var nums = [][]int{num1, num2, num3, num4, num5, num6, num7, num8, num9}

func load() ([]entry, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, fmt.Errorf("couldn't open file: %v", err)
	}
	defer file.Close()

	vals := []entry{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		val := scanner.Text()
		ss := strings.Fields(val)
		e := entry{
			patterns: [][]rune{},
			outputs:  [][]rune{},
			mapping:  map[rune]rune{},
		}

		for i := 0; i < 10; i++ {
			rr := []rune(ss[i])
			sort.Slice(rr, func(i, j int) bool {
				return rr[i] < rr[j]
			})
			e.patterns = append(e.patterns, rr)
		}
		for i := 11; i < 15; i++ {
			rr := []rune(ss[i])
			sort.Slice(rr, func(i, j int) bool {
				return rr[i] < rr[j]
			})
			e.outputs = append(e.outputs, rr)
		}

		vals = append(vals, e)
	}
	return vals, err
}

type entry struct {
	patterns [][]rune
	outputs  [][]rune
	mapping  map[rune]rune
}

func (e *entry) countVals() int {
	count := 0
	for _, output := range e.outputs {
		lo := len(output)
		if lo == len(num1) || lo == len(num4) || lo == len(num7) || lo == len(num8) {
			count++
		}
	}
	return count
}

func getExtra(a, b []rune) []rune {
	extra := []rune{}

	for _, p := range a {
		seen := false
		for _, pp := range b {
			if p == pp {
				seen = true
			}
		}
		if !seen {
			extra = append(extra, p)
		}
	}
	return extra
}

func (e *entry) getEasy() ([]rune, []rune, []rune, []rune) {
	var p1, p4, p7, p8 []rune
	for _, p := range e.patterns {
		if len(p) == len(num1) {
			p1 = p
		}
		if len(p) == len(num4) {
			p4 = p
		}
		if len(p) == len(num7) {
			p7 = p
		}
		if len(p) == len(num8) {
			p8 = p
		}
	}
	return p1, p4, p7, p8
}

func getMap(poss []rune) map[string]int {
	m := map[string]int{}
	for i, n := range nums {
		r := []rune{}
		for _, nn := range n {
			r = append(r, poss[nn])
		}
		sort.Slice(r, func(i, j int) bool {
			return r[i] < r[j]
		})
		m[string(r)] = i + 1
	}
	return m
}

func (e *entry) getOutputVals(m map[string]int) int {
	v := 0
	for i, o := range e.outputs {
		v += m[string(o)] * int(math.Pow(10, float64(3-i)))
	}
	return v
}

func runeIn(r rune, p []rune) bool {
	for _, pp := range p {
		if pp == r {
			return true
		}
	}
	return false
}

func (e *entry) find(n []int, s []rune) []rune {
	for _, p := range e.patterns {
		if len(p) == len(n) {
			seen := true
			for _, pp := range s {
				if !runeIn(pp, p) {
					seen = false
				}
			}
			if seen {
				return p
			}
		}
	}
	return nil
}

func (e *entry) getPoss() []rune {
	p1, p4, p7, p8 := e.getEasy()
	poss := make([][]rune, 7)

	known := append(p1, p4...)
	known = append(known, p7...)

	poss[0] = getExtra(p7, p1)
	poss[1] = getExtra(p4, p1)
	poss[2] = p1
	poss[3] = poss[1]
	poss[4] = getExtra(p8, known)
	poss[5] = p1
	poss[6] = poss[4]

	p2 := e.find(num2, poss[4])

	unknown := getExtra(p2, poss[4])
	unknown = getExtra(unknown, poss[0])
	poss[3] = getExtra(unknown, poss[2])
	poss[1] = getExtra(poss[1], poss[3])

	p3 := e.find(num3, poss[2])

	unknown = getExtra(p3, poss[0])
	unknown = getExtra(unknown, poss[2])
	poss[6] = getExtra(unknown, poss[3])
	poss[4] = getExtra(poss[4], poss[6])

	p5 := e.find(num5, poss[1])

	unknown = getExtra(p5, poss[0])
	unknown = getExtra(unknown, poss[1])
	unknown = getExtra(unknown, poss[3])
	poss[5] = getExtra(unknown, poss[6])
	poss[2] = getExtra(poss[2], poss[5])

	final := []rune{}
	for _, pp := range poss {
		final = append(final, pp[0])
	}

	return final
}

func main() {
	entries, err := load()
	if err != nil {
		log.Fatalf("Couldn't load nums from file: %v", err)
	}

	sum := 0
	for _, e := range entries {
		count := e.countVals()
		sum += count
	}

	sum = 0
	for _, e := range entries {
		poss := e.getPoss()
		m := getMap(poss)
		ovals := e.getOutputVals(m)
		fmt.Println(ovals)
		sum += ovals
	}
	fmt.Println(sum)
}
