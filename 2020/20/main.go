package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
)

const TILES = 10

func load() ([]*Tile, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, fmt.Errorf("couldn't open file: %v", err)
	}
	defer file.Close()

	var t *Tile
	tiles := []*Tile{}
	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		s := scanner.Text()
		if i%(TILES+2) == 0 {
			if t != nil {
				tiles = append(tiles, t)
			}
			re := regexp.MustCompile(`^Tile (\d+):$`)
			vv := re.FindStringSubmatch(s)
			if len(vv) != 2 {
				return nil, fmt.Errorf("couldn't parse id from %s line %d", s, i)
			}
			i, err := strconv.Atoi(vv[1])
			if err != nil {
				return nil, fmt.Errorf("couldn't parse number from id %s", vv[1])
			}
			t = &Tile{
				Id:    i,
				Lines: []string{},
			}
		} else {
			if s != "" {
				t.Lines = append(t.Lines, s)
			}
		}
		i++
	}
	tiles = append(tiles, t)

	return tiles, err
}

type Connection struct {
	T   *Tile
	Val int
	Rot int
}

const UP = 0
const RIGHT = 1
const DOWN = 2
const LEFT = 3

type Tile struct {
	Id        int
	Lines     []string
	Flipped   bool
	Rotations int

	Vals []*Connection

	Conn []*Tile
}

func (t *Tile) Print() {
	for y := range t.Lines {
		t.PrintY(y)
		fmt.Println()
	}
}

func (t *Tile) PrintY(y int) {
	fmt.Printf("%s ", t.Lines[y])
	/*
		fmt.Printf(" %3d(%3d) %s %3d(%3d)",
			t.Vals[LEFT].Val,
			t.Vals[LEFT].Rot,
			t.Lines[y],
			t.Vals[RIGHT].Val,
			t.Vals[RIGHT].Rot)
	*/
}

type Point struct {
	Y int
	X int
}

func Rotator(lines []string) []string {
	runes := [][]rune{}
	newrunes := [][]rune{}
	newLines := make([]string, len(lines))

	for _, line := range lines {
		runes = append(runes, []rune(line))
		newrunes = append(newrunes, make([]rune, len(line)))
	}

	for y := 0; y < len(runes); y++ {
		for x := 0; x < len(runes[y]); x++ {
			np := Point{Y: len(runes[y]) - 1 - x, X: y}

			newrunes[y][x] = runes[np.Y][np.X]
		}
	}

	for y, rline := range newrunes {
		newLines[y] = string(rline)
	}
	return newLines
}

func (t *Tile) DoRotation() {
	t.Rotations--
	t.Lines = Rotator(t.Lines)
}

func reverse(line string) string {
	runes := []rune(line)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func Flipper(lines []string) []string {
	for y, line := range lines {
		lines[y] = reverse(line)
	}
	return lines
}

func (t *Tile) DoFlip() {
	t.Flipped = !t.Flipped

	fmt.Println("******BEFORE FLIP******")
	t.Print()

	t.Lines = Flipper(t.Lines)

	fmt.Println("******AFTER FLIP******")
	t.Print()

}

func (t *Tile) Flip() {
	t.Flipped = !t.Flipped

	for _, c := range t.Vals {
		c.Val, c.Rot = c.Rot, c.Val
	}

	t.Vals[LEFT], t.Vals[RIGHT] = t.Vals[RIGHT], t.Vals[LEFT]
	t.DoFlip()
}

func (t *Tile) matches(needs []int) bool {
	for dir, need := range needs {
		if need == 0 {
			continue
		}
		c := t.Vals[dir]
		if c != nil {
			if c.Val != need && c.Rot != need {
				return false
			}
		}
	}
	return true
}

func (t *Tile) Rotate() {
	t.Rotations++
	t.Rotations %= 3

	t.Vals[UP], t.Vals[RIGHT], t.Vals[DOWN], t.Vals[LEFT] = t.Vals[LEFT], t.Vals[UP], t.Vals[RIGHT], t.Vals[DOWN]

	t.DoRotation()
}

func (t *Tile) RotateTo(needs []int) bool {
	if t.matches(needs) {
		return true
	}

	for i := 0; i < 4; i++ {
		fmt.Println(i)
		t.Rotate()

		// last rotation puts us back to where we started
		if i < 3 {
			if t.matches(needs) {
				return true
			}
		}
	}

	return false
}

func (t Tile) String() string {
	s := fmt.Sprintf("Tile %d", t.Id)

	for i, c := range t.Vals {
		if i == UP {
			s += fmt.Sprintf("(Up %d %d)", c.Val, c.Rot)
		}
		if i == RIGHT {
			s += fmt.Sprintf("(Right %d %d)", c.Val, c.Rot)
		}
		if i == DOWN {
			s += fmt.Sprintf("(Down %d %d)", c.Val, c.Rot)
		}
		if i == LEFT {
			s += fmt.Sprintf("(Left %d %d)", c.Val, c.Rot)
		}
	}
	return s
}

func get(s string) (int, int) {
	r, v := 0, 0

	for i, x := range s {
		j := TILES - 1 - i

		if x == '#' {
			v |= 1 << uint(i)
			r |= 1 << uint(j)
		}
	}

	return r, v
}

func mkStrTopToBottom(ss []string, i int) string {
	s := ""
	for _, sss := range ss {
		s += string(sss[i])
	}
	return s
}

func mkStrBottomToTop(ss []string, i int) string {
	s := ""
	for j := len(ss) - 1; j >= 0; j-- {
		s += string(ss[j][i])
	}
	return s
}

func addVals(tiles []*Tile) []*Tile {
	for _, t := range tiles {
		vals := make([]*Connection, 4)
		// UP - left to right
		r, v := get(t.Lines[0])
		vals[UP] = &Connection{Val: v, Rot: r}

		// RIGHT - top to bottom
		right := mkStrTopToBottom(t.Lines, TILES-1)
		r, v = get(right)
		vals[RIGHT] = &Connection{Val: v, Rot: r}

		// DOWN - right to left
		down := reverse(t.Lines[TILES-1])
		r, v = get(down)
		vals[DOWN] = &Connection{Val: v, Rot: r}

		// LEFT - bottom to top
		left := mkStrBottomToTop(t.Lines, 0)
		r, v = get(left)
		vals[LEFT] = &Connection{Val: v, Rot: r}
		t.Vals = vals
	}
	return tiles
}

func in(tiles []*Tile, t *Tile) bool {
	for _, tt := range tiles {
		if tt.Id == t.Id {
			return true
		}
	}
	return false
}

func getMap(tiles []*Tile) map[int][]*Tile {
	m := map[int][]*Tile{}
	for _, t := range tiles {
		if !in(m[t.Vals[UP].Val], t) {
			m[t.Vals[UP].Val] = append(m[t.Vals[UP].Val], t)
		}
		if !in(m[t.Vals[UP].Rot], t) {
			m[t.Vals[UP].Rot] = append(m[t.Vals[UP].Rot], t)
		}

		if !in(m[t.Vals[RIGHT].Val], t) {
			m[t.Vals[RIGHT].Val] = append(m[t.Vals[RIGHT].Val], t)
		}
		if !in(m[t.Vals[RIGHT].Rot], t) {
			m[t.Vals[RIGHT].Rot] = append(m[t.Vals[RIGHT].Rot], t)
		}

		if !in(m[t.Vals[DOWN].Val], t) {
			m[t.Vals[DOWN].Val] = append(m[t.Vals[DOWN].Val], t)
		}
		if !in(m[t.Vals[DOWN].Rot], t) {
			m[t.Vals[DOWN].Rot] = append(m[t.Vals[DOWN].Rot], t)
		}

		if !in(m[t.Vals[LEFT].Val], t) {
			m[t.Vals[LEFT].Val] = append(m[t.Vals[LEFT].Val], t)
		}
		if !in(m[t.Vals[LEFT].Rot], t) {
			m[t.Vals[LEFT].Rot] = append(m[t.Vals[LEFT].Rot], t)
		}
	}
	return m
}

func checkDirs(d1 []*Tile, d2 []*Tile) []*Tile {
	poss := map[int]*Tile{}
	for _, p := range d1 {
		poss[p.Id] = p
	}

	if len(poss) >= 1 {
		for _, p := range d2 {
			if _, ok := poss[p.Id]; !ok {
				for _, pp := range poss {
					return []*Tile{p, pp}
				}
			}
		}
	}
	return nil
}

func connect(t1, t2 *Tile) {
	if !in(t1.Conn, t2) {
		t1.Conn = append(t1.Conn, t2)
	}
}

func makeGraph(m map[int][]*Tile) {
	for _, mm := range m {
		if len(mm) == 2 {
			connect(mm[0], mm[1])
			connect(mm[1], mm[0])
		}
	}
}

func possFor(m map[int][]*Tile, v int) {
	fmt.Println("poss for", v)
	for vv, mm := range m {
		good := false
		for _, mmm := range mm {
			if mmm.Id == v {
				good = true
			}
		}
		if good {
			for _, mmm := range mm {
				if mmm.Id != v {
					fmt.Println(vv, mmm.Id)
				}
			}
		}
	}
}

func printGraph(tiles []*Tile) {
	for _, t := range tiles {
		fmt.Println(t.Id, "connected to")
		for _, tt := range t.Conn {
			fmt.Printf("%d ", tt.Id)
		}
		fmt.Println()
		fmt.Println("******")
	}
}

func getCorners(tiles []*Tile) []*Tile {
	corners := []*Tile{}
	for _, t := range tiles {
		if len(t.Conn) == 2 {
			corners = append(corners, t)
		}
	}
	return corners
}

func getNext(seen map[int]struct{}, t *Tile) (*Tile, bool) {
	fmt.Println("getNext", t.Id)

	outsideEdges := []*Tile{}
	outsideSides := []*Tile{}
	insideSides := []*Tile{}
	insideEdges := []*Tile{}

	for _, c := range t.Conn {
		if _, ok := seen[c.Id]; ok {
			continue
		}
		// if it has 2 total connections, we know it's an outside edge
		if len(c.Conn) == 2 {
			outsideEdges = append(outsideEdges, c)
		} else if len(c.Conn) == 3 {
			// if it has 3 total connections, we know it's an outside side
			outsideSides = append(outsideSides, c)
		} else {
			count := 0
			for _, cc := range c.Conn {
				if _, ok := seen[cc.Id]; !ok {
					fmt.Println("NOT SEEN", c.Id, "connection", cc.Id, "(from", t.Id)
					count++
				} else {
					fmt.Println("SEEN", c.Id, "connection", cc.Id, "(from", t.Id)
				}
			}
			// 3 4
			fmt.Println(c.Id, count, len(c.Conn))
			// if it has 4 total connections, it's inside:
			if len(c.Conn) != 4 {
				panic(fmt.Sprintf("unexpected result %d count %d total conns %d", c.Id, count, len(c.Conn)))
			}

			// if count is 0 it's the last one left
			if count == 0 {
				insideSides = append(insideSides, c)
			} else if count == 2 {
				// if it's on the outside edge, it has 2 unexplored
				insideSides = append(insideSides, c)
			} else if count == 1 {
				// it it's on the outside edge and it's a corner, it has 1 unexplored
				insideEdges = append(insideEdges, c)
			}
		}

	}

	if len(outsideEdges) > 0 {
		return outsideEdges[0], true
	} else if len(outsideSides) > 0 {
		return outsideSides[0], false
	} else if len(insideEdges) > 0 {
		return insideEdges[0], true
	} else if len(insideSides) > 0 {
		return insideSides[0], false
	}

	return nil, false
}

func addNext(grid [][]*Tile, next *Tile, seen map[int]struct{}, y, x int) (*Tile, bool) {
	var isCorner bool
	next, isCorner = getNext(seen, next)
	if next == nil {
		return nil, false
	}
	grid[y][x] = next
	seen[next.Id] = struct{}{}
	return next, isCorner
}

func getInnerBorder(grid [][]*Tile, seen map[int]struct{}, corner *Tile, min, max int) [][]*Tile {
	y, x := min, min
	grid[y][x] = corner
	seen[corner.Id] = struct{}{}
	x++

	var isCorner bool
	next := corner

	for {
		fmt.Println("************")
		fmt.Println("y", y, "x", x, "min", min, "max", max)
		// we've gone all the way around
		if y == min && x == min {
			break
		}
		// y 11 x 10 min 1 max 11
		// y 11 x 9 min 1 max 11

		// from corner, get first row
		if y == min {
			fmt.Println("from corner, first row")
			next, isCorner = addNext(grid, next, seen, y, x)
			if next == nil {
				break
			}
			x++
			if isCorner {
				y++
				x--
			}
		} else if y > min && x == max-1 {
			// then at next corner, go down, filling in ouside
			fmt.Println("at next corner, go down, filling in outside")
			next, isCorner = addNext(grid, next, seen, y, x)
			if next == nil {
				break
			}
			y++
			if isCorner {
				y--
				x--
			}
		} else if y == max-1 && x > min {
			// next corner, go across, filling in the bottom row
			fmt.Println("at next corner, go back across, filling in the bottom row")
			next, isCorner = addNext(grid, next, seen, y, x)
			if next == nil {
				break
			}
			x--

			if isCorner {
				y--
				x = min
			}
		} else if x == min {
			// then at the next corner, go back up, filling in the last bit of the outside
			fmt.Println("at the next corner, go back up, filling in the last bit of hte outside")
			next, _ = addNext(grid, next, seen, y, x)
			if next == nil {
				break
			}
			y--
		}

	}

	// once we start attacking the inside, we're gonna have to special case when there is 1 in the middle
	// but 12x12 there won't be??
	return grid
}

func countUnfilled(grid [][]*Tile) int {
	count := 0
	for _, row := range grid {
		for _, v := range row {
			if v == nil {
				count++
			}
		}
	}
	return count
}

func printGridNums(grid [][]*Tile) {
	for _, row := range grid {
		for _, v := range row {
			if v == nil {
				fmt.Printf("______|")
			} else {
				fmt.Printf("%6d|", v.Id)
			}
		}
		fmt.Println()
	}
}

func getBorder(corner *Tile, size int) [][]*Tile {
	grid := make([][]*Tile, size)
	for i := 0; i < size; i++ {
		grid[i] = make([]*Tile, size)
	}

	seen := map[int]struct{}{}

	getInnerBorder(grid, seen, corner, 0, size)

	for y := 1; y < size-1; y++ {
		for x := 1; x < size-1; x++ {
			parent := grid[y-1][x]
			for _, t := range parent.Conn {
				if _, ok := seen[t.Id]; !ok {
					grid[y][x] = t
					seen[t.Id] = struct{}{}
				}
			}
			fmt.Println()
		}
	}
	return grid
}

func findMatch(curr, right *Tile) int {
	for _, cv := range curr.Vals {
		for _, rv := range right.Vals {
			if cv.Val == rv.Val {
				return cv.Val
			} else if cv.Val == rv.Rot {
				return cv.Rot
			} else if cv.Rot == rv.Val {
				return cv.Rot
			}
		}
	}
	panic(fmt.Sprintf("oh no we're doomed, no match b/w %d and %d", curr.Id, right.Id))
}

func getNeighbors(grid [][]*Tile, y, x int) []*Tile {
	neighbors := make([]*Tile, 4)

	// UP
	// y - 1, x
	if y > 0 {
		neighbors[UP] = grid[y-1][x]
	}

	// RIGHT
	// y, x + 1
	if x < len(grid[0])-1 {
		neighbors[RIGHT] = grid[y][x+1]
	}

	// DOWN
	// y + 1, x
	if y < len(grid)-1 {
		neighbors[DOWN] = grid[y+1][x]
	}

	// LEFT
	// y, x - 1
	if x > 0 {
		neighbors[LEFT] = grid[y][x-1]
	}

	return neighbors
}

func orient(grid [][]*Tile, tiles []*Tile) {
	lookup := map[int]*Tile{}
	for _, t := range tiles {
		lookup[t.Id] = t
	}

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			curr := grid[y][x]
			needs := make([]int, 4)
			neighbors := getNeighbors(grid, y, x)

			for dir, n := range neighbors {
				if n != nil {
					need := findMatch(curr, n)
					needs[dir] = need
				}
			}
			fmt.Println(needs)

			rotationWorked := curr.RotateTo(needs)
			needsFlip := !rotationWorked

			if needsFlip {
				curr.Flip()
			}

			rotationWorked = curr.RotateTo(needs)
			if !rotationWorked {
				panic("we already flipped it! nothing works!")
			}

			for i, c := range curr.Vals {
				fmt.Println(i)
				if c != nil {
					fmt.Println(c.Val, c.Rot)
				}
			}
		}
	}
}

func reorganize(tiles []*Tile) {
	for _, t := range tiles {
		if t.Flipped {
			t.DoFlip()
		}
		for i := 0; i < t.Rotations; i++ {
			t.DoRotation()
		}
	}
}

func printGrid(grid [][]*Tile) {
	for _, row := range grid {
		for y := -1; y < TILES+1; y++ {
			for _, v := range row {
				if y == -1 {
					//fmt.Printf("        V: %3d, R: %3d       ", v.Vals[UP].Val, v.Vals[UP].Rot)
				} else if y == TILES {
					//fmt.Printf("        V: %3d, R: %3d       ", v.Vals[DOWN].Val, v.Vals[DOWN].Rot)
				} else {
					v.PrintY(y)
				}
			}
			fmt.Println()
		}
		fmt.Println()
	}
}

func printGridClean(grid [][]*Tile) {
	for _, row := range grid {
		for y := 0; y < TILES; y++ {
			for _, v := range row {
				v.PrintY(y)
			}
			fmt.Println()
		}
		fmt.Println()
	}
}

func getImage(grid [][]*Tile) []string {
	image := []string{}
	for _, row := range grid {
		for y := 1; y < TILES-1; y++ {
			rowStr := ""
			for _, v := range row {
				rowStr += string(v.Lines[y][1 : TILES-1])
			}
			image = append(image, rowStr)
		}
	}
	return image
}

func printImage(image []string) {
	for _, l := range image {
		fmt.Println(l)
	}
}

var monster = []Point{
	{Y: 1, X: 0},
	{Y: 2, X: 1},

	{Y: 2, X: 4},
	{Y: 1, X: 5},
	{Y: 1, X: 6},
	{Y: 2, X: 7},

	{Y: 2, X: 10},
	{Y: 1, X: 11},
	{Y: 1, X: 12},
	{Y: 2, X: 13},

	{Y: 2, X: 16},
	{Y: 1, X: 17},
	{Y: 0, X: 18},
	{Y: 1, X: 18},
	{Y: 1, X: 19},
}

func countMonsters(image []string) int {
	count := 0

	width := 20
	height := 3

	for y := 0; y < len(image)-height; y++ {
		for x := 0; x < len(image[y])-width; x++ {
			for i, m := range monster {
				ny, nx := y+m.Y, x+m.X
				if image[ny][nx] != '#' {
					break
				} else if i == len(monster)-1 {
					count++
				}
			}
		}
	}
	return count
}

type Manip func([]string) []string

func countRough(image []string, c int) int {
	count := 0

	for _, line := range image {
		for _, r := range line {
			if r == '#' {
				count++
			}
		}
	}

	return count - (c * len(monster))
}

func main() {
	tiles1, err := load()
	if err != nil {
		log.Fatalf("Couldn't load from file: %v", err)
	}

	tiles := addVals(tiles1)

	m := getMap(tiles)
	makeGraph(m)

	corners := getCorners(tiles)
	size := int(math.Sqrt(float64(len(tiles))))
	grid := getBorder(corners[0], size)

	orient(grid, tiles)

	image := getImage(grid)
	printImage(image)

	manips := []Manip{
		// rotate 3 times
		Rotator,
		Rotator,
		Rotator,

		// then rotate again back to normal and flip
		Rotator,
		Flipper,

		// then rotate 3 more times
		Rotator,
		Rotator,
		Rotator,
	}

	c := 0
	for _, m := range manips {
		image = m(image)
		c = countMonsters(image)
		if c > 0 {
			break
		}
	}
	fmt.Println(c, "monsters")
	rough := countRough(image, c)
	fmt.Println(rough, "roughness")

}
