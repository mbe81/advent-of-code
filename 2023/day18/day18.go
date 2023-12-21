package day18

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/advent-of-code/2023/util"
	"github.com/advent-of-code/2023/util/mathutil"
)

func Run() {
	input, err := util.ReadLines("./day18/input/example.txt")
	if err != nil {
		fmt.Printf("error reading file: %s", err.Error())
		return
	}

	var start time.Time

	start = time.Now()
	util.PrettyPrint(18, 1, Part1(input), start)

	start = time.Now()
	util.PrettyPrint(18, 2, Part2(input), start)
}

type Step struct {
	Direction string
	Distance  int
	Color     string
}

type Vertex struct {
	Y int
	X int
}

func Part1(input []string) int {
	steps := parseSteps(input, 1)

	var vertexes []Vertex
	var posX, posY int

	vertexes = append(vertexes, Vertex{posY, posX})
	for _, step := range steps {
		switch step.Direction {
		case "R":
			posX += step.Distance
		case "D":
			posY += step.Distance
		case "L":
			posX -= step.Distance
		case "U":
			posY -= step.Distance
		}

		vertexes = append(vertexes, Vertex{posY, posX})
	}

	return calculateArea(vertexes)
}

func Part2(input []string) int {
	steps := parseSteps(input, 2)

	var vertexes []Vertex
	var posX, posY int

	vertexes = append(vertexes, Vertex{posY, posX})
	for _, step := range steps {
		switch step.Direction {
		case "R":
			posX += step.Distance
		case "D":
			posY += step.Distance
		case "L":
			posX -= step.Distance
		case "U":
			posY -= step.Distance
		}

		vertexes = append(vertexes, Vertex{posY, posX})
	}

	return calculateArea(vertexes)
}

func parseSteps(input []string, part int) []Step {
	var steps []Step

	for _, l := range input {
		ss := strings.Split(l, " ")

		var step Step
		if part == 1 {
			step.Direction = ss[0]
			step.Distance, _ = strconv.Atoi(ss[1])
		} else if part == 2 {
			switch ss[2][7] {
			case '0':
				step.Direction = "R"
			case '1':
				step.Direction = "D"
			case '2':
				step.Direction = "L"
			case '3':
				step.Direction = "U"

			}
			distance, _ := strconv.ParseInt(ss[2][2:7], 16, 64)
			step.Distance = int(distance)

		}

		steps = append(steps, step)
	}

	return steps
}

func calculateArea(v []Vertex) int {
	boundary := 0
	area := 0
	j := len(v) - 1
	for i := 0; i < len(v); i++ {
		boundary += mathutil.Abs(v[j].X-v[i].X) + mathutil.Abs(v[j].Y-v[i].Y)
		area += v[j].X*v[i].Y - v[i].X*v[j].Y
		j = i
	}
	area = mathutil.Abs(area)
	return boundary/2 + area/2 + 1
}
