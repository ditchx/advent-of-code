package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Lens struct {
	Label       string
	FocalLength uint8
}

func newLens(label string, focalLength uint8) Lens {
	return Lens{
		Label:       label,
		FocalLength: focalLength,
	}
}

func main() {
	file, _ := os.ReadFile(os.Args[1])
	contents := strings.Split(strings.TrimSpace(string(file)), ",")

	var total uint64

	for i := range contents {
		total += hash(contents[i])
	}

	fmt.Println(total)

	fmt.Println(hashmap(contents))
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

func hashmap(data []string) uint64 {
	var totalFocusingPower uint64

	boxes := make([][]Lens, 256)

	for i := range data {
		if strings.Contains(data[i], "=") {
			parts := strings.Split(data[i], "=")

			boxNumber := hash(parts[0])
			lens, _ := strconv.ParseUint(parts[1], 10, 32)

			found := false
			for x := range boxes[boxNumber] {
				if boxes[boxNumber][x].Label == parts[0] {
					boxes[boxNumber][x].FocalLength = uint8(lens)
					found = true
				}
			}

			if false == found {
				boxes[boxNumber] = append(boxes[boxNumber], newLens(parts[0], uint8(lens)))
			}

			continue
		}

		label := strings.TrimRight(data[i], "-")
		boxNumber := hash(label)

		for x := range boxes[boxNumber] {
			if boxes[boxNumber][x].Label == label {
				copy(boxes[boxNumber][x:], boxes[boxNumber][x+1:])
				boxes[boxNumber][len(boxes[boxNumber])-1] = Lens{}
				boxes[boxNumber] = boxes[boxNumber][:len(boxes[boxNumber])-1]
				break
			}
		}

	}

	for boxNumber := range boxes {
		for slotNumber := range boxes[boxNumber] {
			totalFocusingPower += uint64(boxNumber+1) * uint64(slotNumber+1) * uint64(boxes[boxNumber][slotNumber].FocalLength)
		}
	}

	return totalFocusingPower
}
