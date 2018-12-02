package main

import (
	"fmt"
)


const iterations = 5000000
const magic = 2147483647.0
const mask = 65535


func next(i int, f int, c int, ch chan int) int {
    v := i
    for {
        v = (v * f) % magic
        if v % c == 0 {
            ch <- v
        }
    }
}

func compare(a int, b int) bool {
    return (a & mask) == (b & mask)
}


func main() {
    a := 783
    aF := 16807
    b := 325
    bF := 48271

    m := 0

    aC := make(chan int)
    bC := make(chan int)

    go next(a, aF, 4, aC)
    go next(b, bF, 8, bC)

    for i := 0; i < iterations; i++ {
        a, b = <-aC, <-bC
        if compare(a, b) {
            m += 1
        }
    }

    fmt.Println(m)
}
