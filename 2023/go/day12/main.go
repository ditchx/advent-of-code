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

	var total uint64

	for i := range contents {
		x := countCombinations(contents[i])
		total = total + x
	}

	fmt.Println(total)

}

func validate(pattern string, repetitions []uint8) bool {
	pattern = strings.Trim(pattern, ".")
	check := strings.FieldsFunc(pattern, func(c rune) bool { return '#' != c })

	if len(check) != len(repetitions) {
		return false
	}

	for i := range check {
		if len(check[i]) != int(repetitions[i]) {
			return false
		}
	}

	return true
}

func countCombinations(data string) uint64 {
	var total uint64

	parts := strings.Fields(data)

	repetitions := make([]uint8, 0)

	for _, v := range strings.Split(parts[1], ",") {
		ct, _ := strconv.ParseUint(v, 10, 32)
		repetitions = append(repetitions, uint8(ct))
	}

	combinations := expand(parts[0])

	for i := range combinations {
		if validate(combinations[i], repetitions) {
			total++
		}

	}

	return total
}

func expand(pattern string) []string {
	combinations := make([]string, 0)
	qCount := strings.Count(pattern, "?")

	if qCount == 0 {
		return []string{pattern}
	}

	for _, v := range generatePermutations([]string{".", "#"}, qCount) {
		combinations = append(combinations, substitute(pattern, v))
	}

	return combinations
}

func substitute(pattern string, combination string) string {
	for i := range combination {
		pattern = strings.Replace(pattern, "?", string(combination[i]), 1)
	}

	return pattern
}

func generatePermutations(pool []string, length int) []string {
	var result []string
	generatePermutationsRecursive(pool, length, "", &result)
	return result
}

func generatePermutationsRecursive(pool []string, length int, current string, result *[]string) {
	if length == 0 {
		*result = append(*result, current)
		return
	}
	for _, letter := range pool {
		generatePermutationsRecursive(pool, length-1, current+letter, result)
	}
}
