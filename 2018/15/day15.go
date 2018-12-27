package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
)

const (
	wall    = '#'
	nothing = '.'
	goblinR = 'G'
	elfR    = 'E'

	startingAp = 3
	startingHp = 200
)

type justPoint struct {
	x int
	y int
}

type point struct {
	x int
	y int
	r rune
	u *unit
}

func (p *point) String() string {
	return fmt.Sprintf("%d,%d", p.x, p.y)
}

type unit struct {
	p   *point
	ap  int
	hp  int
	elf bool
}

func neighbors(p *point) []point {
	return []point{{
		// above
		x: p.x,
		y: p.y - 1,
	}, {
		// left
		x: p.x - 1,
		y: p.y,
	}, {
		// right
		x: p.x + 1,
		y: p.y,
	}, {
		// below
		x: p.x,
		y: p.y + 1,
	}}
}

func getBoard(vals []string) [][]*point {
	board := [][]*point{}
	for y, v := range vals {
		board = append(board, []*point{})
		for x, vv := range v {
			p := &point{
				x: x,
				y: y,
			}
			board[y] = append(board[y], p)
			if vv == wall {
				p.r = wall
			} else {
				p.r = nothing
				if vv == goblinR {
					p.u = &unit{
						p:   p,
						ap:  startingAp,
						hp:  startingHp,
						elf: false,
					}
				} else if vv == elfR {
					p.u = &unit{
						p:   p,
						ap:  startingAp,
						hp:  startingHp,
						elf: true,
					}
				}
			}
		}
	}
	return board
}

func getTargets(grid [][]*point, u *unit) []*unit {
	needType := !u.elf
	targets := []*unit{}

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid); x++ {
			targetU := grid[y][x].u
			if targetU != nil && targetU.elf == needType {
				targets = append(targets, targetU)
			}
		}
	}
	return targets
}

func getTargetSpaces(grid [][]*point, targets []*unit) []*point {
	spaces := []*point{}
	for _, target := range targets {
		n := neighbors(target.p)
		for _, nn := range n {
			p := grid[nn.y][nn.x]
			if p.r != wall && p.u == nil {
				spaces = append(spaces, p)
			}
		}
	}
	return spaces
}

type explorePoint struct {
	p      *point
	inPath map[justPoint]bool
	path   []*point
}

func copyMap(m map[justPoint]bool) map[justPoint]bool {
	targetMap := make(map[justPoint]bool, len(m))
	for key, value := range m {
		targetMap[key] = value
	}
	return targetMap
}

func copyPath(path []*point) []*point {
	n := make([]*point, len(path))
	copy(n, path)
	return path
}

func shortCircuit(toExplore []explorePoint) (bool, explorePoint) {
	var first *justPoint
	var ep explorePoint
	for i := 0; i < len(toExplore); i++ {
		e := toExplore[i]
		if first == nil {
			if len(e.path) == 0 {
				return false, ep
			}
			first = &justPoint{
				x: e.p.x,
				y: e.p.y,
			}
			ep = e
		} else {
			if e.p.x != first.x || e.p.y != first.y {
				return false, e
			}
		}
	}
	if len(toExplore) == 0 {
		return false, ep
	}
	return true, ep
}

func getShortestPath(grid [][]*point, spaces []*point, u *unit) []*point {
	toExplore := []explorePoint{{
		p:      u.p,
		inPath: map[justPoint]bool{},
		path:   []*point{},
	}}
	dist := 1
	explored := map[justPoint]bool{}
	for {
		if len(toExplore) == 0 {
			return []*point{}
		}
		nextExplore := []explorePoint{}
		goingToExplore := map[justPoint]bool{}

		for _, explore := range toExplore {
			thisOne := justPoint{x: explore.p.x, y: explore.p.y}
			explored[thisOne] = true
			g := grid[explore.p.y][explore.p.x]
			if g.u != nil && g.u.elf == !u.elf {
				return explore.path
			}
			n := neighbors(explore.p)
			for _, nn := range n {
				nnn := justPoint{x: nn.x, y: nn.y}
				if _, ok := explored[nnn]; !ok {

					if _, ok := explore.inPath[nnn]; !ok {
						gg := grid[nn.y][nn.x]
						if gg.r != wall {
							if gg.u == nil || gg.u.elf == !u.elf {
								newMap := copyMap(explore.inPath)
								newMap[nnn] = true

								newPath := copyPath(explore.path)
								newPath = append(newPath, gg)

								if _, ok := goingToExplore[nnn]; !ok {
									goingToExplore[nnn] = true
									nextExplore = append(nextExplore, explorePoint{
										p:      gg,
										inPath: newMap,
										path:   newPath,
									})
								}
							}
						}
					}
				}
			}

		}
		toExplore = nextExplore

		pp := []*point{}
		for _, ppp := range nextExplore {
			pp = append(pp, &point{
				x: ppp.p.x,
				y: ppp.p.y,
			})
		}
		dist++
	}
}

func takeStep(grid [][]*point, u *unit, p *point) {
	grid[u.p.y][u.p.x].u = nil
	grid[p.y][p.x].u = u
	u.p = p
}

func movePhase(grid [][]*point, u *unit) bool {
	targets := getTargets(grid, u)
	if len(targets) == 0 {
		return false
	}
	spaces := getTargetSpaces(grid, targets)
	path := getShortestPath(grid, spaces, u)
	if len(path) == 0 {
		return false
	}
	takeStep(grid, u, path[0])
	return true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func manhattan(p1, p2 *point) int {
	return abs(p1.x-p2.x) + abs(p1.y-p2.y)
}

func isAdj(p1, p2 *point) bool {
	return manhattan(p1, p2) <= 1
}

func canAttack(targets []*unit, u *unit) *unit {
	potential := []*unit{}
	for _, target := range targets {
		if isAdj(target.p, u.p) {
			potential = append(potential, target)
		}
	}
	min := 50000
	var target *unit
	for _, p := range potential {
		if p.hp < min {
			min = p.hp
			target = p
		}
	}
	return target
}

func attack(grid [][]*point, target *unit, u *unit) {
	target.hp -= u.ap
	if target.hp <= 0 {
		grid[target.p.y][target.p.x].u = nil
	}
}

func attackPhase(grid [][]*point, u *unit) bool {
	targets := getTargets(grid, u)
	target := canAttack(targets, u)
	if target != nil {
		attack(grid, target, u)
		return true
	}
	return false
}

func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func numLeft(grid [][]*point) bool {
	someSeen := false
	var this bool
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x].u != nil {
				if !someSeen {
					this = grid[y][x].u.elf
					someSeen = true
				} else if someSeen && grid[y][x].u.elf == !this {
					return true
				}
			}
		}
	}
	return false
}

func remaining(grid [][]*point) int {
	total := 0
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x].u != nil {
				total += grid[y][x].u.hp
			}
		}
	}
	return total
}

func printBoard(grid [][]*point) {
	//clear()
	for y := 0; y < len(grid); y++ {
		if y == 0 {
			fmt.Printf(" ")
			for x := 0; x < len(grid[y]); x++ {
				fmt.Printf("%d", x)
			}
			fmt.Println()
		}
		for x := 0; x < len(grid[y]); x++ {
			if x == 0 {
				fmt.Printf("%d", y)
			}
			g := grid[y][x]
			if g.u != nil {
				if g.u.elf {
					fmt.Printf("%c", elfR)
				} else {
					fmt.Printf("%c", goblinR)
				}
			} else {
				fmt.Printf("%c", g.r)
			}
		}
		fmt.Println()
	}
	fmt.Println()
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			g := grid[y][x]
			if g.u != nil {
				u := g.u
				var r rune
				if u.elf {
					r = 'E'
				} else {
					r = 'G'
				}
				fmt.Printf("(%c) (%d,%d) %d hp\n", r, u.p.x, u.p.y, u.hp)
			}
		}
	}
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

	grid := getBoard(vals)
	printBoard(grid)

	round := 0
	for {
		round++
		fmt.Println("ROUND", round)

		if !numLeft(grid) {
			round--
			goto done
		}

		moved := map[*unit]bool{}
		for y := 0; y < len(grid); y++ {
			for x := 0; x < len(grid[y]); x++ {
				u := grid[y][x].u
				if _, ok := moved[u]; !ok {
					moved[u] = true
					if u != nil {
						if !numLeft(grid) {
							round--
							// I AM ASHAMED
							goto done
						}
						attacked := attackPhase(grid, u)
						if attacked {
							continue
						}

						movePhase(grid, u)
						attackPhase(grid, u)
					}
				}
			}
		}
		printBoard(grid)

		//time.Sleep(1000 * time.Millisecond)
	}
done:
	fmt.Printf("%d rounds\n", round)

	hp := remaining(grid)
	fmt.Printf("%d hp\n", hp)

	fmt.Printf("answer: %d\n", hp*round)
}
