package main

import (
	"fmt"
)

func factors(value int) []int {
	f := []int{}
	for i := 1; i <= value; i++ {
		if value%i == 0 {
			f = append(f, i)
		}
	}
	return f
}

func main() {
	m := 10551354
	factors := factors(m)

	v := 0
	for _, f := range factors {
		v += f
	}
	fmt.Println(v)
}
