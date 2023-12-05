package util

import (
	"fmt"
	"strings"
	"time"
)

var firstOutput = true

func PrettyPrint(day, part, result int, start time.Time) {
	if firstOutput {
		firstOutput = false
		fmt.Println(strings.Repeat("-", 80))
		fmt.Println(" Advent of Code 2023")
		fmt.Println(strings.Repeat("-", 80))
		fmt.Println()
	}

	fmt.Println(" Running day", day, "part", part)
	fmt.Println("  - Result:", result)
	fmt.Println("  - Running time:", time.Since(start))
	fmt.Println()
}
