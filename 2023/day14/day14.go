package day14

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
	"time"

	"github.com/advent-of-code/2023/util"
)

func Run() {
	input, err := util.ReadLines("./day14/input/input.txt")
	if err != nil {
		fmt.Printf("error reading file: %s", err.Error())
		return
	}

	var start time.Time

	start = time.Now()
	util.PrettyPrint(14, 1, Part1(input), start)

	start = time.Now()
	util.PrettyPrint(14, 2, Part2(input), start)
}

const (
	North = "NORTH"
	East  = "EAST"
	South = "SOUTH"
	West  = "WEST"
)

func Part1(input []string) int {
	grid := parseGrid(input)
	grid = tiltRocks(grid, North)

	return calculateLoad(grid)
}

func Part2(input []string) int {
	var totalCycles = 1000000000
	var cycleStart, cycleEnd int

	grid := parseGrid(input)
	hashes := make([]string, 0)

	for i := 0; i < totalCycles; i++ {
		grid = tiltRocks(grid, North)
		grid = tiltRocks(grid, West)
		grid = tiltRocks(grid, South)
		grid = tiltRocks(grid, East)

		hashes, cycleStart = detectCycle(grid, hashes)
		if cycleStart > 0 {
			cycleEnd = i
			break
		}
	}

	cycleLength := cycleEnd - cycleStart
	remainingCycles := totalCycles - cycleEnd
	for i := 0; i < remainingCycles%cycleLength-1; i++ {
		grid = tiltRocks(grid, North)
		grid = tiltRocks(grid, West)
		grid = tiltRocks(grid, South)
		grid = tiltRocks(grid, East)
	}

	return calculateLoad(grid)
}

func parseGrid(input []string) [][]rune {
	grid := make([][]rune, len(input))
	for i := range input {
		grid[i] = []rune(input[i])
	}

	return grid
}

func tiltRocks(grid [][]rune, direction string) [][]rune {
	switch direction {
	case North, South:
		for i := 0; i < len(grid); i++ {
			grid = rollRocks(grid, direction)
		}
	case East, West:
		for i := 0; i < len(grid[0]); i++ {
			grid = rollRocks(grid, direction)
		}
	}
	
	return grid
}

func rollRocks(grid [][]rune, direction string) [][]rune {
	var offsetY, offsetX int
	var opposite bool

	switch direction {
	case North:
		offsetY = -1
		opposite = false
	case East:
		offsetX = 1
		opposite = true
	case South:
		offsetY = 1
		opposite = true
	case West:
		offsetX = -1
		opposite = false
	}

	switch opposite {
	case false:
		for y := 0; y < len(grid); y++ {
			if y+offsetY >= 0 {
				for x := 0; x < len(grid[0]); x++ {
					if x+offsetX >= 0 {
						if grid[y][x] == 'O' && grid[y+offsetY][x+offsetX] == '.' {
							grid[y][x] = '.'
							grid[y+offsetY][x+offsetX] = 'O'
						}
					}
				}
			}
		}
	case true:
		for y := len(grid) - 1; y >= 0; y-- {
			if y+offsetY < len(grid) {
				for x := len(grid[0]) - 1; x >= 0; x-- {
					if x+offsetX < len(grid[0]) {
						if grid[y][x] == 'O' && grid[y+offsetY][x+offsetX] == '.' {
							grid[y][x] = '.'
							grid[y+offsetY][x+offsetX] = 'O'
						}
					}
				}
			}
		}
	}

	return grid
}

func calculateLoad(grid [][]rune) int {
	load := 0
	for i := 0; i < len(grid); i++ {
		rocks := strings.Count(string(grid[i]), "O")
		if rocks > 0 {
			load += rocks * (len(grid) - i)
		}
	}

	return load
}

func detectCycle(grid [][]rune, hashes []string) ([]string, int) {
	hash := calculateHash(grid)
	for i := range hashes {
		if hashes[i] == hash {
			return hashes, i
		}
	}
	hashes = append(hashes, hash)

	return hashes, -1
}

func calculateHash(input [][]rune) string {
	hash := md5.New()
	for i := range input {
		hash.Write([]byte(string(input[i])))
	}

	return hex.EncodeToString(hash.Sum(nil))
}
