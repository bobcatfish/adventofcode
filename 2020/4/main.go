package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func load() ([]map[string]string, error) {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		return nil, fmt.Errorf("couldn't read file: %v", err)
	}

	pp := []map[string]string{}
	for _, v := range strings.Split(string(data), "\n\n") {
		p := map[string]string{}
		for _, vp := range strings.Fields(v) {
			s := strings.Split(vp, ":")
			if len(s) != 2 {
				return nil, fmt.Errorf("got unexpected values in passport: %s", vp)
			}
			p[s[0]] = s[1]
		}
		pp = append(pp, p)

	}
	return pp, nil
}

type vFunc func(string) bool

var required = map[string]vFunc{
	"byr": byr,
	"iyr": iyr,
	"eyr": eyr,
	"hgt": hgt,
	"hcl": hcl,
	"ecl": ecl,
	"pid": pid,
}

func validInt(s string, min, max int) bool {
	y, err := strconv.Atoi(s)
	if err != nil {
		log.Printf("couldn't convert %s", s)
		return false
	}
	return y >= min && y <= max
}

func byr(s string) bool {
	return validInt(s, 1920, 2002)
}

func iyr(s string) bool {
	return validInt(s, 2010, 2020)
}

func eyr(s string) bool {
	return validInt(s, 2020, 2030)
}

func hgt(s string) bool {
	re := regexp.MustCompile(`(\d+)(\w*)`)
	segs := re.FindStringSubmatch(s)
	if len(segs) != 3 {
		log.Printf("couldn't parse height: %s", s)
		return false
	}
	if segs[2] == "cm" {
		return validInt(segs[1], 150, 193)
	}
	if segs[2] == "in" {
		return validInt(segs[1], 59, 76)
	}
	log.Printf("unit wasn't in or cm for %s: %s", s, segs[2])
	return false
}

func validRegex(p, s string) bool {
	pstrict := fmt.Sprintf("^%s$", p)
	matched, err := regexp.MatchString(pstrict, s)
	if err != nil {
		log.Printf("error applying regex to %s: %v", s, err)
		return false
	}
	return matched
}

func hcl(s string) bool {
	return validRegex("#[0-9,a-f]{6}", s)
}

var validEyeColors = map[string]struct{}{
	"amb": struct{}{},
	"blu": struct{}{},
	"brn": struct{}{},
	"gry": struct{}{},
	"grn": struct{}{},
	"hzl": struct{}{},
	"oth": struct{}{},
}

func ecl(s string) bool {
	_, ok := validEyeColors[s]
	return ok
}

func pid(s string) bool {
	return validRegex("[0-9]{9}", s)
}

func countValid(pp []map[string]string) (int, int) {
	c1, c2 := 0, 0
	for _, p := range pp {
		missing, valid := false, true
		for r, f := range required {
			if v, ok := p[r]; !ok {
				missing = true
			} else if !f(v) {
				valid = false
			}
		}
		if !missing {
			c1++
			if valid {
				//fmt.Println(p)
				c2++
			}
		}
	}
	return c1, c2
}

func main() {
	pp, err := load()
	if err != nil {
		log.Fatalf("Couldn't load nums from file: %v", err)
	}

	valid1, valid2 := countValid(pp)
	fmt.Println("part1", valid1)
	fmt.Println("part2", valid2)
}
