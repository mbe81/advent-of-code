package day07

import (
	"fmt"
	"testing"

	"github.com/advent-of-code/2023/util"
)

func TestPart1(t *testing.T) {
	input, err := util.ReadLines("./input/example.txt")
	if err != nil {
		fmt.Printf("error reading file: %s", err.Error())
		return
	}

	got := Part1(input)
	want := 6640
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	input, err := util.ReadLines("./input/example.txt")
	if err != nil {
		fmt.Printf("error reading file: %s", err.Error())
		return
	}

	got := Part2(input)
	want := 5905

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
