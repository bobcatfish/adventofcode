package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type instr struct {
	cmd string
	num int
}

func load() []instr {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("couldn't open file: %v", err)
	}
	defer file.Close()

	vals := []instr{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		val := scanner.Text()
		vv := strings.Split(val, " ")

		inst := instr{
			cmd: vv[0],
		}
		if len(vv) == 2 {
			var err error
			inst.num, err = strconv.Atoi(vv[1])
			if err != nil {
				log.Fatalf("couldn't convert int: %v", err)
			}
		}
		vals = append(vals, inst)
	}
	return vals
}

const wide = 40
const high = 6

func endCycle(cycle, x int) int {
	pos := cycle % wide
	//fmt.Println(cycle, pos)
	if pos == wide-1 {
		fmt.Println()
	}
	if pos >= x-1 && pos <= x+1 {
		fmt.Printf("#")
	} else {
		fmt.Printf(".")
	}
	cycle++
	//fmt.Println("cycle", cycle, "x", x)
	if cycle == 20 || (cycle-20)%40 == 0 {
		//fmt.Println("cycle", cycle, "x", x, "strength", cycle*x)
		return cycle * x
	}
	return 0
}

func main() {
	vals := load()
	fmt.Println(vals)
	x := 1
	num := 0
	index := 0
	sum := 0
	for cycle := 0; ; cycle++ {
		if num != 0 {
			sum += endCycle(cycle, x)
			x += num
			num = 0
			continue
		}

		inst := vals[index]
		index++
		//fmt.Println(inst, index, x)
		if index >= len(vals) {
			sum += endCycle(cycle, x)
			break
		}
		if inst.cmd == "noop" {
			sum += endCycle(cycle, x)
			continue
		}
		if inst.cmd == "addx" {
			num = inst.num
			sum += endCycle(cycle, x)
			continue
		}
		log.Fatalf("unexpected cmd %s", inst.cmd)
	}
	fmt.Println(sum)
}
