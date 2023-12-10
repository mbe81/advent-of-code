package day10

import (
	"fmt"
	"slices"
	"time"

	"github.com/advent-of-code/2023/util"
)

func Run() {
	input, err := util.ReadLines("./day10/input/input.txt")
	if err != nil {
		fmt.Printf("error reading file: %s", err.Error())
		return
	}

	var start time.Time

	start = time.Now()
	util.PrettyPrint(10, 1, Part1(input), start)

	start = time.Now()
	util.PrettyPrint(10, 2, Part2(input), start)
}

type Location struct {
	Visited             bool
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
	var trail = make([]Point, 0)

	locations, start := parseMap(input)
	totalLength := findPath(start, 0, locations, trail)

	return totalLength / 2
}

func parseMap(input []string) (map[Point]Location, Point) {
	var locations = make(map[Point]Location)
	var start Point

	for y := range input {
		for x := range input[y] {
			var directions = make([]string, 0)
			switch input[y][x] {
			case 'S':
				start = Point{y, x}
				directions = append(directions, East, South)
			case '|':
				directions = append(directions, North, South)
			case '-':
				directions = append(directions, East, West)
			case 'L':
				directions = append(directions, North, East)
			case 'J':
				directions = append(directions, North, West)
			case '7':
				directions = append(directions, South, West)
			case 'F':
				directions = append(directions, East, South)
			}

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

	return locations, start
}

func findPath(current Point, step int, locations map[Point]Location, trail []Point) int {
	trail = append(trail, current)
	step++

	locations[current] = Location{
		Visited:             true,
		ConnectedNeighbours: locations[current].ConnectedNeighbours,
	}

	for _, next := range locations[current].ConnectedNeighbours {
		if validDirection(next, locations, trail) {
			step = findPath(next, step, locations, trail)
			return step
		}
	}
	return step
}

func validDirection(next Point, locations map[Point]Location, trail []Point) bool {
	_, exists := locations[next]
	visited := slices.Contains(trail, next)

	return exists && !visited
}

func Part2(input []string) int {
	return 0
}
