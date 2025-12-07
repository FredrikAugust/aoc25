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

	display(m)

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

func display(m [][]byte) {
	for j := range len(m) {
		for i := range len(m[0]) {
			print(string(m[j][i]))
		}
		println()
	}
}

func Day7B() int {
	return 0
}
