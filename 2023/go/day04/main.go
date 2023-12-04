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
	cards := make(map[string]uint)
	var cardNum uint

	for scanner.Scan() {
		line := scanner.Text()
		total += processCards(line)

		countCards(cardNum, line, cards)
		cardNum++
	}

	fmt.Println(total)

	var cardTotal uint
	for _, v := range cards {
		cardTotal += v
	}

	fmt.Println(cardTotal)
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

func countCards(cardNum uint, line string, cards map[string]uint) {
	parts := strings.Split(line, ":")
	sets := strings.Split(parts[1], "|")

	var winning = make(map[string]struct{})
	for _, v := range strings.Fields(sets[0]) {
		winning[v] = struct{}{}
	}

	var index = fmt.Sprint(cardNum)
	val, ok := cards[index]

	if !ok {
		val = 1
		cards[index] = val
	}

	var totalWins = 0
	for _, v := range strings.Fields(sets[1]) {
		if _, ok := winning[v]; ok {
			totalWins++

			nextCard := fmt.Sprint((cardNum + uint(totalWins)))

			nextVal, nextOk := cards[nextCard]

			if !nextOk {
				nextVal = 1
				cards[nextCard] = nextVal
			}

			cards[nextCard] += val

		}
	}

}
