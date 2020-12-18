package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func load() ([]string, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, fmt.Errorf("couldn't open file: %v", err)
	}
	defer file.Close()

	lines := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		lines = append(lines, s)
	}
	return lines, err
}

type do struct {
	Op  rune
	Val int
}

func (d do) String() string {
	s := ""
	if d.Op != 0 {
		s += fmt.Sprintf("Op: %c", d.Op)
	}
	if d.Val != 0 {
		if s != "" {
			s += ", "
		}
		s += fmt.Sprintf("Val: %d", d.Val)
	}
	return s
}

func eval(s string) (int, int) {

	d := []do{}

	i := 0
	for ; i < len(s); i++ {
		r := rune(s[i])

		v := -1
		if r == ' ' {
			continue
		} else if r == '(' {
			var ii int
			v, ii = eval(s[i+1:])
			d = append(d, do{Val: v})
			i += ii
		} else if r == ')' {
			i++
			break
		} else if r == '+' || r == '*' {
			d = append(d, do{Op: r})
		} else {
			var err error
			v, err = strconv.Atoi(string(r))
			if err != nil {
				panic(fmt.Sprintf("im lazy! %c", r))
			}
			d = append(d, do{Val: v})
		}

	}

	d2 := []do{}

	c := -1
	for _, dd := range d {
		if dd.Op == 0 {
			if c == -1 {
				c = dd.Val
			} else {
				c += dd.Val
			}
		} else {
			if dd.Op != '+' {
				d2 = append(d2, do{Val: c})
				d2 = append(d2, do{Op: dd.Op})
				c = -1
			}

		}
	}
	if c != -1 {
		d2 = append(d2, do{Val: c})
	}

	c = -1
	for _, dd := range d2 {
		if dd.Op == 0 {
			if c == -1 {
				c = dd.Val
			} else {
				c *= dd.Val
			}
		}
	}

	return c, i
}

func main() {
	lines, err := load()
	if err != nil {
		log.Fatalf("Couldn't load from file: %v", err)
	}

	sum := 0
	for _, line := range lines {
		v, _ := eval(line)
		fmt.Println(v)
		sum += v
	}
	fmt.Println(sum)
}
