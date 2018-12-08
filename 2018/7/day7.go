package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
)

type node struct {
	id   string
	prev []*node
	next []*node
}

type byID []*node

func (d byID) Len() int           { return len(d) }
func (d byID) Less(i, j int) bool { return d[i].id[0] < d[j].id[0] }
func (d byID) Swap(i, j int)      { d[i], d[j] = d[j], d[i] }

func getNext(s []*node) (*node, []*node) {
	sort.Sort(byID(s))
	return s[0], s[1:]
}

func canVisit(next *node, s []*node, visited map[string]bool) bool {
	for _, ss := range s {
		if ss.id == next.id {
			return false
		}
	}
	for _, prev := range next.prev {
		if _, ok := visited[prev.id]; !ok {
			return false
		}
	}
	return true
}

func traverse(s []*node, visited map[string]bool) []string {
	if len(s) == 0 {
		return []string{}
	}
	n, s := getNext(s)
	order := []string{n.id}
	visited[n.id] = true

	for _, next := range n.next {
		if canVisit(next, s, visited) {
			s = append(s, next)
		}
	}

	nextOrder := traverse(s, visited)
	for _, next := range nextOrder {
		order = append(order, next)
	}
	return order
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

	r, err := regexp.Compile("Step (.*) must be finished before step (.*) can begin.")
	if err != nil {
		log.Fatalf("Couldn't compile regex: %v", err)
	}

	nodeMap := map[string]*node{}
	for _, v := range vals {
		m := r.FindStringSubmatch(v)
		if len(m) != 3 {
			log.Fatalf("Did not find expected values in %q", v)
		}

		afterNode, ok := nodeMap[m[2]]
		if !ok {
			afterNode = &node{
				id:   m[2],
				prev: []*node{},
				next: []*node{},
			}
			nodeMap[m[2]] = afterNode
		}

		beforeNode, ok := nodeMap[m[1]]
		if !ok {
			beforeNode = &node{
				id:   m[1],
				prev: []*node{},
				next: []*node{},
			}
			nodeMap[m[1]] = beforeNode
		}

		beforeNode.next = append(beforeNode.next, afterNode)
		afterNode.prev = append(afterNode.prev, beforeNode)
	}

	var roots []*node
	for _, v := range nodeMap {
		if len(v.prev) == 0 {
			roots = append(roots, v)
		}
	}

	visited := map[string]bool{}
	order := traverse(roots, visited)
	fmt.Println(strings.Join(order, ""))
}
