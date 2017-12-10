package main

import (
	"fmt"
)

const numsLen = 256

const rounds = 64
const denseLen = 16

func twist(nums []int, length int, curr int) {
	for o := 0; o < (length / 2); o++ {
		x := (curr + o) % numsLen
		// Q: Why you have weird % Go? Why
		// A: BECAUSE OF THE SPEC GOD https://groups.google.com/forum/#!topic/golang-nuts/xj7CV857vAg
		// Q: But why the spec that way?
		// A: GOLANG ALWAYS RIGHT SHHHHH
		if x < 0 {
			x += numsLen
		}
		y := ((curr + length - 1) - o) % numsLen
		if y < 0 {
			y += numsLen
		}
		nums[x], nums[y] = nums[y], nums[x]
	}
}

func makeDense(nums []int) [denseLen]int {
	var dense [denseLen]int
	for i := 0; i < denseLen; i++ {
		var v int
		for j := (i * denseLen); j < ((i + 1) * denseLen); j++ {
			v ^= nums[j]
		}
		dense[i] = v
	}
	return dense
}

func hex(nums []int) string {
	var s string
	for _, v := range nums {
		s += fmt.Sprintf("%02x", v)
	}
	return s
}

func main() {
	input := "129,154,49,198,200,133,97,254,41,6,2,1,255,0,191,108"

	var lengths []int
	for i := range input {
		lengths = append(lengths, int(input[i]))
	}
	rando := []int{
		17, 31, 73, 47, 23,
	}
	for _, v := range rando {
		lengths = append(lengths, v)
	}

	nums := [numsLen]int{}
	for i := range nums {
		nums[i] = i
	}

	var curr int
	var skip int

	for i := 0; i < rounds; i++ {
		for _, length := range lengths {
			twist(nums[:], length, curr)
			curr += (length + skip)
			skip++
		}
	}

	dense := makeDense(nums[:])
	hash := hex(dense[:])
	fmt.Println(hash)
}
