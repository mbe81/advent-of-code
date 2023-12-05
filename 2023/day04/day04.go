package day04

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/advent-of-code/2023/util"
)

func Run() {
	input, err := util.ReadLines("./day04/input/input.txt")
	if err != nil {
		fmt.Printf("error reading file: %s", err.Error())
		return
	}

	var start time.Time

	start = time.Now()
	util.PrettyPrint(4, 1, Part1(input), start)

	start = time.Now()
	util.PrettyPrint(4, 2, Part2(input), start)
}

type Card struct {
	Number           int
	WinningNumbers   []int
	ScratchedNumbers []int
	Copies           int
}

func Part1(input []string) int {
	var totalPoints int

	for _, card := range getCards(input) {
		matchingNumbers := countMatchingNumbers(card)
		if matchingNumbers > 0 {
			points := math.Pow(2, float64(matchingNumbers-1))
			totalPoints += int(points)
		}
	}

	return totalPoints
}

func Part2(input []string) int {
	var totalCopies int

	cards := getCards(input)
	for i := range cards {
		matchingNumbers := countMatchingNumbers(cards[i])
		if matchingNumbers > 0 {
			for j := i + 1; j <= min(i+matchingNumbers, len(cards)-1); j++ {
				cards[j].Copies = cards[j].Copies + cards[i].Copies
			}
		}
		totalCopies += cards[i].Copies
	}

	return totalCopies
}

func getCards(input []string) []Card {
	var cards []Card

	for _, l := range input {
		var card Card

		s1 := strings.Split(l, ":")
		s2 := strings.Split(s1[1], "|")

		cardNumberString := s1[0]
		card.Number, _ = strconv.Atoi(strings.Replace(cardNumberString, "Card ", "", -1))
		card.Copies = 1

		winningNumbersString := s2[0]
		for i := 0; i < len(winningNumbersString)/3; i++ {
			winningNumber, _ := strconv.Atoi(strings.Trim(winningNumbersString[i*3:i*3+3], " "))
			card.WinningNumbers = append(card.WinningNumbers, winningNumber)
		}

		scratchedNumbersString := s2[1]
		for i := 0; i < len(scratchedNumbersString)/3; i++ {
			scratchedNumber, _ := strconv.Atoi(strings.Trim(scratchedNumbersString[i*3:i*3+3], " "))
			card.ScratchedNumbers = append(card.ScratchedNumbers, scratchedNumber)
		}

		cards = append(cards, card)
	}

	return cards
}

func countMatchingNumbers(card Card) int {
	var matchingNumbers int

	for _, winningNumber := range card.WinningNumbers {
		if slices.Contains(card.ScratchedNumbers, winningNumber) {
			matchingNumbers++
		}
	}

	return matchingNumbers
}
