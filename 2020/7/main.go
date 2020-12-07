package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func load() ([]string, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, fmt.Errorf("couldn't open file: %v", err)
	}
	defer file.Close()

	vals := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		vals = append(vals, scanner.Text())
	}
	return vals, err
}

type Bag struct {
	Color   string
	Parents []*Bag
}

func parse(v string) (string, []string, error) {
	re := regexp.MustCompile(`^(.+) bags contains? (.+)$`)
	vv := re.FindStringSubmatch(v)
	if len(vv) != 3 {
		return "", nil, fmt.Errorf("unexpected rule %q split into %d pieces %v", v, len(vv), vv)
	}
	p := vv[1]
	c := []string{}

	if vv[2] != "no other bags." {
		for _, cc := range strings.Split(vv[2], ", ") {
			re := regexp.MustCompile(`^(\d+) (.+ .+) bag`)
			segs := re.FindStringSubmatch(cc)
			if len(segs) != 3 {
				return "", nil, fmt.Errorf("couldn't parse bag contents %q got %d from regex", cc, len(segs))
			}
			c = append(c, segs[2])
		}
	}
	return p, c, nil
}

func getBags(vals []string) (map[string]*Bag, error) {
	bm := map[string]*Bag{}
	for _, v := range vals {
		p, c, err := parse(v)
		if err != nil {
			return nil, fmt.Errorf("couldn't get bag from %q: %v", v, err)
		}

		var b *Bag
		var ok bool
		if b, ok = bm[p]; !ok {
			b = &Bag{Color: p, Parents: []*Bag{}}
			bm[p] = b
		}

		for _, cc := range c {
			var cb *Bag
			if cb, ok = bm[cc]; !ok {
				cb = &Bag{Color: cc, Parents: []*Bag{}}
				bm[cc] = cb
			}
			cb.Parents = append(cb.Parents, b)
		}
	}
	return bm, nil
}

func allParents(b *Bag, bm map[string]*Bag) map[*Bag]struct{} {
	parents := map[*Bag]struct{}{}
	for _, b := range b.Parents {
		parents[b] = struct{}{}
	}
	for _, p := range b.Parents {
		pp := allParents(bm[p.Color], bm)

		for b, _ := range pp {
			parents[b] = struct{}{}
		}
	}
	return parents
}

func main() {
	vals, err := load()
	if err != nil {
		log.Fatalf("Couldn't load from file: %v", err)
	}
	bm, err := getBags(vals)
	if err != nil {
		log.Fatalf("Couldn't parse bags: %v", err)
	}
	parents := allParents(bm["shiny gold"], bm)
	fmt.Printf("%d out of %d\n", len(parents), len(bm))

}
