package main

import (
	"fmt"
	"bufio"
	"log"
	"os"
    "strconv"
    "strings"
)


var sound int
var registers map[string] int

type Instr func() int


func getReg(reg string) int {
    if r, ok := registers[reg]; ok {
        return r
    } else {
        registers[reg] = 0
        return 0
    }
}


func send(reg string) Instr {
    return func() int {
        sound = getReg(reg)
        return 1
    }
}

func set(x string, y string) Instr {
    v, err := strconv.Atoi(y)
    if err != nil {
        return func() int {
            value := getReg(y)
            registers[x] = value
            return 1
        }
    }
    return func() int {
        registers[x] = v
        return 1
    }
}

func add(x string, y string) Instr {
    v, err := strconv.Atoi(y)
    if err != nil {
        return func() int {
            value := getReg(y)
            xValue := getReg(x)
            registers[x] = xValue + value
            return 1
        }
    }
    return func() int {
        xValue := getReg(x)
        registers[x] = xValue + v
        return 1
    }
}

func mul(x string, y string) Instr {
    v, err := strconv.Atoi(y)
    if err != nil {
        return func() int {
            value := getReg(y)
            xValue := getReg(x)
            registers[x] = xValue * value
            return 1
        }
    }
    return func() int {
        xValue := getReg(x)
        registers[x] = xValue * v
        return 1
    }
}

func mod(x string, y string) Instr {
    v, err := strconv.Atoi(y)
    if err != nil {
        return func() int {
            value := getReg(y)
            xValue := getReg(x)
            registers[x] = xValue % value
            return 1
        }
    }
    return func() int {
        xValue := getReg(x)
        registers[x] = xValue % v
        return 1
    }
}

func rcv(reg string) Instr {
    return func() int {
        x := getReg(reg)
        if x != 0 {
            fmt.Println(sound)
            os.Exit(0)
        }
        return 0
    }
}

func jgz(x string, y string) Instr {
    v, err := strconv.Atoi(y)
    if err != nil {
        return func() int {
            x := getReg(x)
            if x > 0 {
                return getReg(y)
            }
            return 1
        }
    }
    return func() int {
        x := getReg(x)
        if x > 0 {
            return v
        }
        return 1
    }
}



func getInstr(line string) Instr {
    f := strings.Fields(line)
    switch f[0] {
    case "snd":
        return send(f[1])
    case "set":
        return set(f[1], f[2])
    case "add":
        return add(f[1], f[2])
    case "mul":
        return mul(f[1], f[2])
    case "mod":
        return mod(f[1], f[2])
    case "rcv":
        return rcv(f[1])
    case "jgz":
        return jgz(f[1], f[2])
    }
    log.Fatalf("did not find %v\n", line)
    return nil
}


func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Couldn't open file: %v\n", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

    var instr []Instr
    registers = make(map[string]int)
	for scanner.Scan() {
        line := scanner.Text()
        instr = append(instr, getInstr(line))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

    i := 0
    for {
        i += instr[i]()
        i %= len(instr)
    }
}
