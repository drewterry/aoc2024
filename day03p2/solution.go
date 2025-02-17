package day03p2

import (
	"io"
	"regexp"
	"strconv"
	"strings"

	"aoc/utils"
)

func Solve(r io.Reader) any {
	lines := utils.ReadLines(r)
	re, err := regexp.Compile(`mul\(([0-9]+),([0-9]+)\)|(do\(\))|(don't\(\))`)
	utils.Check(err, "re")

	var result int
	enabled := true
	for _, line := range lines {
		matches := re.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			if match[0] == "do()" {
				enabled = true
				continue
			}

			if match[0] == "don't()" {
				enabled = false
				continue
			}

			if strings.HasPrefix(match[0], "mul") && enabled {
				a, err := strconv.Atoi(match[1])
				utils.Check(err, "a: %s", match[1])

				b, err := strconv.Atoi(match[2])
				utils.Check(err, "a: %s", match[2])

				result += a * b
			}
		}
	}

	return result
}
