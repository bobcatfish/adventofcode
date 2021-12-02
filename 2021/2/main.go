package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type command struct {
	dir string
	v   int
}

func load() ([]command, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, fmt.Errorf("couldn't open file: %v", err)
	}
	defer file.Close()

	commands := []command{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		val := scanner.Text()

		re := regexp.MustCompile(`^(\w*) (\d*)$`)
		vv := re.FindStringSubmatch(val)
		if len(vv) != 3 {
			return nil, fmt.Errorf("couldnt parse %s %v", val)
		}
		i, err := strconv.Atoi(vv[2])
		if err != nil {
			return nil, err
		}
		commands = append(commands, command{dir: vv[1], v: i})
	}
	return commands, err
}

type location struct {
	hor   int
	depth int
	aim   int
}

func (l *location) move1(c command) {
	switch c.dir {
	case "forward":
		l.hor += c.v
	case "down":
		l.depth += c.v
	case "up":
		l.depth -= c.v
	}
}

func (l *location) move2(c command) {
	switch c.dir {
	case "down":
		l.aim += c.v
	case "up":
		l.aim -= c.v
	case "forward":
		l.hor += c.v
		l.depth += l.aim * c.v
	}
}

func main() {
	commands, err := load()
	if err != nil {
		log.Fatalf("Couldn't load nums from file: %v", err)
	}

	l := location{}
	for _, command := range commands {
		l.move1(command)
	}

	fmt.Println(l, l.hor*l.depth)

	l = location{}
	for _, command := range commands {
		l.move2(command)
	}

	fmt.Println(l, l.hor*l.depth)
}
