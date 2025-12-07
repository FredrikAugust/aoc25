package solutions

import (
	"log/slog"
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/fredrikaugust/aoc25/utils"
)

func transpose(input [][]string) [][]string {
	output := make([][]string, len(input[0]))

	for y := range len(input[0]) {
		line := make([]string, len(input))
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
	// input := utils.GetInput(6)
	input := utils.GetSample(6)

	cells := make([][]string, 0)

	for line := range strings.Lines(input) {
		cells = append(cells, strings.Fields(line))
	}

	problems := transpose(cells)
	result := 0

	for _, problem := range problems {
		operand := problem[len(problem)-1]

		numbersStr := problem[:len(problem)-1]
		numbers := make([]int, len(numbersStr))
		for i, numStr := range numbersStr {
			numbers[i], _ = strconv.Atoi(numStr)
		}

		maxNum := slices.Max(numbers)
		maxK := int(math.Floor(math.Log10(float64(maxNum)))) + 1
		// newNumbers := make([]int, len(numbers))
		for k := range maxK {
			digits := make([]int, 0)
			for _, n := range numbers {
				digits = append(digits, getDigitInPos(n, k))
			}

			slog.Info("out", "val", digits, "maxK", maxK, "maxNum", maxNum, "numbers", numbers)
		}

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
