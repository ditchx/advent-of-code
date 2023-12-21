package main

import (
	"fmt"
	"maps"
	"os"
	"strconv"
	"strings"
)

func main() {

	file, _ := os.ReadFile(os.Args[1])
	contents := strings.Split(strings.TrimSpace(string(file)), "\n")

	var total uint64

	for i := range contents {
		x := countArrangements(contents[i])
		total = total + x
	}

	fmt.Println(total)
}

func countArrangements(data string) uint64 {

	parts := strings.Fields(data)

	repetitions := make([]uint8, 0)

	for _, v := range strings.Split(parts[1], ",") {
		ct, _ := strconv.ParseUint(v, 10, 32)
		repetitions = append(repetitions, uint8(ct))
	}

	arrangements := make(map[string]uint8)

	getArrangements(arrangements, []rune(parts[0]), repetitions)

	var total uint64

	for i := range arrangements {
		if arrangements[i] == 2 {
			total++
		}
	}

	return total
}

func getArrangements(arrangements map[string]uint8, chars []rune, repetitions []uint8) {

	if len(chars) == 0 {
		return
	}

	var next []string
	if '?' == chars[0] {
		next = []string{"#", "."}
	} else {
		next = []string{string(chars[0])}
	}

	if len(arrangements) == 0 {
		for i := range next {
			if ok := check(next[i], repetitions); ok > 0 {
				arrangements[next[i]] = ok
			}
		}

		getArrangements(arrangements, chars[1:], repetitions)
		return
	}

	batch := make(map[string]uint8, 0)
	for pattern := range arrangements {
		for k := range next {
			current := pattern + next[k]
			if ok := check(current, repetitions); ok > 0 {
				batch[current] = ok
			}
		}
		delete(arrangements, pattern)
	}

	maps.Copy(arrangements, batch)
	getArrangements(arrangements, chars[1:], repetitions)

}

func check(pattern string, repetitions []uint8) uint8 {
	pattern = strings.Trim(pattern, ".")
	check := strings.FieldsFunc(pattern, func(c rune) bool { return '#' != c })

	less := false

	if len(check) > len(repetitions) {
		return 0
	}

	for i := range check {
		if len(check[i]) > int(repetitions[i]) {
			return 0
		}

		if len(check[i]) < int(repetitions[i]) {
			if less {
				return 0
			}

			less = true
		}

		if less && len(check[i]) == int(repetitions[i]) {
			return 0
		}

	}

	if less {
		return 1
	}

	if len(check) < len(repetitions) {
		return 1
	}

	return 2
}
