package day03

import (
	"fmt"
	"testing"

	"github.com/advent-of-code/2023/util"
)

func TestPart1(t *testing.T) {
	input, err := util.ReadLines("./day03/input/example.txt")
	if err != nil {
		fmt.Printf("Error reading file: %e", err)
		return
	}

	got := Part1(input)
	want := 4361

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestPart2(t *testing.T) {
	input, err := util.ReadLines("./day03/input/example.txt")
	if err != nil {
		fmt.Printf("Error reading file: %e", err)
		return
	}

	if err != nil {
		fmt.Printf("Error reading file: %e", err)
		return
	}

	got := Part1(input)
	want := 467835

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
