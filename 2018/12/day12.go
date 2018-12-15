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
	//numGen = 50000000000
	plant = '#'
	empty = '.'
)

func getInitialState(s string) ([]byte, error) {
	r, err := regexp.Compile("initial state: (.*)")
	if err != nil {
		return []byte{}, fmt.Errorf("couldn't compile initial state regex: %v", err)
	}

	m := r.FindStringSubmatch(s)
	if len(m) != 2 {
		return []byte{}, fmt.Errorf("didn't find expected values in %q", s)
	}

	state := ".." + m[1] + ".."
	stateBytes := []byte{}
	for i := 0; i < len(state); i++ {
		var s string
		if i == 0 {
			s = ".."
			s = s + state[i:i+3]
		} else if i == 1 {
			s = "."
			s = s + state[i-1:i+3]
		} else if i == len(state)-2 {
			s = state[i-2 : i+2]
			s = s + "."
		} else if i == len(state)-1 {
			s = state[i-2 : i+1]
			s = s + ".."
		} else {
			s = state[i-2 : i+3]
		}
		stateBytes = append(stateBytes, patternToByte(s))
	}

	return stateBytes, nil
}

func patternToByte(p string) byte {
	b := byte(0)
	for i := 0; i < len(p); i++ {
		if p[i] == plant {
			b |= byte(1 << uint(i))
		}
	}
	return b
}

func bytesToByte(p []byte) byte {
	b := byte(0)
	for i := 0; i < len(p); i++ {
		b |= byte(p[i] << uint(i))
	}
	return b
}

func getSpread(s string) (byte, byte, error) {
	r, err := regexp.Compile("(.*) => (.*)")
	if err != nil {
		return 0, 0, fmt.Errorf("couldn't compile regex: %v", err)
	}

	m := r.FindStringSubmatch(s)
	if len(m) != 3 {
		return 0, 0, fmt.Errorf("didn't find expected values in %q", s)
	}

	newVal := byte(0)
	if m[2][0] == plant {
		newVal = 1
	}

	return patternToByte(m[1]), newVal, nil
}

func count(s []byte, start int) int {
	count := 0
	for i, r := range s {
		if r == 1 {
			count += start + i
		}
	}
	return count
}

func expandState(state []byte) []byte {
	stateBytes := []byte{}
	for i := 0; i < len(state); i++ {
		var s []byte
		if i == 0 {
			s = []byte{0, 0}
			s = append(s, state[i:i+3]...)
		} else if i == 1 {
			s = []byte{0}
			s = append(s, state[i-1:i+3]...)
		} else if i == len(state)-2 {
			s = state[i-2 : i+2]
			s = append(s, 0)
		} else if i == len(state)-1 {
			s = state[i-2 : i+1]
			s = append(s, 0, 0)
		} else {
			s = state[i-2 : i+3]
		}
		stateBytes = append(stateBytes, bytesToByte(s))
	}
	return stateBytes
}

func mutate(state []byte, spreads map[byte]byte) ([]byte, []byte) {
	newState := []byte{0, 0}

	for i := 0; i < len(state); i++ {
		newState = append(newState, spreads[state[i]])
	}
	newState = append(newState, 0, 0)

	return newState, expandState(newState)
}

func displayBytes(b []byte) {
	for i := 0; i < len(b); i++ {
		if b[i] == byte(1) {
			fmt.Printf("#")
		} else {
			fmt.Printf(".")
		}
	}
	fmt.Println()
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

	spreads := map[byte]byte{}
	for i := 2; i < len(vals); i++ {
		s, r, err := getSpread(vals[i])
		if err != nil {
			log.Fatalf("Couldn't get spread from %q: %v", vals[i], err)
		}
		spreads[s] = r
	}

	state := initial

	start := -2
	var n []byte
	for i := 0; i < numGen; i++ {
		n, state = mutate(state, spreads)
		//displayBytes(n)
		start -= 2
	}

	fmt.Println("start is", start)
	fmt.Println(count(n, start))
}
