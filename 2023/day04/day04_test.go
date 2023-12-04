package day04

import (
	"testing"

	"github.com/advent-of-code/2023/util"
)

func TestPart1(t *testing.T) {
	input, err := util.ReadLines("./input/example.txt")
	if err != nil {
		t.Errorf("error reading file: %s", err.Error())
		return
	}

	got := Part1(input)
	want := 13

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestPart2(t *testing.T) {
	input, err := util.ReadLines("./input/example.txt")
	if err != nil {
		t.Errorf("error reading file: %s", err.Error())
		return
	}

	got := Part2(input)
	want := 30

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
