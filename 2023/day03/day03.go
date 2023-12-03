package day03

import (
	"fmt"
	"strconv"

	"github.com/advent-of-code/2023/util"
)

func Run() {
	input, err := util.ReadLines("./day03/input/puzzle.txt")
	if err != nil {
		fmt.Printf("Error reading file: %e", err)
		return
	}

	fmt.Println("Running Day 3, Part 1")
	fmt.Println("Result: ", Part1(input))

	fmt.Println("Running Day 2, Part 2")
	fmt.Println("Result: ", Part2(input))
}

type PartNumber struct {
	Line   int
	Pos    int
	Length int
	Value  int
}

func Part1(input []string) int {

	var sumNumbers int

	partNumbers := getPartNumbers(input)
	for _, partNumber := range partNumbers {

		// Determine coordinates for adjacent symbols
		xStart := max(partNumber.Pos-1, 0)
		xEnd := min(partNumber.Pos+partNumber.Length, len(input[0])-1)

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
		xStart := max(partNumber.Pos-1, 0)
		xEnd := min(partNumber.Pos+partNumber.Length, len(input[0])-1)

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
	var currentPart *PartNumber = nil

	for y, l := range input {
		for x, c := range l {
			if c >= 48 && c <= 57 {
				// Valid numeric character
				if currentPart == nil {
					// New number, set line and pos
					currentPart = &PartNumber{
						Line: y,
						Pos:  x,
					}
				}
			} else {
				// Invalid numeric character
				if currentPart != nil {
					// End of number, set length and value, add to partNumbers
					currentPart.Length = x - currentPart.Pos
					currentPart.Value, _ = strconv.Atoi(l[currentPart.Pos:x])
					partNumbers = append(partNumbers, *currentPart)
					currentPart = nil
				}
			}
		}
		if currentPart != nil {
			// End of line, set length and value, add to partNumbers
			currentPart.Length = len(l[currentPart.Pos:])
			currentPart.Value, _ = strconv.Atoi(l[currentPart.Pos:])
			partNumbers = append(partNumbers, *currentPart)
			currentPart = nil
		}
	}

	return partNumbers
}
