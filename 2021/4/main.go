package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const size = 5

type board struct {
	vals   [][]int
	marked [][]int
}

func newBoard() board {
	b := board{}
	b.marked = make([][]int, size)
	for i := 0; i < size; i++ {
		b.marked[i] = make([]int, size)
	}
	return b
}

func (b *board) Mark(num int) {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if b.vals[i][j] == num {
				b.marked[i][j] = 1
			}
		}
	}
}

func (b *board) Winner() bool {
	for i := 0; i < size; i++ {
		winner := true
		for j := 0; j < size; j++ {
			if b.marked[i][j] != 1 {
				winner = false
			}
		}
		if winner {
			return true
		}
	}

	for i := 0; i < size; i++ {
		winner := true
		for j := 0; j < size; j++ {
			if b.marked[j][i] != 1 {
				winner = false
			}
		}
		if winner {
			return true
		}
	}
	return false
}

func (b *board) Print() {
	for i := 0; i < size; i++ {
		fmt.Println(b.vals[i])
	}
}

func (b *board) Score() int {
	sum := 0
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if b.marked[i][j] != 1 {
				sum += b.vals[i][j]
			}
		}
	}
	return sum
}

func getNums(s, split string) ([]int, error) {
	var nss []string
	if split == " " {
		nss = strings.Fields(s)
	} else {
		nss = strings.Split(s, split)
	}
	nums := []int{}
	for _, n := range nss {
		i, err := strconv.Atoi(n)
		if err != nil {
			return nil, err
		}
		nums = append(nums, i)
	}
	return nums, nil
}

func load() ([]int, []board, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, nil, fmt.Errorf("couldn't open file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	nums, err := getNums(scanner.Text(), ",")
	if err != nil {
		return nil, nil, err
	}

	boards := []board{}

	i := 0
	for scanner.Scan() {
		t := scanner.Text()
		if i%(size+1) == 0 {
			boards = append(boards, newBoard())
		} else {
			nums, err := getNums(t, " ")
			if err != nil {
				return nil, nil, err
			}
			boards[len(boards)-1].vals = append(boards[len(boards)-1].vals, nums)
		}
		i++
	}
	return nums, boards, err
}

func main() {
	nums, boards, err := load()
	if err != nil {
		log.Fatalf("Couldn't load from file: %v", err)
	}

	winners := make([]bool, len(boards))
	for {
		var winner *board
		winNum := 0
		for _, num := range nums {
			for bi, board := range boards {
				board.Mark(num)

				if board.Winner() && !winners[bi] {
					winner = &board
					winNum = num
					winners[bi] = true
					break
				}
			}
			if winner != nil {
				break
			}
		}
		winner.Print()
		fmt.Println(winner.Score(), winNum, winner.Score()*winNum)
		all := true
		for _, w := range winners {
			if !w {
				all = false
			}
		}
		if all {
			break
		}
	}
}
