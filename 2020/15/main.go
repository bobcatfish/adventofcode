package main

import (
	"fmt"
)

//const max = 2020
const max = 30000000

func main() {
	input := []int{12, 1, 16, 3, 11, 0}
	nums := map[int][]int{}

	var prev int
	for i := 0; i < max; i++ {
		var say int
		if i < len(input) {
			say = input[i]
		} else {

			said := nums[prev]
			if len(said) == 1 {
				say = 0
			} else {
				say = said[1] - said[0]
			}
		}

		tosay := nums[say]
		if tosay == nil {
			tosay = []int{i}
		} else if len(tosay) == 1 {
			tosay = append(tosay, i)
		} else {
			tosay[0], tosay[1] = tosay[1], i
		}

		//fmt.Println(say)
		nums[say] = tosay
		prev = say
	}
	fmt.Println(prev)
}
