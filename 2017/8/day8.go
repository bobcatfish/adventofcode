package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// GLOBAL VARIABLES LIKE A BOSS 😎
var registers map[string]int
var highestHeld int

func getVal(reg string) int {
	if v, ok := registers[reg]; ok {
		return v
	}
	registers[reg] = 0
	return 0
}

type comp struct {
	Reg  string
	comp string
	Val  int
}

type action struct {
	Reg    string
	action string
	Val    int
	C      comp
}

func compMet(c comp) bool {
	lVal := registers[c.Reg]
	rVal := c.Val

	switch c.comp {
	case ">":
		return lVal > rVal
	case "<":
		return lVal < rVal
	case "==":
		return lVal == rVal
	case ">=":
		return lVal >= rVal
	case "<=":
		return lVal <= rVal
	case "!=":
		return lVal != rVal
	}
	log.Fatalf("Did not recognize %v", c.comp)
	return false
}

func doAction(a action) {
	if compMet(a.C) {
		switch a.action {
		case "inc":
			registers[a.Reg] += a.Val
		case "dec":
			registers[a.Reg] -= a.Val
		}
	}
	if registers[a.Reg] > highestHeld {
		highestHeld = registers[a.Reg]
	}
}

func actionFromLine(line string) action {
	fields := strings.Fields(line)
	val, err := strconv.Atoi(fields[2])
	if err != nil {
		log.Fatal(err)
	}
	compVal, err := strconv.Atoi(fields[6])
	if err != nil {
		log.Fatal(err)
	}
	return action{
		Reg:    fields[0],
		action: fields[1],
		Val:    val,
		C: comp{
			Reg:  fields[4],
			comp: fields[5],
			Val:  compVal,
		},
	}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Couldn't open file: %v\n", err)
	}
	defer file.Close()

	registers = make(map[string]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		a := actionFromLine(line)
		doAction(a)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var highest int
	for _, v := range registers {
		if v > highest {
			highest = v
		}
	}
	fmt.Printf("Highest value in register: %v\n", highest)
	fmt.Printf("Highest value ever held: %v\n", highestHeld)
}
