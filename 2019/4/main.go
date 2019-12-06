package main

import (
	"fmt"
	"strings"
)

func print(ii []int) {
	s := strings.Trim(strings.Join(strings.Split(fmt.Sprint(ii), " "), ""), "[]")
	fmt.Println(s)
}

func hasDouble(ii []int) bool {
	for i := 0; i < len(ii); {
		g := 1
		for ; (g + i) < len(ii); g++ {
			if ii[i] != ii[i+g] {
				break
			}
		}
		if g == 2 {
			return true
		}
		i += g
	}
	return false
}

func increasing(ii []int) bool {
	for i := 1; i < len(ii); i++ {
		if ii[i] < ii[i-1] {
			return false
		}
	}
	return true
}

func zero(ii []int, i int) []int {
	for ; i < len(ii); i++ {
		ii[i] = 0
	}
	return ii
}

func isLast(ii []int) bool {
	last := []int{6, 7, 5, 8, 1, 0}
	for i := 0; i < len(ii); i++ {
		if ii[i] > last[i] {
			return true
		}
		if ii[i] < last[i] {
			return false
		}
	}
	return true
}

func getPosNext(ii []int, i int) int {
	c := 0
	if i < 0 {
		return c
	}
	if i >= len(ii) {
		return c
	}
	if isLast(ii) {
		return c
	}

	for {
		if i == 0 || ii[i] >= ii[i-1] {
			if increasing(ii) && hasDouble(ii) {
				c += 1
			}
			c += getPosNext(ii[:], i+1)
		}
		if isLast(ii) {
			return c
		}
		if ii[i] == 9 {
			break
		}
		ii[i] += 1
	}

	iii := ii[:]

	j := i
	for {
		j -= 1
		if j < 0 {
			return c
		}
		if iii[j] != 9 {
			iii[j] += 1
			zero(iii, j+1)
			if isLast(iii) {
				return c
			}
			c += getPosNext(iii, j+1)
			break
		}
	}

	return c
}

func main() {
	i := []int{1, 3, 4, 7, 9, 2}
	c := getPosNext(i, len(i)-1)
	fmt.Println(c)
}
