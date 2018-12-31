package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type node struct {
	r    rune
	next []*node
}

func pop(s []*node) (*node, []*node) {
	return look(s), s[len(s)-1:]
}

func look(s []*node) *node {
	return s[len(s)-1]
}

func push(s []*node, n *node) []*node {
	return append(s, n)
}

func getPaths(s string, curr *node, i int) ([]*node, int) {
	root := curr
	currs := []*node{}
	prevs := []*node{}

	for ; i < len(s); i++ {
		switch s[i] {
		case ')':
			if curr != root {
				prevs = append(prevs, curr)
			}
			return prevs, i
		case '(':
			currs, i = getPaths(s, curr, i+1)
			currs = append([]*node{curr}, currs...)
		case '|':
			currs = []*node{}
			prevs = append(prevs, curr)
			curr = root
		default:
			n := &node{
				r:    rune(s[i]),
				next: []*node{},
			}
			if len(currs) != 0 {
				for _, nn := range currs {
					nn.next = append(nn.next, n)
				}
				currs = []*node{}
			} else if curr != nil {
				curr.next = append(curr.next, n)
			} else if curr == nil {
				root = n
			}
			curr = n
		}
	}
	return []*node{root}, 0
}

func walk(n *node, p string) {
	if len(n.next) == 0 {
		fmt.Println(p)
	}
	for _, next := range n.next {
		pp := fmt.Sprintf("%s%c", p, next.r)
		walk(next, pp)
	}
}

type point struct {
	y int
	x int
}

func step(p point, r rune) (point, point) {
	switch r {
	case 'N':
		return point{x: p.x, y: p.y - 1}, point{x: p.x, y: p.y - 2}
	case 'E':
		return point{x: p.x + 1, y: p.y}, point{x: p.x + 2, y: p.y}
	case 'S':
		return point{x: p.x, y: p.y + 1}, point{x: p.x, y: p.y + 2}
	case 'W':
		return point{x: p.x - 1, y: p.y}, point{x: p.x - 2, y: p.y}
	}
	log.Fatalf("wut")
	return p, p
}

type nodepoint struct {
	p point
	n *node
}

func getPoints(n *node, loc point) []point {
	seenPoints := map[point]struct{}{}
	points := []point{}
	nexts := []nodepoint{{
		p: loc,
		n: n,
	}}
	for {
		if len(nexts) == 0 {
			break
		}
		nextNexts := []nodepoint{}
		for _, nnn := range nexts {
			_, nextPoint := step(nnn.p, nnn.n.r)
			if _, ok := seenPoints[nextPoint]; !ok {
				seenPoints[nextPoint] = struct{}{}
				points = append(points, nextPoint)
				for _, nn := range nnn.n.next {
					nextNexts = append(nextNexts, nodepoint{n: nn, p: nextPoint})
				}
			}
		}
		nexts = nextNexts
	}
	return points
}

func fillIn(grid [][]rune, n *node, minX, minY int) point {
	start := point{x: abs(minX), y: abs(minY)}
	seen := map[point]struct{}{}
	nexts := []nodepoint{{
		p: start,
		n: n,
	}}
	grid[start.y][start.x] = 'X'
	for {
		if len(nexts) == 0 {
			break
		}
		nextNexts := []nodepoint{}
		for _, nnn := range nexts {
			door, nextPoint := step(nnn.p, nnn.n.r)
			if _, ok := seen[nextPoint]; !ok {
				seen[nextPoint] = struct{}{}

				grid[door.y][door.x] = '|'
				grid[nextPoint.y][nextPoint.x] = '.'

				for _, nn := range nnn.n.next {
					nextNexts = append(nextNexts, nodepoint{n: nn, p: nextPoint})
				}
			}
		}
		nexts = nextNexts
	}
	return start
}

func findBounds(points []point) (int, int, int, int) {
	minY, minX, maxY, maxX := 0, 0, 0, 0
	for _, p := range points {
		if p.x < minX {
			minX = p.x
		}
		if p.x > maxX {
			maxX = p.x
		}
		if p.y < minY {
			minY = p.y
		}
		if p.y > maxY {
			maxY = p.y
		}
	}
	return minY, minX, maxY, maxX
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func getGrid(n *node) ([][]rune, point) {
	points := getPoints(n, point{x: 0, y: 0})
	minY, minX, maxY, maxX := findBounds(points)
	height := abs(minY) + maxY + 1
	width := abs(minX) + maxX + 1
	grid := [][]rune{}

	for y := 0; y < height; y++ {
		grid = append(grid, []rune{})
		for x := 0; x < width; x++ {
			grid[y] = append(grid[y], '#')
		}
	}

	start := fillIn(grid, n, minX, minY)
	return grid, start
}

func neighbors(p point) [][]point {
	return [][]point{{
		// above
		{
			x: p.x,
			y: p.y - 1,
		}, {
			x: p.x,
			y: p.y - 2,
		},
	}, {
		// left
		{
			x: p.x - 1,
			y: p.y,
		}, {
			x: p.x - 2,
			y: p.y,
		},
	}, {
		// below
		{
			x: p.x,
			y: p.y + 1,
		}, {
			x: p.x,
			y: p.y + 2,
		},
	}, {
		// right
		{
			x: p.x + 1,
			y: p.y,
		}, {
			x: p.x + 2,
			y: p.y,
		},
	}}
}

type steppoint struct {
	p   point
	len int
}

func findLongest(grid [][]rune, start point) (int, map[point]int) {
	dists := map[point]int{}
	longest := 0
	next := []steppoint{{
		p:   start,
		len: 0,
	}}
	visited := map[point]bool{}

	round := 1
	for {
		if len(next) == 0 {
			break
		}

		nextNext := []steppoint{}
		for _, n := range next {
			neighborss := neighbors(n.p)
			for _, ns := range neighborss {
				if _, ok := visited[ns[1]]; !ok {
					if ns[1].x >= 0 && ns[1].x < len(grid[0]) && ns[1].y >= 0 && ns[1].y < len(grid) {
						if grid[ns[0].y][ns[0].x] == '|' && grid[ns[1].y][ns[1].x] == '.' {
							visited[ns[1]] = true
							s := steppoint{
								p:   ns[1],
								len: n.len + 1,
							}
							if s.len > longest {
								longest = s.len
							}
							if grid[ns[1].y][ns[1].x] == '.' {
								dists[ns[1]] = round
							}
							nextNext = append(nextNext, s)
						}
					}
				}
			}
		}
		next = nextNext
		round++
	}

	return longest, dists
}

func removeStupidRegexMarkers(s string) string {
	return s[1 : len(s)-1]
}

func printGrid(grid [][]rune) {
	for y := -1; y <= len(grid); y++ {
		if y == -1 || y == len(grid) {
			for x := 0; x <= len(grid[0])+1; x++ {
				fmt.Printf("#")
			}
		} else {
			for x := -1; x <= len(grid[y]); x++ {
				if x == -1 || x == len(grid[y]) {
					fmt.Printf("#")
				} else {
					fmt.Printf("%c", grid[y][x])
				}
			}
		}
		fmt.Println()
	}
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

	v := removeStupidRegexMarkers(vals[0])
	fmt.Println(v)

	roots, _ := getPaths(v, nil, 0)
	root := roots[0]

	//walk(root, fmt.Sprintf("%c", root.r))
	grid, start := getGrid(root)
	//fmt.Println("start is", start)
	printGrid(grid)

	longest, dists := findLongest(grid, start)
	fmt.Println(longest)

	c := 0
	for _, dist := range dists {
		if dist >= 1000 {
			c++
		}
	}
	fmt.Println(c)
}
