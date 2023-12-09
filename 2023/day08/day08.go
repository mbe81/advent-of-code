package day08

import (
	"fmt"
	"slices"
	"time"

	"github.com/advent-of-code/2023/util"
	"github.com/advent-of-code/2023/util/math"
)

func Run() {
	input, err := util.ReadLines("./day08/input/input.txt")
	if err != nil {
		fmt.Printf("error reading file: %s", err.Error())
		return
	}

	var start time.Time

	start = time.Now()
	util.PrettyPrint(8, 1, Part1(input), start)

	start = time.Now()
	util.PrettyPrint(8, 2, Part2(input), start)
}

type Node struct {
	Left  string
	Right string
}

func Part1(input []string) int {

	steps, nodes := parseMap(input)

	currentNode := "AAA"
	stepCounter := 0

out1:
	for {
		for i := range steps {
			stepCounter++
			if steps[i] == 'L' {
				currentNode = nodes[currentNode].Left
			} else {
				currentNode = nodes[currentNode].Right
			}

			if currentNode == "ZZZ" {
				break out1
			}
		}
	}

	return stepCounter
}

type Ghost struct {
	CurrentNode string
	PassedNodes []string
	StepsToLoop int
}

func Part2(input []string) int {

	steps, nodes := parseMap(input)

	var ghosts []Ghost
	for key, _ := range nodes {
		if key[2] == 'A' {
			ghost := Ghost{
				CurrentNode: key,
			}
			ghosts = append(ghosts, ghost)
		}
	}

	for i, _ := range ghosts {

		var newNode string
		var stepCounter int

	out2:
		for {
			for j := range steps {
				if steps[j] == 'L' {
					newNode = nodes[ghosts[i].CurrentNode].Left
				} else {
					newNode = nodes[ghosts[i].CurrentNode].Right
				}

				if ghosts[i].CurrentNode[2] == 'Z' && slices.Contains(ghosts[i].PassedNodes, newNode) {
					// Loop found
					ghosts[i].StepsToLoop = stepCounter
					break out2
				}

				ghosts[i].CurrentNode = newNode
				ghosts[i].PassedNodes = append(ghosts[i].PassedNodes, newNode)
				stepCounter++
			}
		}

	}

	var numbers []int
	for _, ghost := range ghosts {
		numbers = append(numbers, ghost.StepsToLoop)
	}

	return math.LCM(numbers[0], numbers[1], numbers[2:]...)
}

func parseMap(input []string) (steps string, nodes map[string]Node) {

	steps = input[0]
	nodes = make(map[string]Node)

	for _, l := range input[2:] {
		key := l[:3]
		nodes[key] = Node{
			Left:  l[7:10],
			Right: l[12:15],
		}
	}

	return steps, nodes
}
