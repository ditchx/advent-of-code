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
	inputFile := os.Args[1]

	fp, _ := os.Open(inputFile)

	scanner := bufio.NewScanner(fp)
	scanner.Split(bufio.ScanLines)

	var row uint
	for scanner.Scan() {
		schematic(partValues, partLocation, &symbols, row, []byte(scanner.Text()))
		row++
	}

	//fmt.Printf("%v\n", partValues)
	//fmt.Printf("%v\n", partLocation)
	//fmt.Printf("%v\n", symbols)

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

		// Lower right
		partID, ok = partLocation[fmt.Sprintf("%d-%d", coords[0]+1, coords[1])]
		if ok {
			added[partID] = struct{}{}
		}

		partID, ok = partLocation[fmt.Sprintf("%d-%d", coords[0]+1, coords[1]+1)]
		if ok {
			added[partID] = struct{}{}
		}

	}

	for k := range added {
		total += partValues[k]
	}

	fmt.Println(total)
}

func schematic(partValues map[string]uint64, partLocation map[string]string, symbols *[][2]uint, row uint, line []byte) {
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
