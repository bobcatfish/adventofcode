package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
    "strconv"
    "strings"
)


type Node struct {
    Num int
    Pipes []*Node
}


var nodes map[int]*Node


func getNode(n int) *Node {
    if v, ok := nodes[n]; ok {
        return v
    }
    nodes[n] = &Node{Num:n,}
    return nodes[n]
}


func nodeFromLine(line string) {
    items := strings.Split(line, " ")
    num, err := strconv.Atoi(items[0])
    if err != nil {
        log.Fatal(err)
    }

    node := getNode(num)

    for i := 2; i < len(items); i++ {
        pipe := strings.Trim(items[i], ",")
        pipeNum, err := strconv.Atoi(pipe)
        if err != nil {
            log.Fatal(err)
        }
        pipeNode := getNode(pipeNum)
        node.Pipes = append(node.Pipes, pipeNode)
    }
    fmt.Println(node)
}

var visited map[*Node]bool

func walk(node *Node) int {
    sum := 0
    if _, ok := visited[node]; !ok {
        sum += 1
        visited[node] = true
        for _, pipe := range node.Pipes {
            sum += walk(pipe)
        }
    }
    return sum
}

func destructiveWalk(node *Node) int {
    sum := 0
    if _, ok := visited[node]; !ok {
        sum += 1
        visited[node] = true
        delete(nodes, node.Num)
        for _, pipe := range node.Pipes {
            sum += walk(pipe)
        }
    }
    return sum
}


func main() {
    nodes = make(map[int]*Node)
    visited = make(map[*Node]bool)

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Couldn't open file: %v\n", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

    var line string
	for scanner.Scan() {
		line = scanner.Text()
        nodeFromLine(line)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

    start := nodes[0]
    numVisited := walk(start)
    fmt.Println(numVisited)

    visited = make(map[*Node]bool)
    totalNum := len(nodes)
    groups := 0
    for ; len(visited) < totalNum ; {
        for i := 0; i < totalNum; i++ {
            if node, ok := nodes[i]; ok {
                if _, ok := visited[node]; !ok {
                    groups += 1
                    destructiveWalk(node)
                }
            }
        }
    }
    fmt.Println(groups)
}
