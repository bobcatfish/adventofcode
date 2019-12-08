package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	root = "COM"
)

type orbit struct {
	name string
	to   []*orbit
	from *orbit
}

type pair struct {
	from string
	to   string
}

func load() ([]pair, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, fmt.Errorf("couldn't open file: %v", err)
	}
	defer file.Close()

	orbits := []pair{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		val := scanner.Text()
		v := strings.Split(val, ")")
		orbits = append(orbits, pair{from: v[0], to: v[1]})
	}
	return orbits, err
}

func makeOrbits(orbits []pair) map[string]*orbit {
	o := map[string]*orbit{}
	for _, p := range orbits {
		o[p.from] = &orbit{name: p.from}
	}

	for _, p := range orbits {
		f, t := p.from, p.to
		var too *orbit
		fo, _ := o[f]
		if to, ok := o[t]; ok {
			too = to
		} else {
			too = &orbit{name: t}
			o[t] = too
		}
		fo.to = append(fo.to, too)
		too.from = fo
	}
	return o
}

func walk(o *orbit) int {
	seen := map[string]struct{}{}
	next := []*orbit{o}
	w := -2

	for {
		nnext := []*orbit{}
		for _, o := range next {
			fmt.Printf(o.name)
			if o.name == "SAN" {
				fmt.Println()
				return w
			}
			seen[o.name] = struct{}{}

			if o.from != nil {
				if _, ok := seen[o.from.name]; !ok {
					nnext = append(nnext, o.from)
				}
			}

			for _, to := range o.to {
				if _, ok := seen[to.name]; !ok {
					nnext = append(nnext, to)
				}
			}
		}
		fmt.Println()
		next = nnext
		w += 1
	}

	return w
}

func main() {
	orbits, err := load()
	if err != nil {
		log.Fatalf("Couldn't load from file: %v", err)
	}

	o := makeOrbits(orbits)

	w := walk(o["YOU"])
	fmt.Println(w)
}
