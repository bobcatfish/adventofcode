package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
    "strings"
)


type HexPos struct {
    Col int
    DiaRow int
}

func (p HexPos) String() string {
    return fmt.Sprintf("Column: %v, DiaRow: %v", p.Col, p.DiaRow)
}

func stepsAway(p HexPos) int {
    var steps int
    for ; p.Col != 0 || p.DiaRow != 0; {
        if p.Col < 0 {
            p.Col += 1
        } else if p.DiaRow < 0 {
            p.DiaRow +=1
        } else if p.Col > 0 {
            p.Col -= 1
        } else if p.DiaRow > 0 {
            p.DiaRow -= 1
        }
        steps++
    }
    return steps
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Couldn't open file: %v\n", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

    var line string
	for scanner.Scan() {
		line = scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

    var max int
    var pos HexPos
    for _, m := range(strings.Split(line, ",")) {
        switch m {
        case "n":
            pos.DiaRow -= 1
        case "ne":
            pos.Col += 1
        case "se":
            pos.Col += 1
            pos.DiaRow += 1
        case "s":
            pos.DiaRow += 1
        case "sw":
            pos.Col -= 1
        case "nw":
            pos.DiaRow -= 1
            pos.Col -= 1
        }
        steps := stepsAway(pos)
        if steps > max {
            max = steps
        }
    }

    fmt.Println(pos)
    fmt.Println(stepsAway(pos))
    fmt.Println(max)
}
