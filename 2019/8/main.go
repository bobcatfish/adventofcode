package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

const WIDTH = 25
const HEIGHT = 6

type Layer [][]byte

func load() ([]Layer, error) {
	val, err := ioutil.ReadFile("input.txt")
	if err != nil {
		return nil, fmt.Errorf("couldn't open file: %v", err)
	}

	layers := []Layer{}
	for i := 0; i < len(val)-1; i += WIDTH * HEIGHT {
		l := [][]byte{}
		for j := 0; j < HEIGHT; j++ {
			s := i + (j * WIDTH)
			l = append(l, val[s:s+WIDTH])
		}
		layers = append(layers, l)
	}

	return layers, err
}

type DigitCount map[int]int

func count(layers []Layer) []DigitCount {
	m := []DigitCount{}

	for _, ll := range layers {
		mm := map[int]int{}
		for _, l := range ll {
			for _, r := range l {
				i := int(r) - int('0')
				mm[i]++
			}
		}
		m = append(m, mm)
	}
	return m
}

func getPix(layers []Layer, y, x int) rune {
	for _, l := range layers {
		layer := l[y]
		r := layer[x]
		if r != '2' {
			if r == '0' {
				return '-'
			} else {
				return 'X'
			}
			return rune(r)
		}
	}
	return ' '
}

func main() {
	vals, err := load()
	if err != nil {
		log.Fatalf("Couldn't load from file: %v", err)
	}
	c := count(vals)

	fi, fc := -1, WIDTH*HEIGHT+1

	for i, cc := range c {
		if cc[0] < fc {
			fc = cc[0]
			fi = i
		}
	}

	fmt.Println(c[fi][1] * c[fi][2])

	for y := 0; y < HEIGHT; y++ {
		for x := 0; x < WIDTH; x++ {
			p := getPix(vals, y, x)
			fmt.Printf("%c", p)
		}
		fmt.Println()
	}
}
