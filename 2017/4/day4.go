package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

// Copied from https://stackoverflow.com/questions/22688651/golang-how-to-sort-string-or-byte
func sortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Couldn't open file: %v\n", err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	count := 0

	for _, row := range lines {
		words := make(map[string]bool)
		valid := true
		for _, w := range strings.Fields(row) {
			s := sortString(w)
			if _, present := words[s]; present {
				valid = false
				break
			} else {
				words[s] = true
			}
		}
		if valid {
			count++
		}
	}
	fmt.Println("Valid passphrases: ", count)
}
