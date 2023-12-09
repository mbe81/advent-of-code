package day09

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/advent-of-code/2023/util"
)

func Run() {
	input, err := util.ReadLines("./day09/input/example.txt")
	if err != nil {
		fmt.Printf("error reading file: %s", err.Error())
		return
	}

	var start time.Time

	start = time.Now()
	util.PrettyPrint(9, 1, Part1(input), start)

	//start = time.Now()
	//util.PrettyPrint(9, 2, Part2(input), start)
}

func Part1(input []string) int {
	var totalPrediction int

	for _, sensorValues := range parseSensorValues(input, false) {
		totalPrediction += predictNextValue(sensorValues)
	}

	return totalPrediction
}

func Part2(input []string) int {
	var totalPrediction int

	for _, sensorValues := range parseSensorValues(input, true) {
		totalPrediction += predictNextValue(sensorValues)
	}

	return totalPrediction
}

func parseSensorValues(input []string, reverse bool) [][]int {
	var sensorValues [][]int

	for _, l := range input {
		var s []int

		for _, val := range strings.Fields(l) {
			sensorValue, _ := strconv.Atoi(val)
			s = append(s, sensorValue)
		}

		if reverse {
			slices.Reverse(s)
		}

		sensorValues = append(sensorValues, s)
	}

	return sensorValues
}

func predictNextValue(sensorValues []int) int {
	var differences [][]int

	currentDiffs := sensorValues
	for {
		var nextDiffs []int
		for j := 0; j < len(currentDiffs)-1; j++ {
			nextDiffs = append(nextDiffs, currentDiffs[j+1]-currentDiffs[j])
		}
		differences = append(differences, nextDiffs)
		currentDiffs = nextDiffs
		if len(nextDiffs) == 1 {
			break
		}
	}

	predictedValue := sensorValues[len(sensorValues)-1] // Last known value
	for j := len(differences) - 1; j >= 0; j-- {
		predictedValue += differences[j][len(differences[j])-1] // Add all increases
	}

	return predictedValue
}
