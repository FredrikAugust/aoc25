package solutions

import (
	"slices"
	"strings"

	"github.com/fredrikaugust/aoc25/utils"
)

func Day7A() int {
	input := utils.GetInput(7)

	m := make([][]byte, 0)

	for line := range strings.Lines(input) {
		m = append(m, []byte(strings.TrimSpace(line)))
	}

	lines := slices.Collect(strings.Lines(input))

	prevLine := lines[0]

	count := 0

	for _, line := range lines[1:] {
		lc := []byte(line)
		for x, c := range line {
			if c == '.' && (prevLine[x] == 'S' || prevLine[x] == '|') {
				lc[x] = '|'
			}

			if c == '^' && prevLine[x] == '|' {
				lc[x-1] = '|'
				lc[x+1] = '|'

				count++
			}
		}
		prevLine = string(lc)
	}

	return count
}

func Day7B() int {
	input := utils.GetInput(7)

	m := make([][]byte, 0)

	for line := range strings.Lines(input) {
		m = append(m, []byte(strings.TrimSpace(line)))
	}

	lines := slices.Collect(strings.Lines(input))

	rootPos := strings.Index(lines[0], "S")

	return traverse(rootPos, lines[1:])
}

type Pos struct {
	x         int
	linesLeft int
}

var cache = make(map[Pos]int, 0)

func traverse(x int, lines []string) int {
	pos := Pos{x, len(lines)}
	val, exists := cache[pos]
	if exists {
		return val
	}
	if len(lines) == 0 {
		return 1
	}
	if lines[0][x] == '^' {
		v := traverse(x-1, lines[1:]) + traverse(x+1, lines[1:])
		cache[pos] = v
		return v
	}
	cache[pos] = traverse(x, lines[1:])
	return traverse(x, lines[1:])
}
