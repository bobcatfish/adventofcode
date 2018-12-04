package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type point struct {
	x int
	y int
}

type rect struct {
	upperLeft  point
	lowerRight point
}

type claim struct {
	num    int
	left   int
	top    int
	width  int
	height int
	rect   rect
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func makeRect(c claim) rect {
	upperLeft := point{
		x: c.left,
		y: c.top,
	}
	lowerRight := point{
		x: c.left + c.width,
		y: c.top + c.height,
	}
	return rect{
		upperLeft:  upperLeft,
		lowerRight: lowerRight,
	}
}

func getIntersection(r1 rect, r2 rect) rect {
	upperLeft := point{
		x: max(r1.upperLeft.x, r2.upperLeft.x),
		y: max(r1.upperLeft.y, r2.upperLeft.y),
	}
	lowerRight := point{
		x: min(r1.lowerRight.x, r2.lowerRight.x),
		y: min(r1.lowerRight.y, r2.lowerRight.y),
	}
	return rect{
		upperLeft:  upperLeft,
		lowerRight: lowerRight,
	}
}

func didIntersect(r rect) bool {
	if r.upperLeft.x >= r.lowerRight.x {
		return false
	}
	if r.upperLeft.y >= r.lowerRight.y {
		return false
	}
	return true
}

func countIt(r rect, counted map[point]bool) int {
	c := 0
	for x := r.upperLeft.x; x < r.lowerRight.x; x++ {
		for y := r.upperLeft.y; y < r.lowerRight.y; y++ {
			p := point{
				x: x,
				y: y,
			}
			if _, ok := counted[p]; !ok {
				counted[p] = true
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
	r, err := regexp.Compile("#(\\d*) @ (\\d*),(\\d*): (\\d*)x(\\d*)")
	if err != nil {
		log.Fatalf("error compiling regex: %v", err)
	}

	claims := []*claim{}
	for _, s := range vals {
		match := r.FindStringSubmatch(s)
		if len(match) != 6 {
			log.Fatalf("didn't get all 6 expected matches for %q, got %d", s, len(match))
		}

		matchInts := []int{}
		for i, m := range match {
			// First match is the entire string
			if i == 0 {
				continue
			}
			i, err := strconv.Atoi(m)
			if err != nil {
				log.Fatalf("error converting %s to int: %v", m, err)
			}
			matchInts = append(matchInts, i)
		}
		claims = append(claims, &claim{
			num:    matchInts[0],
			left:   matchInts[1],
			top:    matchInts[2],
			width:  matchInts[3],
			height: matchInts[4],
		})
	}

	for _, c := range claims {
		c.rect = makeRect(*c)
	}

	n := 0
	counted := map[point]bool{}
	intersections := map[int]bool{}
	for i, c := range claims {
		if i == len(claims)-1 {
			break
		}
		for j := i + 1; j < len(claims); j++ {
			intersect := getIntersection(c.rect, claims[j].rect)
			if didIntersect(intersect) {
				n += countIt(intersect, counted)
				intersections[c.num] = true
				intersections[claims[j].num] = true
			}
		}
	}
	for _, c := range claims {
		if _, ok := intersections[c.num]; !ok {
			fmt.Printf("%d had no overlap\n", c.num)
		}
	}
	fmt.Printf("%d squares overlapped\n", n)
}
