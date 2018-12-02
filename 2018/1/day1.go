package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func getNum(s string) (int, error) {
	pos := (s[0] == '+')
	i, err := strconv.Atoi(s[1:])
	if err != nil {
		return 0, fmt.Errorf("couldn't convert string %q: %v", s, err)
	}
	if pos {
		return i, nil
	}
	return 0 - i, nil
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

	seenFreq := map[int]bool{}
	sum := 0
	for {
		for _, v := range vals {
			n, err := getNum(v)
			if err != nil {
				log.Fatalf("Error getting num: %s", err)
			}
			sum += n
			if _, ok := seenFreq[sum]; ok {
				fmt.Println("Repeated", sum)
				os.Exit(0)
			}
			seenFreq[sum] = true
		}
		fmt.Printf("Freqency: %d\n", sum)
	}
}
