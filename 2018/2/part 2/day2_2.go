package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func measure(s string, m map[string][]string) {
	for i := range s {
		front := s[0:i]
		back := s[i+1:]
		key := fmt.Sprintf("%s_%s", front, back)
		m[key] = append(m[key], s)
	}
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

	m := map[string][]string{}

	for _, s := range vals {
		measure(s, m)
	}

	for k, v := range m {
		if len(v) >= 2 {
			fmt.Println(k)
		}
	}
}
