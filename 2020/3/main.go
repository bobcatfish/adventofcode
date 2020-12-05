package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Map struct {
	M [][]byte
}

const tree = '#'

func (m *Map) IsTree(ry, rx int) bool {
	y, x := ry, rx%len(m.M[0])
	r := m.M[y][x] == tree
	return r
}

func count(m *Map, right, down int) int {
	c := 0
	for y, x := down, right; y < len(m.M); y, x = y+down, x+right {
		if m.IsTree(y, x) {
			c++
		}
	}
	return c
}

func load() (*Map, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, fmt.Errorf("couldn't open file: %v", err)
	}
	defer file.Close()

	m := [][]byte{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		val := scanner.Text()
		m = append(m, []byte(val))
		if err != nil {
			return nil, fmt.Errorf("couldn't parse password and rule from %s: %v", val, err)
		}
	}

	return &Map{M: m}, err
}

func main() {
	m, err := load()
	if err != nil {
		log.Fatalf("Couldn't load nums from file: %v", err)
	}

	c1 := count(m, 3, 1)
	fmt.Println("part1:", c1)

	p2 := make([]int, 5)
	vals := [][]int{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}
	a2 := 1

	for i, v := range vals {
		p2[i] = count(m, v[0], v[1])
		a2 *= p2[i]
	}
	fmt.Println(p2)
	fmt.Println(a2)

}
