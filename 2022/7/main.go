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

type file struct {
	name string
	size int
}

type dir struct {
	name   string
	parent *dir
	dirs   map[string]*dir
	files  []file
	size   int
}

func load() []string {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("couldn't open file: %v", err)
	}
	defer file.Close()

	vals := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		val := scanner.Text()
		vals = append(vals, val)
	}
	return vals
}

func update(v string, d *dir) {
	vv := strings.Split(v, " ")
	if len(vv) != 2 {
		log.Fatalf("expected string to have 2 parts: %s", v)
	}

	if vv[0] == "dir" {
		d.dirs[vv[1]] = &dir{
			name:   vv[1],
			parent: d,
			dirs:   map[string]*dir{},
			files:  []file{},
		}
	} else {
		size, err := strconv.Atoi(vv[0])
		if err != nil {
			log.Fatalf("expected string to start with size: %s", v)
		}
		d.files = append(d.files, file{
			name: vv[1],
			size: size,
		})
	}
}

func doCommand(v string, d *dir) *dir {
	vv := strings.Split(v, " ")
	if len(vv) < 2 {
		log.Fatalf("expected string to have at least 2 parts: %s", v)
	}
	if vv[1] == "ls" {
		return d
	}
	if vv[1] != "cd" || len(vv) != 3 {
		log.Fatalf("did not get expected `cd` command: %s", v)
	}
	if vv[2] == ".." {
		return d.parent
	}
	return d.dirs[vv[2]]
}

func next(v string, d *dir) *dir {
	if v[0] == '$' {
		d = doCommand(v, d)
	} else {
		update(v, d)
	}
	return d
}

func printDir(d *dir, dent string) {
	fmt.Println(dent, d.name)
	for _, dd := range d.dirs {
		printDir(dd, dent+" ")
	}
	for _, f := range d.files {
		fmt.Println(dent+" ", f.name, f.size)
	}
}

func (d *dir) getSize() int {
	if d.size == 0 {
		for _, f := range d.files {
			d.size += f.size
		}
		for _, dd := range d.dirs {
			d.size += dd.getSize()
		}
	}
	return d.size
}

func doSizes(dirs map[string]*dir, threshold int) int {
	sum := 0
	for _, d := range dirs {
		size := d.getSize()
		if size <= threshold {
			sum += size
		}
		sum += doSizes(d.dirs, threshold)
	}
	return sum
}

func getDirs(d *dir) []*dir {
	dirs := []*dir{d}
	for _, dd := range d.dirs {
		dirs = append(dirs, getDirs(dd)...)
	}
	return dirs
}

func main() {
	vals := load()

	root := &dir{
		name:  "/",
		dirs:  map[string]*dir{},
		files: []file{},
	}

	d := root
	for _, v := range vals[1:] {
		d = next(v, d)
	}
	// printDir(root, "")

	size := doSizes(root.dirs, 100000)
	fmt.Println(size)

	totalSize := 70000000
	need := 30000000
	spaceNeeded := need - (totalSize - root.getSize())

	dirs := getDirs(root)
	sort.SliceStable(dirs, func(i, j int) bool {
		return dirs[i].size > dirs[j].size
	})

	smallest := dirs[0].size
	for _, d := range dirs {
		if d.size >= spaceNeeded && d.size < smallest {
			fmt.Println("candidate", d.name, d.size)
			smallest = d.size
		}
	}
	fmt.Println(smallest)
}
