package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strconv"
)

type point struct {
	x int
	y int
}

type star struct {
	p point
	v point
}

const (
	grid = 75
)

func getAdj(p point) []point {
	return []point{
		// above
		{x: p.x, y: p.y + 1},

		// beside
		{x: p.x - 1, y: p.y},
		{x: p.x + 1, y: p.y},

		// below
		{x: p.x, y: p.y - 1},
	}
}

func together(stars []*star) bool {
	left := findFurthestLeft(stars)
	adj := getAdj(left)
	for i := range stars {
		s := stars[i]
		for _, a := range adj {
			if a.x == s.p.x && a.y == s.p.y {
				return true
			}
		}
	}
	return false
}

func findFurthestLeft(stars []*star) point {
	furthest := point{x: 500000000000}

	for _, s := range stars {
		if s.p.x < furthest.x {
			furthest = s.p
		} else if s.p.x == furthest.x && s.p.y < furthest.y {
			furthest = s.p
		}
	}
	return furthest
}

func display(stars []*star) {
	values := [][]string{}

	left := findFurthestLeft(stars)

	y0 := left.y
	yMax := left.y + grid
	x0 := left.x
	xMax := left.x + grid

	// TODO: gonna have to find values to start displaying
	for y := y0; y < yMax; y++ {
		row := []string{}
		for x := x0; x < xMax; x++ {
			row = append(row, ".")
		}
		values = append(values, row)
	}

	for _, star := range stars {
		if star.p.y < yMax && star.p.y >= y0 && star.p.x < xMax && star.p.x >= x0 {
			y, x := star.p.y-y0, star.p.x-x0
			values[y][x] = "#"
		}
	}

	for y := 0; y < grid; y++ {
		for x := 0; x < grid; x++ {
			fmt.Printf(values[y][x])
		}
		fmt.Println()
	}
}

func change(stars []*star) {
	for _, star := range stars {
		star.p.x += star.v.x
		star.p.y += star.v.y
	}
}

func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
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

	signedInt := "([- ])(\\d*)"
	regex := fmt.Sprintf("position=<%s, %s> velocity=<%s, %s>", signedInt, signedInt, signedInt, signedInt)
	r, err := regexp.Compile(regex)
	if err != nil {
		log.Fatalf("Couldn't compile regex: %s", err)
	}

	stars := []*star{}
	for _, v := range vals {
		m := r.FindStringSubmatch(v)
		if len(m) != 9 {
			log.Fatalf("didn't find all expected matches in %q, found %d", v, len(m))
		}

		signs := []string{m[1], m[3], m[5], m[7]}
		pVals := []string{m[2], m[4], m[6], m[8]}
		is := []int{}

		for i := 0; i < len(signs); i++ {
			pVal, err := strconv.Atoi(pVals[i])
			if err != nil {
				log.Fatalf("couldn't convert %q (from %q) to int: %v", pVals[i], v, err)
			}
			if signs[i] == "-" {
				pVal = 0 - pVal
			}
			is = append(is, pVal)
		}
		p := point{
			x: is[0],
			y: is[1],
		}
		v := point{
			x: is[2],
			y: is[3],
		}

		star := star{
			p: p,
			v: v,
		}

		stars = append(stars, &star)
	}

	iteration := 0
	for {
		if together(stars) {
			clear()
			display(stars)
			fmt.Println(iteration)
			break
		}

		change(stars)
		iteration++
	}
}
