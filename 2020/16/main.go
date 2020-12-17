package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func load() ([]*Rule, []int, [][]int, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, nil, nil, fmt.Errorf("couldn't open file: %v", err)
	}
	defer file.Close()

	ruleLines := []string{}
	ticketLines := []string{}
	yourLine := ""
	delimCount := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		if len(s) == 0 {
			delimCount++
		} else if delimCount < 5 && delimCount > 0 {
			delimCount++
			if delimCount == 3 {
				yourLine = s
			}
		} else if delimCount == 5 {
			ticketLines = append(ticketLines, s)
		} else {
			ruleLines = append(ruleLines, s)
		}

	}

	rules, err := parseRules(ruleLines)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("couldn't parse rules: %v", err)
	}
	your, err := parseTicket(yourLine)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("couldn't parse your ticket %s: %v", yourLine, err)
	}
	tickets, err := parseTickets(ticketLines)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("couldn't parse tickets: %v", err)
	}

	return rules, your, tickets, nil
}

type Rule struct {
	Name   string
	Ranges []int
}

func (rule *Rule) Valid(t int) bool {
	return (t >= rule.Ranges[0] && t <= rule.Ranges[1]) ||
		(t >= rule.Ranges[2] && t <= rule.Ranges[3])
}

func parseRules(r []string) ([]*Rule, error) {
	rules := []*Rule{}
	for _, rr := range r {
		re := regexp.MustCompile(`^(.+): (\d+)-(\d+) or (\d+)-(\d+)$`)
		vv := re.FindStringSubmatch(rr)
		if len(vv) != 6 {
			return nil, fmt.Errorf("unexpected rule %q split into %d pieces %v", rr, len(vv), vv)
		}
		ranges := []int{}

		for _, is := range vv[2:] {
			i, err := strconv.Atoi(is)
			if err != nil {
				return nil, fmt.Errorf("unexpected range value %s could not be converted: %v", is, err)
			}
			ranges = append(ranges, i)
		}

		rules = append(rules, &Rule{Name: vv[1], Ranges: ranges})
	}
	return rules, nil
}

func parseTicket(tt string) ([]int, error) {
	s := strings.Split(tt, ",")
	t := []int{}
	for _, ss := range s {
		i, err := strconv.Atoi(ss)
		if err != nil {
			return nil, fmt.Errorf("couldn't convert ticket value %s to int: %v", ss, err)
		}
		t = append(t, i)
	}
	return t, nil
}

func parseTickets(t []string) ([][]int, error) {
	tickets := [][]int{}
	for _, tt := range t {
		t, err := parseTicket(tt)
		if err != nil {
			return nil, fmt.Errorf("couldn't parse %s, %v", tt, err)
		}
		tickets = append(tickets, t)
	}
	return tickets, nil
}

func filterViolations(rules []*Rule, tickets [][]int) ([]int, [][]int) {
	v := []int{}
	good := [][]int{}

	for _, row := range tickets {
		okt := true
		for _, t := range row {
			ok := false
			for _, rule := range rules {
				if rule.Valid(t) {
					ok = true
				}
			}
			if !ok {
				v = append(v, t)
				okt = false
			}
		}
		if okt {
			good = append(good, row)
		}
	}
	return v, good
}

func d(p []*Rule, r *Rule) []*Rule {
	index := -1
	for i, pp := range p {
		if pp == r {
			index = i
			break
		}
	}
	if index != -1 {
		return append(p[:index], p[index+1:]...)
	}
	return p
}

func reduce(poss map[int][]*Rule) {
	for {
		var remove *Rule
		removei := -1
		some := false

		for i, p := range poss {
			if len(p) == 1 {
				remove = p[0]
				removei = i
			} else {
				some = true
			}
		}
		if !some {
			fmt.Println(poss)
			return
		}
		if remove != nil {
			for i, p := range poss {
				if i != removei {
					poss[i] = d(p, remove)
				}
			}
		}
	}
}

func findOrder(rules []*Rule, tickets [][]int) []*Rule {
	poss := map[int][]*Rule{}

	for _, r := range rules {
		for i := range tickets[0] {
			poss[i] = append(poss[i], r)
		}
	}

	for _, row := range tickets {
		for i, t := range row {
			n := []*Rule{}
			p := poss[i]

			for _, pp := range p {
				if pp.Valid(t) {
					n = append(n, pp)
				}
			}
			poss[i] = n
		}
	}

	reduce(poss)

	ordered := make([]*Rule, len(rules))

	for i, rs := range poss {
		ordered[i] = rs[0]
	}
	for i, rs := range ordered {
		fmt.Println(i, rs.Name)
	}

	return ordered
}

func main() {
	rules, your, tickets, err := load()
	if err != nil {
		log.Fatalf("Couldn't load from file: %v", err)
	}
	v, tickets := filterViolations(rules, tickets)
	fmt.Println(v)

	c := 0
	for _, vv := range v {
		c += vv
	}
	fmt.Println(c)

	orderedRules := findOrder(rules, tickets)

	value := 1
	for i, rule := range orderedRules {
		if strings.HasPrefix(rule.Name, "departure") {
			value *= your[i]
		}
	}
	fmt.Println(value)
}
