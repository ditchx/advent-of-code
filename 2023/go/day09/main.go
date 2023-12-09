package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.ReadFile(os.Args[1])
	contents := strings.Split(strings.TrimSpace(string(file)), "\n")

	var extrapolated int64
	for _, line := range contents {
		readings := make([]int64, 0)
		for _, value := range strings.Fields(line) {
			v, _ := strconv.ParseInt(value, 10, 32)
			readings = append(readings, v)
		}

		extrapolated += extrapolate(readings)
	}

	fmt.Println(extrapolated)
}

func extrapolate(readings []int64) int64 {

	intervals := make([]int64, 0)

	keys := make(map[int64]struct{})
	for ct := 1; ct < len(readings); ct++ {
		diff := readings[ct] - readings[ct-1]
		keys[diff] = struct{}{}
		intervals = append(intervals, diff)
	}

	if len(keys) == 1 {
		for k := range keys {
			return readings[len(readings)-1] + k
		}
	}

	return readings[len(readings)-1] + extrapolate(intervals)

}
