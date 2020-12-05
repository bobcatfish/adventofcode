package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type Rule struct {
	Letter rune
	Min    int
	Max    int
}

func (r Rule) Valid(s string) bool {
	c := 0
	for _, v := range s {
		if v == r.Letter {
			c++
		}
	}
	return c >= r.Min && c <= r.Max
}

func (r Rule) Valid2(s []rune) bool {
	i1, i2 := r.Min-1, r.Max-1
	if i1 >= len(s) || i2 >= len(s) {
		return false
	}
	c1, c2 := s[i1], s[i2]
	match1, match2 := c1 == r.Letter, c2 == r.Letter

	return (match1 || match2) && !(match1 && match2)
}

type Password struct {
	P string
	R Rule
}

func (p Password) Valid() bool {
	return p.R.Valid(p.P)
}

func (p Password) Valid2() bool {
	return p.R.Valid2([]rune(p.P))
}

func getPassword(s string) (Password, error) {
	// 9-14 h: hhhhhhhhhhhhhchh
	re := regexp.MustCompile(`(\d+)-(\d+) (\w): (\w+)`)
	segs := re.FindStringSubmatch(s)
	if len(segs) != 5 {
		return Password{}, fmt.Errorf("couldn't parse password and rule from %s, only %d elements matched", s, len(segs))
	}

	min, err := strconv.Atoi(segs[1])
	if err != nil {
		return Password{}, fmt.Errorf("couldn't parse min number from %s", segs[1])
	}
	max, err := strconv.Atoi(segs[2])
	if err != nil {
		return Password{}, fmt.Errorf("couldn't parse max number from %s", segs[2])
	}

	return Password{
		P: segs[4],
		R: Rule{
			Letter: []rune(segs[3])[0],
			Min:    min,
			Max:    max,
		},
	}, nil
}

func load() ([]Password, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, fmt.Errorf("couldn't open file: %v", err)
	}
	defer file.Close()

	pp := []Password{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		val := scanner.Text()

		p, err := getPassword(val)
		if err != nil {
			return nil, fmt.Errorf("couldn't parse password and rule from %s: %v", val, err)
		}
		pp = append(pp, p)
	}

	return pp, err
}

func main() {
	pp, err := load()
	if err != nil {
		log.Fatalf("Couldn't load nums from file: %v", err)
	}

	rule1, rule2 := 0, 0
	for _, p := range pp {
		// fmt.Printf("%d-%d %c %s\n", p.R.Min, p.R.Max, p.R.Letter, p.P)
		if p.Valid() {
			rule1++
		}
		if p.Valid2() {
			rule2++
		}
	}
	fmt.Printf("%d passed rule1, %d passed rule2,out of %d\n", rule1, rule2, len(pp))

}
