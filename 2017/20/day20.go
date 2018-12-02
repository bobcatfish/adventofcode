package main

import (
	"fmt"
	"bufio"
	"log"
	"os"
    "regexp"
    "strconv"
    "math"
)

const r = "p=<(-?\\d*),(-?\\d*),(-?\\d*)>, v=<(-?\\d*),(-?\\d*),(-?\\d*)>, a=<(-?\\d*),(-?\\d*),(-?\\d*)>"

type point struct {
    X int
    Y int
    Z int
}


type particle struct {
    P point
    V point
    A point
}


func getPoint(p []string) point {
    x, err := strconv.Atoi(p[0])
    if err != nil {
        log.Fatal(err)
    }
    y, err := strconv.Atoi(p[1])
    if err != nil {
        log.Fatal(err)
    }
    z, err := strconv.Atoi(p[2])
    if err != nil {
        log.Fatal(err)
    }
    return point{X:x, Y:y, Z:z}
}


func getParticle(line string) particle {
    re := regexp.MustCompile(r)
    values := re.FindStringSubmatch(line)
    p := getPoint(values[1:4])
    v := getPoint(values[4:7])
    a := getPoint(values[7:10])
    return particle{P:p, V:v, A:a}
}

func dist(p point) int {
    return int(math.Abs(float64(p.X)) + math.Abs(float64(p.Y)) + math.Abs(float64(p.Z)))
}

func tick(particles []*particle) {
    for _, p := range particles {
        p.V.X += p.A.X
        p.V.Y += p.A.Y
        p.V.Z += p.A.Z

        p.P.X += p.V.X
        p.P.Y += p.V.Y
        p.P.Z += p.V.Z
    }
}


func closest(particles []*particle) int {
    d := int(int64(^uint(0) >> 1))
    w := -1
    for i, p := range particles {
        distance := dist(p.P)

        //fmt.Println(i, p, distance)

        if distance < d {
            d = distance
            w = i
        }
    }
    return w
}


func removeColl(particles []*particle) []*particle{
    m := make(map[point][]int)
    for i, p := range particles {
        if _, ok := m[p.P]; ok {
            m[p.P] = append(m[p.P], i)
        } else {
            m[p.P] = []int{i}
        }
    }

    var left []*particle
    for _, ps := range m {
        if len(ps) == 1 {
            left = append(left, particles[ps[0]])
        }
    }
    return left
}


const ticks = 1000


func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatalf("Couldn't open file: %v\n", err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    var particles []*particle
    for scanner.Scan() {
        line := scanner.Text()
        p := getParticle(line)
        particles = append(particles, &p)
    }
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    var i int
    for i = 0; i < ticks; i++ {
        tick(particles)
        particles = removeColl(particles)
        fmt.Println(len(particles))
    }

    c := closest(particles)
    fmt.Println("Closest after", i, c)
}

