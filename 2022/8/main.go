package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func load() [][]int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("couldn't open file: %v", err)
	}
	defer file.Close()

	vals := [][]int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		val := scanner.Text()
		v := []int{}
		for _, s := range val {
			i, err := strconv.Atoi(string(s))
			if err != nil {
				log.Fatalf("couldn't convert int: %v", err)
			}
			v = append(v, i)
		}
		vals = append(vals, v)
	}
	return vals
}

func isVisible(vals [][]int, y, x int) bool {
	if y == 0 || x == 0 || y == len(vals)-1 || x == len(vals[0])-1 {
		return true
	}
	vis := []bool{true, true, true, true}
	height := vals[y][x]
	for yy := 0; yy < len(vals); yy++ {
		if yy == y {
			continue
		}
		if vals[yy][x] >= height {
			if yy < y {
				vis[0] = false
			} else {
				vis[1] = false
			}
		}
	}
	for xx := 0; xx < len(vals[0]); xx++ {
		if xx == x {
			continue
		}
		if vals[y][xx] >= height {
			if xx < x {
				vis[2] = false
			} else {
				vis[3] = false
			}
		}
	}
	for _, vis := range vis {
		if vis == true {
			return true
		}
	}
	return false
}

func getScore(vals [][]int, y, x int) int {
	vis := []int{0, 0, 0, 0}
	height := vals[y][x]
	for yy := y - 1; yy >= 0; yy-- {
		vis[0]++
		if vals[yy][x] >= height {
			break
		}
	}
	for yy := y + 1; yy < len(vals); yy++ {
		vis[1]++
		if vals[yy][x] >= height {
			break
		}
	}
	for xx := x - 1; xx >= 0; xx-- {
		vis[2]++
		if vals[y][xx] >= height {
			break
		}
	}
	for xx := x + 1; xx < len(vals); xx++ {
		vis[3]++
		if vals[y][xx] >= height {
			break
		}
	}
	score := 1
	for _, v := range vis {
		score *= v
	}
	return score
}

func main() {
	vals := load()

	count := 0
	for y := 0; y < len(vals); y++ {
		for x := 0; x < len(vals[y]); x++ {
			if isVisible(vals, y, x) {
				count++
			}
		}
	}
	fmt.Println(count)

	max := 0
	for y := 0; y < len(vals); y++ {
		for x := 0; x < len(vals[y]); x++ {
			score := getScore(vals, y, x)
			if score > max {
				max = score
			}
		}
	}
	fmt.Println("max", max)
}
