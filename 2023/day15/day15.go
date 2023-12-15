package day15

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/advent-of-code/2023/util"
)

func Run() {
	input, err := util.ReadLines("./day15/input/input.txt")
	if err != nil {
		fmt.Printf("error reading file: %s", err.Error())
		return
	}

	var start time.Time

	start = time.Now()
	util.PrettyPrint(15, 1, Part1(input), start)

	start = time.Now()
	util.PrettyPrint(15, 2, Part2(input), start)
}

func Part1(input []string) int {
	var totalValue int

	ss := strings.Split(input[0], ",")
	for _, s := range ss {
		totalValue += calculateValue(s)
	}

	return totalValue
}

type Lens struct {
	Label       string
	FocalLength int
}

func Part2(input []string) int {
	var totalPower int

	boxes := make([][]Lens, 256)

	lenses := strings.Split(input[0], ",")
	for _, l := range lenses {
		if l[len(l)-1] == '-' {
			label := l[:len(l)-1]
			boxes = removeFromBox(boxes, label)
			continue
		}

		ss := strings.Split(l, "=")
		label := ss[0]
		focalLength, _ := strconv.Atoi(ss[1])
		boxes = addToBox(boxes, label, focalLength)
	}

	for i := range boxes {
		for j := range boxes[i] {
			totalPower += (i + 1) * (j + 1) * boxes[i][j].FocalLength
		}
	}

	return totalPower
}

func calculateValue(s string) int {
	value := 0
	for i := range s {
		value += int(s[i])
		value *= 17
		value = value % 256
	}
	return value
}

func removeFromBox(boxes [][]Lens, label string) [][]Lens {
	var lenses []Lens

	box := calculateValue(label)

	for i := range boxes[box] {
		if boxes[box][i].Label != label {
			lenses = append(lenses, boxes[box][i])
		}
	}
	boxes[box] = lenses

	return boxes
}

func addToBox(boxes [][]Lens, label string, focalLength int) [][]Lens {
	box := calculateValue(label)
	lensFound := false

	for i := range boxes[box] {
		if boxes[box][i].Label == label {
			boxes[box][i].FocalLength = focalLength
			lensFound = true
		}
	}

	if !lensFound {
		boxes[box] = append(boxes[box], Lens{label, focalLength})
	}

	return boxes
}
