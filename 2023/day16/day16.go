package day16

import (
	"fmt"
	"time"

	"github.com/advent-of-code/2023/util"
)

func Run() {
	input, err := util.ReadLines("./day16/input/input.txt")
	if err != nil {
		fmt.Printf("error reading file: %s", err.Error())
		return
	}

	var start time.Time

	start = time.Now()
	util.PrettyPrint(16, 1, Part1(input), start)

	start = time.Now()
	util.PrettyPrint(16, 2, Part2(input), start)
}

const (
	North = "NORTH"
	East  = "EAST"
	South = "SOUTH"
	West  = "WEST"
)

type Tile struct {
	energized     bool
	seenFromNorth bool
	seenFromEast  bool
	seenFromSouth bool
	seenFromWest  bool
}

func Part1(input []string) int {
	return totalEnergized(input, 0, 0, East)
}

func Part2(input []string) int {
	var maxEnergized int

	for y := range input {
		for x := range input[y] {
			if y == 0 {
				maxEnergized = max(maxEnergized, totalEnergized(input, x, y, South))
			}
			if y == len(input)-1 {
				maxEnergized = max(maxEnergized, totalEnergized(input, x, y, North))
			}
			if x == 0 {
				maxEnergized = max(maxEnergized, totalEnergized(input, x, y, East))
			}
			if x == len(input[y]) {
				maxEnergized = max(maxEnergized, totalEnergized(input, x, y, West))
			}
		}
	}

	return maxEnergized
}

func totalEnergized(grid []string, x, y int, direction string) int {
	tiles := make([][]Tile, len(grid))
	for i := range tiles {
		tiles[i] = make([]Tile, len(grid[i]))
	}

	followBeam(grid, tiles, x, y, direction)

	var energized int
	for i := range tiles {
		for j := range tiles[i] {
			if tiles[i][j].energized {
				energized++
			}
		}
	}
	return energized
}

func followBeam(grid []string, tiles [][]Tile, x, y int, direction string) {
	// Return if beam goes out of grid
	if x < 0 || y < 0 || y > len(grid)-1 || x > len(grid[y])-1 {
		return
	}

	// Return if beam was already entered from same direction
	if direction == North && tiles[y][x].seenFromSouth {
		return
	}
	if direction == East && tiles[y][x].seenFromWest {
		return
	}
	if direction == South && tiles[y][x].seenFromNorth {
		return
	}
	if direction == West && tiles[y][x].seenFromEast {
		return
	}

	// Record entrance of beam into the tile
	tiles[y][x].energized = true

	if direction == North {
		tiles[y][x].seenFromSouth = true
	}
	if direction == East {
		tiles[y][x].seenFromWest = true
	}
	if direction == South {
		tiles[y][x].seenFromNorth = true
	}
	if direction == West {
		tiles[y][x].seenFromEast = true
	}

	// Follow beam
	switch grid[y][x] {
	case '.':
		switch direction {
		case North:
			followBeam(grid, tiles, x, y-1, North)
		case East:
			followBeam(grid, tiles, x+1, y, East)
		case South:
			followBeam(grid, tiles, x, y+1, South)
		case West:
			followBeam(grid, tiles, x-1, y, West)
		}
	case '/':
		switch direction {
		case North:
			followBeam(grid, tiles, x+1, y, East)
		case East:
			followBeam(grid, tiles, x, y-1, North)
		case South:
			followBeam(grid, tiles, x-1, y, West)
		case West:
			followBeam(grid, tiles, x, y+1, South)
		}
	case '\\':
		switch direction {
		case North:
			followBeam(grid, tiles, x-1, y, West)
		case East:
			followBeam(grid, tiles, x, y+1, South)
		case South:
			followBeam(grid, tiles, x+1, y, East)
		case West:
			followBeam(grid, tiles, x, y-1, North)
		}
	case '|':
		switch direction {
		case North:
			followBeam(grid, tiles, x, y-1, North)
		case East:
			followBeam(grid, tiles, x, y-1, North)
			followBeam(grid, tiles, x, y+1, South)
		case South:
			followBeam(grid, tiles, x, y+1, South)
		case West:
			followBeam(grid, tiles, x, y-1, North)
			followBeam(grid, tiles, x, y+1, South)
		}
	case '-':
		switch direction {
		case North:
			followBeam(grid, tiles, x-1, y, West)
			followBeam(grid, tiles, x+1, y, East)
		case East:
			followBeam(grid, tiles, x+1, y, East)
		case South:
			followBeam(grid, tiles, x-1, y, West)
			followBeam(grid, tiles, x+1, y, East)
		case West:
			followBeam(grid, tiles, x-1, y, West)
		}
	}

	return
}
