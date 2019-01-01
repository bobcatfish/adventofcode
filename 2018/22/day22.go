package main

import (
	"fmt"
	"log"
)

const (
	//depth = 510
	depth = 11541
)

type point struct {
	x int
	y int
}

type region struct {
	r       rune
	erosion int
}

//var target = point{x: 10, y: 10}
var target = point{x: 14, y: 778}

func getIndex(p point, grid [][]region) int {
	if (p.x == 0 && p.y == 0) || p == target {
		return 0
	}
	if p.y == 0 {
		return p.x * 16807
	}
	if p.x == 0 {
		return p.y * 48271
	}
	return grid[p.y-1][p.x].erosion * grid[p.y][p.x-1].erosion
}

func getErosion(p point, grid [][]region) int {
	index := getIndex(p, grid)
	return (index + depth) % 20183
}

func getType(erosion int) rune {
	switch erosion % 3 {
	case 0:
		return '.'
	case 1:
		return '='
	case 2:
		return '|'
	}
	log.Fatalf("IMPOSSIBLE :O")
	return 0
}

func draw(grid [][]region) {
	for y := 0; y <= target.y; y++ {
		for x := 0; x <= target.x; x++ {
			var r rune
			if x == 0 && y == 0 {
				r = 'M'
			} else if x == target.x && y == target.y {
				r = 'T'
			} else {
				r = grid[y][x].r
			}
			fmt.Printf("%c", r)
		}
		fmt.Println()
	}
}

func getRisk(grid [][]region) int {
	risk := 0
	for y := 0; y <= target.y; y++ {
		for x := 0; x <= target.x; x++ {
			switch grid[y][x].r {
			case '=':
				risk++
			case '|':
				risk += 2
			}
		}
	}
	return risk
}

func main() {
	grid := [][]region{}

	for y := 0; y <= target.y; y++ {
		grid = append(grid, []region{})
		for x := 0; x <= target.x; x++ {
			p := point{x: x, y: y}
			erosion := getErosion(p, grid)
			t := getType(erosion)
			grid[y] = append(grid[y], region{
				erosion: erosion,
				r:       t,
			})
		}
	}

	draw(grid)

	risk := getRisk(grid)
	fmt.Println(risk)
}
