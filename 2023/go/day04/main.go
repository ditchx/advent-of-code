package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	inputFile := os.Args[1]

	fp, _ := os.Open(inputFile)

	scanner := bufio.NewScanner(fp)
	scanner.Split(bufio.ScanLines)

	var total uint64
	for scanner.Scan() {
		total += processCards(scanner.Text())
	}

	fmt.Println(total)

}

func processCards(line string) uint64 {
	var points uint64

	parts := strings.Split(line, ":")
	sets := strings.Split(parts[1], "|")

	var winning = make(map[string]struct{})
	for _, v := range strings.Fields(sets[0]) {
		winning[v] = struct{}{}
	}

	points = 0
	for _, v := range strings.Fields(sets[1]) {
		if _, ok := winning[v]; ok {
			if points == 0 {
				points = 1
			} else {
				points *= 2
			}
		}

	}

	return points
}
