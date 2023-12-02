package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var cube = map[string]uint{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func main() {
	inputFile := os.Args[1]

	fp, _ := os.Open(inputFile)

	scanner := bufio.NewScanner(fp)
	scanner.Split(bufio.ScanLines)

	var total uint
	for scanner.Scan() {
		total += validateGame([]byte(scanner.Text()))
	}

	fmt.Println(total)

}

func validateGame(line []byte) uint {
	var gameNumber uint

	game := strings.Split(string(line), ":")

	fmt.Sscanf(game[0], "Game %d", &gameNumber)
	sets := strings.Split(game[1], ";")

	for _, s := range sets {
		cubes := strings.Split(s, ",")
		for _, c := range cubes {
			var count uint
			var color string
			fmt.Sscanf(c, "%d %s", &count, &color)

			if count > cube[color] {
				return 0
			}
		}
	}

	return gameNumber
}
