package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type HandType uint8

const (
	HighCard HandType = iota
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

var cardRank = map[rune]uint{
	'2': 0,
	'3': 1,
	'4': 2,
	'5': 3,
	'6': 4,
	'7': 5,
	'8': 6,
	'9': 7,
	'T': 8,
	'J': 9,
	'Q': 10,
	'K': 11,
	'A': 12,
}

func main() {
	inputFile := os.Args[1]

	fp, _ := os.Open(inputFile)

	scanner := bufio.NewScanner(fp)
	scanner.Split(bufio.ScanLines)

	hands := make(map[string]uint64)
	ranking := make([]string, 0)

	for scanner.Scan() {
		line := strings.Fields(scanner.Text())

		hands[line[0]], _ = strconv.ParseUint(line[1], 10, 32)
		ranking = append(ranking, line[0])
	}

	sort.Slice(ranking, func(i, j int) bool {
		rankI := checkHandType(ranking[i])
		rankJ := checkHandType(ranking[j])

		if rankI != rankJ {
			return rankI < rankJ
		}

		for ct := 0; ct < 5; ct++ {
			runeI := rune(ranking[i][ct])
			runeJ := rune(ranking[j][ct])

			if cardRank[runeI] != cardRank[runeJ] {
				return cardRank[runeI] < cardRank[runeJ]
			}
		}

		return false
	})

	var winnings uint64

	for i, v := range ranking {
		winnings += uint64(i+1) * hands[v]
	}

	fmt.Println(winnings)
}

func checkHandType(hand string) HandType {
	letters := make(map[rune]uint)
	var maxCount uint
	for _, v := range hand {
		count, ok := letters[v]

		if !ok {
			count = 0
		}

		letters[v] = count + 1

		if letters[v] >= maxCount {
			maxCount = letters[v]
		}
	}

	if maxCount == 3 {
		if len(letters) == 2 {
			return FullHouse
		}

		return ThreeOfAKind
	}

	if maxCount == 2 {
		if len(letters) == 3 {
			return TwoPair
		}
		return OnePair
	}

	if len(letters) == 2 && maxCount == 4 {
		return FourOfAKind
	}

	if len(letters) == 1 {
		return FiveOfAKind
	}

	return HighCard
}
