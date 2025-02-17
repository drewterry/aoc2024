package day04p1

import (
	"strings"
	"testing"

	"aoc/utils"

	"github.com/stretchr/testify/assert"
)

var testInput = `MMMSXXMASM
								 MSAMXMSMSA
								 AMXSXMAAMM
								 MSAMASMSMX
								 XMASAMXAMM
								 XXAMMXXAMA
								 SMSMSASXSS
								 SAXAMASAAA
								 MAMMMXMMMM
								 MXMXAXMASX`

func TestSolve(t *testing.T) {
	tests := []struct {
		input  string
		answer int
	}{
		{testInput, 18},
	}

	if testing.Verbose() {
		utils.Verbose = true
	}

	for _, test := range tests {
		r := strings.NewReader(test.input)

		result := Solve(r).(int)

		if result != test.answer {
			t.Errorf("Expected %d, got %d", test.answer, result)
		}
	}
}

func TestCheckHorizontal(t *testing.T) {
	// Forwards
	assert.True(t, CheckHorizontal([][]rune{{'X', 'X', 'M', 'A', 'S'}}, 0, 1))
	assert.True(t, CheckHorizontal([][]rune{{'X', 'M', 'A', 'S', 'S'}}, 0, 0))

	// Backwards
	assert.True(t, CheckHorizontal([][]rune{{'S', 'A', 'M', 'X', 'X'}}, 0, 0))
	assert.False(t, CheckHorizontal([][]rune{{'S', 'A', 'M', 'X', 'X'}}, 0, 1))

	// Wrong letters
	assert.False(t, CheckHorizontal([][]rune{{'X', 'M', 'A', 'S', 'S'}}, 0, 1))

	// Out of range
	assert.False(t, CheckHorizontal([][]rune{{'X', 'M', 'A', 'S', 'S'}}, 0, 2))
	assert.False(t, CheckHorizontal([][]rune{{'X', 'M', 'A', 'S', 'S'}}, 1, 0))
}

func TestCheckVertical(t *testing.T) {
	// Forwards
	assert.True(t, CheckVertical([][]rune{{'X'}, {'X'}, {'M'}, {'A'}, {'S'}}, 1, 0))
	assert.True(t, CheckVertical([][]rune{{'X'}, {'M'}, {'A'}, {'S'}, {'S'}}, 0, 0))

	// Backwards
	assert.True(t, CheckVertical([][]rune{{'S'}, {'A'}, {'M'}, {'X'}, {'X'}}, 0, 0))
	assert.False(t, CheckVertical([][]rune{{'S'}, {'A'}, {'M'}, {'X'}, {'X'}}, 1, 0))

	// Wrong letters
	assert.False(t, CheckVertical([][]rune{{'X'}, {'M'}, {'A'}, {'S'}, {'S'}}, 1, 0))

	// Out of range
	assert.False(t, CheckVertical([][]rune{{'X'}, {'X'}, {'M'}, {'A'}, {'S'}}, 2, 0))
	assert.False(t, CheckVertical([][]rune{{'X'}, {'X'}, {'M'}, {'A'}, {'S'}}, 0, 1))
}
