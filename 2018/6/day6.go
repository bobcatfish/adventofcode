package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type point struct {
	x int
	y int
}

type startingPoint struct {
	// Adjadcent points, indexed by their distance away.
	// At each distance, we have a map where the keys are the
	// adjacent points.
	adjByDist map[int]map[point]bool
	area      int
	p         point
}

// foundIntersection keeps track of points that have been found at a distance
// from a startingPoint.
type foundIntersection struct {
	startingArea *int
	startPoint   point
	dist         int
}

func getAdj(p point) []point {
	return []point{
		// above
		{x: p.x, y: p.y + 1},

		// beside
		{x: p.x - 1, y: p.y},
		{x: p.x + 1, y: p.y},

		// below
		{x: p.x, y: p.y - 1},
	}

}

func stepOut(start *startingPoint, dist int, adj map[point]bool, found map[point]*foundIntersection) {
	start.adjByDist[dist] = map[point]bool{}
	for p := range adj {
		for _, adjP := range getAdj(p) {
			start.adjByDist[dist][adjP] = true
			prev, ok := found[adjP]
			if ok {
				if prev.startPoint != start.p && prev.dist >= dist {
					if prev.startingArea != nil {
						*prev.startingArea--
						prev.startingArea = nil
					}
				}
			} else {
				start.area++
				found[adjP] = &foundIntersection{
					startPoint:   start.p,
					startingArea: &start.area,
					dist:         dist,
				}
			}
		}
	}
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

	starts := []startingPoint{}
	for _, v := range vals {
		vs := strings.Split(v, ", ")
		x, err := strconv.Atoi(vs[0])
		if err != nil {
			log.Fatalf("Error converting %q to point: %v", v, err)
		}
		y, err := strconv.Atoi(vs[1])
		if err != nil {
			log.Fatalf("Error converting %q to point: %v", v, err)
		}
		startPoint := point{
			x: x,
			y: y,
		}
		start := startingPoint{
			p: startPoint,
			adjByDist: map[int]map[point]bool{
				0: map[point]bool{
					startPoint: true,
				},
			},
			area: 1,
		}
		starts = append(starts, start)
	}

	// The keys here are points that are adjacent to other points.
	// The values are the maps from that starting point that is adjacent,
	// so that if we find another adjacent point, we can remove them
	found := map[point]*foundIntersection{}

	// seed the known points in
	for i := range starts {
		start := starts[i]
		found[start.p] = &foundIntersection{
			startPoint:   start.p,
			startingArea: &start.area,
			dist:         0,
		}
	}
	// this will be used to detect values that aren't infinitely increasing
	prevCounts := map[point]int{}
	for i := 1; ; i++ {
		for j := range starts {
			start := &starts[j]
			stepOut(start, i, start.adjByDist[i-1], found)
		}

		// This is the "best" idea I have for detecting infinity
		fmt.Printf("%d...", i)
		if i%10 == 0 {
			fmt.Println()
			candidates := []int{}
			for _, start := range starts {
				if prevCounts[start.p] == start.area {
					candidates = append(candidates, start.area)
				}
			}
			if len(candidates) > 0 {
				sort.Ints(candidates)
				fmt.Println("potential part 1 answer:", candidates[len(candidates)-1])
			}
		}

		for _, start := range starts {
			prevCounts[start.p] = start.area
		}

	}
}
