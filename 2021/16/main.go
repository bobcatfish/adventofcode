package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const maxint = int(^uint(0) >> 1)

func load() ([]rune, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, fmt.Errorf("couldn't open file: %v", err)
	}
	defer file.Close()

	vals := []rune{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		val := scanner.Text()
		for _, r := range val {

			vals = append(vals, r)
		}
	}
	return vals, err
}

func getBits(vals []rune) (string, error) {
	bits := ""
	for _, r := range vals {
		i, err := strconv.ParseInt(string(r), 16, 64)
		if err != nil {
			return "", err
		}
		s := fmt.Sprintf("%04s", strconv.FormatInt(int64(i), 2))
		bits = bits + s
	}
	return bits, nil
}

type packet struct {
	Version int
	Type    int
	Val     int
	Sub     []packet
}

func getInt(bits string, consumed, l int) (int, int, error) {
	vv := bits[consumed : consumed+l]
	v, err := strconv.ParseInt(vv, 2, 64)
	if err != nil {
		return 0, 0, err
	}

	return int(v), consumed + l, nil
}

func getLiteral(bits string, consumed int) (int, int, error) {
	si := ""
	si2 := ""
	last := false

	for !last {
		chunk := bits[consumed : consumed+5]
		consumed += 5
		if chunk[0] == '0' {
			last = true
		}
		si += string(chunk[1:])
		si2 += string(chunk)
	}
	v, err := strconv.ParseInt(si, 2, 64)
	if err != nil {
		return 0, 0, err
	}
	return int(v), consumed, nil
}

func getPacket(bits string, consumed int) (packet, int, error) {
	p := packet{}

	var err error
	p.Version, consumed, err = getInt(bits, consumed, 3)
	if err != nil {
		return p, 0, err
	}
	p.Type, consumed, err = getInt(bits, consumed, 3)
	if err != nil {
		return p, 0, err
	}
	if p.Type == 4 {
		var i int
		i, consumed, err = getLiteral(bits, consumed)
		p.Val = i
	} else {
		l := bits[consumed]
		consumed++
		p.Sub = []packet{}
		var allPackets func() bool
		if l == '0' {
			var subLen int
			subLen, consumed, err = getInt(bits, consumed, 15)
			if err != nil {
				return p, 0, err
			}
			end := consumed + subLen
			allPackets = func() bool {
				return consumed >= end
			}
		} else if l == '1' {
			var subCount int
			subCount, consumed, err = getInt(bits, consumed, 11)
			allPackets = func() bool {
				return len(p.Sub) == subCount
			}
		} else {
			return p, 0, fmt.Errorf("unknown length type id %s", string(l))
		}
		for !allPackets() {
			var subPacket packet
			subPacket, consumed, err = getPacket(bits, consumed)
			if err != nil {
				return p, 0, err
			}
			p.Sub = append(p.Sub, subPacket)
		}
	}

	return p, consumed, nil
}

func zeros(bits string, consumed int) bool {
	for _, r := range bits[consumed:] {
		if r != '0' {
			return false
		}
	}
	return true
}

func getTotal(packets []packet) int {
	vTotal := 0
	for _, p := range packets {
		vTotal += p.Version
		vTotal += getTotal(p.Sub)
	}
	return vTotal
}

func sum(p packet) int {
	v := 0
	for _, p := range p.Sub {
		v += p.Value()
	}
	return v
}

func product(p packet) int {
	v := 1
	for _, p := range p.Sub {
		v *= p.Value()
	}
	return v
}

func minimum(p packet) int {
	min := maxint
	for _, p := range p.Sub {
		v := p.Value()
		if v < min {
			min = v
		}
	}
	return min
}

func maximum(p packet) int {
	max := -1
	for _, p := range p.Sub {
		v := p.Value()
		if v > max {
			max = v
		}
	}
	return max
}

func greaterThan(p packet) int {
	if p.Sub[0].Value() > p.Sub[1].Value() {
		return 1
	}
	return 0
}

func lessThan(p packet) int {
	if p.Sub[0].Value() < p.Sub[1].Value() {
		return 1
	}
	return 0
}

func equal(p packet) int {
	if p.Sub[0].Value() == p.Sub[1].Value() {
		return 1
	}
	return 0
}

func (p packet) Value() int {
	switch p.Type {
	case 0:
		return sum(p)
	case 1:
		return product(p)
	case 2:
		return minimum(p)
	case 3:
		return maximum(p)
	case 4:
		return p.Val
	case 5:
		return greaterThan(p)
	case 6:
		return lessThan(p)
	case 7:
		return equal(p)
	}
	return -1
}
func main() {
	vals, err := load()
	if err != nil {
		log.Fatalf("Couldn't load nums from file: %v", err)
	}
	bits, err := getBits(vals)
	if err != nil {
		log.Fatalf("Couldn't get bits: %v", err)
	}

	p, _, err := getPacket(bits, 0)
	if err != nil {
		log.Fatalf("Couldn't get packets: %v", err)
	}
	vTotal := getTotal([]packet{p})
	fmt.Println(vTotal)

	value := p.Value()
	fmt.Println(value)
}
