package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	a = 1
	b = 2
	c = 3
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
		return []int{}, fmt.Errorf("")
	}
	newReg := copyReg(reg)
	newReg[input[c]] = input[a]
	return newReg, nil
}

type test func(int, int) bool

func ir(reg []int, input []int, t test) ([]int, error) {
	if input[c] >= len(reg) || input[b] >= len(reg) {
		return []int{}, fmt.Errorf("")
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
		return []int{}, fmt.Errorf("")
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
		return []int{}, fmt.Errorf("")
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

var opcodes = []opcode{
	addr, addi, multr, multi, banr, bani, borr, bori, setr, seti, gtir, gtri, gtrr, eqir, eqri, eqrr,
}

type sample struct {
	before []int
	after  []int
	input  []int
}

func parseRegisters(s string) ([]int, error) {
	r, err := regexp.Compile(".*: *\\[(\\d), (\\d), (\\d), (\\d)\\]")
	if err != nil {
		return []int{}, fmt.Errorf("couldn't compile register regex: %v", err)
	}
	m := r.FindStringSubmatch(s)
	if len(m) != 5 {
		return []int{}, fmt.Errorf("didn't find all expected values in %q", s)
	}

	is := []int{}
	for i := 1; i < 5; i++ {
		v, err := strconv.Atoi(m[i])
		if err != nil {
			return []int{}, fmt.Errorf("couldn't parse int from %q", m[i])
		}
		is = append(is, v)
	}
	return is, nil
}

func parseInput(m []string) ([]int, error) {
	is := []int{}
	for _, v := range m {
		i, err := strconv.Atoi(v)
		if err != nil {
			return []int{}, fmt.Errorf("couldn't parse int from %q", v)
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

func main() {
	file, err := os.Open("input_1.txt")
	if err != nil {
		log.Fatalf("Couldn't open file: %v", err)
	}
	defer file.Close()

	vals := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		vals = append(vals, scanner.Text())
	}

	samples := []sample{}
	for i := 0; i < len(vals); i += 4 {
		before, err := parseRegisters(vals[i])
		if err != nil {
			log.Fatalf("error parsing before from %q: %v", vals[i], err)
		}
		input, err := parseInput(strings.Split(vals[i+1], " "))
		if err != nil {
			log.Fatalf("error parsing input from %q: %v", vals[i+1], err)
		}
		after, err := parseRegisters(vals[i+2])
		if err != nil {
			log.Fatalf("error parsing after from %q: %v", vals[i+2], err)
		}
		sample := sample{
			before: before,
			input:  input,
			after:  after,
		}
		samples = append(samples, sample)
	}

	multis := 0

	for _, sample := range samples {
		count := 0
		for _, opc := range opcodes {
			after, err := opc(sample.before, sample.input)
			if err != nil {
				continue
			}
			if registerEq(after, sample.after) {
				count++
			}
		}
		fmt.Println("count", count)
		if count >= 3 {
			multis++
		}
	}

	fmt.Println("total samples", len(samples))
	fmt.Println("total multis", multis)
}
