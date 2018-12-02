package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func count(s string, counts map[int]int) {
	seen := map[rune]int{}
	for _, c := range s {
		seen[c] = seen[c] + 1
	}
	for _, v := range seen {
		if v == 2 {
			counts[2]++
			break
		}
	}
	for _, v := range seen {
		if v == 3 {
			counts[3]++
			break
		}
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

	counts := map[int]int{
		2: 0,
		3: 0,
	}

	for _, s := range vals {
		count(s, counts)
	}

	fmt.Println(counts[2] * counts[3])
}
