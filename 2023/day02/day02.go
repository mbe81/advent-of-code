package day02

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/advent-of-code/2023/util"
)

func Run() {
	input, err := util.ReadLines("./day02/input/puzzle.txt")
	if err != nil {
		fmt.Printf("Error reading file: %e", err)
		return
	}

	fmt.Println("Running Day 2, Part 1")
	fmt.Println("Result: ", Part1(input))

	fmt.Println("Running Day 2, Part 2")
	fmt.Println("Result: ", Part2(input))
}

type Game struct {
	ID       int
	CubeSets []CubeSet
}

type CubeSet struct {
	RedCubes   int
	GreenCubes int
	BlueCubes  int
}

func Part1(input []string) int {
	var sumIDs int

	for _, l := range input {

		game := parseGame(l)
		var gameImpossible bool

		for _, cubeSet := range game.CubeSets {
			if cubeSet.RedCubes > 12 || cubeSet.GreenCubes > 13 || cubeSet.BlueCubes > 14 {
				gameImpossible = true
			}
		}

		if !gameImpossible {
			sumIDs += game.ID
		}
	}

	return sumIDs
}

func Part2(input []string) int {
	var sumPower int

	for _, l := range input {

		game := parseGame(l)
		var reqRedCubes, reqGreenCubes, reqBlueCubes int

		for _, cubeSet := range game.CubeSets {
			if cubeSet.RedCubes > reqRedCubes {
				reqRedCubes = cubeSet.RedCubes
			}
			if cubeSet.GreenCubes > reqGreenCubes {
				reqGreenCubes = cubeSet.GreenCubes
			}
			if cubeSet.BlueCubes > reqBlueCubes {
				reqBlueCubes = cubeSet.BlueCubes
			}
		}

		sumPower += reqRedCubes * reqGreenCubes * reqBlueCubes
	}

	return sumPower
}

func parseGame(s string) Game {

	var game Game

	lineSplit := strings.Split(s, ":")
	game.ID, _ = strconv.Atoi(strings.Replace(lineSplit[0], "Game ", "", 1))

	sets := strings.Split(lineSplit[1], ";")

	for i := range sets {

		var cubeSet CubeSet

		cubes := strings.Split(sets[i], ",")
		for j := range cubes {
			cubeSplit := strings.Split(cubes[j], " ")

			// Start by 1 because the string starts with a space
			cubeCount, _ := strconv.Atoi(cubeSplit[1])
			cubeColor := cubeSplit[2]

			switch cubeColor {
			case "red":
				cubeSet.RedCubes += cubeCount
			case "green":
				cubeSet.GreenCubes += cubeCount
			case "blue":
				cubeSet.BlueCubes += cubeCount
			}
		}

		game.CubeSets = append(game.CubeSets, cubeSet)
	}

	return game
}
