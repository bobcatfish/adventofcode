package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
    "strconv"
    "strings"
)

const numWalls = 100

func mod(x int, y int) int {
    v := x % y
    if v < 0 {
        v += y
    }
    return v
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Couldn't open file: %v\n", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

    var walls [numWalls]int
    var line string
	for scanner.Scan() {
		line = scanner.Text()
        f := strings.Fields(line)
        i, err := strconv.Atoi(strings.Trim(f[0], ":"))
        if err != nil {
            log.Fatal(err)
        }
        r, err := strconv.Atoi(f[1])
        if r > 2 {
            r += (r-2)
        }

        if err != nil {
            log.Fatal(err)
        }
        walls[i] = r
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

    var delay int
    var buzz bool
    for ;; delay++ {
        for i := 0; i < numWalls; i++ {
            if walls[i] != 0 {
                hit := mod(delay + i, walls[i])
                if hit == 0 {
                    buzz = true
                    break
                }
            }
        }
        if !buzz {
            fmt.Println(delay)
            break
        } else {
            buzz = false
        }
    }
}
