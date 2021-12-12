package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

type node struct {
	Name  string
	Conns []*node
	Big   bool
}

func getNode(name string, nodes map[string]*node) *node {
	var ok bool
	var n *node
	if n, ok = nodes[name]; !ok {
		n = &node{Name: name, Conns: []*node{}}
		n.Big = unicode.IsUpper([]rune(name)[0])
		nodes[name] = n
	}
	return n
}

func load() (map[string]*node, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, fmt.Errorf("couldn't open file: %v", err)
	}
	defer file.Close()

	m := map[string]*node{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		val := scanner.Text()
		ss := strings.Split(val, "-")
		node1 := getNode(ss[0], m)
		node2 := getNode(ss[1], m)

		node1.Conns = append(node1.Conns, node2)
		node2.Conns = append(node2.Conns, node1)
	}
	return m, err
}

func copySeen(s map[*node]int) map[*node]int {
	seen := make(map[*node]int, len(s))
	for k, v := range s {
		seen[k] = v
	}
	return seen
}

func getPaths(curr *node, paths [][]*node, nodes map[string]*node, alreadySeen map[*node]int, small int) [][]*node {
	seen := copySeen(alreadySeen)
	if !curr.Big {
		if curr.Name == "start" || curr.Name == "end" {
			seen[curr] = small
		} else {
			seen[curr] += 1
			if seen[curr] == small {
				small = 1
			}
		}
	}

	newPaths := [][]*node{}
	if len(paths) == 0 {
		newPaths = [][]*node{{curr}}
	} else {
		for _, p := range paths {
			newPath := make([]*node, len(p))
			copy(newPath, p)
			newPath = append(newPath, curr)
			newPaths = append(newPaths, newPath)
		}
	}

	nextPaths := [][]*node{}
	if curr.Name != "end" {
		for _, n := range curr.Conns {
			if i, _ := seen[n]; i < small {
				nextPaths = append(nextPaths, getPaths(n, newPaths, nodes, seen, small)...)
			}
		}
	}
	if len(nextPaths) == 0 {
		return newPaths
	}
	return nextPaths
}

func filterValidPaths(paths [][]*node) [][]*node {
	valid := [][]*node{}
	for _, p := range paths {
		if p[len(p)-1].Name == "end" {
			valid = append(valid, p)
		}
	}
	return valid
}

func main() {
	nodes, err := load()
	if err != nil {
		log.Fatalf("Couldn't load nums from file: %v", err)
	}

	paths := getPaths(nodes["start"], [][]*node{}, nodes, map[*node]int{}, 1)
	paths = filterValidPaths(paths)
	fmt.Println(len(paths))

	paths = getPaths(nodes["start"], [][]*node{}, nodes, map[*node]int{}, 2)
	paths = filterValidPaths(paths)
	fmt.Println(len(paths))
}
