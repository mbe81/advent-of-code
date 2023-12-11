package day11

import (
	"fmt"
	"strings"
	"time"

	"github.com/advent-of-code/2023/util"
	"github.com/advent-of-code/2023/util/math"
)

func Run() {
	input, err := util.ReadLines("./day11/input/example.txt")
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
	var totalSteps = 0
	for i, g := range galaxies {
		for j, h := range galaxies {
			if i < j { // only calculate each distance once
				minSteps := math.Abs(g.y-h.y) + math.Abs(g.x-h.x)
				totalSteps += minSteps
				//fmt.Println("Calculate g1:", g, ", g2:", h, ", steps", minSteps)
			}
		}
	}
	return totalSteps
}

func Part2(input []string) int {

	return 0
}

type Galaxy struct {
	y int
	x int
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
			part1 := input[:y]
			part2 := input[y:]
			input = append(part1, strings.Repeat(".", len(input[y])))
			input = append(input, part2...)
			y = y + 1
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
				input[y] = input[y][:x] + "." + input[y][x:]
			}
			x = x + 1
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
