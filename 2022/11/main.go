package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type monkey struct {
	id          int
	items       []int
	op          func(int) int
	divis       int
	trueMonkey  int
	falseMonkey int
	inspections int
}

func load() []*monkey {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("couldn't open file: %v", err)
	}
	defer file.Close()

	vals := []*monkey{}
	scanner := bufio.NewScanner(file)
	i := 0
	var m *monkey
	for scanner.Scan() {
		val := scanner.Text()
		ii := i % 7
		switch ii {
		case 0:
			m = &monkey{id: i / 7, items: []int{}}
			vals = append(vals, m)
		case 1:
			if ii == 1 {
				val = strings.Replace(val, ",", "", -1)
				vv := strings.Split(val, ":")
				vvv := strings.Split(vv[1], " ")
				for _, vs := range vvv {
					if len(vs) > 0 {
						item, err := strconv.Atoi(vs)
						if err != nil {
							log.Fatalf("couldn't convert item: %v", err)
						}
						m.items = append(m.items, item)
					}
				}
			}
		case 2:
			var amount int
			re := regexp.MustCompile(`Operation: new = old (.) (.*)`)
			vv := re.FindStringSubmatch(val)
			if vv[2] != "old" {
				amountint, err := strconv.Atoi(vv[2])
				if err != nil {
					log.Fatalf("couldn't convert amount: %v", err)
				}
				amount = amountint
			}
			m.op = func(item int) int {
				if vv[2] == "old" {
					amount = item
				}
				switch vv[1] {
				case "*":
					return item * amount
				case "+":
					return item + amount
				default:
					log.Fatalf("unknown operation %c", vv[1])
				}
				return 0
			}
		case 3, 4, 5:
			vv := strings.Split(val, " ")
			amount, err := strconv.Atoi(vv[len(vv)-1])
			if err != nil {
				log.Fatalf("couldn't convert amount: %v", err)
			}
			switch ii {
			case 3:
				m.divis = amount
			case 4:
				m.trueMonkey = amount
			case 5:
				m.falseMonkey = amount
			}
		}
		i++
	}
	return vals
}

func main() {
	vals := load()
	for _, v := range vals {
		fmt.Println(*v)
	}

	for i := 0; i < 20; i++ {
		for _, m := range vals {
			for _, item := range m.items {
				m.inspections++
				item = m.op(item)
				item /= 3
				var newMonkey int
				mod := item % m.divis
				if mod == 0 {
					newMonkey = m.trueMonkey
				} else {
					newMonkey = m.falseMonkey
				}
				vals[newMonkey].items = append(vals[newMonkey].items, item)
			}
			m.items = []int{}
		}
	}
	for _, m := range vals {
		fmt.Println(m.id, m.inspections)
	}
	sort.SliceStable(vals, func(i, j int) bool {
		return vals[i].inspections > vals[j].inspections
	})
	fmt.Println(vals[0].inspections, vals[1].inspections, vals[0].inspections*vals[1].inspections)
}
