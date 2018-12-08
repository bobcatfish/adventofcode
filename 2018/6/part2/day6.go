package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	gridSize     = 1000
	boundaryDist = 10000
)

type point struct {
	x int
	y int
}

func count(distances []int) int {
	c := 0
	for _, d := range distances {
		if d < boundaryDist {
			c++
		}
	}
	return c
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func distance(p1, p2 point) int {
	return abs(p1.x-p2.x) + abs(p1.y-p2.y)
}

func getDist(p point, starts []point) int {
	d := 0
	for _, startP := range starts {
		d += distance(p, startP)
	}
	return d
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Couldn't open file: %v", err)
	}
	defer file.Close()

	vals := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		vals = append(vals, scanner.Text())
	}

	starts := []point{}
	for _, v := range vals {
		vs := strings.Split(v, ", ")
		x, err := strconv.Atoi(vs[0])
		if err != nil {
			log.Fatalf("Error converting %q to point: %v", v, err)
		}
		y, err := strconv.Atoi(vs[1])
		if err != nil {
			log.Fatalf("Error converting %q to point: %v", v, err)
		}
		startPoint := point{
			x: x,
			y: y,
		}
		starts = append(starts, startPoint)
	}

	distances := []int{}
	for y := 0; y < gridSize; y++ {
		for x := 0; x < gridSize; x++ {
			p := point{x: x, y: y}
			d := getDist(p, starts)
			distances = append(distances, d)
		}
	}

	c := count(distances)
	fmt.Println(c)
}
