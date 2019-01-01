package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	a = 0
	b = 1
	c = 2
)

type opcode func([]int, []int) ([]int, error)
type op func(int, int) int

func copyReg(reg []int) []int {
	newReg := make([]int, len(reg))
	copy(newReg, reg)
	return newReg
}

func regOp(reg []int, input []int, o op) ([]int, error) {
	if input[c] >= len(reg) || input[a] >= len(reg) || input[b] >= len(reg) {
		return []int{}, fmt.Errorf("")
	}
	newReg := copyReg(reg)
	newReg[input[c]] = o(reg[input[a]], reg[input[b]])
	return newReg, nil
}

func valOp(reg []int, input []int, o op) ([]int, error) {
	if input[c] >= len(reg) || input[a] >= len(reg) {
		return []int{}, fmt.Errorf("")
	}
	newReg := copyReg(reg)
	newReg[input[c]] = o(reg[input[a]], input[b])
	return newReg, nil
}

func addr(reg []int, input []int) ([]int, error) {
	return regOp(reg, input, func(a int, b int) int { return a + b })
}

func addi(reg []int, input []int) ([]int, error) {
	return valOp(reg, input, func(a int, b int) int { return a + b })
}

func multr(reg []int, input []int) ([]int, error) {
	return regOp(reg, input, func(a int, b int) int { return a * b })
}

func multi(reg []int, input []int) ([]int, error) {
	return valOp(reg, input, func(a int, b int) int { return a * b })
}

func banr(reg []int, input []int) ([]int, error) {
	return regOp(reg, input, func(a int, b int) int { return a & b })
}

func bani(reg []int, input []int) ([]int, error) {
	return valOp(reg, input, func(a int, b int) int { return a & b })
}

func borr(reg []int, input []int) ([]int, error) {
	return regOp(reg, input, func(a int, b int) int { return a | b })
}

func bori(reg []int, input []int) ([]int, error) {
	return valOp(reg, input, func(a int, b int) int { return a | b })
}

func setr(reg []int, input []int) ([]int, error) {
	if input[c] >= len(reg) || input[a] >= len(reg) {
		return []int{}, fmt.Errorf("")
	}
	newReg := copyReg(reg)
	newReg[input[c]] = reg[input[a]]
	return newReg, nil
}

func seti(reg []int, input []int) ([]int, error) {
	if input[c] >= len(reg) {
		return []int{}, fmt.Errorf("%d is outside of registers %v", input[c], reg)
	}
	newReg := copyReg(reg)
	newReg[input[c]] = input[a]
	return newReg, nil
}

type test func(int, int) bool

func ir(reg []int, input []int, t test) ([]int, error) {
	if input[c] >= len(reg) || input[b] >= len(reg) {
		return []int{}, fmt.Errorf("%d or %d is outside of registers %v", input[c], input[b], reg)
	}
	newReg := copyReg(reg)
	if t(input[a], reg[input[b]]) {
		newReg[input[c]] = 1
	} else {
		newReg[input[c]] = 0
	}
	return newReg, nil
}

func ri(reg []int, input []int, t test) ([]int, error) {
	if input[c] >= len(reg) || input[a] >= len(reg) {
		return []int{}, fmt.Errorf("%d or %d is outside of registers %v", input[c], input[a], reg)
	}
	newReg := copyReg(reg)
	if t(reg[input[a]], input[b]) {
		newReg[input[c]] = 1
	} else {
		newReg[input[c]] = 0
	}
	return newReg, nil
}

func rr(reg []int, input []int, t test) ([]int, error) {
	if input[c] >= len(reg) || input[a] >= len(reg) || input[b] >= len(reg) {
		return []int{}, fmt.Errorf("%d or %d or %d is outside of registers %v", input[c], input[a], input[b], reg)
	}
	newReg := copyReg(reg)
	if t(reg[input[a]], reg[input[b]]) {
		newReg[input[c]] = 1
	} else {
		newReg[input[c]] = 0
	}
	return newReg, nil
}

func gtir(reg []int, input []int) ([]int, error) {
	return ir(reg, input, func(a, b int) bool { return a > b })
}

func gtri(reg []int, input []int) ([]int, error) {
	return ri(reg, input, func(a, b int) bool { return a > b })
}

func gtrr(reg []int, input []int) ([]int, error) {
	return rr(reg, input, func(a, b int) bool { return a > b })
}

func eqir(reg []int, input []int) ([]int, error) {
	return ir(reg, input, func(a, b int) bool { return a == b })
}

func eqri(reg []int, input []int) ([]int, error) {
	return ri(reg, input, func(a, b int) bool { return a == b })
}

func eqrr(reg []int, input []int) ([]int, error) {
	return rr(reg, input, func(a, b int) bool { return a == b })
}

var opcodes = map[string]opcode{
	"addr": addr,
	"addi": addi,
	"mulr": multr,
	"muli": multi,
	"banr": banr,
	"bani": bani,
	"borr": borr,
	"bori": bori,
	"setr": setr,
	"seti": seti,
	"gtir": gtir,
	"gtri": gtri,
	"gtrr": gtrr,
	"eqir": eqir,
	"eqri": eqri,
	"eqrr": eqrr,
}

func parseInput(m []string) ([]int, error) {
	is := []int{}
	for i := 1; i < len(m); i++ {
		i, err := strconv.Atoi(m[i])
		if err != nil {
			return []int{}, fmt.Errorf("couldn't parse int from %q", m[i])
		}
		is = append(is, i)
	}
	return is, nil
}

func registerEq(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, aa := range a {
		if b[i] != aa {
			return false
		}
	}
	return true
}

func readFile(name string) []string {
	file, err := os.Open(name)
	if err != nil {
		log.Fatalf("Couldn't open file: %v", err)
	}
	defer file.Close()

	vals := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		vals = append(vals, scanner.Text())
	}
	return vals
}

func getInstrReg(s string) (int, error) {
	vs := strings.Split(s, " ")
	if len(vs) != 2 {
		return 0, fmt.Errorf("splitting %s didn't work", s)
	}
	i, err := strconv.Atoi(vs[1])
	if err != nil {
		return 0, fmt.Errorf("error converting %q (from %q) to int: %v", vs[1], s, err)
	}
	return i, nil
}

type instruction struct {
	op    string
	input []int
}

func main() {
	vals := readFile("input.txt")
	instrReg, err := getInstrReg(vals[0])
	if err != nil {
		log.Fatalf("couldn't get instr register: %v", err)
	}

	instr := []instruction{}

	for i := 1; i < len(vals); i++ {
		ss := strings.Split(vals[i], " ")
		input, err := parseInput(ss)
		if err != nil {
			log.Fatalf("couldn't parse %q: %v", vals[i], err)
		}
		instr = append(instr, instruction{
			op:    ss[0],
			input: input,
		})
	}

	registers := []int{0, 0, 0, 0, 0, 0}
	for {
		next := registers[instrReg]
		if next >= len(instr) || next < 0 {
			break
		}
		todo := instr[next]
		registers, err = opcodes[todo.op](registers, todo.input)
		if err != nil {
			log.Fatalf("error executing %d (reg: %v): %v", next, registers, err)
		}
		registers[instrReg]++
	}
	fmt.Println(registers)
}
