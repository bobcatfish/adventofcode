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

type point struct {
	x int
	y int
	z int
}

type bot struct {
	p point
	r int
}

func getBot(s string) (bot, error) {
	b := bot{}
	r, err := regexp.Compile("pos=<(.*)>, r=(.*)")
	if err != nil {
		return b, fmt.Errorf("couldn't compile regex: %v", err)
	}
	m := r.FindStringSubmatch(s)
	if len(m) != 3 {
		return b, fmt.Errorf("didnt find expected values in %q: %v", s, err)
	}
	coorStrs := strings.Split(m[1], ",")
	coors := []int{}
	for _, s := range coorStrs {
		i, err := strconv.Atoi(s)
		if err != nil {
			return b, fmt.Errorf("coudln't convert %q to int: %v", s, err)
		}
		coors = append(coors, i)
	}
	b.p = point{
		x: coors[0],
		y: coors[1],
		z: coors[2],
	}

	rr, err := strconv.Atoi(m[2])
	if err != nil {
		return b, fmt.Errorf("couldnt convert range %q to int: %v", m[2], err)
	}
	b.r = rr
	return b, nil
}

func findStrongest(bots []bot) *bot {
	var strongest *bot
	for i, bot := range bots {
		if strongest == nil || bot.r > strongest.r {
			strongest = &bots[i]
		}
	}
	return strongest
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func within(strongest *bot, b bot) bool {
	dist := abs(strongest.p.x-b.p.x) + abs(strongest.p.y-b.p.y) + abs(strongest.p.z-b.p.z)
	return dist <= strongest.r
}

func findBotsInRange(strongest *bot, bots []bot) []bot {
	rangeBots := []bot{}

	for _, b := range bots {
		if within(strongest, b) {
			rangeBots = append(rangeBots, b)
		}
	}

	return rangeBots
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

	bots := []bot{}
	for _, v := range vals {
		bot, err := getBot(v)
		if err != nil {
			log.Fatalf("Couldn't get bot from %q: %v", v, err)
		}
		bots = append(bots, bot)
	}

	strongest := findStrongest(bots)
	fmt.Printf("Strongest bot (%d) at: %d,%d,%d\n", strongest.r, strongest.p.x, strongest.p.y, strongest.p.z)

	within := findBotsInRange(strongest, bots)
	fmt.Println(len(within))
}
