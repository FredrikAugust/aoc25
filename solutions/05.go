package solutions

import (
	"slices"
	"strconv"
	"strings"

	"github.com/fredrikaugust/aoc25/utils"
)

func Day5A() int {
	input := utils.GetInput(5)

	rangesStr, ingredientsStr, _ := strings.Cut(input, "\n\n")

	type Range struct {
		start int
		end   int
	}

	ranges := make([]Range, 0)

	for line := range strings.Lines(rangesStr) {
		line = strings.TrimSpace(line)
		start, end, _ := strings.Cut(line, "-")
		s, _ := strconv.Atoi(start)
		e, _ := strconv.Atoi(end)
		ranges = append(ranges, Range{start: s, end: e})
	}

	count := 0
	for ingredientIDStr := range strings.Lines(ingredientsStr) {
		ingredientIDStr = strings.TrimSpace(ingredientIDStr)
		ingredientID, _ := strconv.Atoi(ingredientIDStr)
		if slices.ContainsFunc(ranges, func(r Range) bool {
			return r.start <= ingredientID && r.end >= ingredientID
		}) {
			count++
		}
	}

	return count
}

type Range struct {
	Start int
	End   int
}

func Day5B() int {
	input := utils.GetInput(5)

	rangesStr, _, _ := strings.Cut(input, "\n\n")

	ranges := make([]Range, 0)

	for line := range strings.Lines(rangesStr) {
		line = strings.TrimSpace(line)
		start, end, _ := strings.Cut(line, "-")
		s, _ := strconv.Atoi(start)
		e, _ := strconv.Atoi(end)
		ranges = append(ranges, Range{Start: s, End: e})
	}

	return SolveDay05(ranges)
}

func SolveDay05(ranges []Range) int {
	slices.SortFunc(ranges, func(r1 Range, r2 Range) int {
		return r1.Start - r2.Start
	})

	count := 0

	idx := 0
	pos := ranges[idx].Start
	lastReachablePos := ranges[idx].End

	for {
		if idx+1 == len(ranges) {
			count += lastReachablePos - pos + 1
			break
		}

		next := ranges[idx+1]

		if next.Start <= lastReachablePos {
			count += next.Start - pos
			pos = next.Start
			if next.End > lastReachablePos {
				lastReachablePos = next.End
			}
		} else {
			count += lastReachablePos - pos + 1
			pos = next.Start
			lastReachablePos = next.End
		}
		idx++
	}

	return count
}
