package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	file, _ := os.ReadFile(os.Args[1])
	contents := strings.Split(strings.TrimSpace(string(file)), ",")

	var total uint64

	for i := range contents {
		total += hash(contents[i])
	}

	fmt.Println(total)
}

func hash(data string) uint64 {
	var current uint64

	for i := range data {
		current += uint64(data[i])
		current *= 17
		current = current % 256
	}

	return current
}
