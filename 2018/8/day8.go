package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type node struct {
	children []node
	m        []int
}

func get(vs []int, i, c, m int) (int, []node, []int) {
	if i >= len(vs) {
		fmt.Println("im surprised we got here")
		return 0, []node{}, []int{}
	}
	newi := i
	children := []node{}
	for ci := 0; ci < c; ci++ {
		cc, cm := vs[newi], vs[newi+1]
		newi += 2
		var nextChildren []node
		var nextMetadata []int
		newi, nextChildren, nextMetadata = get(vs, newi, cc, cm)
		child := node{children: nextChildren, m: nextMetadata}
		children = append(children, child)
	}

	nextMeta := []int{}
	for mi := 0; mi < m; mi++ {
		nextMeta = append(nextMeta, vs[newi])
		newi++
	}
	return newi, children, nextMeta
}

func count(n node, sum int) int {
	for _, c := range n.children {
		sum = count(c, sum)
	}
	for _, m := range n.m {
		sum += m
	}
	return sum
}

func fancyCount(n node, sum int) int {
	if len(n.children) == 0 {
		for _, m := range n.m {
			sum += m
		}
	} else {
		for _, m := range n.m {
			if m == 0 || m > len(n.children) {
				continue
			}
			sum = fancyCount(n.children[m-1], sum)
		}
	}
	return sum
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Couldn't open file: %v", err)
	}
	defer file.Close()

	vals := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		vals = append(vals, scanner.Text())
	}
	vsStr := strings.Split(vals[0], " ")
	vs := []int{}
	for _, v := range vsStr {
		i, err := strconv.Atoi(v)
		if err != nil {
			log.Fatalf("error converting to int %q: %v", v, err)
		}
		vs = append(vs, i)
	}

	c, m := vs[0], vs[1]
	_, children, metadata := get(vs, 2, c, m)
	n := node{children: children, m: metadata}

	sum := count(n, 0)
	fmt.Println("part 1", sum)

	sum = fancyCount(n, 0)
	fmt.Println("part 2", sum)
}
