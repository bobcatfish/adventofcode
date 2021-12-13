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
	X int
	Y int
}

type fold struct {
	V      int
	AlongX bool
}

func load() (map[point]struct{}, int, int, []fold, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, 0, 0, nil, fmt.Errorf("couldn't open file: %v", err)
	}
	defer file.Close()

	points := map[point]struct{}{}
	folds := []fold{}
	pPoints := true
	maxY, maxX := -1, -1

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		val := scanner.Text()
		if val == "" {
			pPoints = false
			continue
		}

		if pPoints {
			ss := strings.Split(val, ",")
			x, err := strconv.Atoi(ss[0])
			if err != nil {
				return nil, 0, 0, nil, err
			}
			y, err := strconv.Atoi(ss[1])
			if err != nil {
				return nil, 0, 0, nil, err
			}
			if x > maxX {
				maxX = x
			}
			if y > maxY {
				maxY = y
			}
			points[point{X: x, Y: y}] = struct{}{}
		} else {
			f := fold{}
			ss := strings.Split(val, "=")
			v, err := strconv.Atoi(ss[1])
			if err != nil {
				return nil, 0, 0, nil, err
			}
			f.V = v

			ss = strings.Split(ss[0], " ")
			if ss[len(ss)-1] == "x" {
				f.AlongX = true
			}
			folds = append(folds, f)
		}

	}
	return points, maxY, maxX, folds, err
}

func countGrid(points map[point]struct{}, maxY, maxX int) int {
	count := 0
	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			p := point{X: x, Y: y}
			if _, ok := points[p]; ok {
				count++
			}
		}
	}
	return count
}

func printGrid(points map[point]struct{}, maxY, maxX int) {
	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			p := point{X: x, Y: y}
			if _, ok := points[p]; ok {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func doFold(points map[point]struct{}, maxY, maxX int, f fold) (map[point]struct{}, int, int) {
	newPoints := map[point]struct{}{}
	if f.AlongX {
		for y := 0; y <= maxY; y++ {
			for x := 0; x < f.V; x++ {
				p := point{Y: y, X: x}
				if _, ok := points[p]; ok {
					newPoints[p] = struct{}{}
				}
				reverseP := point{Y: y, X: (f.V * 2) - x}
				if _, ok := points[reverseP]; ok {
					newPoints[p] = struct{}{}
				}
			}
		}
		return newPoints, maxY, f.V - 1
	}
	for y := 0; y < f.V; y++ {
		for x := 0; x <= maxX; x++ {
			p := point{Y: y, X: x}
			if _, ok := points[p]; ok {
				newPoints[p] = struct{}{}
			}
			reverseP := point{Y: (f.V * 2) - y, X: x}
			if _, ok := points[reverseP]; ok {
				newPoints[p] = struct{}{}
			}
		}
	}
	return newPoints, f.V - 1, maxX
}

func main() {
	points, maxY, maxX, folds, err := load()
	if err != nil {
		log.Fatalf("Couldn't load nums from file: %v", err)
	}

	for i, f := range folds {
		points, maxY, maxX = doFold(points, maxY, maxX, f)
		c := countGrid(points, maxY, maxX)
		fmt.Println("fold", i+1, c, "dots")
	}
	printGrid(points, maxY, maxX)
}
