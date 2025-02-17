package day04p1

import (
	"fmt"
	"io"
	"strings"

	"aoc/utils"
)

func CheckHorizontal(puzzle [][]rune, row int, col int) bool {
	if row > len(puzzle)-1 {
		return false
	}

	if col+4 > len(puzzle[row]) {
		return false
	}

	forward := puzzle[row][col] == 'X' &&
		puzzle[row][col+1] == 'M' &&
		puzzle[row][col+2] == 'A' &&
		puzzle[row][col+3] == 'S'

	backward := puzzle[row][col] == 'S' &&
		puzzle[row][col+1] == 'A' &&
		puzzle[row][col+2] == 'M' &&
		puzzle[row][col+3] == 'X'

	return forward || backward
}

func CheckVertical(puzzle [][]rune, row int, col int) bool {
	if row+4 > len(puzzle) {
		return false
	}

	if col > len(puzzle[row])-1 {
		return false
	}

	forward := puzzle[row][col] == 'X' &&
		puzzle[row+1][col] == 'M' &&
		puzzle[row+2][col] == 'A' &&
		puzzle[row+3][col] == 'S'

	backward := puzzle[row][col] == 'S' &&
		puzzle[row+1][col] == 'A' &&
		puzzle[row+2][col] == 'M' &&
		puzzle[row+3][col] == 'X'

	return forward || backward
}

func CheckRightDiagonal(puzzle [][]rune, row int, col int) bool {
	if row+4 > len(puzzle) || col+4 > len(puzzle[row]) {
		return false
	}

	if row > len(puzzle)-1 || col > len(puzzle[row])-1 {
		return false
	}

	forward := puzzle[row][col] == 'X' &&
		puzzle[row+1][col+1] == 'M' &&
		puzzle[row+2][col+2] == 'A' &&
		puzzle[row+3][col+3] == 'S'

	backward := puzzle[row][col] == 'S' &&
		puzzle[row+1][col+1] == 'A' &&
		puzzle[row+2][col+2] == 'M' &&
		puzzle[row+3][col+3] == 'X'

	return forward || backward
}

func CheckLeftDiagonal(puzzle [][]rune, row int, col int) bool {
	if row+4 > len(puzzle) || col-3 < 0 {
		return false
	}

	if row > len(puzzle)-1 || col > len(puzzle[row])-1 {
		return false
	}

	forward := puzzle[row][col] == 'X' &&
		puzzle[row+1][col-1] == 'M' &&
		puzzle[row+2][col-2] == 'A' &&
		puzzle[row+3][col-3] == 'S'

	backward := puzzle[row][col] == 'S' &&
		puzzle[row+1][col-1] == 'A' &&
		puzzle[row+2][col-2] == 'M' &&
		puzzle[row+3][col-3] == 'X'

	return forward || backward
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

	fmt.Println("Puzzle:")
	for _, line := range puzzle {
		fmt.Printf("%c\n", line)
	}

	var found int
	for row := range puzzle {
		for col := range puzzle[row] {
			if CheckHorizontal(puzzle, row, col) {
				found++
				solution[row][col] = puzzle[row][col]
				solution[row][col+1] = puzzle[row][col+1]
				solution[row][col+2] = puzzle[row][col+2]
				solution[row][col+3] = puzzle[row][col+3]
			}

			if CheckVertical(puzzle, row, col) {
				found++
				solution[row][col] = puzzle[row][col]
				solution[row+1][col] = puzzle[row+1][col]
				solution[row+2][col] = puzzle[row+2][col]
				solution[row+3][col] = puzzle[row+3][col]
			}

			if CheckRightDiagonal(puzzle, row, col) {
				found++
				solution[row][col] = puzzle[row][col]
				solution[row+1][col+1] = puzzle[row+1][col+1]
				solution[row+2][col+2] = puzzle[row+2][col+2]
				solution[row+3][col+3] = puzzle[row+3][col+3]
			}

			if CheckLeftDiagonal(puzzle, row, col) {
				found++
				solution[row][col] = puzzle[row][col]
				solution[row+1][col-1] = puzzle[row+1][col-1]
				solution[row+2][col-2] = puzzle[row+2][col-2]
				solution[row+3][col-3] = puzzle[row+3][col-3]
			}
		}
	}

	fmt.Println("\nSolution:")
	for _, line := range solution {
		fmt.Printf("%c\n", line)
	}

	return found
}
