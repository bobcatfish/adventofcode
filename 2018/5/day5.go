package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

func react(s string) string {
	for i, r := range s {
		if i == len(s)-1 {
			return s
		}
		reaction := false
		if unicode.IsLower(r) && unicode.IsUpper(rune(s[i+1])) && r == unicode.ToLower(rune(s[i+1])) {
			reaction = true

		} else if unicode.IsUpper(r) && unicode.IsLower(rune(s[i+1])) && r == unicode.ToUpper(rune(s[i+1])) {
			reaction = true
		}

		if reaction {
			front := s[:i]
			back := s[(i + 2):]
			return front + back
		}

	}
	return s
}

func getTypes(s string) []rune {
	types := map[rune]bool{}
	for _, r := range s {
		if _, ok := types[unicode.ToLower(r)]; !ok {
			types[unicode.ToLower(r)] = true
		}
	}
	keys := make([]rune, 0, len(types))
	for k := range types {
		keys = append(keys, k)
	}
	return keys
}

func remove(t rune, s string) string {
	removedS := strings.Replace(s, string(t), "", -1)
	removedS = strings.Replace(removedS, string(unicode.ToUpper(t)), "", -1)
	return removedS
}

func allReactions(p string) int {
	for {
		new := react(p)
		if new == p {
			break
		}
		p = new
	}
	return len(p)
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

	p := vals[0]
	l := allReactions(p)
	fmt.Println("part 1", l)

	types := getTypes(p)

	min := 50000
	for _, t := range types {
		newP := remove(t, p)
		newL := allReactions(newP)
		if newL < min {
			min = newL
		}
	}

	fmt.Println("part 2", min)
}
