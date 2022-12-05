package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type Stack struct {
	s []byte
}

func (s *Stack) Push(b []byte) {
	s.s = append(s.s, b...)
}

func (s *Stack) Pop(i int) []byte {
	b := s.s[len(s.s)-i:]
	s.s = s.s[:len(s.s)-i]
	return b
}

type Instr struct {
	count int
	from  int
	to    int
}

func move2(stacks []*Stack, instrs []Instr) {
	for _, instr := range instrs {
		b := stacks[instr.from].Pop(instr.count)
		stacks[instr.to].Push(b)
	}
}

func move(stacks []*Stack, instrs []Instr) {
	for _, instr := range instrs {
		for i := 0; i < instr.count; i++ {
			b := stacks[instr.from].Pop(1)
			stacks[instr.to].Push(b)
		}
	}
}

func parseNums(ss []string) []int {
	nums := []int{}
	for _, sss := range ss {
		i, err := strconv.Atoi(sss)
		if err != nil {
			log.Fatalf("couldn't convert %s to int: %v", sss, err)
		}
		nums = append(nums, i)
	}
	return nums
}

func getInstrs(strs []string) []Instr {
	instrs := []Instr{}
	for _, s := range strs {
		re := regexp.MustCompile(`move (\d*) from (\d*) to (\d*)`)
		vv := re.FindStringSubmatch(s)
		vvv := parseNums(vv[1:])
		instr := Instr{
			count: vvv[0],
			from:  vvv[1] - 1,
			to:    vvv[2] - 1,
		}
		instrs = append(instrs, instr)
	}
	return instrs
}

func getStacks(strs []string) []*Stack {
	numStacks := (len(strs[len(strs)-1]) + 1) / 4

	backwardStacks := [][]byte{}
	for i := 0; i < numStacks; i++ {
		backwardStacks = append(backwardStacks, []byte{})
	}
	for _, s := range strs[:len(strs)-1] {
		for i := 0; i < numStacks; i++ {
			index := 1 + (4 * i)
			if len(s) > index && s[index] != ' ' {
				backwardStacks[i] = append(backwardStacks[i], s[index])
			}
		}
	}
	stacks := []*Stack{}
	for i := 0; i < numStacks; i++ {
		stacks = append(stacks, &Stack{s: []byte{}})
		bstack := backwardStacks[i]
		for j := len(bstack) - 1; j >= 0; j-- {
			stacks[i].Push([]byte{bstack[j]})
		}
	}
	return stacks
}

func load() ([]string, []string) {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("couldn't open file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	strs := []string{}
	breakIndex := -1
	for scanner.Scan() {
		val := scanner.Text()
		if len(val) > 1 && val[1] == '1' {
			breakIndex = len(strs)
		}
		strs = append(strs, val)
	}
	return strs[:breakIndex+1], strs[breakIndex+2:]
}

func printTops(stacks []*Stack) {
	for _, s := range stacks {
		fmt.Printf("%c", s.s[len(s.s)-1])
	}
	fmt.Println()
}

func main() {
	stackStrs, inStrs := load()
	instrs := getInstrs(inStrs)
	stacks := getStacks(stackStrs)

	move(stacks, instrs)
	printTops(stacks)

	stacks2 := getStacks(stackStrs)
	move2(stacks2, instrs)
	printTops(stacks2)
}
