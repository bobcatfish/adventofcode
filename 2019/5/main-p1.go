package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func getOp(i int) (int, []int) {
	modes := make([]int, 3)

	modes[2] = i / 10000 % 10
	modes[1] = i / 1000 % 10
	modes[0] = i / 100 % 10
	op := i % 100

	return op, modes
}

func doOp(vals []int, i int, input int) (int, []int) {
	op, modes := getOp(vals[i])

	switch op {
	case 1:
	case 2:
	case 3:
	case 4:
	default:
		return -1, vals
	}

	val3 := func(vals []int, modes []int, i int) (int, int, int) {
		var x, y int
		if modes[0] == 1 {
			x = vals[i+1]
		} else {
			x = vals[vals[i+1]]
		}
		if modes[1] == 1 {
			y = vals[i+2]
		} else {
			y = vals[vals[i+2]]
		}
		newLoc := vals[i+3]
		return x, y, newLoc
	}
	val1 := func(vals []int, modes []int, i int) int {
		var reg int
		if modes[0] == 1 {
			reg = vals[i+1]
		} else {
			reg = vals[vals[i+1]]
		}
		return reg
	}

	inc := 0
	switch v := op; v {
	case 1:
		x, y, newLoc := val3(vals, modes, i)
		vals[newLoc] = x + y
		inc = 4
	case 2:
		x, y, newLoc := val3(vals, modes, i)
		vals[newLoc] = x * y
		inc = 4
	case 3:
		//reg := val1(vals, modes, i)
		reg := vals[i+1]
		vals[reg] = input
		inc = 2
	case 4:
		reg := val1(vals, modes, i)
		fmt.Println("output:", reg)
		inc = 2
	default:
		return -1, vals
	}
	return i + inc, vals
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
		v := strings.Split(val, ",")
		for _, vv := range v {
			i, err := strconv.Atoi(vv)
			if err != nil {
				return nil, err
			}
			vals = append(vals, i)
		}
	}
	return vals, err
}

func main() {
	vals, err := loadNums()
	if err != nil {
		log.Fatalf("Couldn't load nums from file: %v", err)
	}

	input := 1

	i := 0
	for {
		i, vals = doOp(vals, i, input)
		if i < 0 {
			break
		}
	}
	//fmt.Println(vals)
}
