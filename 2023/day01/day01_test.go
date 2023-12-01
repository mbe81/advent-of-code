package day01

import (
	"fmt"
	"testing"

	"github.com/advent-of-code/2023/util"
)

func TestPart1(t *testing.T) {
	input, err := util.ReadLines("./day01/input/example-part1.txt")
	if err != nil {
		fmt.Printf("Error reading file: %e", err)
		return
	}

	got := Part1(input)
	want := 142

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestPart2(t *testing.T) {
	input, err := util.ReadLines("./day01/input/example-part2.txt")
	if err != nil {
		fmt.Printf("Error reading file: %e", err)
		return
	}

	if err != nil {
		fmt.Printf("Error reading file: %e", err)
		return
	}

	got := Part1(input)
	want := 281

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
