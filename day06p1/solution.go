package day06p1

import (
	"fmt"
	"io"
	"strings"

	"aoc/utils"
)

func Solve(r io.Reader) any {
	lines := utils.ReadLines(r)
	_ = lines

	var dir utils.Point
	var guard utils.Point
	lab := make([][]string, 0, len(lines))
	for y, row := range lines {
		cells := strings.Split(row, "")
		lab = append(lab, cells)

		for x, cell := range cells {
			switch cell {
			case "^":
				dir = utils.South
				guard = utils.Point{X: x, Y: y}

			case ">":
				dir = utils.East
				guard = utils.Point{X: x, Y: y}

			case "v":
				dir = utils.North
				guard = utils.Point{X: x, Y: y}

			case "<":
				dir = utils.West
				guard = utils.Point{X: x, Y: y}

			default:

			}
		}
	}

	printLab(lab)

	for {
		lab[guard.Y][guard.X] = "X"

		next := guard.Add(dir)
		if outOfBounds(lab, next) {
			break
		}

		if lab[next.Y][next.X] == "#" {
			dir = dir.Left()
		}

		guard = guard.Add(dir)
		if outOfBounds(lab, guard) {
			break
		}
	}

	printLab(lab)

	var locations int
	for _, row := range lab {
		for _, cell := range row {
			if cell == "X" {
				locations++
			}
		}
	}

	return locations
}

func outOfBounds(lab [][]string, point utils.Point) bool {
	return point.Y < 0 || point.Y > len(lab)-1 || point.X < 0 || point.X > len(lab[point.Y])-1
}

func printLab(lab [][]string) {
	for _, line := range lab {
		fmt.Println(line)
	}
}
