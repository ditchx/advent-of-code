package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	file, _ := os.ReadFile(os.Args[1])
	contents := strings.Split(strings.TrimSpace(string(file)), "\n")

	maxWeight := len(contents)
	var total uint64
	for j := range contents[0] {
		var load uint64
		slot := make([]int, 0)
		for i := range contents {
			slot = append(slot, maxWeight-i)
			if contents[i][j] == 'O' {
				load += uint64(slot[0])
				slot = slot[1:]
			}

			if contents[i][j] == '#' {
				slot = nil
				slot = make([]int, 0)
			}
		}
		total += load
	}

	fmt.Println(total)
}
