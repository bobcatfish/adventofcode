package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func load() (int, []int, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return 0, nil, fmt.Errorf("couldn't open file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	if !scanner.Scan() {
		return 0, nil, fmt.Errorf("couldn't read line with timestamp")
	}
	tss := scanner.Text()

	ts, err := strconv.Atoi(tss)
	if err != nil {
		return 0, nil, fmt.Errorf("couldn't convert timestamp %s: %v", tss, err)
	}

	if !scanner.Scan() {
		return 0, nil, fmt.Errorf("couldn't read line with buses")
	}
	bs := scanner.Text()
	b := []int{}

	for _, s := range strings.Split(bs, ",") {
		if s == "x" {
			b = append(b, 0)
			continue
		}
		bt, err := strconv.Atoi(s)
		if err != nil {
			return 0, nil, fmt.Errorf("couldn't convert bus %s: %v", s, err)
		}
		b = append(b, bt)
	}

	return ts, b, nil
}

func findDepartures(ts int, b []int) map[int]int {
	d := map[int]int{}
	for _, v := range b {
		if v == 0 {
			continue
		}
		if ts%v == 0 {
			d[v] = ts
		} else {
			m := ts / v
			d[v] = v * (m + 1)
		}
	}
	return d
}

func f(first, second, offset int) (int, int) {
	//fmt.Println(first, second, offset)
	vals := []int{}

	for j := 0; ; j++ {
		val := first * j

		if (val+offset)%second == 0 {
			vals = append(vals, val)
		}
		if len(vals) == 2 {
			break
		}
	}
	return vals[0], vals[1] - vals[0]
}

func keep(num, offset, nstart, noffset int) (int, int) {
	vals := []int{}

	for poss := nstart; ; poss += noffset {
		val := poss - offset

		if val%num == 0 {
			vals = append(vals, val)
		}
		if len(vals) == 2 {
			break
		}

	}
	return vals[0], vals[1] - vals[0]
}

func find(ts int, b []int) int {
	vals := [][]int{}

	for i, v := range b {
		if v != 0 {
			vals = append(vals, []int{v, i})
		}
	}
	first, second := vals[len(vals)-2], vals[len(vals)-1]
	start, offset := f(first[0], second[0], second[1]-first[1])

	for i := len(vals) - 3; i >= 0; i-- {
		next := vals[i]
		start, offset = keep(next[0], first[1]-next[1], start, offset)
		first = next
	}

	fmt.Println(start, offset)
	return start
}

func main() {
	ts, b, err := load()
	if err != nil {
		log.Fatalf("Couldn't load from file: %v", err)
	}
	fmt.Println(ts, b)

	d := findDepartures(ts, b)

	eBus, eTime := -1, int(^uint(0)>>1)
	for bus, time := range d {
		if time < eTime {
			eBus, eTime = bus, time
		}
	}
	minutes := eTime - ts
	fmt.Println(eBus, eTime, minutes, eBus*minutes)

	p2 := find(ts, b)
	fmt.Println(p2)

}
