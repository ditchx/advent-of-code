package main

import (
	"fmt"
	"os"
	"strings"
)

type Node struct {
	Name  string
	Left  *Node
	Right *Node
}

func newNode(name string) *Node {
	return &Node{
		Name: name,
	}
}

func main() {
	file, _ := os.ReadFile(os.Args[1])

	content := strings.ReplaceAll(string(file), "(", "")
	content = strings.ReplaceAll(content, ")", "")
	content = strings.ReplaceAll(content, ",", "")
	content = strings.ReplaceAll(content, "=", "")
	lines := strings.Split(content, "\n")

	steps := lines[0]

	lines = lines[2:]

	nodeMap := make(map[string]*Node)

	for _, line := range lines {
		data := strings.Fields(line)

		if len(data) == 0 {
			continue
		}

		if _, ok := nodeMap[data[0]]; !ok {
			nodeMap[data[0]] = newNode(data[0])
		}

		leftNode, lefOk := nodeMap[data[1]]
		if !lefOk {
			leftNode = newNode(data[1])
		}
		nodeMap[data[0]].Left = leftNode

		rightNode, rightOk := nodeMap[data[2]]
		if !rightOk {
			rightNode = newNode(data[2])
		}
		nodeMap[data[0]].Right = rightNode

	}

	var current = nodeMap["AAA"].Name

	var found = false
	var stepCount uint64
	for !found {
		for _, j := range steps {
			if 'L' == j {
				current = nodeMap[current].Left.Name
			} else {
				current = nodeMap[current].Right.Name
			}

			stepCount += 1
			if current == "ZZZ" {
				found = true
				break
			}
		}
	}

	fmt.Println(stepCount)
}
