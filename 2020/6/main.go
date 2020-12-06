package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func load() ([][]string, error) {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		return nil, fmt.Errorf("couldn't read file: %v", err)
	}

	pp := [][]string{}
	for _, v := range strings.Split(string(data), "\n\n") {
		pp = append(pp, strings.Fields(v))
	}
	return pp, nil
}

func getCounts1(a [][]string) []map[rune]struct{} {
	mm := []map[rune]struct{}{}
	for _, ss := range a {
		m := map[rune]struct{}{}
		for _, s := range ss {
			for _, r := range s {
				m[r] = struct{}{}
			}
		}
		mm = append(mm, m)
	}
	return mm
}

func getCounts2(a [][]string) []map[rune]struct{} {
	mm := []map[rune]struct{}{}
	for _, ss := range a {
		m := map[rune]struct{}{}
		for i, s := range ss {
			if i == 0 {
				for _, r := range s {
					m[r] = struct{}{}
				}
			} else {
				for r, _ := range m {
					found := false
					for _, rr := range s {
						if r == rr {
							found = true
							break
						}
					}
					if !found {
						delete(m, r)
					}
				}
			}
		}
		mm = append(mm, m)
	}
	return mm
}

func sum(counts []map[rune]struct{}) int {
	c := 0
	for _, m := range counts {
		c += len(m)
	}
	return c
}

func main() {
	a, err := load()
	if err != nil {
		log.Fatalf("Couldn't load nums from file: %v", err)
	}
	counts1 := getCounts1(a)
	fmt.Println(sum(counts1))

	counts2 := getCounts2(a)
	fmt.Println(sum(counts2))
}
