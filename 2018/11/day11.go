package main

import (
	"fmt"
	"log"
	"strconv"
)

const (
	size  = 300
	magic = 5

	serialNum = 1718
)

func power(x, y int) (int, error) {
	rackID := x + 10
	power := rackID * y
	power += serialNum
	power *= rackID

	hundreds := strconv.Itoa(power / 100)
	hundo := hundreds[len(hundreds)-1:]

	i, err := strconv.Atoi(hundo)
	if err != nil {
		return 0, fmt.Errorf("couldn't convert hundreds digit of %d: %v", power, err)
	}
	return i - magic, nil
}

func getCube(powers [][]int, cubeSize, x, y int) [][]int {
	cube := [][]int{}

	for row := y; row < y+cubeSize; row++ {
		cube = append(cube, powers[row][x:x+cubeSize])
	}
	return cube
}

func getCubeCount(cube [][]int) int {
	sum := 0
	for y := 0; y < len(cube); y++ {
		for x := 0; x < len(cube[y]); x++ {
			sum += cube[y][x]
		}
	}
	return sum
}

func findBest(powers [][]int, cubeSize int) (int, int, int) {
	best, bestX, bestY := 0, 0, 0
	for y := 0; y < size-cubeSize; y++ {
		for x := 0; x < size-cubeSize; x++ {
			cube := getCube(powers, cubeSize, x, y)
			next := getCubeCount(cube)
			if next > best {
				best, bestX, bestY = next, x, y
			}
		}
	}
	return best, bestX, bestY
}

func main() {
	powers := [][]int{}

	for y := 0; y < size; y++ {
		powers = append(powers, []int{})
		for x := 0; x < size; x++ {
			p, err := power(x, y)
			if err != nil {
				log.Fatalf("Couldn't get power for (%d, %d): %v", x, y, err)
			}
			powers[y] = append(powers[y], p)
		}
	}

	overallBest, bestSize, ox, oy := 0, 0, 0, 0
	for i := 0; i < size; i++ {
		best, x, y := findBest(powers, i)
		fmt.Printf("best at (%d,%d,%d): %d\n", x, y, i, best)
		fmt.Printf("overall best at (%d,%d,%d): %d\n", ox, oy, bestSize, overallBest)
		if best > overallBest {
			overallBest, bestSize, ox, oy = best, i, x, y
		}
	}

	fmt.Printf("overall best at (%d,%d,%d): %d\n", ox, oy, bestSize, overallBest)
}
