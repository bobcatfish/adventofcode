package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func load() string {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("couldn't open file: %v", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	return scanner.Text()
}

func isMarker(s string, markerLen int) bool {
	m := map[rune]struct{}{}
	for _, c := range s {
		m[c] = struct{}{}
	}
	return len(m) == markerLen
}

func main() {
	buff := load()
	fmt.Println(buff)
	for _, markerLen := range []int{4, 14} {
		for i := 0; i < len(buff)-markerLen; i++ {
			ss := buff[i : i+markerLen]
			if isMarker(ss, markerLen) {
				fmt.Println(ss, i+markerLen)
				break
			}
		}
	}
}
