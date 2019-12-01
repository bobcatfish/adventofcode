package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func getMass(i int) int {
	return (i / 3) - 2
}

func getFuel(m int) int {
	sum := 0
	n := m

	for {
		n = getMass(n)
		if n <= 0 {
			break
		}
		sum += n
	}
	return sum
}

func loadNums() ([]int, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, fmt.Errorf("couldn't open file: %v", err)
	}
	defer file.Close()

	vals := []int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		val := scanner.Text()
		i, err := strconv.Atoi(val)
		if err != nil {
			return nil, fmt.Errorf("couldn't convert %s to int", i)
		}
		vals = append(vals, i)
	}
	return vals, err
}

func main() {
	vals, err := loadNums()
	if err != nil {
		log.Fatalf("Couldn't load nums from file: %v", err)
	}

	masses := make([]int, len(vals))
	sum := 0
	for i := range vals {
		//masses[i] = getMass(vals[i])
		masses[i] = getFuel(vals[i])
		sum += masses[i]
	}

	fmt.Println(sum)
}
