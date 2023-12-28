package main

import (
	"fmt"
	"os"
	"strings"
)

const up string = "UP"
const down string = "DOWN"
const left string = "LEFT"
const right string = "RIGHT"

type Ray struct {
	X         int
	Y         int
	Direction string
}

var errOutOfBounds error = fmt.Errorf("Ray out of bounds")

func (r Ray) Next() Ray {
	nextRay := newRay(r.X, r.Y, r.Direction)

	switch r.Direction {
	case up:
		nextRay.X--
	case down:
		nextRay.X++
	case left:
		nextRay.Y--
	case right:
		nextRay.Y++
	}

	return nextRay
}

func newRay(x int, y int, direction string) Ray {
	return Ray{
		X:         x,
		Y:         y,
		Direction: direction,
	}
}

func main() {
	file, _ := os.ReadFile(os.Args[1])
	grid := strings.Split(strings.TrimSpace(string(file)), "\n")

	tiles := make(map[string][]Ray)

	initial := newRay(0, 0, right)
	beam(grid, tiles, initial)

	fmt.Println(len(tiles))

	var highest int

	rightMost := len(grid[0]) - 1
	bottomMost := len(grid) - 1
	for i := 0; i < len(grid); i++ {
		tiles = nil
		tiles = make(map[string][]Ray)
		initial := newRay(i, 0, right)
		beam(grid, tiles, initial)

		current := len(tiles)
		highest = max(highest, current)

		tiles = nil
		tiles = make(map[string][]Ray)
		initial = newRay(i, rightMost, left)
		beam(grid, tiles, initial)

		current = len(tiles)
		highest = max(highest, current)

	}

	for i := 0; i < len(grid[0]); i++ {
		tiles = nil
		tiles = make(map[string][]Ray)
		initial := newRay(0, i, down)
		beam(grid, tiles, initial)

		current := len(tiles)
		highest = max(highest, current)

		tiles = nil
		tiles = make(map[string][]Ray)
		initial = newRay(bottomMost, i, up)
		beam(grid, tiles, initial)

		current = len(tiles)
		highest = max(highest, current)

	}

	fmt.Println(highest)
}

func beam(grid []string, tiles map[string][]Ray, r Ray) {

	if !checkBoundary(r, grid) {
		return
	}

	k := makeKey(r.X, r.Y)
	// If tile already contains
	// a ray with same direction,
	// this means we already passed this
	// and should exit to stop cycling/looping
	if rayList, ok := tiles[k]; ok {
		for i := range rayList {
			if rayList[i].Direction == r.Direction {
				return
			}
		}
	}
	tiles[k] = append(tiles[k], r)

	currentTile := grid[r.X][r.Y]

	if '.' == currentTile {
		nextRay := r.Next()
		beam(grid, tiles, nextRay)
		return
	}

	if '|' == currentTile {
		if left == r.Direction || right == r.Direction {

			upRay := newRay(r.X-1, r.Y, up)
			downRay := newRay(r.X+1, r.Y, down)

			beam(grid, tiles, upRay)
			beam(grid, tiles, downRay)
		} else {
			nextRay := r.Next()
			beam(grid, tiles, nextRay)
		}

		return
	}

	if '-' == currentTile {
		if up == r.Direction || down == r.Direction {

			leftRay := newRay(r.X, r.Y-1, left)
			rightRay := newRay(r.X, r.Y+1, right)

			beam(grid, tiles, leftRay)
			beam(grid, tiles, rightRay)

		} else {
			nextRay := r.Next()
			beam(grid, tiles, nextRay)

		}
		return
	}

	if '/' == currentTile {
		nextRay := newRay(r.X, r.Y, r.Direction)
		switch nextRay.Direction {
		case up:
			nextRay.Direction = right
			nextRay.Y++
		case down:
			nextRay.Direction = left
			nextRay.Y--
		case left:
			nextRay.Direction = down
			nextRay.X++
		case right:
			nextRay.Direction = up
			nextRay.X--
		}

		beam(grid, tiles, nextRay)
		return
	}

	if '\\' == currentTile {
		nextRay := newRay(r.X, r.Y, r.Direction)
		switch nextRay.Direction {
		case up:
			nextRay.Direction = left
			nextRay.Y--
		case down:
			nextRay.Direction = right
			nextRay.Y++
		case left:
			nextRay.Direction = up
			nextRay.X--
		case right:
			nextRay.Direction = down
			nextRay.X++
		}

		beam(grid, tiles, nextRay)
		return
	}

}

func makeKey(x, y int) string {
	return fmt.Sprintf("%d-%d", x, y)
}

func checkBoundary(r Ray, grid []string) bool {
	if r.X < 0 {
		return false
	}

	if r.X > len(grid)-1 {
		return false
	}

	if r.Y < 0 {
		return false
	}

	if r.Y > len(grid[0])-1 {
		return false
	}

	return true
}
