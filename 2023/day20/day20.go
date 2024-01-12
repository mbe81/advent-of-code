package day20

import (
	"fmt"
	"time"

	"github.com/advent-of-code/2023/util"
)

func Run() {
	input, err := util.ReadLines("./day20/input/input.txt")
	if err != nil {
		fmt.Printf("error reading file: %s", err.Error())
		return
	}

	var start time.Time

	start = time.Now()
	util.PrettyPrint(20, 1, Part1(input), start)

	start = time.Now()
	util.PrettyPrint(20, 2, Part2(input), start)
}

func Part1(input []string) int {

	return 0
}

func Part2(input []string) int {

	return 0
}