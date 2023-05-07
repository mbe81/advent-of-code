package main

import (
	"fmt"
	"github.com/mbe81/advent-of-code/util"
	"strings"
)

func main() {
	lines := util.ReadFile("./2022/day03/input.txt")

	// Part One
	priority := 0
	for _, line := range lines {
		l := len(line)
		compOne := line[0 : l/2]
		compTwo := line[l/2 : l]
		duplicateItems := findItemsInBothCompartments(compOne, compTwo)
		priority = priority + calculatePriority(duplicateItems)
	}

	fmt.Println("The sum of the priorities which appear in both compartments is", priority)

	// Part Two
	l := len(lines)
	priority = 0
	for i := 0; i < l; i = i + 3 {
		rucksackOne := lines[i]
		rucksackTwo := lines[i+1]
		rucksackThree := lines[i+2]
		groupBadge := findGroupBadge(rucksackOne, rucksackTwo, rucksackThree)
		priority = priority + calculatePriority(groupBadge)
	}

	fmt.Println("The sum of the priorities for all group badges is", priority)

}

func findItemsInBothCompartments(compOne string, compTwo string) string {
	var duplicates string
	for _, item := range compOne {
		for _, item2 := range compTwo {
			if item == item2 {
				// duplicate found
				duplicates = duplicates + string(item)
				compTwo = strings.ReplaceAll(compTwo, string(item), "")
				break
			}
		}
	}
	return duplicates
}

func findGroupBadge(rucksackOne string, rucksackTwo string, rucksackThree string) string {
	for _, item := range rucksackOne {
		for _, item2 := range rucksackTwo {
			if item == item2 {
				for _, item3 := range rucksackThree {
					if item == item3 {
						// found the group badge
						return string(item)
					}
				}
			}
		}
	}
	return ""
}

func calculatePriority(items string) int {
	var total int32
	for _, item := range items {
		if item >= 65 && item <= 90 {
			total = total + (item - 64 + 26)
		} else if item >= 97 && item <= 1022 {
			total = total + (item - 96)
		}
	}
	return int(total)
}
