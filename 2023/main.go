package main

import (
	"flag"
	"fmt"

	"github.com/advent-of-code/2023/day01"
)

func main() {

	var day *int
	day = flag.Int("day", 1, "day to run the puzzle for")
	flag.Parse()

	switch *day {
	case 1:
		day01.Run()
	default:
		fmt.Println("Use 'go run main.go -day <day> to run a day")
	}
}
