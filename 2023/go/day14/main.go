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

	states := make(map[string]uint64)
	steps := make([]string, 0)
	var start, end uint64
	present := false
	dish := strings.TrimSpace(string(file))
	for {
		dish = cycle(dish, 1)
		start, present = states[dish]

		steps = append(steps, dish)
		if present {
			break
		}

		states[dish] = end
		end++
	}

	index := ((1000000000 - start) % (end - start)) + start - 1
	final := steps[index]

	contents = strings.Split(final, "\n")
	maxWeight = len(contents)
	total = 0

	for i := range contents {
		total += uint64(maxWeight-i) * uint64(strings.Count(contents[i], "O"))
	}

	fmt.Println(total)

}

func cycle(dish string, count int) string {
	var result string
	result = dish
	for ct := 0; ct < count; ct++ {
		result = north(result)
		result = west(result)
		result = south(result)
		result = east(result)
	}

	return result
}

func north(dish string) string {
	contents := strings.Split(strings.TrimSpace(dish), "\n")

	shifted := make([][]rune, 0)

	for range contents {
		shifted = append(shifted, []rune(strings.Repeat(".", len(contents[0]))))
	}

	for j := range contents[0] {
		slot := make([]int, 0)
		for i := range contents {
			shifted[i][j] = rune(contents[i][j])
			if contents[i][j] == '.' {
				slot = append(slot, i)
			}

			if contents[i][j] == '#' {
				slot = nil
				slot = make([]int, 0)
			}

			if contents[i][j] == 'O' {
				if len(slot) > 0 {
					shifted[slot[0]][j] = 'O'
					shifted[i][j] = '.'
					slot = slot[1:]
					slot = append(slot, i)
				}
			}
		}
	}

	result := make([]string, 0)
	for i := range shifted {
		result = append(result, string(shifted[i]))
	}

	return strings.Join(result, "\n")

}

func west(dish string) string {
	contents := strings.Split(strings.TrimSpace(dish), "\n")

	shifted := make([][]rune, 0)

	for range contents {
		shifted = append(shifted, []rune(strings.Repeat(".", len(contents[0]))))
	}

	for i := range contents {
		slot := make([]int, 0)
		for j := range contents[i] {
			shifted[i][j] = rune(contents[i][j])
			if contents[i][j] == '.' {
				slot = append(slot, j)
			}

			if contents[i][j] == '#' {
				slot = nil
				slot = make([]int, 0)
			}

			if contents[i][j] == 'O' {
				if len(slot) > 0 {
					shifted[i][slot[0]] = 'O'
					shifted[i][j] = '.'
					slot = slot[1:]
					slot = append(slot, j)
				}
			}
		}
	}

	result := make([]string, 0)
	for i := range shifted {
		result = append(result, string(shifted[i]))
	}

	return strings.Join(result, "\n")
}

func south(dish string) string {
	contents := strings.Split(strings.TrimSpace(dish), "\n")

	shifted := make([][]rune, 0)

	for range contents {
		shifted = append(shifted, []rune(strings.Repeat(".", len(contents[0]))))
	}

	for j := range contents[0] {
		slot := make([]int, 0)
		for i := len(contents) - 1; i >= 0; i-- {
			shifted[i][j] = rune(contents[i][j])
			if contents[i][j] == '.' {
				slot = append(slot, i)
			}

			if contents[i][j] == '#' {
				slot = nil
				slot = make([]int, 0)
			}

			if contents[i][j] == 'O' {
				if len(slot) > 0 {
					shifted[slot[0]][j] = 'O'
					shifted[i][j] = '.'
					slot = slot[1:]
					slot = append(slot, i)
				}
			}
		}
	}

	result := make([]string, 0)
	for i := range shifted {
		result = append(result, string(shifted[i]))
	}

	return strings.Join(result, "\n")

}

func east(dish string) string {
	contents := strings.Split(strings.TrimSpace(dish), "\n")

	shifted := make([][]rune, 0)

	for range contents {
		shifted = append(shifted, []rune(strings.Repeat(".", len(contents[0]))))
	}

	for i := range contents {
		slot := make([]int, 0)
		for j := len(contents[i]) - 1; j >= 0; j-- {
			shifted[i][j] = rune(contents[i][j])
			if contents[i][j] == '.' {
				slot = append(slot, j)
			}

			if contents[i][j] == '#' {
				slot = nil
				slot = make([]int, 0)
			}

			if contents[i][j] == 'O' {
				if len(slot) > 0 {
					shifted[i][slot[0]] = 'O'
					shifted[i][j] = '.'
					slot = slot[1:]
					slot = append(slot, j)
				}
			}
		}
	}

	result := make([]string, 0)
	for i := range shifted {
		result = append(result, string(shifted[i]))
	}

	return strings.Join(result, "\n")
}
