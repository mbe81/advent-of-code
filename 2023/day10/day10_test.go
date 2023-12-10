package day10

import (
	"fmt"
	"testing"

	"github.com/advent-of-code/2023/day10/backup"
	"github.com/advent-of-code/2023/util"
)

func TestPart1(t *testing.T) {
	input, err := util.ReadLines("./input/example-part1.txt")
	if err != nil {
		fmt.Printf("error reading file: %s", err.Error())
		return
	}

	got := backup.Part1(input)
	want := 8
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestPart2(t *testing.T) {
	input, err := util.ReadLines("./input/example.txt")
	if err != nil {
		fmt.Printf("error reading file: %s", err.Error())
		return
	}

	got := backup.Part2(input)
	want := 71503

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
