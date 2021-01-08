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

func load() ([]*Rule, []string, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, nil, fmt.Errorf("couldn't open file: %v", err)
	}
	defer file.Close()

	ruleLines := []string{}
	messageLines := []string{}
	delimCount := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		if len(s) == 0 {
			delimCount++
		} else if delimCount == 1 {
			messageLines = append(messageLines, s)
		} else {
			ruleLines = append(ruleLines, s)
		}

	}

	rules, err := parseRules(ruleLines)
	if err != nil {
		return nil, nil, fmt.Errorf("couldn't parse rules: %v", err)
	}

	return rules, messageLines, nil
}

func parseOptions(s string) ([][]int, error) {
	o := [][]int{}
	ss := strings.Split(s, "|")
	for _, sss := range ss {
		oo := []int{}
		f := strings.Fields(sss)
		for _, ff := range f {
			i, err := strconv.Atoi(ff)
			if err != nil {
				return nil, fmt.Errorf("could not parse option from %s: %v", ff, err)
			}
			oo = append(oo, i)
		}
		o = append(o, oo)
	}
	return o, nil
}

func parseRules(l []string) ([]*Rule, error) {
	r := make([]*Rule, 131)

	for _, ll := range l {
		var idstr string
		rule := &Rule{}
		re := regexp.MustCompile(`^(.+): "(\w)"$`)
		vv := re.FindStringSubmatch(ll)
		if len(vv) != 3 {
			re := regexp.MustCompile(`^(.+): (.*)$`)
			vv := re.FindStringSubmatch(ll)
			if len(vv) != 3 {
				return nil, fmt.Errorf("couldn't parse either kind of rule from %s", ll)
			}
			idstr = vv[1]
			o, err := parseOptions(vv[2])
			if err != nil {
				return nil, fmt.Errorf("couldn't parse options from %v: %v", vv[2], err)
			}
			rule.Options = o
		} else {
			idstr = vv[1]
			rule.Val = rune(vv[2][0])
		}

		i, err := strconv.Atoi(idstr)
		if err != nil {
			return nil, fmt.Errorf("couldn't convert id to int %s: %v", idstr, err)
		}
		rule.Id = i
		r[i] = rule
	}

	return r, nil
}

type Rule struct {
	Options [][]int
	Val     rune
	Id      int
}

func isMatch(rules []*Rule, id int, t string, i int) (bool, []int) {
	rule := rules[id]
	fmt.Printf("checking %c, %v, %s, %d\n", rule.Val, rule.Options, t, i)

	if i >= len(t) {
		// !!!!!!!!!!!1
		return false, []int{i}
	}

	if rule.Val != 0 {
		if rune(t[i]) != rule.Val {
			fmt.Printf("failed %d: %c %c\n", id, rune(t[i]), rule.Val)
			return false, []int{i + 1}
		} else {
			fmt.Printf("match %c %c\n", rune(t[i]), rule.Val)
			return true, []int{i + 1}
		}

	} else {
		fmt.Println("looking at options, i is", i)

		matches := []int{}
		for _, option := range rule.Options {
			fmt.Println("about to check option", option)
			match := true
			iii := []int{i}
			for _, o := range option {
				niii := []int{}
				for _, ii := range iii {
					var ok bool
					fmt.Println("loop option", o, t, ii)
					ok, nii := isMatch(rules, o, t, ii)
					if ok {
						niii = append(niii, nii...)
					}
				}
				iii = niii
			}
			if match {
				matches = append(matches, iii...)
			}
		}
		if len(matches) > 0 {
			return true, matches
		}

	}

	fmt.Printf("failed %d:  %c, %v, %s, %d\n", id, rule.Val, rule.Options, t, i)
	return false, nil
}

func doit(rules []*Rule, tickets []string) {
	count := 0
	matches := []string{}
	for _, t := range tickets {
		if ok, iii := isMatch(rules, 0, t, 0); ok {
			for _, i := range iii {
				if i == len(t) {
					count++
					matches = append(matches, t)
				}
			}
		}
	}
	fmt.Println("COUNT", count)
	for _, m := range matches {
		fmt.Println(m)
	}
}

func main() {
	rules, tickets, err := load()
	if err != nil {
		log.Fatalf("Couldn't load from file: %v", err)
	}
	//doit(rules, tickets)

	rules[8] = &Rule{
		Options: [][]int{
			{42},
			{42, 8},
		},
	}
	rules[11] = &Rule{
		Options: [][]int{
			{42, 31},
			{42, 11, 31},
		},
	}
	doit(rules, tickets)
}
