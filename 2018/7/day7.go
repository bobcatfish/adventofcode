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

const (
	numWorkers = 5
	baseSec    = 60
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

func getAllAvailable(s []*node, numWorkers int) ([]*node, []*node) {
	sort.Sort(byID(s))
	if len(s) > numWorkers {
		return s[:numWorkers], s[numWorkers:]
	}
	return s, []*node{}
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

func getReady(candidates []*node, inProgress map[rune]*node, finished map[rune]*node) (*node, []*node) {
	for _, n := range finished {
		for _, candidate := range n.next {
			alreadyPresent := false
			for _, c := range candidates {
				if c.id == candidate.id {
					alreadyPresent = true
				}
			}
			_, alreadyFinished := finished[rune(candidate.id[0])]
			_, currentlyUnderway := inProgress[rune(candidate.id[0])]

			if !alreadyPresent && !alreadyFinished && !currentlyUnderway {
				parentsFinished := true
				for _, parent := range candidate.prev {
					_, parentFinished := finished[rune(parent.id[0])]
					if !parentFinished {
						parentsFinished = false
					}
				}
				if parentsFinished {
					candidates = append(candidates, candidate)
				}
			}
		}
	}
	if len(candidates) > 0 {
		return getNext(candidates)
	}
	return nil, candidates
}

type work struct {
	w    int
	id   rune
	node *node
}

func build(s []*node, totalLen int) int {
	finished := map[rune]*node{}
	inProgress := map[rune]*node{}
	workers := make([]*work, numWorkers)

	for i := 0; i < numWorkers; i++ {
		workers[i] = &work{}
	}

	t := 1
	for ; ; t++ {
		for i, w := range workers {
			if w.w > 0 {
				w.w--
			} else {
				var next *node
				next, s = getReady(s, inProgress, finished)
				if next != nil {
					inProgress[rune(next.id[0])] = next
					workers[i] = &work{
						w:    int(rune(next.id[0])-('A'-1)) + baseSec,
						id:   rune(next.id[0]),
						node: next,
					}
					// Count the current second as work
					workers[i].w--
				}
			}
		}
		for _, w := range workers {
			if w.w == 0 && w.node != nil {
				finished[w.id] = w.node
			}
		}
		if len(finished) == totalLen {
			break
		}
	}
	return t
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
	fmt.Println("part 1", strings.Join(order, ""))

	time := build(roots, len(order))
	fmt.Println("part 2", time)
}
