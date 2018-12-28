package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const (
	open   = '.'
	tree   = '|'
	lumber = '#'

	rounds = 1000000000
	size   = 50
)

type point struct {
	x int
	y int
}

func neighbors(p point, points map[point]rune) []rune {
	ps := []point{{
		// above
		x: p.x,
		y: p.y - 1,
	}, {
		// above left
		x: p.x - 1,
		y: p.y - 1,
	}, {
		// above right
		x: p.x + 1,
		y: p.y - 1,
	}, {
		// left
		x: p.x - 1,
		y: p.y,
	}, {
		// right
		x: p.x + 1,
		y: p.y,
	}, {
		// below
		x: p.x,
		y: p.y + 1,
	}, {
		// below left
		x: p.x - 1,
		y: p.y + 1,
	}, {
		// below right
		x: p.x + 1,
		y: p.y + 1,
	}}
	rs := []rune{}

	for _, p := range ps {
		if p.x >= 0 && p.x < size && p.y >= 0 && p.y < size {
			r, ok := points[p]
			if ok {
				rs = append(rs, r)
			} else {
				rs = append(rs, open)
			}
		}
	}
	return rs
}

type next func([]rune) rune

func countLumbers(neighbors []rune) int {
	lumbers := 0
	for _, r := range neighbors {
		if r == lumber {
			lumbers++
		}
	}
	return lumbers
}

func countTrees(neighbors []rune) int {
	trees := 0
	for _, r := range neighbors {
		if r == tree {
			trees++
		}
	}
	return trees
}

func nextLumber(neighbors []rune) rune {
	trees := countTrees(neighbors)
	lumbers := countLumbers(neighbors)

	if trees > 0 && lumbers > 0 {
		return lumber
	}
	return open
}

func nextOpen(neighbors []rune) rune {
	trees := countTrees(neighbors)
	if trees >= 3 {
		return tree
	}
	return open
}

func nextTree(neighbors []rune) rune {
	lumbers := countLumbers(neighbors)
	if lumbers >= 3 {
		return lumber
	}
	return tree
}

func getPoints(vals []string) map[point]rune {
	points := map[point]rune{}
	for y, v := range vals {
		for x, r := range v {
			if r != open {
				p := point{x: x, y: y}
				points[p] = r
			}
		}
	}
	return points
}

func count(points map[point]rune) (int, int, int) {
	lumberCount := 0
	treeCount := 0
	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			p := point{x: x, y: y}
			r, ok := points[p]
			if ok {
				if r == lumber {
					lumberCount++
				} else if r == tree {
					treeCount++
				}
			}
		}
	}

	return lumberCount, treeCount, lumberCount * treeCount
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

	nexts := map[rune]next{
		lumber: nextLumber,
		open:   nextOpen,
		tree:   nextTree,
	}
	points := getPoints(vals)
	scores := []int{}

	for i := 0; i < 1000; i++ {
		nextPoints := map[point]rune{}
		for y := 0; y < size; y++ {
			for x := 0; x < size; x++ {
				var next rune

				p := point{x: x, y: y}
				n := neighbors(p, points)
				r, ok := points[p]

				if ok {
					next = nexts[r](n)
				} else {
					next = nexts[open](n)
				}

				if next != open {
					nextPoints[p] = next
				}
			}
		}
		points = nextPoints
		l, t, tot := count(points)
		fmt.Printf("Round %d: %d lumberyards x %d trees = %d\n", i, l, t, tot)
		scores = append(scores, tot)
	}

	// pattern repeats every 28 rounds after some point
	offset := (rounds - 1000) % 28
	fmt.Printf("final score would be: %d\n", scores[999-offset])
}
