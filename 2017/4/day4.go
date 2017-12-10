package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

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

			// Thanks @BlueMonday!
			r := []rune(w)
			sort.Slice(r, func(i, j int) bool { return r[i] < r[j] })
			s := string(r)

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
