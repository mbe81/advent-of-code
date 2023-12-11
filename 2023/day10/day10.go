package day10

import (
	"fmt"
	"slices"
	"strings"
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
	locations, start := parseMap(input)
	totalLength, _ := findPath(start, 0, locations, make([]Point, 0))

	return totalLength / 2
}

func Part2(input []string) int {

	locations, start := parseMap(input)
	_, trail := findPath(start, 0, locations, make([]Point, 0))

	// Create expanded map
	tiles := make([]string, len(input)*2)
	for y := range tiles {
		tiles[y] = strings.Repeat(" ", len(input[0])*2)
	}

	// Expand original pipes
	for _, v := range trail {
		tiles[v.Y*2] = tiles[v.Y*2][:v.X*2] + string(input[v.Y][v.X]) + tiles[v.Y*2][v.X*2+1:]
	}

	// Connect expanded pipes
	for y := 0; y < len(tiles); y = y + 2 {
		for x := 0; x < len(tiles[y]); x = x + 2 {
			var directions []string
			if tiles[y][x] != ' ' {
				directions = possibleDirections(Point{y / 2, x / 2}, input)
			}

			for _, d := range directions {
				if d == North && y > 0 {
					if tiles[y-1][x] == ' ' {
						tiles[y-1] = tiles[y-1][:x] + "|" + tiles[y-1][x+1:]
					}
				}
				if d == East && x < len(tiles[y])-1 {
					if tiles[y][x+1] == ' ' {
						tiles[y] = tiles[y][:x+1] + "-" + tiles[y][x+2:]
					}
				}
				if d == South && y < len(tiles)-1 {
					if tiles[y+1][x] == ' ' {
						tiles[y+1] = tiles[y+1][:x] + "|" + tiles[y+1][x+1:]
					}
				}
				if d == West && x > 0 {
					if tiles[y][x-1] == ' ' {
						tiles[y] = tiles[y][:x-1] + "-" + tiles[y][x:]
					}
				}
			}
		}
	}

	// Check if point is inside loop at Point{y+1, x+1} so we can count '|' as wall
	var counter = 0

	for y := 0; y < len(tiles); y = y + 2 {
		for x := 0; x < len(tiles[y]); x = x + 2 {
			if tiles[y][x] == ' ' && insideLoop(y+1, x+1, tiles) {
				counter++
			}
		}
	}

	return counter
}

func parseMap(input []string) (map[Point]Location, Point) {
	var locations = make(map[Point]Location)
	var start Point

	for y := range input {
		for x := range input[y] {
			if input[y][x] == 'S' {
				start = Point{y, x}
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

	return locations, start
}

func possibleDirections(point Point, input []string) []string {
	var directions []string

	y := point.Y
	x := point.X

	switch input[y][x] {
	case 'S':
		if y > 0 && (input[y-1][x] == '|' || input[y-1][x] == 'F' || input[y-1][x] == '7') {
			directions = append(directions, North)
		}
		if x < len(input[y])-1 && (input[y][x+1] == '-' || input[y][x+1] == '7' || input[y][x+1] == 'J') {
			directions = append(directions, East)
		}
		if y < len(input)-1 && (input[y+1][x] == '|' || input[y+1][x] == 'L' || input[y+1][x] == 'J') {
			directions = append(directions, South)
		}
		if x > 0 && (input[y][x-1] == '-' || input[y][x-1] == 'F' || input[y][x-1] == 'L') {
			directions = append(directions, West)
		}
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

	return directions
}

func findPath(current Point, step int, locations map[Point]Location, trail []Point) (int, []Point) {
	trail = append(trail, current)
	step++

	for _, next := range locations[current].ConnectedNeighbours {
		if validDirection(next, locations, trail) {
			step, trail = findPath(next, step, locations, trail)

			return step, trail
		}
	}

	return step, trail
}

func validDirection(next Point, locations map[Point]Location, trail []Point) bool {
	_, exists := locations[next]
	visited := slices.Contains(trail, next)

	return exists && !visited
}

func insideLoop(y, x int, tiles []string) bool {
	var counter = 0

	for i := 0; i < len(tiles[y])-x; i++ {
		if tiles[y][x+i] == '|' {
			counter++
		}
	}

	return counter%2 == 1
}
