package day17

import (
	"fmt"
	"strconv"
	"time"

	"github.com/advent-of-code/2023/util"
)

func Run() {
	input, err := util.ReadLines("./day17/input/example.txt")
	if err != nil {
		fmt.Printf("error reading file: %s", err.Error())
		return
	}

	var start time.Time

	start = time.Now()
	util.PrettyPrint(17, 1, Part1(input), start)

	start = time.Now()
	util.PrettyPrint(17, 2, Part2(input), start)
}

type Tile struct {
	energized     bool
	seenFromNorth bool
	seenFromEast  bool
	seenFromSouth bool
	seenFromWest  bool
}

//var i int

type Direction int

const (
	North Direction = 1
	East            = 2
	South           = 4
	West            = 8
)

type Distance int

func Part1(input []string) int {
	graph, seen := parseGraph(input)

	var distance, minDistance Distance
	var valid bool

	distance, valid = findPath(graph, seen, 0, 0, len(graph), len(graph[0]), 0, East, 1)
	if valid && (distance < minDistance || minDistance == 0) {
		minDistance = distance
	}
	distance, valid = findPath(graph, seen, 0, 0, len(graph), len(graph[0]), 0, South, 1)
	if valid && (distance < minDistance || minDistance == 0) {
		minDistance = distance
	}

	fmt.Println(distance)
	return 0
}

func Part2(input []string) int {
	return 0
}

func parseGraph(input []string) ([][]Distance, [][]Direction) {
	graph := make([][]Distance, len(input))
	seen := make([][]Direction, len(input))

	for y := range input {
		graph[y] = make([]Distance, len(input[y]))
		seen[y] = make([]Direction, len(input[y]))
		for x := range input[y] {
			d, _ := strconv.Atoi(input[y][x:x])
			graph[y][x] = Distance(d)
		}
	}

	return graph, seen
}

func findPath(graph [][]Distance, seen [][]Direction, posY, posX, destY, destX int, distance Distance, dir Direction, dirSteps int) (Distance, bool) {
	var newDistance, minDistance Distance
	var valid bool

	if posY < 0 || posX < 0 || posY > len(graph)-1 || posX > len(graph[0])-1 {
		return 0, false
	}
	oppositeDir := opposite(dir)
	if seen[posY][posX]%oppositeDir*2 >= oppositeDir {
		//fmt.Println(seen)
		//fmt.Println("Loop detected")
		return 0, false
	}
	seen[posY][posX] += dir
	//fmt.Println(seen)
	//i++
	//if i >= 6 {
	//	os.Exit(0)
	//}

	distance += graph[posY][posX]
	if posY == destY && destY == destX {
		valid = true
		return distance, valid
	}

	var newY, newX int
	var newDir Direction

	if dirSteps <= 2 {
		newY, newX, newDir = moveForward(posY, posX, dir)
		newDistance, valid = findPath(graph, seen, newY, newX, destY, destX, distance, newDir, dirSteps+1)
		if valid && (newDistance < minDistance || minDistance == 0) {
			minDistance = newDistance
		}
	}

	newY, newX, newDir = moveLeft(posY, posX, dir)
	newDistance, valid = findPath(graph, seen, newY, newX, destY, destX, distance, newDir, 1)
	if valid && (newDistance < minDistance || minDistance == 0) {
		minDistance = newDistance
	}

	newY, newX, newDir = moveRight(posY, posX, dir)
	newDistance, valid = findPath(graph, seen, newY, newX, destY, destX, distance, newDir, 1)
	if valid && (newDistance < minDistance || minDistance == 0) {
		minDistance = newDistance
	}

	if minDistance > 0 {
		fmt.Println(distance, minDistance)
		return distance + minDistance, true
	}

	return minDistance, false
}

func moveForward(posY, posX int, dir Direction) (int, int, Direction) {
	switch dir {
	case North:
		posY = posY - 1
	case East:
		posX = posX + 1
	case South:
		posY = posY + 1
	case West:
		posX = posX - 1
	}
	return posY, posX, dir
}

func moveLeft(posY, posX int, dir Direction) (int, int, Direction) {
	switch dir {
	case North:
		dir = West
		posX = posX - 1
	case East:
		dir = North
		posY = posY - 1
	case South:
		dir = East
		posX = posX + 1
	case West:
		dir = South
		posY = posY + 1
	}
	return posY, posX, dir
}

func moveRight(posY, posX int, dir Direction) (int, int, Direction) {
	switch dir {
	case North:
		dir = East
		posX = posX + 1
	case East:
		dir = South
		posY = posY + 1
	case South:
		dir = West
		posX = posX - 1
	case West:
		dir = North
		posY = posY - 1
	}
	return posY, posX, dir
}

func opposite(dir Direction) Direction {
	switch dir {
	case North:
		return South
	case East:
		return West
	case South:
		return North
	case West:
		return East
	}
	return 0
}

func seenFrom(seen [][]Direction, posY, posX int, dir Direction) ([][]Direction, bool) {
	if seen[posY][posX]%dir*2 >= dir {
		return seen, true
	}
	seen[posY][posX] += dir
	return seen, false
}
