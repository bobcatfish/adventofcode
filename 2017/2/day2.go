package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func checksum1(row string) int64 {
	largest := int64(0)
	smallest := int64(^uint(0) >> 1)
	for _, n := range strings.Fields(row) {
		v, err := strconv.ParseInt(n, 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		if v < smallest {
			smallest = v
		}
		if v > largest {
			largest = v
		}
	}
	return largest - smallest
}

func checksum2(row string) int64 {
	var nums []int64
	for _, n := range strings.Fields(row) {
		v, err := strconv.ParseInt(n, 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		nums = append(nums, v)
	}

	// n squared, i'm crying
	for i, n := range nums {
		for j := i + 1; j < len(nums); j++ {
			if n%nums[j] == 0 {
				return n / nums[j]
			}
			if nums[j]%n == 0 {
				return nums[j] / n
			}
		}
	}
	log.Fatalf("Didn't find divisible values in %v", row)
	return 0
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

	var sum int64

	for _, row := range lines {
		sum += checksum2(row)
	}
	fmt.Printf("Total is %d\n", sum)
}
