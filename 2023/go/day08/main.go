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

	startingNodes := make([]string, 0)

	for _, line := range lines {
		data := strings.Fields(line)

		if len(data) == 0 {
			continue
		}

		if 'A' == data[0][2] {
			startingNodes = append(startingNodes, data[0])
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

	endPoints := make([]uint64, 0)

	for _, current = range startingNodes {
		found = false
		stepCount = 0
		for !found {
			for _, j := range steps {
				if 'L' == j {
					current = nodeMap[current].Left.Name
				} else {
					current = nodeMap[current].Right.Name
				}

				stepCount += 1
				if 'Z' == current[2] {
					found = true
					break
				}
			}
		}
		endPoints = append(endPoints, stepCount)
	}

	stepCount = LCM(endPoints[0], endPoints[1], endPoints[2:]...)
	fmt.Println(stepCount)
}

// The following functions were shamelessly copied
// and modifed from https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/
func GCD(a, b uint64) uint64 {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b uint64, integers ...uint64) uint64 {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
