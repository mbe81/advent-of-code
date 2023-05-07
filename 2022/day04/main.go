package main

import (
	"fmt"
	"github.com/mbe81/advent-of-code/util"
	"strconv"
	"strings"
)

func main() {
	lines := util.ReadFile("./2022/day04/input.txt")

	var countFullOverlap, countPartialOverlap int

	for _, line := range lines {

		s := strings.Split(line, ",")
		sectionOne, sectionTwo := s[0], s[1]

		s1 := strings.Split(sectionOne, "-")
		sectionOneMin, _ := strconv.Atoi(s1[0])
		sectionOneMax, _ := strconv.Atoi(s1[1])

		s2 := strings.Split(sectionTwo, "-")
		sectionTwoMin, _ := strconv.Atoi(s2[0])
		sectionTwoMax, _ := strconv.Atoi(s2[1])

		// Part One
		if sectionOneMin >= sectionTwoMin && sectionOneMax <= sectionTwoMax {
			countFullOverlap++
		} else if sectionTwoMin >= sectionOneMin && sectionTwoMax <= sectionOneMax {
			countFullOverlap++
		}

		//Part Two
		if sectionOneMin <= sectionTwoMax && sectionOneMax >= sectionTwoMin {
			countPartialOverlap++
		} else if sectionTwoMin <= sectionOneMax && sectionTwoMax >= sectionOneMin {
			countPartialOverlap++
		}

	}

	fmt.Println("The number of fully overlapping assignments is", countFullOverlap)
	fmt.Println("The number of pairs with overlapping ranges is", countPartialOverlap)
}
