package day23

import (
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/advent-of-code/2023/util"
)

func Run() {
	input, err := util.ReadLines("./day23/input/input.txt")
	if err != nil {
		fmt.Printf("error reading file: %s", err.Error())
		return
	}

	var start time.Time

	start = time.Now()
	util.PrettyPrint(23, 1, Part1(input), start)

	start = time.Now()
	util.PrettyPrint(23, 2, Part2(input), start)
}

type Location struct {
	ConnectedNeighbours []Point
}

type Point struct {
	Y int
	X int
}

const (
	North = "NORTH"
	East  = "EAST"
	South = "SOUTH"
	West  = "WEST"
)

func Part1(input []string) int {

	locations, start, dest := parseMap(input)

	solutionChannel := make(chan int)
	resultChannel := make(chan int)
	go watchSolution(solutionChannel, resultChannel)

	findPath(start, dest, locations, make([]Point, 0), solutionChannel)
	solutionChannel <- -1

	longestPath := <-resultChannel

	return longestPath
}

func Part2(input []string) int {
	// While running fast on the example input the real puzzle takes some hours to solve :-(

	newInput := make([]string, len(input))
	for i := range input {
		newInput[i] = input[i]
		newInput[i] = strings.Replace(newInput[i], "<", ".", -1)
		newInput[i] = strings.Replace(newInput[i], "^", ".", -1)
		newInput[i] = strings.Replace(newInput[i], ">", ".", -1)
		newInput[i] = strings.Replace(newInput[i], "v", ".", -1)
	}

	locations, start, dest := parseMap(newInput)

	solutionChannel := make(chan int)
	resultChannel := make(chan int)
	go watchSolution(solutionChannel, resultChannel)

	findPath(start, dest, locations, make([]Point, 0), solutionChannel)
	solutionChannel <- -1

	longestPath := <-resultChannel

	return longestPath
}

func parseMap(input []string) (locations map[Point]Location, start, dest Point) {
	locations = make(map[Point]Location)

	for y := range input {
		for x := range input[y] {
			if y == 0 && input[y][x] == '.' {
				start = Point{y, x}
			}
			if y == len(input)-1 && input[y][x] == '.' {
				dest = Point{y, x}
			}

			directions := possibleDirections(Point{y, x}, input)

			var neighbours = make([]Point, 0)
			for _, d := range directions {
				if d == North && y > 0 {
					neighbours = append(neighbours, Point{y - 1, x})
				}
				if d == East && x < len(input[y])-1 {
					neighbours = append(neighbours, Point{y, x + 1})
				}
				if d == South && y < len(input)-1 {
					neighbours = append(neighbours, Point{y + 1, x})
				}
				if d == West && x > 0 {
					neighbours = append(neighbours, Point{y, x - 1})
				}
			}

			locations[Point{y, x}] = Location{
				ConnectedNeighbours: neighbours,
			}
		}
	}

	return locations, start, dest
}

func possibleDirections(point Point, input []string) []string {
	var directions []string

	y := point.Y
	x := point.X

	switch input[y][x] {
	case '.':
		if y > 0 && (input[y-1][x] != '#' && input[y-1][x] != 'v') {
			directions = append(directions, North)
		}
		if x < len(input[y])-1 && (input[y][x+1] != '#' && input[y][x+1] != '<') {
			directions = append(directions, East)
		}
		if y < len(input)-1 && (input[y+1][x] != '#' && input[y+1][x] != '^') {
			directions = append(directions, South)
		}
		if x > 0 && (input[y][x-1] != '#' && input[y][x-1] != '>') {
			directions = append(directions, West)
		}
	case '^':
		if y > 0 && (input[y-1][x] != '#' && input[y-1][x] != 'v') {
			directions = append(directions, North)
		}
	case '>':
		if x < len(input[y])-1 && (input[y][x+1] != '#' && input[y][x+1] != '<') {
			directions = append(directions, East)
		}
	case 'v':
		if y < len(input)-1 && (input[y+1][x] != '#' && input[y+1][x] != '^') {
			directions = append(directions, South)
		}
	case '<':
		if x > 0 && (input[y][x-1] != '#' && input[y][x-1] != '>') {
			directions = append(directions, West)
		}
	}

	return directions
}

func watchSolution(solutions, result chan int) {
	var longestPath int

out:
	for {
		select {
		case steps := <-solutions:
			if steps == -1 {
				// No more solutions
				break out
			}
			if steps > longestPath {
				longestPath = steps
			}
			fmt.Println("Solution found in", steps, "steps. New longest path in", longestPath, "steps.")
		}
	}
	result <- longestPath
}

func findPath(current Point, dest Point, locations map[Point]Location, trail []Point, solutions chan int) {
	if current == dest {
		solutions <- len(trail)
		return
	}

	trail = append(trail, current)

	for _, next := range locations[current].ConnectedNeighbours {
		if validDirection(next, locations, trail) {
			findPath(next, dest, locations, trail, solutions)
		}
	}
}

func validDirection(next Point, locations map[Point]Location, trail []Point) bool {
	_, exists := locations[next]
	visited := slices.Contains(trail, next)

	return exists && !visited
}
