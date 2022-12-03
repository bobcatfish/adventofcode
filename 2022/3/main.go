package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func getCompartments(s string) []string {
	comps := make([]string, 2)
	comps[0] = s[:len(s)/2]
	comps[1] = s[len(s)/2:]
	return comps
}

func common(s []string) (rune, error) {
	elements := s[0]
	for i := 1; i < len(s); i++ {
		n := []rune{}
		for _, c1 := range elements {
			for _, c2 := range s[i] {
				if c1 == c2 {
					n = append(n, c1)
				}
			}
		}
		elements = string(n)
	}

	if len(elements) == 0 {
		return 0, fmt.Errorf("no common elements")
	}

	for i := 1; i < len(elements); i++ {
		if elements[0] != elements[i] {
			return 0, fmt.Errorf("unexpected common elements %s", elements)
		}
	}

	return rune(elements[0]), nil

}

func getPriority(r rune) int {
	if r >= 97 {
		return int(r) - 96
	}
	return int(r) - 38
}

func load() ([]string, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, fmt.Errorf("couldn't open file: %v", err)
	}
	defer file.Close()

	vals := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		val := scanner.Text()
		vals = append(vals, val)
	}
	return vals, err
}

func main() {
	vals, err := load()
	if err != nil {
		log.Fatalf("Couldn't load from file: %v", err)
	}

	sum := 0
	for _, v := range vals {
		comps := getCompartments(v)
		c, err := common(comps)
		if err != nil {
			log.Fatalf("error finding common item %s %s: %v", comps[0], comps[1], err)
		}
		sum += getPriority(c)
	}
	fmt.Println(sum)

	sum2 := 0
	for i := 0; i < len(vals); i += 3 {
		v := vals[i : i+3]
		c, err := common(v)
		if err != nil {
			log.Fatalf("error finding common item %v %v", v, err)
		}
		sum2 += getPriority(c)
	}
	fmt.Println(sum2)
}
