package main

import (
	"fmt"
	"github.com/mbe81/advent-of-code/util"
	"sort"
)

import (
	"log"
	"strconv"
)

type elf struct {
	calories int
}

var elves []elf

func main() {

	// Read the file and sum up calories per elf
	lines := util.ReadFile("./2022/day01/input.txt")
	calories := 0
	for _, line := range lines {
		if line == "" {
			elves = append(elves, elf{calories: calories})
			calories = 0
		} else {
			newCalories, err := strconv.Atoi(line)
			if err != nil {
				log.Fatal("Invalid input")
			}
			calories = calories + newCalories
		}
	}

	// Print the number of calories carried by the elf with the most calories
	for _, elf := range elves {
		if elf.calories > calories {
			calories = elf.calories
		}
	}
	fmt.Println("The elf with the most calories is carrying", calories, "calories")

	// Print the total number of calories carries by the three elves carrying the most
	sort.Slice(elves, func(i, j int) bool {
		return elves[i].calories > elves[j].calories
	})
	sumCalories := 0
	for i := 0; i < 3; i++ {
		sumCalories = sumCalories + elves[i].calories
	}
	fmt.Println("The three elves with the most calories are carrying", sumCalories, "calories in total")
}
