package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const bursts = 10000

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

func mod(x int, y int) int {
	v := x % y
	if v < 0 {
		v += y
	}
	return v
}

func grid(lines []string) (map[point]bool, location) {
	m := make(map[point]bool)

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
				m[point{X: x, Y: y}] = true
			}
		}
	}

	return m, loc
}

func burst(m map[point]bool, loc location) (int, location) {
	var count int
	var d int

	if _, ok := m[loc.Loc]; ok {
		d = mod(loc.Dir-1, 4)
		delete(m, loc.Loc)
	} else {
		d = mod(loc.Dir+1, 4)
		m[loc.Loc] = true
		count += 1
	}

	fmt.Println("turning", d)

	loc.Loc.X += dirs[d].X
	loc.Loc.Y += dirs[d].Y
	loc.Dir = d

	fmt.Println(loc)

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
	fmt.Println(loc)

	var infections int
	var c int
	for i := 0; i < bursts; i++ {
		c, loc = burst(m, loc)
		infections += c
	}
	fmt.Println(infections)
}
