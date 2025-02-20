package day05p1

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"aoc/utils"
)

func Solve(r io.Reader) any {
	lines := utils.ReadLines(r)
	_ = lines

	rules := make(map[int]map[int]struct{})
	prints := make([][]int, 0, 256)
	section := false

	for _, line := range lines {
		if line == "" {
			section = true
			continue
		}

		if section {

			pages := strings.Split(line, ",")
			print := make([]int, 0, len(pages))

			for _, page := range pages {
				pageInt, err := strconv.Atoi(page)
				utils.Check(err, "print order: %s", page)

				print = append(print, pageInt)
			}

			prints = append(prints, print)

		} else {
			pages := strings.Split(line, "|")

			page1, err := strconv.Atoi(pages[0])
			utils.Check(err, "rules")
			page2, err := strconv.Atoi(pages[1])
			utils.Check(err, "rules")

			rule, exists := rules[page1]
			if !exists {
				rule = make(map[int]struct{})
				rules[page1] = rule
			}

			rule[page2] = struct{}{}
		}
	}

	var result int
	for _, print := range prints {

		correct := true
		for i, page := range print {
			if i == 0 {
				continue
			}

			rule, exists := rules[page]
			if !exists {
				continue
			}

			for _, prevPage := range print[:i] {
				_, exists := rule[prevPage]
				if exists {
					correct = false
					break
				}
			}
		}

		if correct {
			middle := (len(print) - 1) / 2
			result += print[middle]

			fmt.Println(print)
			fmt.Println(print[middle])
		}
	}

	return result
}
