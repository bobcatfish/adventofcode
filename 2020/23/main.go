package main

import (
	"fmt"
)

const total = 1000000

//const total = 9

const moves = 10000000

//const moves = 100

type Cups struct {
	C      *Cup
	ValMap map[int]*Cup
	MaxVal int
}

func (cups Cups) String() string {
	s := ""
	c := cups.C
	curr := cups.C
	for {
		s += fmt.Sprintf("%d ", curr.Val)
		curr = curr.Next
		if curr == c {
			break
		}
	}
	return s
}

type Cup struct {
	Next *Cup
	Prev *Cup
	Val  int
}

func (cups *Cups) add(x, y int) int {
	return (x + y) % cups.MaxVal
}

func (cups *Cups) subtract(x, y int) int {
	z := x - y
	if z == 0 {
		return cups.MaxVal
	}
	return z
}

func (cups *Cups) grab3() []*Cup {
	grab := make([]*Cup, 3)

	curr := cups.C
	for i := 0; i < len(grab); i++ {
		next := curr.Next
		grab[i] = next
		curr = next
	}
	cups.C.Next = curr.Next
	curr.Prev = cups.C

	return grab
}

func (cups *Cups) move(grab []*Cup, destLabel int) {
	before := cups.ValMap[destLabel]
	after := before.Next

	before.Next = grab[0]
	grab[0].Prev = before

	grab[2].Next = after
	after.Prev = grab[2]
}

func in(x int, l []*Cup) bool {
	for _, v := range l {
		if x == v.Val {
			return true
		}
	}
	return false
}

func (cups *Cups) getLabel(grab []*Cup) int {
	destLabel := cups.subtract(cups.C.Val, 1)
	for {
		if in(destLabel, grab) {
			destLabel = cups.subtract(destLabel, 1)
		} else {
			break
		}
	}
	return destLabel
}

func (cups *Cups) Go() {
	grab := cups.grab3()
	destLabel := cups.getLabel(grab)
	cups.move(grab, destLabel)
	cups.C = cups.C.Next
}

func (cups *Cups) Label() string {
	s := ""
	one := cups.ValMap[1]
	for curr := one.Next; curr != one; curr = curr.Next {
		s += fmt.Sprintf("%d", curr.Val)
	}
	return s
}

func (cups *Cups) TwoCups() (int, int) {
	one := cups.ValMap[1]
	oneplus := one.Next
	twoplus := oneplus.Next
	return oneplus.Val, twoplus.Val
}

func Setup(seed []int) Cups {
	cups := Cups{
		C:      &Cup{Val: seed[0]},
		MaxVal: 9,
		ValMap: map[int]*Cup{},
	}
	cups.ValMap[seed[0]] = cups.C
	first := cups.C

	for i := 1; i < len(seed); i++ {
		next := &Cup{
			Val:  seed[i],
			Prev: cups.C,
		}
		cups.ValMap[seed[i]] = next
		cups.C.Next = next
		cups.C = next
	}
	if total > 9 {
		for i := cups.MaxVal + 1; i <= total; i++ {
			cups.MaxVal = i
			next := &Cup{
				Val:  i,
				Prev: cups.C,
			}
			cups.ValMap[i] = next
			cups.C.Next = next
			cups.C = next
		}
	}
	first.Prev = cups.C
	cups.C.Next = first
	cups.C = first

	return cups
}

func main() {
	seed := []int{9, 1, 6, 4, 3, 8, 2, 7, 5}
	cups := Setup(seed)

	for i := 0; i < moves; i++ {
		if i%100000 == 0 {
			fmt.Println("round", i)
		}
		cups.Go()
	}

	c1, c2 := cups.TwoCups()
	fmt.Println(c1, c2, c1*c2)
}
