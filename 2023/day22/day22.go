package day22

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/advent-of-code/2023/util"
)

func Run() {
	input, err := util.ReadLines("./day22/input/input.txt")
	if err != nil {
		fmt.Printf("error reading file: %s", err.Error())
		return
	}

	var start time.Time

	start = time.Now()
	util.PrettyPrint(22, 1, Part1(input), start)

	start = time.Now()
	util.PrettyPrint(22, 2, Part2(input), start)
}

type Coordinate struct {
	X int
	Y int
	Z int
}

type Brick struct {
	ID    int
	Start Coordinate
	End   Coordinate
}

func Part1(input []string) int {

	bricks := parseBricks(input)
	simulateBrickFall(bricks, 0)

	var bricksToRemove int
	for i := 0; i < len(bricks); i++ {
		newBricks := make([]Brick, len(bricks))
		copy(newBricks, bricks)

		fallenBricks := simulateBrickFall(newBricks, bricks[i].ID)
		if fallenBricks == 0 {
			bricksToRemove++
		}
	}

	return bricksToRemove
}

func Part2(input []string) int {

	bricks := parseBricks(input)
	simulateBrickFall(bricks, 0)

	var totalFallen int
	for i := 0; i < len(bricks); i++ {
		newBricks := make([]Brick, len(bricks))
		copy(newBricks, bricks)

		fallenBricks := simulateBrickFall(newBricks, bricks[i].ID)
		totalFallen += fallenBricks
	}

	return totalFallen
}

func simulateBrickFall(bricks []Brick, removedBrickID int) int {
	stackPile := createStackPile(10, 10, 300)

	fallenBricks := 0
	for i := range bricks {
		if bricks[i].ID != removedBrickID {

			// Calculate how far the brick can fall
			var zFallen = 0
		out:
			for z := bricks[i].Start.Z; z > 0; z-- {

				xMin := min(bricks[i].Start.X, bricks[i].End.X)
				xMax := max(bricks[i].Start.X, bricks[i].End.X)
				for x := xMin; x <= xMax; x++ {

					yMin := min(bricks[i].Start.Y, bricks[i].End.Y)
					yMax := max(bricks[i].Start.Y, bricks[i].End.Y)
					for y := yMin; y <= yMax; y++ {

						if stackPile[x][y][z-1] != 0 && stackPile[x][y][z-1] != removedBrickID {
							break out
						}
					}
				}

				zFallen++
			}

			// Move the brick
			if zFallen > 0 {
				bricks[i].Start.Z = bricks[i].Start.Z - zFallen
				bricks[i].End.Z = bricks[i].End.Z - zFallen
				fallenBricks++
			}

			// Add the brick to the stack pile
			for x := bricks[i].Start.X; x <= bricks[i].End.X; x++ {
				for y := bricks[i].Start.Y; y <= bricks[i].End.Y; y++ {
					for z := bricks[i].Start.Z; z <= bricks[i].End.Z; z++ {
						stackPile[x][y][z] = bricks[i].ID
					}
				}
			}
		}
	}

	return fallenBricks
}

func parseBricks(input []string) []Brick {
	var bricks []Brick

	for i, l := range input {
		var b Brick

		ss := strings.Split(l, "~")

		b.ID = i + 1
		b.Start = parseCoordinate(ss[0])
		b.End = parseCoordinate(ss[1])

		// Always have the coordinate with the lowest value as starting coordinate
		if b.Start.X > b.End.X {
			t := b.End.X
			b.End.X = b.Start.X
			b.Start.X = t
		}
		if b.Start.Y > b.End.Y {
			t := b.End.Y
			b.End.Y = b.Start.Y
			b.Start.Y = t
		}
		if b.Start.Z > b.End.Z {
			t := b.End.Z
			b.End.Z = b.Start.Z
			b.Start.Z = t
		}

		bricks = append(bricks, b)
	}

	// Sort bricks so that bricks with the lowest Z-value are on top of the slice
	sort.Slice(bricks, func(i, j int) bool {
		return bricks[i].Start.Z < bricks[j].Start.Z
	})

	return bricks
}

func parseCoordinate(s string) Coordinate {
	var c Coordinate

	ss := strings.Split(s, ",")
	c.X, _ = strconv.Atoi(ss[0])
	c.Y, _ = strconv.Atoi(ss[1])
	c.Z, _ = strconv.Atoi(ss[2])

	return c
}

func createStackPile(x, y, z int) [][][]int {
	stackPile := make([][][]int, x)
	for i := 0; i < x; i++ {
		stackPile[i] = make([][]int, y)
		for j := 0; j < y; j++ {
			stackPile[i][j] = make([]int, z)
		}
	}
	return stackPile
}
