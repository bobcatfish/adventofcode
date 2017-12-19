package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Couldn't open file: %v\n", err)
	}
	defer file.Close()

	var instr []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		instrStr := scanner.Text()
		i, err := strconv.Atoi(instrStr)
		if err != nil {
			log.Fatal(err)
		}
		instr = append(instr, i)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	c := 0
	iter := 0
	for ; c < len(instr) && c >= 0; iter++ {
		offset := instr[c]
		if instr[c] >= 3 {
			instr[c]--
		} else {
			instr[c]++
		}
		c = c + offset
	}
	fmt.Println(iter)
}
