package main

import (
	"fmt"
	"bufio"
	"log"
	"os"
    "unicode"
)


var cont_hor = byte("|"[0])
var cont_ver = byte("-"[0])
var turn = byte("+"[0])
var count int

type point struct {
	X int
	Y int
}

var up = point{
    X: 0,
    Y: -1,
}
var left = point{
    X: -1,
    Y: 0,
}
var right = point{
    X: 1,
    Y: 0,
}
var down = point{
    X: 0,
    Y: 1,
}


func findStart(lines []string) point {
    for i := 0; i < len(lines[0]); i++ {
        if lines[0][i] == cont_hor {
            return point{X: i, Y: 0}
        }
    }
    log.Fatalf("did not find start %v", lines[0])
    return point{}
}

func getPoint(lines []string, p point) byte {
    return lines[p.Y][p.X]
}

func nextVer(lines []string, p point) point {
    nextL := point{X: p.X + left.X, Y: p.Y + left.Y}
    nextR := point{X: p.X + right.X, Y: p.Y + right.Y}

    if getPoint(lines, nextL) == cont_ver {
        return left
    } else if getPoint(lines, nextR) == cont_ver {
        return right
    }
    fmt.Println("Nothing next vert from %v, %v\n", p, lines[p.Y])
    os.Exit(0)
    return point{}
}

func nextHor(lines []string, p point) point {
    nextUp := point{X: p.X + up.X, Y: p.Y + up.Y}
    nextDown := point{X: p.X + down.X, Y: p.Y + down.Y}

    if getPoint(lines, nextUp) == cont_hor {
        return up
    } else if getPoint(lines, nextDown) == cont_hor {
        return down
    }
    fmt.Println("Nothing next hor from %v\n", p)
    os.Exit(0)
    return point{}
}

func follow(lines []string, start point, seen string, dir point) string {
    count += 1
    next := point{X: start.X + dir.X, Y: start.Y + dir.Y}
    v := getPoint(lines, next)
    if v == " "[0] {
        fmt.Println(seen)
        fmt.Println(count)
        os.Exit(0)
    }
    if v != turn {
        if unicode.IsLetter(rune(v)) {
            seen += string(v)
        }
        return follow(lines, next, seen, dir)
    } else {
        if dir == up || dir == down {
            dir = nextVer(lines, next)
            return follow(lines, next, seen, dir)
        } else {
            dir = nextHor(lines, next)
            return follow(lines, next, seen, dir)
        }
    }
    log.Fatalf("Can't go to next %v\n", string(v))
    return ""
}

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatalf("Couldn't open file: %v\n", err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    var lines []string
    for scanner.Scan() {
        line := scanner.Text()
        lines = append(lines, line)
    }
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
    /*
    for i, l := range lines {
        fmt.Println(i, l)
    }
    os.Exit(0)
    */
    start := findStart(lines)

    seen := follow(lines, start, "", down)
    fmt.Println(seen)
}

