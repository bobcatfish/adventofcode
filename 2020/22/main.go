package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Cards struct {
	S []int
}

func (s *Cards) Pop() int {
	var x int
	x, s.S = s.S[0], s.S[1:len(s.S)]
	return x
}

func (s *Cards) AddBottom(ii []int) {
	s.S = append(s.S, ii...)
}

func (s *Cards) Copy(max int) *Cards {
	c := &Cards{}
	for i, v := range s.S {
		if max != -1 && i >= max {
			break
		}
		c.S = append(c.S, v)
	}
	return c
}

func (s *Cards) Snapshot() string {
	return fmt.Sprintf("%v", s.S)
}

func load() ([]*Cards, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, fmt.Errorf("couldn't open file: %v", err)
	}
	defer file.Close()

	cardss := []*Cards{}
	delimCount := 0
	cards := &Cards{}

	scanner := bufio.NewScanner(file)
	firstLine := true
	for scanner.Scan() {
		s := scanner.Text()
		if firstLine {
			firstLine = false
			continue
		}

		if len(s) == 0 || (delimCount > 0 && delimCount < 2) {
			delimCount++
			if delimCount == 2 {
				cardss = append(cardss, cards)
				cards = &Cards{}
			}
			continue
		}

		i, err := strconv.Atoi(s)
		if err != nil {
			return nil, fmt.Errorf("couldn't convert value %s: %v", s, err)
		}
		cards.S = append(cards.S, i)
	}
	cardss = append(cardss, cards)
	return cardss, err
}

func score(s []int) int {
	c := 0

	for i := 1; i < len(s)+1; i++ {
		j := len(s) - i
		c += i * s[j]
	}
	return c
}

func combat(crab, me *Cards) {
	for {
		if len(me.S) == 0 || len(crab.S) == 0 {
			break
		}
		meNext, crabNext := me.Pop(), crab.Pop()

		if meNext > crabNext {
			me.AddBottom([]int{meNext, crabNext})
		} else if crabNext > meNext {
			crab.AddBottom([]int{crabNext, meNext})
		} else {
			panic("ITS A TIE???")
		}
	}
}

func printWinnerScore(crab, me *Cards) {
	var c int
	if len(me.S) > 0 {
		c = score(me.S)
	} else {
		c = score(crab.S)
	}
	fmt.Println(c)
}

const CRAB = 0
const ME = 1

func Snapshot(crab, me *Cards) string {
	return crab.Snapshot() + me.Snapshot()
}

func rcombat(crab, me *Cards) int {
	rounds := map[string]struct{}{}

	var winner int
	for {
		// if we are done
		if len(me.S) == 0 || len(crab.S) == 0 {
			break
		}

		// has this round happened before?
		r := Snapshot(crab, me)
		if _, ok := rounds[r]; ok {
			return CRAB
		}
		rounds[r] = struct{}{}

		meNext, crabNext := me.Pop(), crab.Pop()

		if len(me.S) >= meNext && len(crab.S) >= crabNext {
			winner = rcombat(crab.Copy(crabNext), me.Copy(meNext))
		} else {
			if meNext > crabNext {
				winner = ME
			} else if crabNext > meNext {
				winner = CRAB
			} else {
				panic("ITS A TIE???")
			}
		}

		if winner == CRAB {
			crab.AddBottom([]int{crabNext, meNext})
		} else {
			me.AddBottom([]int{meNext, crabNext})
		}
	}
	return winner
}

func main() {
	cards, err := load()
	if err != nil {
		log.Fatalf("Couldn't load from file: %v", err)
	}
	crab, me := cards[CRAB].Copy(-1), cards[ME].Copy(-1)
	combat(crab, me)
	printWinnerScore(crab, me)

	crab2, me2 := cards[CRAB].Copy(-1), cards[ME].Copy(-1)
	rcombat(crab2, me2)
	printWinnerScore(crab2, me2)

}
