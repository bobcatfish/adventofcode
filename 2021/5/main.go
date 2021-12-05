package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x int
	y int
}

type line struct {
	start point
	end   point
}

func parseCoor(s string) (point, error) {
	ss := strings.Split(s, ",")
	p := point{}
	var err error
	p.x, err = strconv.Atoi(ss[0])
	if err != nil {
		return p, err
	}
	p.y, err = strconv.Atoi(ss[1])
	if err != nil {
		return p, err
	}
	return p, nil
}

func isDiagonal(p1, p2 point) bool {
	return p1.x != p2.x && p1.y != p2.y
}

func load() ([]line, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, fmt.Errorf("couldn't open file: %v", err)
	}
	defer file.Close()

	lines := []line{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		val := scanner.Text()
		ss := strings.Fields(val)
		start, err := parseCoor(ss[0])
		if err != nil {
			return nil, err
		}
		end, err := parseCoor(ss[2])
		if err != nil {
			return nil, err
		}
		lines = append(lines, line{start: start, end: end})
	}
	return lines, err
}

func cover(covered map[point]int, l line) {
	curr := l.start
	p := point{x: curr.x, y: curr.y}
	covered[p] = covered[p] + 1

	for {
		if l.end.x != curr.x {
			if l.end.x > curr.x {
				curr.x += 1
			} else {
				curr.x -= 1
			}
		}

		if l.end.y != curr.y {
			if l.end.y > curr.y {
				curr.y += 1
			} else {
				curr.y -= 1
			}
		}
		p := point{x: curr.x, y: curr.y}

		covered[p] = covered[p] + 1
		if curr.x == l.end.x && curr.y == l.end.y {
			break
		}
	}
}

func count(covered map[point]int) int {
	v := 0
	for _, c := range covered {
		if c >= 2 {
			v++
		}
	}
	return v
}

func main() {
	lines, err := load()
	if err != nil {
		log.Fatalf("Couldn't load nums from file: %v", err)
	}
	covered := map[point]int{}
	for _, l := range lines {
		if !isDiagonal(l.start, l.end) {
			cover(covered, l)
		}
	}
	v := count(covered)
	fmt.Println(v)

	covered = map[point]int{}
	for _, l := range lines {
		cover(covered, l)
	}
	v = count(covered)
	fmt.Println(v)
}
