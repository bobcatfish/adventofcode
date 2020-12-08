package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func load() ([]Instr, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, fmt.Errorf("couldn't open file: %v", err)
	}
	defer file.Close()

	vals := []Instr{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		ss := strings.Fields(s)
		if len(ss) != 2 {
			return nil, fmt.Errorf("unexpected format in line %s, didn't split as expected %v", s, ss)
		}
		i, err := strconv.Atoi(ss[1])
		if err != nil {
			return nil, fmt.Errorf("unexpected format in line %s, couldn't convert offset", s, ss[1])
		}
		vals = append(vals, Instr{
			Type:   ss[0],
			Offset: i,
		})
	}
	return vals, err
}

type Instr struct {
	Type   string
	Offset int
}

type instr func(i, a, offset int) (int, int)

func acc(i, a, offset int) (int, int) {
	return i + 1, a + offset
}

func jmp(i, a, offset int) (int, int) {
	return i + offset, a
}

func nop(i, a, offset int) (int, int) {
	return i + 1, a
}

var instrs = map[string]instr{
	"acc": acc,
	"jmp": jmp,
	"nop": nop,
}

func run(ii []Instr) bool {
	a := 0
	seen := map[int]struct{}{}
	for i := 0; i < len(ii) && i >= 0; {
		fmt.Printf("i: %d, a: %d\n", i, a)
		if _, ok := seen[i]; ok {
			fmt.Printf("about to run %d again, accumulator is %d\n", i, a)
			return false
		}
		seen[i] = struct{}{}

		curr := ii[i]
		i, a = instrs[(curr.Type)](i, a, curr.Offset)
	}
	return true
}

func main() {
	ii, err := load()
	if err != nil {
		log.Fatalf("Couldn't load from file: %v", err)
	}

	for i := 0; i < len(ii); i++ {
		if ii[i].Type == "jmp" {
			c := make([]Instr, len(ii))
			copy(c, ii)
			c[i] = Instr{Type: "nop"}

			if run(c) {
				fmt.Println("found it!")
				break
			}
		}
	}
}
