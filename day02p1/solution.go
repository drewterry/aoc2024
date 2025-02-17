package day02p1

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

		diffs := make([]int, len(fields)-1)
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

		if gra && (inc || dec) {
			safe++
		}
	}

	return safe
}
