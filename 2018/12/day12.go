package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

const (
	numGen = 20
	plant  = '#'
	empty  = '.'
)

func getInitialState(s string) (string, error) {
	r, err := regexp.Compile("initial state: (.*)")
	if err != nil {
		return "", fmt.Errorf("couldn't compile initial state regex: %v", err)
	}

	m := r.FindStringSubmatch(s)
	if len(m) != 2 {
		return "", fmt.Errorf("didn't find expected values in %q", s)
	}

	return m[1], nil
}

func getSpread(s string) (string, rune, error) {
	r, err := regexp.Compile("(.*) => (.*)")
	if err != nil {
		return "", ' ', fmt.Errorf("couldn't compile regex: %v", err)
	}

	m := r.FindStringSubmatch(s)
	if len(m) != 3 {
		return "", ' ', fmt.Errorf("didn't find expected values in %q", s)
	}

	return m[1], rune(m[2][0]), nil
}

func count(s string, start int) int {
	count := 0
	for i, r := range s {
		if r == plant {
			count += start + i
		}
	}
	return count
}

func mutate(state string, spreads map[string]rune) string {
	state = fmt.Sprintf("%c%c%c%c%s%c%c%c%c", empty, empty, empty, empty, state, empty, empty, empty, empty)
	newState := ""
	for i := 2; i < len(state)-2; i++ {
		s := state[i-2 : i+3]
		if r, ok := spreads[s]; ok != false {
			newState += string(r)
		} else {
			newState += string(r)
		}
	}
	return newState
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Couldn't open file: %v", err)
	}
	defer file.Close()

	vals := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		vals = append(vals, scanner.Text())
	}

	initial, err := getInitialState(vals[0])
	if err != nil {
		log.Fatalf("Couldn't get initial state: %v", err)
	}
	fmt.Println(initial)

	spreads := map[string]rune{}
	for i := 2; i < len(vals); i++ {
		// TODO
		s, r, err := getSpread(vals[i])
		if err != nil {
			log.Fatalf("Couldn't get spread from %q: %v", vals[i], err)
		}
		spreads[s] = r
	}

	state := initial
	start := 0
	for i := 0; i < numGen; i++ {
		state = mutate(state, spreads)
		fmt.Println(state)
		start -= 2
	}

	fmt.Println("start is", start)
	fmt.Println(count(state, start))
}
