package day01p1

import (
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"

	"aoc/utils"
)

func CheckAdd(a, b int) (int, bool) {
	sum := a + b
	// If overflow occurred, sum + (-a) != b
	if (sum - a) != b {
		return sum, true
	}
	return sum, false
}

func Solve(r io.Reader) any {
	lines := utils.ReadLines(r)

	l0 := make([]int, len(lines))
	l1 := make([]int, len(lines))

	var err error
	for i, line := range lines {
		f := strings.Fields(line)

		l0[i], err = strconv.Atoi(f[0])
		utils.Check(err, "l0 %d", i)

		l1[i], err = strconv.Atoi(f[1])
		utils.Check(err, "l1 %d", i)
	}

	sort.Ints(l0)
	sort.Ints(l1)

	var total int
	var overflow bool
	for i := range lines {
		distance := l0[i] - l1[i]

		if distance < 0 {
			distance = -distance
		}

		total, overflow = CheckAdd(total, distance)

		if overflow {
			fmt.Println("overflow")
		}
	}

	return total
}
