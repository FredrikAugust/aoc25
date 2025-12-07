package solutions

import (
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/fredrikaugust/aoc25/utils"
)

func transpose[E string | byte](input [][]E) [][]E {
	output := make([][]E, len(input[0]))

	for y := range len(input[0]) {
		line := make([]E, len(input))
		for x := range len(input) {
			line[x] = input[x][y]
		}
		output[y] = line
	}

	return output
}

func Day6A() int {
	input := utils.GetInput(6)
	// input := utils.GetSample(6)

	cells := make([][]string, 0)

	for line := range strings.Lines(input) {
		cells = append(cells, strings.Fields(line))
	}

	problems := transpose(cells)
	result := 0

	for _, problem := range problems {
		slices.Reverse(problem)
		operand := problem[0]

		switch operand {
		case "*":
			res := 1
			for _, num := range problem[1:] {
				n, _ := strconv.Atoi(num)
				res *= n
			}
			result += res
		case "+":
			res := 0
			for _, num := range problem[1:] {
				n, _ := strconv.Atoi(num)
				res += n
			}
			result += res

		}
	}

	return result
}

func getDigitInPos(num, pos int) int {
	return int(math.Floor(float64(num)/math.Pow10(pos))) % 10
}

func Day6B() int {
	input := utils.GetInput(6)
	// input := utils.GetSample(6)
	lines := slices.Collect(strings.Lines(input))

	longestLineLen := 0
	for _, l := range lines {
		if len(l) > longestLineLen {
			longestLineLen = len(l)
		}
	}

	blocks := make([][]string, 0)

	buf := make([]string, 0)
	for i := range longestLineLen {
		var column strings.Builder
		for _, line := range lines {
			if i >= len(line) {
				column.WriteString(" ")
			} else {
				column.WriteByte(line[i])
			}
		}
		// if all spaces i.e. finished block
		if strings.TrimSpace(column.String()) == "" {
			rows := make([]string, 0)
			for j := range len(buf[0]) {
				var row strings.Builder
				for i := range len(buf) {
					row.WriteByte(buf[i][j])
				}
				rows = append(rows, row.String())
			}
			buf = make([]string, 0)
			blocks = append(blocks, rows)
		} else {
			buf = append(buf, column.String())
		}
	}

	sum := 0
	for _, block := range blocks {
		operand := string(block[len(block)-1][0])
		numStrs := block[0 : len(block)-1]
		numBlock := make([][]byte, 0)
		for _, str := range numStrs {
			numBlock = append(numBlock, []byte(str))
		}
		numBlock = transpose(numBlock)
		numbers := make([]int, 0)
		for _, n := range numBlock {
			var s strings.Builder
			s.Write(n)
			i, _ := strconv.Atoi(strings.TrimSpace(s.String()))
			numbers = append(numbers, i)
		}

		switch operand {
		case "*":
			value := 1
			for _, n := range numbers {
				value *= n
			}
			sum += value
		case "+":
			value := 0
			for _, n := range numbers {
				value += n
			}
			sum += value
		}
	}

	return sum
}
