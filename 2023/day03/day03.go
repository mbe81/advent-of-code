package day03

import (
	"fmt"
	"regexp"
	"strconv"
	"time"

	"github.com/advent-of-code/2023/util"
)

func Run() {
	input, err := util.ReadLines("./day03/input/puzzle.txt")
	if err != nil {
		fmt.Printf("error reading file: %s", err.Error())
		return
	}

	var start time.Time

	start = time.Now()
	util.PrettyPrint(3, 1, Part1(input), start)

	start = time.Now()
	util.PrettyPrint(3, 2, Part2(input), start)
}

type PartNumber struct {
	Line     int
	StartPos int
	EndPos   int
	Value    int
}

func Part1(input []string) int {

	var sumNumbers int

	partNumbers := getPartNumbers(input)
	for _, partNumber := range partNumbers {

		// Determine coordinates for adjacent symbols
		xStart := max(partNumber.StartPos-1, 0)
		xEnd := min(partNumber.EndPos+1, len(input[0])-1)

		yStart := max(partNumber.Line-1, 0)
		yEnd := min(partNumber.Line+1, len(input)-1)

	out:
		for y := yStart; y <= yEnd; y++ {
			for x := xStart; x <= xEnd; x++ {
				c := input[y][x]
				if !(c >= 48 && c <= 57) && c != 46 {
					// No numeric character, no dot --> must be a symbol
					sumNumbers += partNumber.Value
					break out
				}
			}
		}
	}

	return sumNumbers
}

type Gear struct {
	Line int
	Pos  int
}

func Part2(input []string) int {

	var sumGearRatio int

	partNumbers := getPartNumbers(input)
	partNumbersToMultiply := make(map[Gear][]PartNumber)

	for _, partNumber := range partNumbers {

		// Determine coordinates for adjacent symbols
		xStart := max(partNumber.StartPos-1, 0)
		xEnd := min(partNumber.EndPos+1, len(input[0])-1)

		yStart := max(partNumber.Line-1, 0)
		yEnd := min(partNumber.Line+1, len(input)-1)

		for y := yStart; y <= yEnd; y++ {
			for x := xStart; x <= xEnd; x++ {
				c := input[y][x]
				if string(c) == "*" {
					gear := Gear{
						Line: y,
						Pos:  x,
					}

					partNumbersToMultiply[gear] = append(partNumbersToMultiply[gear], partNumber)
				}
			}
		}
	}

	for _, parts := range partNumbersToMultiply {
		if len(parts) == 2 {
			sumGearRatio += parts[0].Value * parts[1].Value
		}
	}

	return sumGearRatio
}

func getPartNumbers(input []string) []PartNumber {
	var partNumbers []PartNumber

	re := regexp.MustCompile(`\d+`)
	for lineNum, lineStr := range input {

		numbers := re.FindAllString(lineStr, -1)
		indexes := re.FindAllStringIndex(lineStr, -1)

		for i, number := range numbers {
			value, _ := strconv.Atoi(number)
			partNumbers = append(partNumbers, PartNumber{
				Line:     lineNum,
				StartPos: indexes[i][0],
				EndPos:   indexes[i][1] - 1,
				Value:    value,
			})
		}
	}

	return partNumbers
}
