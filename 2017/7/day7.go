package main

import (
	"bufio"
	"log"
	"os"
    "fmt"
    "strings"
    "strconv"
)

type Node struct {
    Name string
    Nodes []*Node
    Weight int
}

func getTower(node *Node) int {
    total := node.Weight
    for _, n := range node.Nodes {
        total += getTower(n)
    }
    return total
}

func followProblem(node *Node) {
    var val = 0
    var other = 0

    fmt.Printf("%v children for %v weight %v\n", len(node.Nodes), node.Name, node.Weight)
    var valChild *Node
    var otherChild *Node
    nextBad := false
    for _, child := range node.Nodes {
        tower := getTower(child)
        fmt.Printf("%v is at %v\n", child.Name, tower)

        if nextBad && val != tower {
            fmt.Printf("Majority are %v, this(%v) is %v, weight is %v\n", val, child.Name, tower, child.Weight)
            followProblem(child)
        } else if val == tower {
            nextBad = true
        } else if val == 0 && other == 0 {
            val = tower
            valChild = child
        } else if other == 0 {
            other = tower
            otherChild = child
        } else {
            if tower == val {
                fmt.Printf("Majority are %v, this(%v) is %v, weight is %v\n", val, otherChild.Name, other, otherChild.Weight)
                followProblem(otherChild)
                break
            } else {
                fmt.Printf("Majority are %v, this(%v) is %v, weight is %v\n", other, valChild.Name, val, valChild.Weight)
                followProblem(valChild)
                break
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

    tree := make(map[string]*Node)
    roots := make(map[string]bool)
    notRoots := make(map[string]bool)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
        line := scanner.Text()
        fields := strings.Fields(line)
        name := fields[0]
        weightStr := strings.TrimLeft(strings.TrimRight(fields[1], ")"), "(")
        weight, err := strconv.Atoi(weightStr)
        if err != nil {
            log.Fatal(err)
        }

        if _, ok := tree[name]; !ok {
            newNode := Node{Name: name}
            tree[name] = &newNode
        }
        node := tree[name]
        node.Weight = weight
        if len(fields) > 2 {
            for i := 3; i < len(fields); i++ {
                childName := strings.TrimRight(fields[i], ",")
                if _, ok := tree[childName]; !ok {
                    newChild := Node{Name: childName,}
                    tree[childName] = &newChild
                }
                child := tree[childName]
                node.Nodes = append(node.Nodes, child)

                notRoots[childName] = true
            }
            roots[name] = true
        }

	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

    var root *Node
    for k, _ := range roots {
        if _, present := notRoots[k]; !present {
            root = tree[k]
        }
    }

    fmt.Println(root)
    followProblem(root)
}
