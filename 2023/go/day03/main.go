package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	var partValues = make(map[string]uint64)
	var partLocation = make(map[string]string)
	var symbols = make([][2]uint, 0)
	var gears = make([][2]uint, 0)
	inputFile := os.Args[1]

	fp, _ := os.Open(inputFile)

	scanner := bufio.NewScanner(fp)
	scanner.Split(bufio.ScanLines)

	var row uint
	for scanner.Scan() {
		schematic(partValues, partLocation, &symbols, &gears, row, []byte(scanner.Text()))
		row++
	}

	//fmt.Printf("%v\n", partValues)
	//fmt.Printf("%v\n", partLocation)
	//fmt.Printf("%v\n", symbols)
	//fmt.Printf("%v\n", gears)

	var total uint64
	added := make(map[string]struct{})
	for _, coords := range symbols {

		var partID string
		var ok bool

		// Upper left
		partID, ok = partLocation[fmt.Sprintf("%d-%d", coords[0]-1, coords[1]-1)]
		if ok {
			added[partID] = struct{}{}
		}

		// Top
		partID, ok = partLocation[fmt.Sprintf("%d-%d", coords[0]-1, coords[1])]
		if ok {
			added[partID] = struct{}{}
		}

		// Upper right
		partID, ok = partLocation[fmt.Sprintf("%d-%d", coords[0]-1, coords[1]+1)]
		if ok {
			added[partID] = struct{}{}
		}

		// Left
		partID, ok = partLocation[fmt.Sprintf("%d-%d", coords[0], coords[1]-1)]
		if ok {
			added[partID] = struct{}{}
		}

		// Right
		partID, ok = partLocation[fmt.Sprintf("%d-%d", coords[0], coords[1]+1)]
		if ok {
			added[partID] = struct{}{}
		}

		// Lower left
		partID, ok = partLocation[fmt.Sprintf("%d-%d", coords[0]+1, coords[1]-1)]
		if ok {
			added[partID] = struct{}{}
		}

		// Down
		partID, ok = partLocation[fmt.Sprintf("%d-%d", coords[0]+1, coords[1])]
		if ok {
			added[partID] = struct{}{}
		}

		// Lower right
		partID, ok = partLocation[fmt.Sprintf("%d-%d", coords[0]+1, coords[1]+1)]
		if ok {
			added[partID] = struct{}{}
		}

	}

	for k := range added {
		total += partValues[k]
	}

	fmt.Println(total)

	var totalRatio uint64

	for _, coords := range gears {

		var partID string
		var ok bool
		var currentRatio uint64 = 1
		var gearList map[string]struct{}

		gearList = map[string]struct{}{}

		// Upper left
		partID, ok = partLocation[fmt.Sprintf("%d-%d", coords[0]-1, coords[1]-1)]
		if ok {
			gearList[partID] = struct{}{}
		}

		// Top
		partID, ok = partLocation[fmt.Sprintf("%d-%d", coords[0]-1, coords[1])]
		if ok {
			gearList[partID] = struct{}{}
		}

		// Upper right
		partID, ok = partLocation[fmt.Sprintf("%d-%d", coords[0]-1, coords[1]+1)]
		if ok {
			gearList[partID] = struct{}{}
		}

		// Left
		partID, ok = partLocation[fmt.Sprintf("%d-%d", coords[0], coords[1]-1)]
		if ok {
			gearList[partID] = struct{}{}
		}

		// Right
		partID, ok = partLocation[fmt.Sprintf("%d-%d", coords[0], coords[1]+1)]
		if ok {
			gearList[partID] = struct{}{}
		}

		// Lower left
		partID, ok = partLocation[fmt.Sprintf("%d-%d", coords[0]+1, coords[1]-1)]
		if ok {
			gearList[partID] = struct{}{}
		}

		// Down
		partID, ok = partLocation[fmt.Sprintf("%d-%d", coords[0]+1, coords[1])]
		if ok {
			gearList[partID] = struct{}{}
		}

		// Lower right
		partID, ok = partLocation[fmt.Sprintf("%d-%d", coords[0]+1, coords[1]+1)]
		if ok {
			gearList[partID] = struct{}{}
		}

		if len(gearList) != 2 {
			continue
		}

		for k := range gearList {
			currentRatio *= partValues[k]
		}

		totalRatio += currentRatio
	}

	fmt.Println(totalRatio)
}

func schematic(partValues map[string]uint64, partLocation map[string]string, symbols *[][2]uint, gears *[][2]uint, row uint, line []byte) {
	var locs []string
	var currentNumber string

	for i, v := range line {
		if unicode.IsDigit(rune(v)) {
			currentNumber = currentNumber + string(v)
			locs = append(locs, fmt.Sprintf("%d-%d", row, i))
			continue
		}

		if len(currentNumber) > 0 {
			var partID = strings.Join(locs, "_")
			partValues[partID], _ = strconv.ParseUint(currentNumber, 10, 32)

			for _, l := range locs {
				partLocation[l] = partID
			}

			currentNumber = ""
			locs = make([]string, 0)
		}

		if string(v) != "." {
			*symbols = append(*symbols, [2]uint{uint(row), uint(i)})
		}

		if string(v) == "*" {
			*gears = append(*gears, [2]uint{uint(row), uint(i)})
		}

	}

	if len(currentNumber) > 0 {
		var partID = strings.Join(locs, "_")
		partValues[partID], _ = strconv.ParseUint(currentNumber, 10, 32)

		for _, l := range locs {
			partLocation[l] = partID
		}

		currentNumber = ""
		locs = make([]string, 0)
	}

}
