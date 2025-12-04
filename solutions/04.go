package solutions

import (
	"github.com/fredrikaugust/aoc25/utils"
)

func Day4A() int {
	input := utils.GetInput(4)

	m := utils.InputTo2DRuneMap(input)

	// x, y
	neighbourIndices := [][]int{
		{-1, -1}, {0, -1}, {1, -1},
		{-1, 0} /*     */, {1, 0},
		{-1, 1}, {0, 1}, {1, 1},
	}

	counter := 0

	for y, l := range m {
		for x, c := range l {
			if c == rune('.') {
				continue
			}
			neighbouringRollsOfPaper := 0
			for _, neighbour := range neighbourIndices {
				xOffset, yOffset := neighbour[0], neighbour[1]

				nX, nY := x+xOffset, y+yOffset
				if nY < 0 || nY >= len(m) {
					continue
				}
				if nX < 0 || nX >= len(m[0]) {
					continue
				}

				if n := m[nY][nX]; n == rune('@') {
					neighbouringRollsOfPaper++
				}
			}
			if neighbouringRollsOfPaper < 4 {
				counter++
			}
		}
	}

	return counter
}

// x, y
var neighbourIndices = [][]int{
	{-1, -1}, {0, -1}, {1, -1},
	{-1, 0} /*     */, {1, 0},
	{-1, 1}, {0, 1}, {1, 1},
}

func Day4B() int {
	input := utils.GetInput(4)

	m := utils.InputTo2DRuneMap(input)

	counter := 0

	neighbourCountMap := calculateNeighbours(m)

	for hasRemovable(&neighbourCountMap) {
		for y, l := range neighbourCountMap {
			for x, c := range l {
				if c != -1 && c < 4 {
					neighbourCountMap[y][x] = -1
					counter++
				}
			}
		}
		recalculateNeighbours(&neighbourCountMap)
	}

	return counter
}

func hasRemovable(m *[][]int) bool {
	for _, l := range *m {
		for _, c := range l {
			if c != -1 && c < 4 {
				return true
			}
		}
	}
	return false
}

func recalculateNeighbours(m *[][]int) {
	for y, l := range *m {
		for x, c := range l {
			if c == -1 {
				continue
			}
			neighbouringRollsOfPaper := 0
			for _, neighbour := range neighbourIndices {
				xOffset, yOffset := neighbour[0], neighbour[1]

				nX, nY := x+xOffset, y+yOffset
				if nY < 0 || nY >= len(*m) {
					continue
				}
				if nX < 0 || nX >= len((*m)[0]) {
					continue
				}

				if n := (*m)[nY][nX]; n != -1 {
					neighbouringRollsOfPaper++
				}
			}
			(*m)[y][x] = neighbouringRollsOfPaper
		}
	}
}

func calculateNeighbours(m [][]rune) [][]int {
	neighbourCountMap := make([][]int, len(m))
	for y, l := range m {
		line := make([]int, len(l))
		for x, c := range l {
			if c == rune('.') {
				line[x] = -1
				continue
			}
			neighbouringRollsOfPaper := 0
			for _, neighbour := range neighbourIndices {
				xOffset, yOffset := neighbour[0], neighbour[1]

				nX, nY := x+xOffset, y+yOffset
				if nY < 0 || nY >= len(m) {
					continue
				}
				if nX < 0 || nX >= len(m[0]) {
					continue
				}

				if n := m[nY][nX]; n == rune('@') {
					neighbouringRollsOfPaper++
				}
			}
			line[x] = neighbouringRollsOfPaper
		}
		neighbourCountMap[y] = line
	}
	return neighbourCountMap
}
