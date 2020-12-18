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

func eval(s string) (int, int) {
	c := -1
	op := ' '
	for i := 0; i < len(s); i++ {
		r := rune(s[i])

		v := -1
		if r == ' ' {
			continue
		} else if r == '(' {
			var ii int
			v, ii = eval(s[i+1:])
			i += ii
		} else if r == ')' {
			return c, i + 1
		} else if r == '+' || r == '*' {
			op = r
		} else {
			var err error
			v, err = strconv.Atoi(string(r))
			if err != nil {
				panic(fmt.Sprintf("im lazy! %c", r))
			}
		}

		if v != -1 {
			if op == '*' {
				c *= v
			} else if op == '+' {
				c += v
			} else {
				c = v
			}
		}
	}
	return c, -1
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
