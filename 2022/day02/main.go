package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type choice int

const (
	ROCK     choice = 1
	PAPER    choice = 2
	SCISSORS choice = 3
)

type result int

const (
	WIN  result = 6
	LOSE result = 0
	DRAW result = 3
)

func main() {

	// Read the file and sum up the scores per game
	file, err := os.Open("./2022/day02/input.txt")
	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(file)

	// Calculate the scores per strategy
	totalScorePartOne := 0
	totalScorePartTwo := 0
	letters := []string{}
	for scanner.Scan() {
		letters = strings.Split(scanner.Text(), " ")
		totalScorePartOne = totalScorePartOne + calculateScorePartOne(letters[0], letters[1])
		totalScorePartTwo = totalScorePartTwo + calculateScorePartTwo(letters[0], letters[1])
	}
	_ = file.Close()

	fmt.Println("Total score according to strategy guide in Part One", totalScorePartOne)
	fmt.Println("Total score according to strategy guide in Part Two", totalScorePartTwo)
}

func calculateScorePartOne(hisLetter string, myLetter string) int {

	var hisChoice choice
	switch hisLetter {
	case "A":
		hisChoice = ROCK
	case "B":
		hisChoice = PAPER
	case "C":
		hisChoice = SCISSORS
	}

	var myChoice choice
	switch myLetter {
	case "X":
		myChoice = ROCK
	case "Y":
		myChoice = ROCK
	case "Z":
		myChoice = SCISSORS
	}

	// Draw
	if myChoice == hisChoice {
		return int(DRAW) + int(myChoice)
	}

	// Win
	if myChoice == PAPER && hisChoice == ROCK {
		return int(WIN) + int(myChoice)
	}
	if myChoice == ROCK && hisChoice == SCISSORS {
		return int(WIN) + int(myChoice)
	}
	if myChoice == SCISSORS && hisChoice == PAPER {
		return int(WIN) + int(myChoice)
	}

	// Lose
	if myChoice == ROCK && hisChoice == PAPER {
		return int(LOSE) + int(myChoice)
	}
	if myChoice == SCISSORS && hisChoice == ROCK {
		return int(LOSE) + int(myChoice)
	}
	if myChoice == PAPER && hisChoice == SCISSORS {
		return int(LOSE) + int(myChoice)
	}

	return 0

}

func calculateScorePartTwo(hisLetter string, myLetter string) int {

	var hisChoice choice
	switch hisLetter {
	case "A":
		hisChoice = ROCK
	case "B":
		hisChoice = PAPER
	case "C":
		hisChoice = SCISSORS
	}

	var myResult result
	switch myLetter {
	case "X":
		myResult = LOSE
	case "Y":
		myResult = DRAW
	case "Z":
		myResult = WIN
	}

	// Draw
	if myResult == DRAW {
		return int(DRAW) + int(hisChoice)
	}

	// Win
	if myResult == WIN && hisChoice == ROCK {
		return int(WIN) + int(PAPER)
	}
	if myResult == WIN && hisChoice == SCISSORS {
		return int(WIN) + int(ROCK)
	}
	if myResult == WIN && hisChoice == PAPER {
		return int(WIN) + int(SCISSORS)
	}

	// Lose
	if myResult == LOSE && hisChoice == ROCK {
		return int(LOSE) + int(SCISSORS)
	}
	if myResult == LOSE && hisChoice == SCISSORS {
		return int(LOSE) + int(PAPER)
	}
	if myResult == LOSE && hisChoice == PAPER {
		return int(LOSE) + int(ROCK)
	}

	return 0
}
