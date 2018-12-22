package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

const (
	clay  = '#'
	empty = '.'
	water = '~'
	wet   = '|'
)

type point struct {
	x          int
	y          int
	goingLeft  bool
	goingRight bool
	wasUp      bool
	wasDown    bool
}

var waterLoc = point{
	x: 500,
	y: 0,
}

func getXY(v string) ([]point, error) {
	r, err := regexp.Compile("y=(.*), x=(.*)")
	if err != nil {
		return []point{}, fmt.Errorf("couldn't compile y first regex: %v", err)
	}
	m := r.FindStringSubmatch(v)
	if len(m) == 3 {
		y, err := strconv.Atoi(m[1])
		if err != nil {
			return []point{}, fmt.Errorf("couldn't convert %q to int: %v", m[1], err)
		}
		xs := strings.Split(m[2], "..")
		if len(xs) != 2 {
			return []point{}, fmt.Errorf("didn't find expected range in %q", m[2], err)
		}

		xstart, err := strconv.Atoi(xs[0])
		if err != nil {
			return []point{}, fmt.Errorf("couldn't convert %q to int: %v", xs[0], err)
		}

		xend, err := strconv.Atoi(xs[1])
		if err != nil {
			return []point{}, fmt.Errorf("couldn't convert %q to int: %v", xs[1], err)
		}

		p := []point{}
		for x := xstart; x <= xend; x++ {
			p = append(p, point{x: x, y: y})
		}
		return p, nil
	}

	r, err = regexp.Compile("x=(.*), y=(.*)")
	if err != nil {
		return []point{}, fmt.Errorf("couldn't compile x first regex: %v", err)
	}
	m = r.FindStringSubmatch(v)

	if len(m) != 3 {
		return []point{}, fmt.Errorf("couldn't find any matches in %q", v)
	}

	x, err := strconv.Atoi(m[1])
	if err != nil {
		return []point{}, fmt.Errorf("couldn't convert %q to int: %v", m[1], err)
	}
	ys := strings.Split(m[2], "..")
	if len(ys) != 2 {
		return []point{}, fmt.Errorf("didn't find expected range in %q", m[2], err)
	}

	ystart, err := strconv.Atoi(ys[0])
	if err != nil {
		return []point{}, fmt.Errorf("couldn't convert %q to int: %v", ys[0], err)
	}

	yend, err := strconv.Atoi(ys[1])
	if err != nil {
		return []point{}, fmt.Errorf("couldn't convert %q to int: %v", ys[1], err)
	}

	p := []point{}
	for y := ystart; y <= yend; y++ {
		p = append(p, point{x: x, y: y})
	}
	return p, nil
}

func getSize(ps []point) (int, int) {
	maxX, maxY := 0, 0

	for _, p := range ps {
		if p.x > maxX {
			maxX = p.x
		}
		if p.y > maxY {
			maxY = p.y
		}
	}
	return maxX + 1, maxY + 1
}

func getMins(ps []point) (int, int) {
	minX, minY := 9999, 9999

	for _, p := range ps {
		if p.x < minX {
			minX = p.x
		}
		if p.y < minY {
			minY = p.y
		}
	}
	return minX, minY
}

func pastEnd(grid [][]rune, p point) bool {
	if p.y >= len(grid) || p.y < 0 || p.x < 0 || p.x >= len(grid[0]) {
		return true
	}
	return false
}

func atEnd(grid [][]rune, p point) bool {
	if p.y == len(grid)-1 || p.x == len(grid[0])-1 {
		return true
	}
	return false
}

func canGo(grid [][]rune, p point) bool {
	if pastEnd(grid, p) {
		return false
	}
	item := grid[p.y][p.x]
	blocked := item == water || item == clay
	return !blocked
}

func down(loc point) point {
	return point{
		x: loc.x,
		y: loc.y + 1,
	}
}

func left(loc point) point {
	return point{
		x: loc.x - 1,
		y: loc.y,
	}
}

func right(loc point) point {
	return point{
		x: loc.x + 1,
		y: loc.y,
	}
}

func up(loc point) point {
	return point{
		x: loc.x,
		y: loc.y - 1,
	}
}

func getNext(grid [][]rune, loc point) []point {
	d := down(loc)
	if canGo(grid, d) {
		d.wasDown = true
		return []point{d}
	}
	ps := []point{}
	if loc.goingLeft || loc.wasDown || loc.wasUp {
		l := left(loc)
		if canGo(grid, l) {
			l.goingLeft = true
			ps = append(ps, l)
		}
	}
	if loc.wasUp || loc.goingRight || loc.wasDown {
		r := right(loc)
		if canGo(grid, r) {
			r.goingRight = true
			ps = append(ps, r)
		}
	}
	if len(ps) > 0 {
		return ps
	}
	u := up(loc)
	if canGo(grid, u) && atRest(grid, loc) {
		u.wasUp = true
		return []point{u}
	}
	return []point{loc}
}

func downBlocked(grid [][]rune, p point) bool {
	d := down(p)
	if pastEnd(grid, d) {
		return false
	}
	dItem := grid[d.y][d.x]
	return dItem == water || dItem == clay
}

func fullLeft(grid [][]rune, p point) bool {
	for x := p.x; x >= 0; x-- {
		pp := point{
			y: p.y,
			x: x,
		}
		if grid[pp.y][pp.x] == water || grid[pp.y][pp.x] == clay {
			return true
		}
		if !downBlocked(grid, pp) {
			return false
		}
	}
	return false
}

func fullRight(grid [][]rune, p point) bool {
	for x := p.x; x < len(grid[0]); x++ {
		pp := point{
			y: p.y,
			x: x,
		}
		if grid[pp.y][pp.x] == water || grid[pp.y][pp.x] == clay {
			return true
		}
		if !downBlocked(grid, pp) {
			return false
		}
	}
	return false
}

func atRest(grid [][]rune, p point) bool {
	if downBlocked(grid, p) {
		if fullLeft(grid, p) {
			return fullRight(grid, p)
		}
	}
	return false
}

func fillRest(grid [][]rune, p point) []point {
	for x := p.x; x >= 0; x-- {
		pp := point{
			y: p.y,
			x: x,
		}
		if grid[pp.y][pp.x] == water || grid[pp.y][pp.x] == clay {
			break
		}
		grid[pp.y][pp.x] = water
	}

	for x := p.x + 1; x < len(grid[0]); x++ {
		pp := point{
			y: p.y,
			x: x,
		}
		if grid[pp.y][pp.x] == water || grid[pp.y][pp.x] == clay {
			break
		}
		grid[pp.y][pp.x] = water
	}
	return []point{{
		y:     p.y - 1,
		x:     p.x,
		wasUp: true,
	}}
}

func findWater(points []point, grid [][]rune, loc point) {
	todo := []point{loc}
	todoMap := map[point]bool{
		loc: true,
	}
	loc.wasDown = true

	for i := 0; ; i++ {
		if len(todo) == 0 {
			break
		}
		visit := todo[0]
		v2 := point{
			x: visit.x,
			y: visit.y,
		}
		delete(todoMap, v2)
		todo = todo[1:]

		var nexts []point
		if atRest(grid, visit) {
			nexts = fillRest(grid, visit)
			//grid[visit.y][visit.x] = water
		} else {
			grid[visit.y][visit.x] = wet
			nexts = getNext(grid, visit)
		}
		if atEnd(grid, visit) {
			continue
		}

		for _, next := range nexts {
			nItem := grid[next.y][next.x]
			if nItem == water || nItem == clay {
				break
			}
			if next.x == visit.x && next.y == visit.y {
				break
			}
			if pastEnd(grid, next) {
				continue
			}
			n2 := point{
				x: next.x,
				y: next.y,
			}
			if _, ok := todoMap[n2]; ok {
				break
			}
			todo = append(todo, next)
			todoMap[n2] = true
		}
		/*
			printGrid(points, grid, todoMap)
			time.Sleep(200 * time.Millisecond)
		*/
	}
}

func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func printGrid(points []point, grid [][]rune, todo map[point]bool) {
	clear()
	minX, minY := getMins(points)
	for y := minY; y < len(grid); y++ {
		for x := minX; x < len(grid[y]); x++ {
			if _, ok := todo[point{y: y, x: x}]; ok {
				fmt.Printf("%c", '+')
			} else {
				fmt.Printf("%c", grid[y][x])
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func count(points []point, grid [][]rune) int {
	c := 0
	_, minY := getMins(points)
	for y := minY; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == wet || grid[y][x] == water {
				c++
			}
		}
	}
	return c
}

func countLeft(points []point, grid [][]rune) int {
	c := 0
	_, minY := getMins(points)
	for y := minY; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == water {
				c++
			}
		}
	}
	return c
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

	points := []point{}
	for _, v := range vals {
		ps, err := getXY(v)
		if err != nil {
			log.Fatalf("Couldn't get x and y values from %q: %v", v, err)
		}
		points = append(points, ps...)
	}

	minX, minY := getMins(points)
	fmt.Println("MINX", minX, "MINY", minY)
	width, height := getSize(points)
	grid := [][]rune{}
	for y := 0; y < height; y++ {
		grid = append(grid, []rune{})
		for x := 0; x < width; x++ {
			grid[y] = append(grid[y], empty)
		}
	}

	for _, p := range points {
		grid[p.y][p.x] = clay
	}

	findWater(points, grid, waterLoc)

	printGrid(points, grid, map[point]bool{})
	c := count(points, grid)
	fmt.Println(c)

	c2 := countLeft(points, grid)
	fmt.Println(c2)
}
