package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	file, _ := os.ReadFile(os.Args[1])
	contents := strings.Split(strings.TrimSpace(string(file)), "\n")

	coords := make([][2]uint64, 0)

	expandedRowCount := make([]uint64, 0)
	expandedColCount := make([]uint64, 0)

	var colCount uint64
	var rowCount uint64
	for j := range contents[0] {
		expandCol := true
		for i := range contents {

			if len(contents) > len(expandedRowCount) {
				if strings.IndexRune(contents[i], '#') == -1 {
					rowCount++
				}
				expandedRowCount = append(expandedRowCount, rowCount)
			}

			if '#' == contents[i][j] {
				expandCol = false
			}
		}

		if expandCol {
			colCount++
		}

		expandedColCount = append(expandedColCount, colCount)
	}

	for i := range contents {
		for j := range contents[i] {
			if '#' == rune(contents[i][j]) {
				coords = append(coords, [2]uint64{
					uint64(i) + uint64(expandedRowCount[i]),
					uint64(j) + uint64(expandedColCount[j]),
				})
			}
		}
	}

	var sum uint64

	for i := 0; i < len(coords)-1; i++ {
		for j := i + 1; j < len(coords); j++ {
			sum += distance(coords[i], coords[j])
		}
	}

	fmt.Println(sum)
}

func distance(a, b [2]uint64) uint64 {

	return uint64(math.Abs(float64(b[0])-float64(a[0]))) + uint64(math.Abs(float64(b[1])-float64(a[1])))
}

func pair(p []string) {
	for i := 0; i < len(p)-1; i++ {
		for j := i + 1; j < len(p); j++ {
			fmt.Printf("%s %s\n", p[i], p[j])
		}
	}
}
