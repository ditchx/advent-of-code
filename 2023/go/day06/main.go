package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputFile := os.Args[1]

	fp, _ := os.Open(inputFile)

	scanner := bufio.NewScanner(fp)
	scanner.Split(bufio.ScanLines)

	var time []uint64
	var distance []uint64

	scanner.Scan()
	for _, f := range strings.Fields(scanner.Text()) {
		value, err := strconv.ParseUint(f, 10, 32)
		if err == nil {
			time = append(time, value)
		}
	}

	scanner.Scan()
	for _, f := range strings.Fields(scanner.Text()) {
		value, err := strconv.ParseUint(f, 10, 32)
		if err == nil {
			distance = append(distance, value)
		}
	}

	var raceCount = len(time)
	var puzzle1 uint64 = 1

	for raceNumber := 0; raceNumber < raceCount; raceNumber++ {
		var chances uint
		var milli uint64
		for milli = 0; milli < time[raceNumber]; milli++ {

			if (milli * (time[raceNumber] - milli)) > distance[raceNumber] {
				chances++
			}
		}
		puzzle1 *= uint64(chances)
	}

	fmt.Println(puzzle1)

}
