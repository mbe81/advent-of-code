package day03

import (
	"os"
	"testing"

	"github.com/advent-of-code/2023/util"
)

func TestPart1(t *testing.T) {
	input, err := util.ReadLines("./input/example.txt")
	path, err := os.Getwd()
	t.Log(path)
	if err != nil {
		t.Errorf("error reading file: %s", err.Error())
		return
	}

	got := Part1(input)
	want := 4361

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
	want := 467835
	t.Log(got, want)

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
