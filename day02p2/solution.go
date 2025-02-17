package day02p2

import (
	"io"
	"strconv"
	"strings"

	"aoc/utils"
)

func Solve(r io.Reader) any {
	lines := utils.ReadLines(r)
	_ = lines

	var safe int
	var err error
	for _, line := range lines {
		fields := strings.Fields(line)

		line := make([]int, len(fields))
		for i, s := range fields {
			line[i], err = strconv.Atoi(s)
			utils.Check(err, "atoi: %d", i)
		}

		if checkSafe(line) {
			safe++
			continue
		}

		for i := range line {
			lineMod := make([]int, 0, len(line)-1)
			lineMod = append(lineMod, line[:i]...)
			lineMod = append(lineMod, line[i+1:]...)

			if checkSafe(lineMod) {
				safe++
				break
			}
		}
	}

	return safe
}

func checkSafe(line []int) bool {
	diffs := make([]int, len(line)-1)
	for i := 1; i < len(line); i++ {
		diffs[i-1] = line[i-1] - line[i]
	}

	gra := true
	inc := true
	dec := true
	for _, diff := range diffs {
		gra = gra && diff <= 3 && diff >= -3
		inc = inc && diff > 0
		dec = dec && diff < 0
	}

	return gra && (inc || dec)
}
