package day09

import (
	"fmt"
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

	start = time.Now()
	util.PrettyPrint(9, 2, Part2(input), start)
}

func Part1(input []string) int {
	sensorValues := parseSensorValues(input)
	totalPredictedValue := 0

	for _, sensorRow := range sensorValues {
		differences := calculateDifferences(sensorRow)

		nextIncrease := 0
		for j := len(differences) - 1; j > 0; j-- {
			nextIncrease += differences[j][len(differences[j])-1]
		}

		lastValue := sensorRow[len(sensorRow)-1]
		predictedValue := lastValue + nextIncrease

		totalPredictedValue += predictedValue
	}

	return totalPredictedValue
}

func Part2(input []string) int {
	sensorValues := parseSensorValues(input)
	totalPredictedValue := 0

	for _, sensorRow := range sensorValues {
		differences := calculateDifferences(sensorRow)

		nextDecrease := 0
		for j := len(differences) - 1; j > 0; j-- {
			previousDecrease := nextDecrease
			nextDecrease = differences[j][0] - previousDecrease
		}

		firstValue := sensorRow[0]
		predictedValue := firstValue - nextDecrease

		totalPredictedValue += predictedValue
	}

	return totalPredictedValue
}

func parseSensorValues(input []string) [][]int {
	var sensorValues [][]int

	for _, l := range input {
		var s []int
		sString := strings.Fields(l)
		for i := range sString {
			sensorValue, _ := strconv.Atoi(sString[i])
			s = append(s, sensorValue)
		}
		sensorValues = append(sensorValues, s)
	}

	return sensorValues
}

func calculateDifferences(sensorValues []int) [][]int {

	var differences [][]int
	differences = append(differences, sensorValues)
	i := 0
	for {
		var differenceRow []int
		for j := 0; j < len(differences[i])-1; j++ {
			differenceRow = append(differenceRow, differences[i][j+1]-differences[i][j])
		}
		differences = append(differences, differenceRow)
		i++

		if len(differenceRow) == 1 {
			break
		}
	}

	return differences
}
