package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var steps = 100

func load() ([][]int, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, fmt.Errorf("couldn't open file: %v", err)
	}
	defer file.Close()

	vals := [][]int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		val := scanner.Text()
		row := []int{}
		for _, s := range val {
			i, err := strconv.Atoi(string(s))
			if err != nil {
				return nil, err
			}
			row = append(row, i)
		}
		vals = append(vals, row)
	}
	return vals, err
}

type point struct {
	x int
	y int
}

func flashNeighbors(p point, os [][]int) {
	n := []point{
		{x: p.x, y: p.y + 1},
		{x: p.x, y: p.y - 1},
		{x: p.x + 1, y: p.y},
		{x: p.x - 1, y: p.y},
		{x: p.x - 1, y: p.y - 1},
		{x: p.x + 1, y: p.y + 1},
		{x: p.x - 1, y: p.y + 1},
		{x: p.x + 1, y: p.y - 1},
	}

	for _, nn := range n {
		if nn.y >= 0 && nn.y < len(os) && nn.x >= 0 && nn.x < len(os[nn.y]) {
			os[nn.y][nn.x] += 1
		}
	}
}

func getFlashes(os [][]int) int {
	for y := 0; y < len(os); y++ {
		for x := 0; x < len(os[y]); x++ {
			os[y][x] += 1
		}
	}
	flashed := map[point]struct{}{}
	flashes := 0
	for {
		fCount := 0
		for y := 0; y < len(os); y++ {
			for x := 0; x < len(os[y]); x++ {
				p := point{x: x, y: y}
				if os[y][x] > 9 {
					if _, ok := flashed[p]; !ok {
						flashed[p] = struct{}{}
						fCount++
						flashNeighbors(p, os)
					}
				}
			}
		}
		if fCount == 0 {
			break
		}
		flashes += fCount
	}

	for p, _ := range flashed {
		os[p.y][p.x] = 0
	}

	return flashes
}

func display(os [][]int) {
	for y := 0; y < len(os); y++ {
		fmt.Println(os[y])
	}
}

func main() {
	os, err := load()
	if err != nil {
		log.Fatalf("Couldn't load nums from file: %v", err)
	}

	flashes := 0
	i := 0
	for {
		f := getFlashes(os)
		i++
		if f == (len(os) * len(os[0])) {
			fmt.Println("all at", i)
			break
		}
		flashes += f
		if i == steps {
			fmt.Println(flashes)
		}
	}
}
