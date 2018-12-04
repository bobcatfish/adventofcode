package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"time"
)

func isNewGuard(s string) (bool, int, error) {
	r, err := regexp.Compile("Guard #(\\d*) begins shift")
	if err != nil {
		return false, 0, fmt.Errorf("error compiling regex: %v", err)
	}
	m := r.FindStringSubmatch(s)
	if len(m) > 0 {
		i, err := strconv.Atoi(m[1])
		if err != nil {
			return false, 0, fmt.Errorf("error converting guard id %q to int: %v", m[1], err)
		}
		return true, i, nil
	}
	return false, 0, nil
}

type entry struct {
	date  time.Time
	entry string
}

type byDate []entry

func (d byDate) Len() int           { return len(d) }
func (d byDate) Less(i, j int) bool { return d[i].date.Before(d[j].date) }
func (d byDate) Swap(i, j int)      { d[i], d[j] = d[j], d[i] }

type minSpan struct {
	sleep int
	wake  int
}

type shift struct {
	// do not need start
	start    int
	sleeping []*minSpan
}

func constructShifts(entries []entry) (map[int]*shift, error) {
	shifts := map[int]*shift{}
	var g int
	var span = &minSpan{}
	for _, e := range entries {
		isNew, id, err := isNewGuard(e.entry)
		if err != nil {
			log.Fatalf("error checking for guard swap %q: %v", e.entry, err)
		}
		if isNew {
			g = id
			if _, ok := shifts[g]; !ok {
				shifts[g] = &shift{
					// this is not quite right
					start:    e.date.Minute(),
					sleeping: []*minSpan{},
				}
			}
		} else if e.entry == "falls asleep" {
			span = &minSpan{}
			shifts[g].sleeping = append(shifts[g].sleeping, span)
			span.sleep = e.date.Minute()
		} else if e.entry == "wakes up" {
			span.wake = e.date.Minute()
		} else {
			return map[int]*shift{}, fmt.Errorf("didn't recognize entry %q", e.entry)
		}
	}
	return shifts, nil
}

func findMax(shifts map[int]*shift) (int, *shift) {
	maxMin := 0
	maxGuard := 0
	var maxShift *shift

	for k, v := range shifts {
		min := 0

		for _, s := range v.sleeping {
			min += (s.wake - s.sleep)
		}

		if min > maxMin {
			maxMin = min
			maxGuard = k
			maxShift = v
		}
	}

	return maxGuard, maxShift
}

func findMostMin(hackMap map[int]int) (int, int) {
	maxK, maxV := 0, 0
	for k, v := range hackMap {
		if v > maxV {
			maxK = k
			maxV = v
		}
	}
	return maxK, maxV
}

func getGuardMinMap(v *shift) map[int]int {
	hackMap := map[int]int{}

	for _, s := range v.sleeping {
		m := s.sleep
		for ; m < s.wake; m++ {
			if _, ok := hackMap[m]; !ok {
				hackMap[m] = 0
			}
			hackMap[m]++
		}
	}
	return hackMap
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

	r, err := regexp.Compile("\\[(.*)\\] (.*)")
	if err != nil {
		log.Fatalf("error compiling regex: %v", err)
	}

	entries := []entry{}
	for _, v := range vals {
		m := r.FindStringSubmatch(v)
		if len(m) == 0 {
			log.Fatalf("error parsing string %q for sorting: %v", v, err)
		}
		layout := "2006-01-02 15:04"
		t, err := time.Parse(layout, m[1])
		if err != nil {
			log.Fatalf("error parsing time %q: %v", v, err)
		}
		entries = append(entries, entry{
			date:  t,
			entry: m[2],
		})
	}

	sort.Sort(byDate(entries))

	shifts, err := constructShifts(entries)
	if err != nil {
		log.Fatalf("error constructing shifts: %v", err)
	}

	maxGuard, maxShift := findMax(shifts)
	fmt.Println("guard", maxGuard)
	hackMap := getGuardMinMap(maxShift)

	maxMin, minCount := findMostMin(hackMap)
	fmt.Println("max min", maxMin, "count", minCount)
	fmt.Println("part 1 solution", maxGuard*maxMin)

	minMaps := map[int]map[int]int{}

	for guard, shift := range shifts {
		hackMap := getGuardMinMap(shift)
		minMaps[guard] = hackMap
	}

	maxMinGuard := 0
	maxMinCount := 0
	maxMinMin := 0
	for guard, minMap := range minMaps {
		for minute, count := range minMap {
			if count > maxMinCount {
				maxMinCount = count
				maxMinMin = minute
				maxMinGuard = guard
			}
		}
	}

	fmt.Println("max min min", maxMinMin)
	fmt.Println("max min guard", maxMinGuard)
	fmt.Println("part 2 solution", maxMinMin*maxMinGuard)
}
