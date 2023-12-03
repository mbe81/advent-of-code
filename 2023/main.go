package main

import (
	"flag"
	"fmt"

	"github.com/advent-of-code/2023/day01"
	"github.com/advent-of-code/2023/day02"
	"github.com/advent-of-code/2023/day03"
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
	default:
		fmt.Println("Use 'go run main.go -day <day> to run a day")
	}
}
