package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Couldn't read file: %v\n", err)
	}
	var sum int64
	// Skip linefeed at the end of the input
	n := len(input) - 1
	for i := 0; i < n; i++ {
		//next := (i + 1) % n
		next := (i + (n / 2)) % n

		cInt := int64(input[i] - '0')
		nextInt := int64(input[next] - '0')
		if cInt == nextInt {
			sum += cInt
		}
	}
	fmt.Printf("Total is %d\n", sum)
}
