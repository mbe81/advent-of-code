package day13

import (
	"fmt"
	"time"

	"github.com/advent-of-code/2023/util"
	"github.com/advent-of-code/2023/util/stringutil"
)

func Run() {
	input, err := util.ReadLines("./day13/input/input.txt")
	if err != nil {
		fmt.Printf("error reading file: %s", err.Error())
		return
	}

	var start time.Time

	start = time.Now()
	util.PrettyPrint(13, 1, Part1(input), start)

	start = time.Now()
	util.PrettyPrint(13, 2, Part2(input), start)
}

func Part1(input []string) int {
	var totalScore int

	patterns := parsePatterns(input)
	for _, p := range patterns {
		var score int

		// Find horizontal match
		score = findReflection(p, 0)
		if score > 0 {
			totalScore += score * 100
			continue
		}

		// Find vertical match
		p = stringutil.Transpose(p)
		score = findReflection(p, 0)
		if score > 0 {
			totalScore += score
			continue
		}
	}

	return totalScore
}

func Part2(input []string) int {
	var totalScore int

	patterns := parsePatterns(input)

	for _, p := range patterns {
		var score int

		// Find horizontal match
		score = findReflection(p, 1)
		if score > 0 {
			totalScore += score * 100
			continue
		}

		// Find vertical match
		p = stringutil.Transpose(p)
		score = findReflection(p, 1)
		if score > 0 {
			totalScore += score
			continue
		}
	}

	return totalScore
}

func parsePatterns(input []string) [][]string {
	var patterns [][]string
	var p []string

	for _, l := range input {
		if l == "" {
			patterns = append(patterns, p)
			p = make([]string, 0)
		} else {
			p = append(p, l)
		}
	}
	patterns = append(patterns, p)

	return patterns
}

func findReflection(pattern []string, maxDifference int) int {
	for i := 0; i < len(pattern)-1; i++ {
		smudges := 0
		for j := 0; j <= i; j++ {
			if i-j < 0 || i+j+1 > len(pattern)-1 {
				break // End of pattern reached
			}
			smudges += countSmudges(pattern[i-j], pattern[i+j+1])
		}
		if smudges == maxDifference {
			return i + 1
		}
	}
	return -1
}

func countSmudges(a, b string) int {
	var smudges int
	for i := range a {
		if a[i] != b[i] {
			smudges++
		}
	}

	return smudges
}
