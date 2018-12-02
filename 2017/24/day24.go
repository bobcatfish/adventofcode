package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const numPorts = 57

type port struct {
	A int
	B int
	T int
}

func isUsed(i uint, used uint64) bool {
	u := used & (1 << i)
	return u == 0
}

func setUsed(i uint, used uint64) uint64 {
	newUsed := used
	newUsed ^= (1 << i)
	return newUsed
}

func getNeed(p port, need int) int {
	if need == p.A {
		return p.B
	}
	return p.A
}

func explore(ports [numPorts]port, m map[int][]uint, used uint64, need int, total int, seen map[uint64]bool, len int) (int, int) {
	var totals []int
	var lens []int

	for _, v := range m[need] {
		if !isUsed(v, used) {
			newNeed := getNeed(ports[v], need)
			newUsed := setUsed(v, used)

			_, present := seen[newUsed]
			if !present {
				seen[newUsed] = true
				newLen, newTotal := explore(ports, m, newUsed, newNeed, total+ports[v].T, seen, len+1)
				totals = append(totals, newTotal)
				lens = append(lens, newLen)
			}
		}
	}

	maxLen := len
	maxI := -1
	for i, l := range lens {
		if l > maxLen {
			maxLen = l
			maxI = i
		} else if l == maxLen {
			if totals[i] > totals[maxI] {
				maxI = i
			}
		}
	}
	if maxI == -1 {
		return len, total
	}

	return maxLen, totals[maxI]
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Couldn't open file: %v\n", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var ports [numPorts]port
	m := make(map[int][]uint)
	i := 0

	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Split(line, "/")

		a, err := strconv.Atoi(values[0])
		if err != nil {
			log.Fatal(err)
		}

		b, err := strconv.Atoi(values[1])
		if err != nil {
			log.Fatal(err)
		}
		ports[i] = port{A: a, B: b, T: a + b}

		aa, ok := m[a]
		if ok {
			m[a] = append(aa, uint(i))
		} else {
			m[a] = []uint{uint(i)}
		}
		bb, ok := m[b]
		if ok {
			m[b] = append(bb, uint(i))
		} else {
			m[b] = []uint{uint(i)}
		}
		i++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var used uint64
	used = 0xFFFFFFFFFFFFFFFF
	seen := make(map[uint64]bool)

	l, v := explore(ports, m, used, 0, 0, seen, 1)

	fmt.Println(l)
	fmt.Println(v)
}
