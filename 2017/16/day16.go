package main

import (
	"fmt"
	"bufio"
	"log"
	"os"
    "strconv"
    "strings"
)


type MoveFunc func([]byte) []byte


func spin(move string) func([]byte) []byte {
    n, err := strconv.Atoi(move)
    if err != nil {
        log.Fatal(err)
    }
    return func(d []byte) []byte {
        for i := 0; i < n; i++ {
            d = append(d[len(d)-1:len(d)],  d[:len(d)-1]...)
        }
        return d
    }
}

func exchange(move string) func([]byte) []byte {
    s := strings.Split(move, "/")
    a, err := strconv.Atoi(s[0])
    if err != nil {
        log.Fatal(err)
    }
    b, err := strconv.Atoi(s[1])
    if err != nil {
        log.Fatal(err)
    }
    return func(d []byte) []byte {
        d[a], d[b] = d[b], d[a]
        return d
    }
}

func partner(move string) func([]byte) []byte {
    a, b := move[0], move[2]
    return func(d []byte) []byte {
        ai, bi := -1, -1

        for i := 0; i < len(d); i++ {
            if d[i] == a {
                ai = i
            }
            if d[i] == b {
                bi = i
            }
        }
        d[ai], d[bi] = d[bi], d[ai]
        return d
    }
}


func getFunc(move string) MoveFunc {
    command := move[0]
    move = move[1:]

    switch(command) {
    case "s"[0]:
        return spin(move)
    case "x"[0]:
        return exchange(move)
    case "p"[0]:
        return partner(move)
    }
    log.Fatalf("Did not recognize command", command)
    return nil
}

const iterations = 1000000000
const interval = 60

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Couldn't open file: %v\n", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

    d := []byte("abcdefghijklmnop")
    //org := string(d)

    var moves []string
	for scanner.Scan() {
        line := scanner.Text()
        moves = strings.Split(line, ",")
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

    var fMoves []MoveFunc

    for _, m := range moves {
        fMoves = append(fMoves, getFunc(m))
    }

    iters := iterations % interval
    for i := 0; i < iters; i++ {
        for _, m := range fMoves {
            d = m(d)
        }
        /*
        fmt.Println(i, string(d))
        if string(d) == org {
            fmt.Println(i)
            if i > 200 {
                break
            }
        }
        */
    }

    fmt.Println(string(d))
}
