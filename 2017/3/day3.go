package main

import (
	"fmt"
	"math"
)

func main() {
	i := 312051

	multOfFour := ((i - 1) / 4) + 1
	ring := int(math.Floor(math.Sqrt(float64(multOfFour))))

	offset := i - 1
	for j := 0; j < ring; j++ {
		capacity := (((j * 2) - 1) * 4) + 4
		if offset >= capacity {
			offset -= capacity
		}
	}

	// Size of each side of a ring follows a pattern 0, 1, 3, 5, 7, ... + 2
	height := (ring * 2) + 1
	// Sides are 0 = right, 1 = top, 3 = left, 4 = bottom.
	for side := 0; side < 4; side++ {
		if offset < (height - 1) {
			break
		}
		offset -= height - 1
	}

	middle := (height / 2)
	result := ring + int(math.Abs(float64(offset-middle)))
	fmt.Println(result)
}
