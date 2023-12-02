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

	var puzzle1 uint
	var puzzle2 uint

	for scanner.Scan() {
		line := []byte(scanner.Text())
		puzzle1 += validateGame(line)
		puzzle2 += getPower(line)
	}

	fmt.Printf("Puzzle 1: %d\n", puzzle1)
	fmt.Printf("Puzzle 2: %d\n", puzzle2)

}

func getPower(line []byte) uint {
	var minCubes = map[string]uint{
		"red":   0,
		"green": 0,
		"blue":  0,
	}

	game := strings.Split(string(line), ":")
	sets := strings.Split(game[1], ";")

	for _, s := range sets {
		cubes := strings.Split(s, ",")
		for _, c := range cubes {
			var count uint
			var color string
			fmt.Sscanf(c, "%d %s", &count, &color)

			if count >= minCubes[color] {
				minCubes[color] = count
			}
		}
	}

	return minCubes["red"] * minCubes["green"] * minCubes["blue"]
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
