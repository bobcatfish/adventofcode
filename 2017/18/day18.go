package main

import (
	"fmt"
	"bufio"
	"log"
	"os"
    "strconv"
    "strings"
    "sync"
    "runtime/debug"
)


type Instr func() int

var count int


func getReg(reg string, registers map[string] int) int {
    _, err := strconv.Atoi(reg)
    if err == nil {
        log.Fatal(string(debug.Stack()))
    }
    if r, ok := registers[reg]; ok {
        return r
    } else {
        fmt.Println("new reg", reg)
        registers[reg] = 0
        return 0
    }
}


func send(reg string, registers map[string]int, c chan int) Instr {
    var doCount bool
    if registers["p"] == 1 {
        doCount = true
    }
    return func() int {
        sound := getReg(reg, registers)
        if doCount {
            count += 1
            fmt.Println(count)
        }
        c <- sound
        return 1
    }
}

func set(x string, y string, registers map[string]int) Instr {
    v, err := strconv.Atoi(y)
    if err != nil {
        return func() int {
            value := getReg(y, registers)
            registers[x] = value
            return 1
        }
    }
    return func() int {
        registers[x] = v
        return 1
    }
}

func add(x string, y string, registers map[string]int) Instr {
    v, err := strconv.Atoi(y)
    if err != nil {
        return func() int {
            value := getReg(y, registers)
            xValue := getReg(x, registers)
            registers[x] = xValue + value
            return 1
        }
    }
    return func() int {
        xValue := getReg(x, registers)
        registers[x] = xValue + v
        return 1
    }
}

func mul(x string, y string, registers map[string]int) Instr {
    v, err := strconv.Atoi(y)
    if err != nil {
        return func() int {
            value := getReg(y, registers)
            xValue := getReg(x, registers)
            registers[x] = xValue * value
            return 1
        }
    }
    return func() int {
        xValue := getReg(x, registers)
        registers[x] = xValue * v
        return 1
    }
}

func mod(x string, y string, registers map[string]int) Instr {
    v, err := strconv.Atoi(y)
    if err != nil {
        return func() int {
            value := getReg(y, registers)
            xValue := getReg(x, registers)
            if value == 0 {
                registers[x] = 0
            } else {
                registers[x] = xValue % value
            }
            return 1
        }
    }
    return func() int {
        xValue := getReg(x, registers)
        registers[x] = xValue % v
        return 1
    }
}

func rcv(reg string, registers map[string]int, c chan int) Instr {
    return func() int {
        v := <-c
        registers[reg] = v
        return 1
    }
}

func jgz(x string, y string, registers map[string]int) Instr {
    v, err := strconv.Atoi(y)
    if err != nil {
        return func() int {
            xv, err := strconv.Atoi(x)
            if err != nil {
                x := getReg(x, registers)
                if x > 0 {
                    return getReg(y, registers)
                }
                return 1
            } else {
                if xv > 0 {
                    return getReg(y, registers)
                }
                return 1
            }
        }
    }
    return func() int {
        xv, err := strconv.Atoi(x)
        if err != nil {
            x := getReg(x, registers)
            if x > 0 {
                return v
            }
            return 1
        } else {
            if xv > 0 {
                return v
            }
            return 1
        }
    }
}



func getInstr(line string, registers map[string]int, mine chan int, other chan int) Instr {
    f := strings.Fields(line)
    switch f[0] {
    case "snd":
        return send(f[1], registers, other)
    case "set":
        return set(f[1], f[2], registers)
    case "add":
        return add(f[1], f[2], registers)
    case "mul":
        return mul(f[1], f[2], registers)
    case "mod":
        return mod(f[1], f[2], registers)
    case "rcv":
        return rcv(f[1], registers, mine)
    case "jgz":
        return jgz(f[1], f[2], registers)
    }
    log.Fatalf("did not find %v\n", line)
    return nil
}

func doit(instr []Instr, wg sync.WaitGroup, registers map[string]int) {
    i := 0
    for {
        i += instr[i]()
        //i %= len(instr)
        if i >= len(instr) {
            break
        }
        if i < 0 {
            log.Fatalf("i less than 0 %v\n", i)
        }
        //fmt.Println(registers)
    }
    wg.Done()
    os.Exit(0)
}


func main() {
    progs := [2]int{0, 1}
    c1 := make(chan int, 10000)
    c2 := make(chan int, 10000)
    var registers [2]map[string]int

    var instr [2][]Instr

    for _, p := range(progs) {
        file, err := os.Open("input.txt")
        if err != nil {
            log.Fatalf("Couldn't open file: %v\n", err)
        }
        defer file.Close()

        scanner := bufio.NewScanner(file)

        registers[p] = make(map[string]int)
        registers[p]["p"] = p

        for scanner.Scan() {
            line := scanner.Text()
            var n Instr
            if p == 0 {
                n = getInstr(line, registers[p], c1, c2)
            } else {
                n = getInstr(line, registers[p], c2, c1)
            }
            instr[p] = append(instr[p], n)
        }
        if err := scanner.Err(); err != nil {
            log.Fatal(err)
        }
    }
    var wg sync.WaitGroup
    wg.Add(2)
    for i, _ := range(progs) {
        go doit(instr[i], wg, registers[i])
    }
    wg.Wait()
}
