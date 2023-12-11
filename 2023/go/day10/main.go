package main

import (
	"fmt"
	"os"
	"strings"
)

type Path struct {
	Distance uint64
	prevRow  int
	prevCol  int
	row      int
	column   int
}

func newPath(prevRow, prevCol, row, column int) *Path {
	return &Path{
		Distance: 1,
		prevRow:  prevRow,
		prevCol:  prevCol,
		row:      row,
		column:   column,
	}
}

func (p *Path) Next(maze [][]rune) string {
	current := maze[p.row][p.column]

	switch current {
	case '|':
		if p.prevRow < p.row { // Going down
			p.prevCol = p.column
			p.prevRow = p.row
			p.row++
		}

		if p.prevRow > p.row { // Going Up
			p.prevCol = p.column
			p.prevRow = p.row
			p.row--
		}

	case '-':
		if p.prevCol > p.column { // Left
			p.prevCol = p.column
			p.prevRow = p.row
			p.column--
		} else { // Right
			p.prevCol = p.column
			p.prevRow = p.row
			p.column++
		}
	case 'L':
		if p.prevRow < p.row {
			p.prevCol = p.column
			p.prevRow = p.row
			p.column++
		}

		if p.prevCol > p.column {
			p.prevCol = p.column
			p.prevRow = p.row
			p.row--
		}

	case 'J':
		if p.prevRow < p.row { // Down-Left
			p.prevCol = p.column
			p.prevRow = p.row
			p.column--
		}
		if p.prevCol < p.column { // Right-up
			p.prevCol = p.column
			p.prevRow = p.row
			p.row--
		}
	case '7':
		if p.prevCol < p.column { // Right-Down
			p.prevCol = p.column
			p.prevRow = p.row
			p.row++
		}
		if p.prevRow > p.row { // Up-left
			p.prevCol = p.column
			p.prevRow = p.row
			p.column--
		}
	case 'F':
		if p.prevCol > p.column {
			p.prevCol = p.column
			p.prevRow = p.row
			p.row++
		}

		if p.prevRow > p.row {
			p.prevCol = p.column
			p.prevRow = p.row
			p.column++
		}
	}

	p.Distance++

	return fmt.Sprintf("%d_%d", p.row, p.column)
}

func main() {
	file, _ := os.ReadFile(os.Args[1])
	contents := strings.Split(strings.TrimSpace(string(file)), "\n")

	columnCount := len(contents[0])
	rowCount := len(contents)

	maze := make([][]rune, rowCount)
	var startRow, startColumn int

	for i, line := range contents {
		maze[i] = make([]rune, columnCount)
		for j, c := range line {
			maze[i][j] = c

			if 'S' == c {
				startRow = i
				startColumn = j
			}
		}
	}

	paths := make([]*Path, 0)

	// Up
	if startRow > 0 && (strings.IndexRune("7|F", maze[startRow-1][startColumn]) != -1) {
		paths = append(paths, newPath(startRow, startColumn, startRow-1, startColumn))
	}

	// Down
	if startRow < rowCount && (strings.IndexRune("J|L", maze[startRow+1][startColumn]) != -1) {
		paths = append(paths, newPath(startRow, startColumn, startRow+1, startColumn))
	}

	// Left
	if startColumn > 0 && (strings.IndexRune("-LF", maze[startRow][startColumn-1]) != -1) {
		paths = append(paths, newPath(startRow, startColumn, startRow, startColumn-1))
	}

	// Right
	if startColumn < columnCount && (strings.IndexRune("-J7", maze[startRow][startColumn+1]) != -1) {
		paths = append(paths, newPath(startRow, startColumn, startRow, startColumn+1))
	}

	passed := make(map[string]uint64)
	pipe := make(map[string]struct{})
	pipe[fmt.Sprintf("%d_%d", startRow, startColumn)] = struct{}{}

	var greatest uint64
	for ct := 0; ct < (rowCount * columnCount); ct++ {
		res1 := paths[0].Next(maze)
		res2 := paths[1].Next(maze)

		pipe[res1] = struct{}{}
		pipe[res2] = struct{}{}

		_, ok := passed[res1]
		if res1 == res2 && !ok && res1 != fmt.Sprintf("%d_%d", startRow, startColumn) {
			passed[res1] = paths[0].Distance
			greatest = max(greatest, passed[res1])
		}

	}

	fmt.Println(greatest)

	passed[fmt.Sprintf("%d_%d", startRow, startColumn)] = 0

	// Ray-polygon intersection
	var enclosed uint64
	for i := 0; i < rowCount; i++ {
		for j := 0; j < columnCount; j++ {
			if _, inPipe := pipe[fmt.Sprintf("%d_%d", i, j)]; !inPipe {
				var crossings uint64
				for next := j + 1; next < columnCount; next++ {

					if _, partOfPipe := pipe[fmt.Sprintf("%d_%d", i, next)]; !partOfPipe {
						continue
					}

					if strings.IndexRune("J|L", maze[i][next]) != -1 {
						crossings++
					}
				}
				if crossings%2 == 1 {
					enclosed++
				}
			}
		}
	}

	fmt.Println(enclosed)
}
