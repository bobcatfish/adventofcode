package main

import (
	"fmt"
    "math/bits"
    "strconv"
    "log"
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

func hash(input string) string {
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
    return hash
}

type point struct {
    X int
    Y int
}

func walk(lines [][]string, y int, x int, m string) {
    lines[y][x] = m
    var neighbours = []point{
		point{
			Y: -1,
			X: 0,
		},
		point{
			Y: 0,
			X: -1,
		},
		point{
			Y: 0,
			X: 1,
		},
		point{
			Y: 1,
			X: 0,
		},
	}
    for _, n := range neighbours {
        ny := y + n.Y
        nx := x + n.X
        if ny >= 0 && nx >= 0 && ny < len(lines) && nx < len(lines[ny]) && lines[ny][nx] == "#" {
            walk(lines, ny, nx, m)
        }
    }
}


func doIt(lines [][]string) int {
    fmt.Println(lines)
    regions := 1

    for y := 0; y < len(lines); y++ {
        for x := 0; x < len(lines[y]); x++ {
            if lines[y][x] == "#" {
                walk(lines, y, x, strconv.FormatInt(int64(regions), 10))
                regions++
            }
        }
    }

    return regions - 1
}

const size = 128

func main() {
    input := "ffayrhll"

    count := 0
    lines := make([][]string, size)
    for i := 0; i < size; i++ {
        s := fmt.Sprintf("%s-%d", input, i)
        h := hash(s)
        line := ""

        for j := 0; j < len(h); j++ {
            v, err := strconv.ParseInt(string(h[j]), 16, 8)
            if err != nil {
                log.Fatal(err)
            }
            count += bits.OnesCount64(uint64(v))

            s := fmt.Sprintf("%04b", v)
            line += s
        }
        lines[i] = make([]string, len(line))
        for j := 0; j < len(line); j++ {
            if string(line[j]) == "1" {
                lines[i][j] = "#"
            } else {
                lines[i][j] = "-"
            }
        }
    }

    fmt.Println(count)

    regions := doIt(lines)
    fmt.Println(regions)
}
