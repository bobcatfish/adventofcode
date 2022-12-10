package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type pos struct {
	y int
	x int
}

type move struct {
	dir  byte
	num  int
	poop int
}

func load() []move {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("couldn't open file: %v", err)
	}
	defer file.Close()

	vals := []move{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		val := scanner.Text()
		vv := strings.Split(val, " ")
		i, err := strconv.Atoi(vv[1])
		if err != nil {
			log.Fatalf("couldn't convert int: %v", err)
		}
		vals = append(vals, move{
			dir: vv[0][0],
			num: i,
		})
	}
	return vals
}

func adjacent(t, h *pos) bool {
	return math.Abs(float64(t.y-h.y)) <= 1 && math.Abs(float64(t.x-h.x)) <= 1
}

func moveTail(h, t *pos) {
	if !adjacent(t, h) {
		if h.y > t.y {
			t.y++
		} else if h.y < t.y {
			t.y--
		}
		if h.x > t.x {
			t.x++
		} else if h.x < t.x {
			t.x--
		}
	}
}

func doMove(m move, knots []*pos, seen map[pos]struct{}) {
	for i := 0; i < m.num; i++ {
		if m.dir == 'R' {
			knots[0].x++
		} else if m.dir == 'L' {
			knots[0].x--
		} else if m.dir == 'U' {
			knots[0].y++
		} else if m.dir == 'D' {
			knots[0].y--
		}
		for j := 0; j < len(knots)-1; j++ {
			moveTail(knots[j], knots[j+1])
		}
		seen[*knots[len(knots)-1]] = struct{}{}
	}
}

func doKnots(vals []move, numKnots int) {
	knots := []*pos{}
	for i := 0; i < numKnots; i++ {
		knots = append(knots, &pos{y: 0, x: 0})
	}
	seen := map[pos]struct{}{
		*knots[0]: struct{}{},
	}

	for _, m := range vals {
		doMove(m, knots, seen)
	}
	fmt.Println(len(seen))
}

func main() {
	vals := load()
	doKnots(vals, 2)
	doKnots(vals, 10)
}
