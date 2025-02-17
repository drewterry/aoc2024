package day03p1

import (
	"io"
	"regexp"
	"strconv"

	"aoc/utils"
)

func Solve(r io.Reader) any {
	lines := utils.ReadLines(r)
	re, err := regexp.Compile(`mul\(([0-9]+),([0-9]+)\)`)
	utils.Check(err, "re")

	var result int
	for _, line := range lines {
		matches := re.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			a, err := strconv.Atoi(match[1])
			utils.Check(err, "a: %s", match[1])

			b, err := strconv.Atoi(match[2])
			utils.Check(err, "a: %s", match[2])

			result += a * b
		}
	}

	return result
}
