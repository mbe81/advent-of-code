package day12

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/advent-of-code/2023/util"
)

func Run() {
	input, err := util.ReadLines("./day12/input/example.txt")
	if err != nil {
		fmt.Printf("error reading file: %s", err.Error())
		return
	}

	var start time.Time

	start = time.Now()
	util.PrettyPrint(12, 1, Part1(input), start)

	start = time.Now()
	util.PrettyPrint(12, 2, Part2(input), start)
}

func Part1(input []string) int {
	var totalValid = 0
	for _, l := range input {

		s := strings.Split(l, " ")
		record := s[0]

		s = strings.Split(s[1], ",")

		var springs []int
		for i := range s {
			n, _ := strconv.Atoi(s[i])
			springs = append(springs, n)
		}

		//fmt.Println(record, springs)

		var replaceable []int
		for i := range record {
			if record[i] == '?' {
				replaceable = append(replaceable, i)
			}
		}
		// Dots are replaced with a space to be able to use strings.fields
		record = strings.Replace(record, ".", " ", -1)

		countValid := 0
		possibilities := replaceUnknown(record, 0, replaceable, make([]string, 0))

		for _, p := range possibilities {
			if validateSprings(p, springs) {
				//fmt.Println(p, countValid, "VALID")
				countValid++
			} else {
				//fmt.Println(p, countValid, "INVALID")
			}
		}

		totalValid += countValid
	}

	return totalValid
}

func Part2(input []string) int {

	return 0
}

func replaceUnknown(record string, pos int, replaceable []int, options []string) []string {
	// Replace with space
	strpos := replaceable[pos]
	record = record[:strpos] + " " + record[strpos+1:]

	// Replace the next character. After the last possible replacement add the possible record to the options
	if pos < len(replaceable)-1 {
		options = replaceUnknown(record, pos+1, replaceable, options)
	} else {
		if !slices.Contains(options, record) {
			options = append(options, record)
		}
	}

	// Replace with #
	strpos = replaceable[pos]
	record = record[:strpos] + "#" + record[strpos+1:]

	if pos < len(replaceable)-1 {
		options = replaceUnknown(record, pos+1, replaceable, options)
	} else {
		if !slices.Contains(options, record) {
			options = append(options, record)
		}
	}
	return options
}
func validateSprings(record string, wells []int) bool {
	if strings.Contains(record, "?") {
		return false
	}

	found := strings.Fields(record)

	if len(found) != len(wells) {
		return false
	}

	var valid = true
	for i := range found {
		if len(found[i]) != wells[i] {
			valid = false
		}
	}
	return valid
}
