package main

import (
	"fmt"
)


type Node struct {
    V int
    N *Node
}

const insertions = 50000000
const size = 314
const valueAfter = 2017

func main() {
    root := &Node{V:0}
    root.N = root

    node := root
    var length int

    for i := 1; i < insertions + 1; i++ {
        for j := 0; j < size; j++ {
            node = node.N
        }
        node.N = &Node{V:i, N:node.N}
        node = node.N
        length += 1

        if i % 1000000 == 0 {
            fmt.Println(i)
        }
    }

    node = root
    for i := 0; i < insertions + 1; i++ {
        node = node.N
        if node.V == valueAfter {
            fmt.Println(node.N.V)
            break
        }
    }
}
