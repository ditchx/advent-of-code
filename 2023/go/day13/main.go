package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	file, _ := os.ReadFile(os.Args[1])
	contents := strings.Split(strings.TrimSpace(string(file)), "\n\n")

	var total uint64
	smudged := false

	original := make(map[string][2]uint64)
	for _, c := range contents {

		v := vertical(c, smudged)
		h := horizontal(c, smudged)

		original[c] = [2]uint64{v, h}
		total = total + v + h
	}

	fmt.Println(total)

	total = 0
	smudged = true
	for _, c := range contents {

		v := vertical(c, smudged)
		h := horizontal(c, smudged)

		if original[c][0] == v && original[c][1] != h {
			v = 0
		}

		if original[c][0] != v && original[c][1] == h {
			h = 0
		}

		total = total + v + h
	}

	fmt.Println(total)
}

func vertical(data string, smudged bool) uint64 {
	var value uint64
	lines := strings.Split(data, "\n")
	canditates := make(map[int]uint)

	for i := range lines {
		var previous = make([]rune, 0)
		for j := range lines[i] {
			if len(previous) == 0 {
				previous = append(previous, rune(lines[i][j]))
				continue
			}

			if previous[len(previous)-1] == rune(lines[i][j]) {
				previous = append(previous, rune(lines[i][j]))

				index := j
				left := index - 1
				right := index
				mirrored := true
				for {
					if left < 0 {
						break
					}
					if right >= len(lines[i]) {
						break
					}

					if lines[i][left] != lines[i][right] {
						mirrored = false
						break
					}

					left--
					right++
				}

				if mirrored {
					canditates[index]++

					if len(lines) == int(canditates[index]) {
						value = uint64(index)
					}
				}

				continue
			}

			previous = nil
			previous = append([]rune{}, rune(lines[i][j]))

		}
	}

	if smudged {
		for k, v := range canditates {
			if int(v) == len(lines)-1 {
				return uint64(k)
			}
		}
	}

	return value
}

func horizontal(data string, smudged bool) uint64 {
	var value uint64
	lines := strings.Split(data, "\n")
	canditates := make(map[int]uint)

	for j := range lines[0] {
		var previous = make([]rune, 0)
		for i := range lines {
			if len(previous) == 0 {
				previous = append(previous, rune(lines[i][j]))
				continue
			}

			if previous[len(previous)-1] == rune(lines[i][j]) {
				previous = append(previous, rune(lines[i][j]))

				index := i

				up := index - 1
				down := index
				mirrored := true
				for {
					if up < 0 {
						break
					}

					if down >= len(lines) {
						break
					}

					if lines[up][j] != lines[down][j] {
						mirrored = false
					}

					up--
					down++
				}

				if mirrored {
					canditates[index]++

					if len(lines[0]) == int(canditates[index]) {
						value = uint64(index * 100)
					}

				}

				continue
			}

			previous = nil
			previous = append([]rune{}, rune(lines[i][j]))

		}
	}

	if smudged {
		for k, v := range canditates {
			if int(v) == len(lines[0])-1 {
				return uint64(k * 100)
			}
		}
	}

	return value

}
