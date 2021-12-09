package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

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
		for _, vv := range val {
			i, err := strconv.Atoi(string(vv))
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

func neighbors(p point, hm [][]int) []point {
	n := []point{
		point{x: p.x, y: p.y - 1},
		point{x: p.x, y: p.y + 1},
		point{x: p.x - 1, y: p.y},
		point{x: p.x + 1, y: p.y},
	}
	fn := []point{}
	for _, p := range n {
		if p.x >= 0 && p.x < len(hm[0]) && p.y >= 0 && p.y < len(hm) {
			fn = append(fn, p)
		}
	}
	return fn
}

func getBasin(p point, hm [][]int) []int {
	basin := []int{hm[p.y][p.x]}
	seen := map[point]struct{}{
		p: struct{}{},
	}
	toSee := neighbors(p, hm)

	for {
		next := []point{}
		for _, n := range toSee {
			if _, ok := seen[n]; !ok {
				seen[n] = struct{}{}
				val := hm[n.y][n.x]
				if val != 9 {
					basin = append(basin, val)
					next = append(next, neighbors(n, hm)...)
				}
			}
		}
		toSee = next

		if len(toSee) == 0 {
			break
		}
	}
	return basin
}

func main() {
	hm, err := load()
	if err != nil {
		log.Fatalf("Couldn't load nums from file: %v", err)
	}

	low := []point{}
	for y := 0; y < len(hm); y++ {
		for x := 0; x < len(hm[y]); x++ {
			p := point{x: x, y: y}
			n := neighbors(p, hm)

			val := hm[y][x]
			isLow := true
			for _, nn := range n {
				if hm[nn.y][nn.x] <= val {
					isLow = false
				}
			}
			if isLow {
				low = append(low, p)
			}
		}
	}
	sum := 0
	for _, l := range low {
		sum += hm[l.y][l.x] + 1
	}
	fmt.Println(sum)

	basins := [][]int{}
	for _, l := range low {
		basin := getBasin(l, hm)
		basins = append(basins, basin)
	}

	sizes := []int{}
	for _, b := range basins {
		sizes = append(sizes, len(b))
	}
	sort.Ints(sizes)
	result := 1
	for _, s := range sizes[len(sizes)-3:] {
		result *= s
	}

	fmt.Println(result)
}
