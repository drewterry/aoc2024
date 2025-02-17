package day01p2

import (
	"fmt"
	"io"
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

	m1 := make(map[int]int)
	for _, x := range l1 {
		m1[x]++
	}

	var similarity int
	var overflow bool
	for _, x := range l0 {
		similarity, overflow = CheckAdd(similarity, x*m1[x])

		if overflow {
			fmt.Println("overflow")
		}
	}

	return similarity
}
