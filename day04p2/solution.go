package day04p2

import (
	"io"
	"strings"

	"aoc/utils"
)

func CheckRightDiagonal(puzzle [][]rune, row int, col int) bool {
	if row+3 > len(puzzle) || col+3 > len(puzzle[row]) {
		return false
	}

	if row > len(puzzle)-1 || col > len(puzzle[row])-1 {
		return false
	}

	if puzzle[row+1][col+1] != 'A' {
		return false
	}

	solutions := [][]rune{
		{'M', 'M', 'S', 'S'},
		{'M', 'S', 'M', 'S'},
		{'S', 'S', 'M', 'M'},
		{'S', 'M', 'S', 'M'},
	}

	for _, solution := range solutions {
		if puzzle[row][col] == solution[0] &&
			puzzle[row][col+2] == solution[1] &&
			puzzle[row+2][col] == solution[2] &&
			puzzle[row+2][col+2] == solution[3] {
			return true
		}
	}

	return false
}

func Solve(r io.Reader) any {
	lines := utils.ReadLines(r)

	puzzle := make([][]rune, len(lines))
	solution := make([][]rune, len(lines))
	for i, line := range lines {
		line = strings.TrimSpace(line)
		puzzle[i] = []rune(line)

		solution[i] = make([]rune, len(line))
		for j := range line {
			solution[i][j] = '·'
		}
	}

	// fmt.Println("Puzzle:")
	// for _, line := range puzzle {
	// 	fmt.Printf("%c\n", line)
	// }

	var found int
	for row := range puzzle {
		for col := range puzzle[row] {
			if CheckRightDiagonal(puzzle, row, col) {
				found++
				solution[row][col] = puzzle[row][col]
				solution[row][col+2] = puzzle[row][col+2]
				solution[row+1][col+1] = puzzle[row+1][col+1]
				solution[row+2][col] = puzzle[row+2][col]
				solution[row+2][col+2] = puzzle[row+2][col+2]
			}
		}
	}

	// fmt.Println("\nSolution:")
	// for _, line := range solution {
	// 	fmt.Printf("%c\n", line)
	// }

	return found
}
