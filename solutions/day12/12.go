package day12

import (
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/fredrikaugust/aoc25/utils"
)

func Day12A() int {
	input := utils.GetInput(12)

	blockRe := regexp.MustCompile(`\d:\n([#.]+\n)+`)

	blocksStr := blockRe.FindAllString(input, -1)
	blocks := make([][][]bool, 0)

	for _, block := range blocksStr {
		b := make([][]bool, 0)
		lines := slices.Collect(strings.Lines(strings.TrimSpace(block)))
		lines = lines[1:]
		for _, line := range lines {
			l := make([]bool, 0)
			for _, c := range strings.TrimSpace(line) {
				if c == '#' {
					l = append(l, true)
				} else {
					l = append(l, false)
				}
			}
			b = append(b, l)
		}
		blocks = append(blocks, b)
	}

	inputRe := regexp.MustCompile(`\d+x\d+: .+`)

	inputsStr := inputRe.FindAllString(input, -1)
	type Problem struct {
		X int
		Y int

		Counts []int
	}
	inputs := make([]Problem, 0)
	for _, is := range inputsStr {
		size, counts, _ := strings.Cut(is, ": ")

		xS, yS, _ := strings.Cut(size, "x")
		x, _ := strconv.Atoi(xS)
		y, _ := strconv.Atoi(yS)

		problem := Problem{
			X: x,
			Y: y,
		}

		cs := utils.ListOfIntsFromString(counts, " ")
		problem.Counts = cs

		inputs = append(inputs, problem)
	}

	naiveSolutions := 0
	for _, input := range inputs {
		size := input.X * input.Y

		naiveRequired := 0
		for _, count := range input.Counts {
			naiveRequired += count
		}
		naiveRequired *= 9

		if naiveRequired <= size {
			naiveSolutions += 1
		}
	}

	return naiveSolutions
}
