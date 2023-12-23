package main

import (
	"flag"
	"fmt"

	"github.com/advent-of-code/2023/day01"
	"github.com/advent-of-code/2023/day02"
	"github.com/advent-of-code/2023/day03"
	"github.com/advent-of-code/2023/day04"
	"github.com/advent-of-code/2023/day05"
	"github.com/advent-of-code/2023/day06"
	"github.com/advent-of-code/2023/day07"
	"github.com/advent-of-code/2023/day08"
	"github.com/advent-of-code/2023/day09"
	"github.com/advent-of-code/2023/day10"
	"github.com/advent-of-code/2023/day11"
	"github.com/advent-of-code/2023/day12"
	"github.com/advent-of-code/2023/day13"
	"github.com/advent-of-code/2023/day14"
	"github.com/advent-of-code/2023/day15"
	"github.com/advent-of-code/2023/day16"
	"github.com/advent-of-code/2023/day17"
	"github.com/advent-of-code/2023/day18"
	"github.com/advent-of-code/2023/day19"
	"github.com/advent-of-code/2023/day22"
	"github.com/advent-of-code/2023/day23"
)

func main() {

	var day *int
	day = flag.Int("day", 0, "day to run the puzzle for")
	flag.Parse()

	switch *day {
	case 1:
		day01.Run()
	case 2:
		day02.Run()
	case 3:
		day03.Run()
	case 4:
		day04.Run()
	case 5:
		day05.Run()
	case 6:
		day06.Run()
	case 7:
		day07.Run()
	case 8:
		day08.Run()
	case 9:
		day09.Run()
	case 10:
		day10.Run()
	case 11:
		day11.Run()
	case 12:
		day12.Run()
	case 13:
		day13.Run()
	case 14:
		day14.Run()
	case 15:
		day15.Run()
	case 16:
		day16.Run()
	case 17:
		day17.Run()
	case 18:
		day18.Run()
	case 19:
		day19.Run()
	case 22:
		day22.Run()
	case 23:
		day23.Run()
	case -1:
		day01.Run()
		day02.Run()
		day03.Run()
		day04.Run()
		day05.Run()
		day07.Run()
		day08.Run()
		day09.Run()
		day10.Run()
		day11.Run()
		day12.Run()
		day13.Run()
		day14.Run()
		day15.Run()
		day16.Run()
		day17.Run()
		day18.Run()
		day19.Run()
		day22.Run()
		day23.Run()
	default:
		fmt.Println("Use 'go run main.go -day <day>' to run a day")
	}
}
