package main

import (
	"bufio"
	"container/list"
	"fmt"
	"log"
	"os"
	"strconv"
)

const window = 25

func load() ([]int, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, fmt.Errorf("couldn't open file: %v", err)
	}
	defer file.Close()

	vals := []int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		i, err := strconv.Atoi(s)
		if err != nil {
			return nil, fmt.Errorf("couldn't convert %s to a num: %v", s, err)
		}
		vals = append(vals, i)
	}
	return vals, err
}

type BoundedSet struct {
	Q *list.List
	S map[int]struct{}
}

func NewBoundedSet() *BoundedSet {
	return &BoundedSet{
		Q: list.New(),
		S: map[int]struct{}{},
	}
}

func (b *BoundedSet) Add(i int) {
	if b.Q.Len() == window {
		e := b.Q.Front()
		delete(b.S, e.Value.(int))
		b.Q.Remove(e)

	}
	b.Q.PushBack(i)
	b.S[i] = struct{}{}
}

func Valid(i int, b *BoundedSet) bool {
	for e := b.Q.Front(); e != nil; e = e.Next() {
		if _, ok := b.S[i-e.Value.(int)]; ok {
			return true
		}
	}
	return false
}

type Sum struct {
	Val int
	Num int
}

func FindSum(s int, ii []int) (int, int) {
	sums := map[int]*Sum{}

	for i, v := range ii {
		sums[i] = &Sum{Val: v}
		for j := 0; j < i; j++ {
			sums[j].Val += v
			sums[j].Num++
			if sums[j].Val == s {
				return j, sums[j].Num
			}
		}
	}
	return -1, -1
}

func main() {
	ii, err := load()
	if err != nil {
		log.Fatalf("Couldn't load from file: %v", err)
	}

	b := NewBoundedSet()
	n := -1
	for i, v := range ii {
		if i >= window {
			if !Valid(v, b) {
				n = v
				break
			}
		}
		b.Add(v)
	}
	fmt.Println(n)

	index, num := FindSum(n, ii)
	if index == -1 {
		log.Fatalf("did not find!!!")
	}

	smallest := int((^uint(0)) >> 1)
	largest := -smallest - 1
	for i := index; i < index+num+1; i++ {
		v := ii[i]
		if v > largest {
			largest = v
		}
		if v < smallest {
			smallest = v
		}
	}
	fmt.Printf("solution is from %d to %d\n", index, index+num)
	fmt.Printf("smallest %d, largest %d, sum %d\n", smallest, largest, smallest+largest)
}
