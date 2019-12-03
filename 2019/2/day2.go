package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func doOp(vals []int, i int) (int, []int) {
	switch v := vals[i]; v {
	case 1:
	case 2:
	default:
		return -1, vals
	}
	xLoc, yLoc := vals[i+1], vals[i+2]
	newLoc := vals[i+3]
	x, y := vals[xLoc], vals[yLoc]
	switch v := vals[i]; v {
	case 1:
		vals[newLoc] = x + y
	case 2:
		vals[newLoc] = x * y
	default:
		return -1, vals
	}
	return i + 4, vals
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
	valsBackup := make([]int, len(vals))
	copy(valsBackup, vals)

	for k := 0; k <= 99; k++ {
		for j := 0; j < 99; j++ {
			copy(vals, valsBackup)
			vals[1] = k
			vals[2] = j
			i := 0
			for {
				i, vals = doOp(vals, i)
				if i < 0 {
					break
				}
			}
			if vals[0] == 19690720 {
				fmt.Println("WINNER", k, j)
				fmt.Println((100 * k) + j)
				os.Exit(0)
			}
		}
	}
	fmt.Println("END")
}
