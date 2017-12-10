package main

import (
	"fmt"
)

const numLen = 256

func twist(nums []int, length int, curr int) {
	for o := 0; o < (length / 2); o++ {
		x := (curr + o) % numLen
		// Q: Why you have weird % Go? Why
		// A: BECAUSE OF THE SPEC GOD https://groups.google.com/forum/#!topic/golang-nuts/xj7CV857vAg
		// Q: But why the spec that way?
		// A: GOLANG ALWAYS RIGHT SHHHHH
		if x < 0 {
			x += numLen
		}
		y := ((curr + length - 1) - o) % numLen
		if y < 0 {
			y += numLen
		}
		nums[x], nums[y] = nums[y], nums[x]
	}
}

func main() {
	lengths := []int{
		129, 154, 49, 198, 200, 133, 97, 254, 41, 6, 2, 1, 255, 0, 191, 108,
	}
	nums := [numLen]int{}
	for i := range nums {
		nums[i] = i
	}

	var curr int
	for skip, length := range lengths {
		twist(nums[:], length, curr)
		curr += (length + skip)
	}
	fmt.Println(nums)
	fmt.Println(nums[0] * nums[1])
}
