package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func load() []int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("couldn't open file: %v", err)
	}
	defer file.Close()

	vals := []int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		val := scanner.Text()
		i, err := strconv.Atoi(val)
		if err != nil {
			log.Fatalf("couldn't convert int: %v", err)
		}
		vals = append(vals, i)
	}
	return vals
}

func main() {
	vals := load()
	fmt.Println(vals)
}
