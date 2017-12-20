package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
    "strconv"
    "strings"
)

type wall struct {
    Range int
    Loc int
    Down bool
}

func (w wall) String() string {
    return fmt.Sprintf("Range: %v Loc %v", w.Range, w.Loc)
}


const numWalls = 100


func tick(walls [numWalls]*wall) {
    for i := 0; i < len(walls); i++ {
        if walls[i] != nil {
            if walls[i].Loc == 0 && walls[i].Down {
                walls[i].Down = false
            } else if walls[i].Loc == walls[i].Range - 1 && !walls[i].Down {
                walls[i].Down = true
            }

            if walls[i].Down {
                walls[i].Loc -= 1
            } else {
                walls[i].Loc += 1
            }
        }
    }
}


func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Couldn't open file: %v\n", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

    var walls [numWalls]*wall
    var line string
	for scanner.Scan() {
		line = scanner.Text()
        f := strings.Fields(line)
        i, err := strconv.Atoi(strings.Trim(f[0], ":"))
        if err != nil {
            log.Fatal(err)
        }
        r, err := strconv.Atoi(f[1])
        if err != nil {
            log.Fatal(err)
        }
        walls[i] = &wall{Range: r,}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

    sev := 0

    for i := 0; i < len(walls); i++ {
        if walls[i] != nil {
            if walls[i].Loc == 0 {
                sev += (i * walls[i].Range)
            }
        }
        tick(walls)
    }

    fmt.Println(sev)
}
