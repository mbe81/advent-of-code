package day02

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/advent-of-code/2023/util"
)

func Run() {
	input, err := util.ReadLines("./day02/input/puzzle.txt")
	if err != nil {
		fmt.Printf("error reading file: %s", err.Error())
		return
	}

	util.PrettyPrint(2, 1, Part1(input), time.Now())
	util.PrettyPrint(2, 2, Part2(input), time.Now())
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

		game := parseGameString(l)
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

		game := parseGameString(l)
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

func parseGameString(s string) Game {

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
