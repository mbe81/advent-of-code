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
	case -1:
		day01.Run()
		day02.Run()
		day03.Run()
		day04.Run()
		day05.Run()
		day07.Run()
	default:
		fmt.Println("Use 'go run main.go -day <day> to run a day")
	}
}
