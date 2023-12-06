package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	var line []string
	scanner.Scan()
	line = strings.Fields(scanner.Text())
	line = line[1:]
	for _, f := range line {
		value, err := strconv.ParseUint(f, 10, 32)
		if err == nil {
			time = append(time, value)
		}
	}

	longRaceTime, _ := strconv.ParseUint(strings.Join(line, ""), 10, 32)

	scanner.Scan()
	line = strings.Fields(scanner.Text())
	line = line[1:]
	for _, f := range line {
		value, err := strconv.ParseUint(f, 10, 32)
		if err == nil {
			distance = append(distance, value)
		}
	}
	longRaceDistance, _ := strconv.ParseUint(strings.Join(line, ""), 10, 64)

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

	x := sort.Search(int(longRaceTime), func(i int) bool {
		return uint64(i)*uint64(int(longRaceTime)-i) > uint64(longRaceDistance)
	})

	y := sort.Search(int(longRaceTime), func(i int) bool {
		return uint64(i)*uint64(int(longRaceTime)-i) <= uint64(longRaceDistance)
	})

	puzzle2 := y - x
	fmt.Printf("%d\n", puzzle2)
}
