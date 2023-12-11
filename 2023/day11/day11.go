package day11

import (
	"fmt"
	"strings"
	"time"

	"github.com/advent-of-code/2023/util"
)

func Run() {
	input, err := util.ReadLines("./day11/input/input.txt")
	if err != nil {
		fmt.Printf("error reading file: %s", err.Error())
		return
	}

	var start time.Time

	start = time.Now()
	util.PrettyPrint(9, 1, Part1(input), start)

	start = time.Now()
	util.PrettyPrint(9, 2, Part2(input), start)
}

func Part1(input []string) int {
	galaxies := parseImage(input)

	expansionFactor := 2
	totalSteps := 0

	for i, g := range galaxies {
		for j, h := range galaxies {
			if i < j { // only calculate each connection once
				totalSteps += calculateDistance(g, h, expansionFactor, input)
			}
		}
	}

	return totalSteps
}

func Part2(input []string) int {
	galaxies := parseImage(input)

	expansionFactor := 1000000
	totalSteps := 0

	for i, g := range galaxies {
		for j, h := range galaxies {
			if i < j { // only calculate each connection once
				totalSteps += calculateDistance(g, h, expansionFactor, input)
			}
		}
	}

	return totalSteps
}

type Galaxy struct {
	Y int
	X int
}

func parseImage(input []string) []Galaxy {
	// duplicate rows without galaxies
	for y := 0; y < len(input); y++ {
		var galaxyFound = false
		for x := 0; x < len(input[y]); x++ {
			if input[y][x] == '#' {
				galaxyFound = true
			}
		}
		if !galaxyFound {
			input[y] = strings.Repeat("-", len(input[y]))
		}
		galaxyFound = false
	}

	// duplicate columns without galaxies
	for x := 0; x < len(input[0]); x++ {
		var galaxyFound = false
		for y := 0; y < len(input); y++ {
			if input[y][x] == '#' {
				galaxyFound = true
			}
		}
		if !galaxyFound {
			for y := 0; y < len(input); y++ {
				if input[y][x] == '.' {
					input[y] = input[y][:x] + "|" + input[y][x+1:]
				} else if input[y][x] == '-' {
					input[y] = input[y][:x] + "+" + input[y][x+1:]
				}
			}
		}
		galaxyFound = false
	}

	var galaxies []Galaxy

	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[y]); x++ {
			if input[y][x] == '#' {
				galaxies = append(galaxies, Galaxy{y, x})
			}
		}
	}
	return galaxies
}

func calculateDistance(a, b Galaxy, expansionFactor int, input []string) int {
	var steps int

	for k := min(a.Y, b.Y); k < max(a.Y, b.Y); k++ {
		if input[k][a.X] == '-' || input[k][a.X] == '+' {
			steps = steps + expansionFactor
		} else {
			steps++
		}
	}

	for k := min(a.X, b.X); k < max(a.X, b.X); k++ {
		if input[a.Y][k] == '|' || input[a.Y][k] == '+' {
			steps = steps + expansionFactor
		} else {
			steps++
		}
	}
	return steps
}
