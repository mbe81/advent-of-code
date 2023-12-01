package day01

import (
	"fmt"
	"strings"

	"github.com/advent-of-code/2023/util"
)

func Run() {
	input, err := util.ReadLines("./day01/input/puzzle.txt")
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

	return calibrationValue
}

func Part2(inputLines []string) int {
	var calibrationValue int

	for _, l := range inputLines {

		l = strings.Replace(l, "one", "one1one", -1)
		l = strings.Replace(l, "two", "two2two", -1)
		l = strings.Replace(l, "three", "three3three", -1)
		l = strings.Replace(l, "four", "four4four", -1)
		l = strings.Replace(l, "five", "five5five", -1)
		l = strings.Replace(l, "six", "six6six", -1)
		l = strings.Replace(l, "seven", "seven7seven", -1)
		l = strings.Replace(l, "eight", "eight8eight", -1)
		l = strings.Replace(l, "nine", "nine9nine", -1)

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
