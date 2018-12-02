package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"runtime/debug"
	"strconv"
	"strings"
)

// I'm pretty sure the solution doens't work, who knows how I found the right answer!!

type Instr func() int

var count int

func getReg(reg string, registers map[string]int) int {
	_, err := strconv.Atoi(reg)
	if err == nil {
		log.Fatal(string(debug.Stack()))
	}
	if r, ok := registers[reg]; ok {
		return r
	} else {
		fmt.Println("new reg", reg)
		registers[reg] = 0
		return 0
	}
}

func printReg(registers map[string]int) {
	alpha := [8]string{"a", "b", "c", "d", "e", "f", "g", "h"}
	for _, a := range alpha {
		fmt.Printf("%v:%v, ", a, registers[a])
	}
	fmt.Println()
}

func magic1(x string, y string, registers map[string]int) Instr {
	return func() int {
		registers["e"] = registers["b"]
		registers["f"] = 0
		registers["g"] = 0
		return 1
	}
}

var happened int

func set(x string, y string, registers map[string]int) Instr {
	v, err := strconv.Atoi(y)
	if err != nil {
		return func() int {
			value := getReg(y, registers)
			registers[x] = value
			return 1
		}
	}
	return func() int {
		registers[x] = v
		return 1
	}
}

func sub(x string, y string, registers map[string]int) Instr {
	v, err := strconv.Atoi(y)
	if err != nil {
		return func() int {
			value := getReg(y, registers)
			xValue := getReg(x, registers)
			registers[x] = xValue - value
			return 1
		}
	}
	return func() int {
		if x == "d" {
			printReg(registers)
		}

		xValue := getReg(x, registers)
		registers[x] = xValue - v
		return 1
	}
}

func mul(x string, y string, registers map[string]int) Instr {
	v, err := strconv.Atoi(y)
	if err != nil {
		return func() int {
			count += 1
			value := getReg(y, registers)
			xValue := getReg(x, registers)
			registers[x] = xValue * value
			return 1
		}
	}
	return func() int {
		xValue := getReg(x, registers)
		registers[x] = xValue * v
		return 1
	}
}

func jnz(x string, y string, registers map[string]int) Instr {
	v, err := strconv.Atoi(y)
	if err != nil {
		return func() int {
			xv, err := strconv.Atoi(x)
			if err != nil {
				x := getReg(x, registers)
				if x != 0 {
					return getReg(y, registers)
				}
				return 1
			} else {
				if xv != 0 {
					return getReg(y, registers)
				}
				return 1
			}
		}
	}
	return func() int {
		xv, err := strconv.Atoi(x)
		if err != nil {
			x := getReg(x, registers)
			if x != 0 {
				return v
			}
			return 1
		} else {
			if xv != 0 {
				return v
			}
			return 1
		}
	}
}

func getInstr(line string, registers map[string]int) Instr {
	f := strings.Fields(line)
	switch f[0] {
	case "set":
		return set(f[1], f[2], registers)
	case "sub":
		return sub(f[1], f[2], registers)
	case "mul":
		return mul(f[1], f[2], registers)
	case "jnz":
		return jnz(f[1], f[2], registers)
	case "magic1":
		return magic1(f[1], f[2], registers)
	}
	log.Fatalf("did not find %v\n", line)
	return nil
}

func doit(instr []Instr, registers map[string]int) {
	i := 0
	for {
		i += instr[i]()
		if i >= len(instr) {
			break
		}
		if i < 0 {
			log.Fatalf("i less than 0 %v\n", i)
		}
	}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Couldn't open file: %v\n", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var instr []Instr
	registers := make(map[string]int)
	registers["a"] = 1

	for scanner.Scan() {
		line := scanner.Text()
		n := getInstr(line, registers)
		instr = append(instr, n)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	doit(instr, registers)
	fmt.Println(count)
}
