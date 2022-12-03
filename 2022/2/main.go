package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Choice struct {
	points int
}

func newChoice(r byte) (*Choice, error) {
	// rock / lose
	if r == 'A' || r == 'X' {
		return &Choice{points: 1}, nil
	}
	// paper / draw
	if r == 'B' || r == 'Y' {
		return &Choice{points: 2}, nil
	}
	// scissors / win
	if r == 'C' || r == 'Z' {
		return &Choice{points: 3}, nil
	}
	return nil, fmt.Errorf("unknown selection %c", r)
}

func (c *Choice) Wins() *Choice {
	if c.points == 1 {
		return &Choice{points: 3}
	}
	return &Choice{points: c.points - 1}
}

func (c *Choice) Loses() *Choice {
	if c.points == 3 {
		return &Choice{points: 1}
	}
	return &Choice{points: c.points + 1}
}

func (c *Choice) Play(o *Choice) int {
	var cc *Choice
	cc = c.Wins()
	if o.points == cc.points {
		return 6 + c.points
	}
	if o.points == c.points {
		return 3 + c.points
	}
	return c.points
}

func (c *Choice) PlayFancy(o *Choice) int {
	var cc *Choice
	cc = o
	if c.points == 1 {
		cc = o.Wins()
	} else if c.points == 3 {
		cc = o.Loses()
	}
	return cc.Play(o)
}

type Play struct {
	left  *Choice
	right *Choice
}

func load() ([]*Play, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, fmt.Errorf("couldn't open file: %v", err)
	}
	defer file.Close()

	vals := []*Play{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		val := scanner.Text()
		left, err := newChoice(val[0])
		if err != nil {
			return nil, fmt.Errorf("couldn't parse play: %v", err)
		}
		right, err := newChoice(val[2])
		if err != nil {
			return nil, fmt.Errorf("couldn't parse play: %v", err)
		}

		p := &Play{
			left:  left,
			right: right,
		}
		vals = append(vals, p)
	}
	return vals, err
}

func main() {
	plays, err := load()
	if err != nil {
		log.Fatalf("Couldn't load from file: %v", err)
	}
	p1 := 0
	for _, p := range plays {
		p1 += p.right.Play(p.left)
	}
	fmt.Println(p1)

	p2 := 0
	for _, p := range plays {
		p2 += p.right.PlayFancy(p.left)
	}
	fmt.Println(p2)
}
