package day07

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/advent-of-code/2023/util"
)

func Run() {
	input, err := util.ReadLines("./day07/input/input.txt")
	if err != nil {
		fmt.Printf("error reading file: %s", err.Error())
		return
	}

	var start time.Time

	start = time.Now()
	util.PrettyPrint(7, 1, Part1(input), start)

	start = time.Now()
	util.PrettyPrint(7, 2, Part2(input), start)
}

type Hand struct {
	CardList       string
	CardQuantities []CardQuantity
	Cards          []CardType
	HandResult     HandResult
	Bid            int
}

type CardQuantity struct {
	Type  CardType
	Count int
}

type CardType int

const (
	CardTypeAce   CardType = 14
	CardTypeKing           = 13
	CardTypeQueen          = 12
	CardTypeJack           = 11
	CardTypeTen            = 10
	CardTypeNine           = 9
	CardTypeEight          = 8
	CardTypeSeven          = 7
	CardTypeSix            = 6
	CardTypeFive           = 5
	CardTypeFour           = 4
	CardTypeThree          = 3
	CardTypeTwo            = 2
	CardTypeJoker          = 1
)

type HandResult int

const (
	HandResultFiveOfAKind  HandResult = 7
	HandResultFourOfAKind             = 6
	HandResultFullHouse               = 5
	HandResultThreeOfAKind            = 4
	HandResultTwoPair                 = 3
	HandResultOnePair                 = 2
	HandResultHighCard                = 1
)

func Part1(input []string) int {
	var totalWinnings int

	hands := parseHands(input, 1)

	for i, _ := range hands {
		switch hands[i].CardQuantities[0].Count {
		case 5:
			hands[i].HandResult = HandResultFiveOfAKind
		case 4:
			hands[i].HandResult = HandResultFourOfAKind
		case 3:
			if hands[i].CardQuantities[1].Count == 2 {
				hands[i].HandResult = HandResultFullHouse
			} else {
				hands[i].HandResult = HandResultThreeOfAKind
			}
		case 2:
			if hands[i].CardQuantities[1].Count == 2 {
				hands[i].HandResult = HandResultTwoPair
			} else {
				hands[i].HandResult = HandResultOnePair
			}
		case 1:
			hands[i].HandResult = HandResultHighCard
		}
	}

	// Sort hands from low to high
	sort.Slice(hands, func(i, j int) bool {
		s1 := int(hands[i].HandResult)*10000000000 +
			int(hands[i].Cards[0])*100000000 +
			int(hands[i].Cards[1])*1000000 +
			int(hands[i].Cards[2])*10000 +
			int(hands[i].Cards[3])*100 +
			int(hands[i].Cards[4])

		s2 := int(hands[j].HandResult)*10000000000 +
			int(hands[j].Cards[0])*100000000 +
			int(hands[j].Cards[1])*1000000 +
			int(hands[j].Cards[2])*10000 +
			int(hands[j].Cards[3])*100 +
			int(hands[j].Cards[4])

		return s1 < s2
	})

	// Calculate score
	for i, h := range hands {
		totalWinnings += (i + 1) * h.Bid
	}

	return totalWinnings
}

func Part2(input []string) int {
	var totalWinnings int

	hands := parseHands(input, 2)

	for i, _ := range hands {

		var numberOfJokers int

		var quantitiesWithoutJoker []CardQuantity
		quantitiesWithoutJoker = hands[i].CardQuantities
		for j, cq := range hands[i].CardQuantities {
			if cq.Type == CardTypeJoker {
				numberOfJokers = cq.Count
				quantitiesWithoutJoker = append(hands[i].CardQuantities[:j], hands[i].CardQuantities[j+1:]...)
				break
			}

		}

		switch numberOfJokers {
		case 5:
			hands[i].HandResult = HandResultFiveOfAKind

		case 4:
			hands[i].HandResult = HandResultFiveOfAKind

		case 3:
			switch quantitiesWithoutJoker[0].Count {
			case 2:
				hands[i].HandResult = HandResultFiveOfAKind
			case 1:
				hands[i].HandResult = HandResultFourOfAKind
			}
		case 2:
			switch quantitiesWithoutJoker[0].Count {
			case 3:
				hands[i].HandResult = HandResultFiveOfAKind
			case 2:
				hands[i].HandResult = HandResultFourOfAKind
			case 1:
				if quantitiesWithoutJoker[1].Count == 2 {
					hands[i].HandResult = HandResultFullHouse
				} else {
					hands[i].HandResult = HandResultThreeOfAKind
				}
			}
		case 1:
			switch quantitiesWithoutJoker[0].Count {
			case 4:
				hands[i].HandResult = HandResultFiveOfAKind
			case 3:
				hands[i].HandResult = HandResultFourOfAKind
			case 2:
				if quantitiesWithoutJoker[1].Count == 2 {
					hands[i].HandResult = HandResultFullHouse
				} else {
					hands[i].HandResult = HandResultThreeOfAKind
				}
			case 1:
				if quantitiesWithoutJoker[1].Count == 2 {
					hands[i].HandResult = HandResultTwoPair
				} else {
					hands[i].HandResult = HandResultOnePair
				}
			}
		case 0:
			switch quantitiesWithoutJoker[0].Count {
			case 5:
				hands[i].HandResult = HandResultFiveOfAKind
			case 4:
				hands[i].HandResult = HandResultFourOfAKind
			case 3:
				if hands[i].CardQuantities[1].Count == 2 {
					hands[i].HandResult = HandResultFullHouse
				} else {
					hands[i].HandResult = HandResultThreeOfAKind
				}
			case 2:
				if quantitiesWithoutJoker[1].Count == 2 {
					hands[i].HandResult = HandResultTwoPair
				} else {
					hands[i].HandResult = HandResultOnePair
				}
			case 1:
				hands[i].HandResult = HandResultHighCard
			}
		}
	}

	// Sort hands from low to high
	sort.Slice(hands, func(i, j int) bool {
		s1 := int(hands[i].HandResult)*10000000000 +
			int(hands[i].Cards[0])*100000000 +
			int(hands[i].Cards[1])*1000000 +
			int(hands[i].Cards[2])*10000 +
			int(hands[i].Cards[3])*100 +
			int(hands[i].Cards[4])

		s2 := int(hands[j].HandResult)*10000000000 +
			int(hands[j].Cards[0])*100000000 +
			int(hands[j].Cards[1])*1000000 +
			int(hands[j].Cards[2])*10000 +
			int(hands[j].Cards[3])*100 +
			int(hands[j].Cards[4])

		return s1 < s2
	})

	// Calculate score
	for i, h := range hands {
		totalWinnings += (i + 1) * h.Bid
	}

	return totalWinnings
}

func parseHands(input []string, part int) []Hand {
	var hands []Hand
	for _, l := range input {
		var h Hand

		s := strings.Fields(l)

		// Set cards
		h.CardList = s[0]
		cardList := s[0]
		for i := 0; i < len(cardList); i++ {
			var cardType CardType
			switch cardList[i : i+1] {
			case "A":
				cardType = CardTypeAce
			case "K":
				cardType = CardTypeKing
			case "Q":
				cardType = CardTypeQueen
			case "J":
				switch part {
				case 1:
					cardType = CardTypeJack
				case 2:
					cardType = CardTypeJoker
				}
			case "T":
				cardType = CardTypeTen
			case "9":
				cardType = CardTypeNine
			case "8":
				cardType = CardTypeEight
			case "7":
				cardType = CardTypeSeven
			case "6":
				cardType = CardTypeSix
			case "5":
				cardType = CardTypeFive
			case "4":
				cardType = CardTypeFour
			case "3":
				cardType = CardTypeThree
			case "2":
				cardType = CardTypeTwo
			}
			h.Cards = append(h.Cards, cardType)
		}

		// Set card quantities
		var cardQuantityMap = make(map[CardType]int)
		for _, cardType := range h.Cards {
			cardQuantityMap[cardType]++
		}
		for t, c := range cardQuantityMap {
			h.CardQuantities = append(h.CardQuantities, CardQuantity{t, c})
		}
		sort.Slice(h.CardQuantities, func(i, j int) bool {
			return h.CardQuantities[i].Count > h.CardQuantities[j].Count
		})

		// Set bid
		h.Bid, _ = strconv.Atoi(s[1])

		hands = append(hands, h)
	}
	return hands
}
