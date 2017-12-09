package main

import (
	"fmt"
	"github.com/golang-collections/collections/stack"
	"io/ioutil"
	"log"
)

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Couldn't read file: %v\n", err)
	}

	s := stack.New()
	s.Push(0)
	total := 0
	garbage := false
	gCount := 0
	cancelled := false
	for _, c := range input {
		var topScore int
		if s.Len() > 0 {
			topScore = s.Peek().(int)
		}
		switch string(c) {
		case "{":
			if !garbage {
				s.Push(topScore + 1)
			} else if garbage && cancelled {
				cancelled = false
			} else if garbage {
				gCount++
			}
		case "}":
			if !garbage {
				total += s.Pop().(int)
			} else if garbage && cancelled {
				cancelled = false
			} else if garbage {
				gCount++
			}
		case "<":
			if !garbage {
				garbage = true
			} else if garbage && cancelled {
				cancelled = false
			} else if garbage {
				gCount++
			}
		case "!":
			if garbage && !cancelled {
				cancelled = true
			} else if garbage && cancelled {
				cancelled = false
			}
		case ">":
			if garbage && !cancelled {
				garbage = false
			} else if garbage && cancelled {
				cancelled = false
			}
		default:
			if garbage && cancelled {
				cancelled = false
			} else if garbage {
				gCount++
			}
		}
	}
	for s.Len() > 0 {
		total += s.Pop().(int)
	}
	fmt.Println(total)
	fmt.Println(gCount)
}
