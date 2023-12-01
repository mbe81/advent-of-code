package day01

import (
	"fmt"

	"github.com/advent-of-code/2023/util"
)

func Run() {
	input, err := util.ReadLines("./day01/input.txt")
	if err != nil {
		fmt.Printf("Error reading file: %e", err)
		return
	}

	fmt.Println("Running Day 1, Part 1")
	fmt.Println("Result: ", Part1(input))

	fmt.Println("Running Day 1, Part 2")
	fmt.Println("Result: ", Part2(input))
}

func Part1(input []string) int {
	var calibrationValue int

	for _, l := range input {
		firstNumber, lastNumber := getFirstAndLastNumber(l)
		calibrationValue += firstNumber*10 + lastNumber
	}

	fmt.Println("Calibration Value: ", calibrationValue)
	return calibrationValue
}

func Part2(inputLines []string) int {
	var calibrationValue int

	for _, l := range inputLines {

		for i := 0; i < len(l); i++ {

			// Only replace the first character of the written number with the real number so eightwothree will convert
			// to 8igh23three instead of 8wo3 and all possible numbers are kept.

			if len(l)-i >= 3 {
				switch l[i : i+3] {
				case "one":
					l = l[:i] + "1" + l[i+1:]
				case "two":
					l = l[:i] + "2" + l[i+1:]
				case "six":
					l = l[:i] + "6" + l[i+1:]
				}
			}

			if len(l)-i >= 4 {
				switch l[i : i+4] {
				case "four":
					l = l[:i] + "4" + l[i+1:]
				case "five":
					l = l[:i] + "5" + l[i+1:]
				case "nine":
					l = l[:i] + "9" + l[i+1:]
				}
			}

			if len(l)-i >= 5 {
				switch l[i : i+5] {
				case "three":
					l = l[:i] + "3" + l[i+1:]
				case "seven":
					l = l[:i] + "7" + l[i+1:]
				case "eight":
					l = l[:i] + "8" + l[i+1:]
				}
			}

		}

		firstNumber, lastNumber := getFirstAndLastNumber(l)
		calibrationValue += firstNumber*10 + lastNumber
	}

	return calibrationValue
}

func getFirstAndLastNumber(s string) (firstNumber, lastNumber int) {
	var setFirst bool
	for _, c := range s {
		if c >= 48 && c <= 57 {
			currentNumber := int(c - 48)
			if !setFirst {
				firstNumber = currentNumber
				setFirst = true
			}
			lastNumber = currentNumber
		}
	}

	return firstNumber, lastNumber
}
