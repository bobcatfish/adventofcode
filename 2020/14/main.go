package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func load() ([]string, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, fmt.Errorf("couldn't open file: %v", err)
	}
	defer file.Close()

	a := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		a = append(a, s)
	}
	return a, err
}

func getMask(v string) string {
	vv := strings.Split(v, " ")
	return vv[2]
}

func applyMask(val int, mask string) int {
	for i, r := range mask {
		v := 1 << uint(len(mask)-1-i)

		switch r {
		case 'X':
			continue
		case '1':
			val |= v
		case '0':
			val &= ^v
		}
	}
	return val
}

func apply(mem map[int]int, mask string, add, val int) {
	mem[add] = applyMask(val, mask)
}

func makeMasks(floating map[int]struct{}) []string {
	masks := []string{""}
	for i := 0; i < 36; i++ {
		if _, ok := floating[i]; ok {
			new := []string{}
			for j := range masks {
				new = append(new, masks[j]+"0")
				masks[j] = masks[j] + "1"
			}
			masks = append(masks, new...)
		} else {
			for j := range masks {
				masks[j] = masks[j] + "X"
			}
		}
	}
	return masks
}

func applyMask2(mem int, mask string) []int {
	floating := map[int]struct{}{}
	for i, r := range mask {
		v := 1 << uint(len(mask)-1-i)
		switch r {
		case 'X':
			floating[i] = struct{}{}
		case '1':
			mem |= v
		case '0':
			continue
		}
	}
	masks := makeMasks(floating)
	mems := []int{}
	for _, m := range masks {
		mems = append(mems, applyMask(mem, m))
	}

	return mems
}

func apply2(mem map[int]int, mask string, add, val int) {
	mems := applyMask2(add, mask)
	for _, m := range mems {
		mem[m] = val
	}
}

func extract(v string) (int, int) {
	re := regexp.MustCompile(`mem\[(\d+)\] = (\d+)$`)
	segs := re.FindStringSubmatch(v)
	add, _ := strconv.Atoi(segs[1])
	val, _ := strconv.Atoi(segs[2])
	return add, val
}

type applyFunc func(mem map[int]int, mask string, add, val int)

func doit(a []string, f applyFunc) {
	mem := map[int]int{}
	var mask string

	for _, v := range a {
		if strings.HasPrefix(v, "mask") {
			mask = getMask(v)
		} else {
			add, val := extract(v)
			f(mem, mask, add, val)
		}
	}

	total := 0
	for _, v := range mem {
		total += v
	}
	fmt.Println(total)
}

func main() {
	a, err := load()
	if err != nil {
		log.Fatalf("Couldn't load from file: %v", err)
	}
	doit(a, apply)
	doit(a, apply2)
}
