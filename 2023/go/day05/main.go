package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Mapping struct {
	DestinationRangeStart uint64
	SourceRangeStart      uint64
	RangeLength           uint64
}

var ErrLessThan = fmt.Errorf("Target is less than range")
var ErrGreaterThan = fmt.Errorf("Target is greater than range")

func newMapping(dest uint64, src uint64, length uint64) *Mapping {
	return &Mapping{
		DestinationRangeStart: dest,
		SourceRangeStart:      src,
		RangeLength:           length,
	}
}

func (m *Mapping) Map(target uint64) (uint64, error) {
	if m.SourceRangeStart > target {
		return 0, ErrLessThan
	}

	if target > (m.SourceRangeStart + m.RangeLength) {
		return 0, ErrGreaterThan
	}

	mapped := target - m.SourceRangeStart + m.DestinationRangeStart

	return mapped, nil
}

type AlmanacMap struct {
	name     string
	mappings []*Mapping
}

func newAlmanacMap(name string) *AlmanacMap {
	return &AlmanacMap{
		name: name,
	}
}

func (a *AlmanacMap) AddMapping(m *Mapping) {
	a.mappings = append(a.mappings, m)
}

func (a *AlmanacMap) Map(target uint64) uint64 {

	for _, m := range a.mappings {
		if found, err := m.Map(target); err == nil {
			return found
		}
	}

	return target
}

func main() {
	var seeds []uint64
	var mappingTypes []string
	almanac := make(map[string]*AlmanacMap)

	inputFile := os.Args[1]

	fp, _ := os.Open(inputFile)

	scanner := bufio.NewScanner(fp)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	for _, s := range strings.Fields(scanner.Text()) {
		num, err := strconv.ParseUint(s, 10, 32)
		if err == nil {
			seeds = append(seeds, num)
		}
	}

	var lastMapType string
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())

		if len(line) == 0 {
			lastMapType = ""
			continue
		}

		if line[1] == "map:" {
			lastMapType = line[0]
			mappingTypes = append(mappingTypes, lastMapType)
			almanac[lastMapType] = newAlmanacMap(lastMapType)
			continue
		}

		dest, _ := strconv.ParseUint(line[0], 10, 32)
		src, _ := strconv.ParseUint(line[1], 10, 32)
		length, _ := strconv.ParseUint(line[2], 10, 32)

		almanac[lastMapType].AddMapping(newMapping(dest, src, length))

	}

	var location = puzzle1(almanac, mappingTypes, seeds)
	fmt.Println(location)

	location = puzzle2(almanac, mappingTypes, seeds)
	fmt.Println(location)

}

func puzzle1(almanac map[string]*AlmanacMap, mappingTypes []string, seeds []uint64) uint64 {
	var location uint64 = 0
	for i, seed := range seeds {
		l := seed
		for _, mapType := range mappingTypes {
			l = almanac[mapType].Map(l)
		}

		if i == 0 {
			location = l
			continue
		}

		if location >= l {
			location = l
		}
	}

	return location
}

// Brute-force! lol
// @TODO: Optimize or think of a different approach
func puzzle2(almanac map[string]*AlmanacMap, mappingTypes []string, seeds []uint64) uint64 {
	var lowest uint64

	c := make(chan uint64)
	var rangeCount int = len(seeds) / 2

	for ct := 0; ct < len(seeds); ct = ct + 2 {

		go func(start, end uint64) {
			var init bool = false
			var location uint64 = 0

			for start < end {

				l := start
				start++
				for _, mapType := range mappingTypes {
					l = almanac[mapType].Map(l)
				}

				if !init {
					location = l
					init = true
					continue
				}

				if location >= l {
					location = l
				}

			}
			c <- location
		}(seeds[ct], seeds[ct]+seeds[ct+1])
	}

	for ct := 0; ct < rangeCount; ct++ {
		l := <-c
		if ct == 0 {
			lowest = l
		}
		if lowest >= l {
			lowest = l
		}
	}

	return lowest
}
