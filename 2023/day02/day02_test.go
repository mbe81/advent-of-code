package day02

import (
	"fmt"
	"testing"

	"github.com/advent-of-code/2023/util"
)

func TestPart1(t *testing.T) {
	input, err := util.ReadLines("./day02/input/example.txt")
	if err != nil {
		fmt.Printf("error reading file: %s", err.Error())
		return
	}

	got := Part1(input)
	want := 8

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestPart2(t *testing.T) {
	input, err := util.ReadLines("./day02/input/example.txt")
	if err != nil {
		fmt.Printf("error reading file: %s", err.Error())
		return
	}

	got := Part1(input)
	want := 2286

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
