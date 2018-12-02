package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const bursts = 10000000

type point struct {
	X int
	Y int
}

var up = point{X: 0, Y: -1}
var left = point{X: -1, Y: 0}
var down = point{X: 0, Y: 1}
var right = point{X: 1, Y: 0}
var dirs = [4]point{
	up, left, down, right,
}

type location struct {
	Loc point
	Dir int
}

const clean = 0
const weak = 1
const infect = 2
const flag = 3

func mod(x int, y int) int {
	v := x % y
	if v < 0 {
		v += y
	}
	return v
}

func grid(lines []string) (map[point]int, location) {
	m := make(map[point]int)

	loc := location{
		Loc: point{
			X: (len(lines[0]) / 2),
			Y: (len(lines) / 2),
		},
		Dir: 0,
	}

	for y := 0; y < len(lines); y++ {
		for x := 0; x < len(lines[y]); x++ {
			if lines[y][x] == '#' {
				m[point{X: x, Y: y}] = infect
			}
		}
	}

	return m, loc
}

func burst(m map[point]int, loc location) (int, location) {
	var count int
	var d int

	v, ok := m[loc.Loc]
	if !ok {
		v = clean
	}

	if v == clean {
		d = mod(loc.Dir+1, 4)
		m[loc.Loc] = weak
	} else if v == weak {
		d = loc.Dir
		m[loc.Loc] = infect
		count += 1
	} else if v == infect {
		d = mod(loc.Dir-1, 4)
		m[loc.Loc] = flag
	} else if v == flag {
		d = mod(loc.Dir+2, 4)
		delete(m, loc.Loc)
	}

	loc.Loc.X += dirs[d].X
	loc.Loc.Y += dirs[d].Y
	loc.Dir = d

	return count, loc
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Couldn't open file: %v\n", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	m, loc := grid(lines)

	var infections int
	var c int
	for i := 0; i < bursts; i++ {
		c, loc = burst(m, loc)
		infections += c
	}
	fmt.Println(infections)
}
