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
	dist := manhattan(strongest.p, b.p)
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

func neighbors(p point) []point {
	return []point{{
		// above
		x: p.x,
		y: p.y - 1,
		z: p.z,
	}, {
		// below
		x: p.x,
		y: p.y + 1,
		z: p.z,
	}, {
		// left
		x: p.x - 1,
		y: p.y,
		z: p.z,
	}, {
		// right
		x: p.x + 1,
		y: p.y,
		z: p.z,
	}, {
		// in front
		x: p.x,
		y: p.y,
		z: p.z + 1,
	}, {
		// behind
		x: p.x,
		y: p.y,
		z: p.z - 1,
	}}
}

func manhattan(p1, p2 point) int {
	return abs(p1.x-p2.x) + abs(p1.y-p2.y) + abs(p1.z-p2.z)
}

func getPointsInRange(b bot) []point {
	points := []point{}
	currPoints := []point{b.p}
	seen := map[point]bool{}

	for {
		//fmt.Println("RANGE", b.r, "SEEN", len(seen))
		nextPoints := []point{}
		for _, p := range currPoints {
			n := neighbors(p)
			for _, nn := range n {
				if _, ok := seen[nn]; !ok {
					if manhattan(nn, b.p) <= b.r {
						nextPoints = append(nextPoints, nn)
						points = append(points, nn)
						seen[nn] = true
					}
				}
			}
		}
		currPoints = nextPoints
		if len(currPoints) == 0 {
			break
		}
	}

	return points
}

func getRangeCount(bots []bot) map[point]int {
	rangeCount := map[point]int{}

	for _, b := range bots {
		fmt.Println("BOT", b.p)
		points := getPointsInRange(b)
		for _, p := range points {
			c := rangeCount[p]
			rangeCount[p] = c + 1
		}
	}
	return rangeCount
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
	within := findBotsInRange(strongest, bots)
	fmt.Println("Bots within range of strongest bot:", len(within))

	rangeCount := getRangeCount(bots)

	maxC := 0
	var maxP point
	for p, count := range rangeCount {
		if count > maxC {
			maxC = count
			maxP = p
		}
	}

	fmt.Printf("point closest to all (to %d) is %d,%d,%d\n", maxC, maxP.x, maxP.y, maxP.z)
}
