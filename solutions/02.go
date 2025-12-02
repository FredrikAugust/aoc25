package solutions

import (
	"log/slog"
	"strconv"
	"strings"

	"github.com/fredrikaugust/aoc25/utils"
)

func Day2A() {
	input := utils.GetInput(2)

	invalidInputs := make([]int, 0)

	for interval := range strings.SplitSeq(input, ",") {
		parts := strings.Split(interval, "-")
		start, _ := strconv.Atoi(parts[0])
		end, _ := strconv.Atoi(parts[1])

		for n := range end - start + 1 {
			n += start

			nStr := strconv.Itoa(n)

			// odd length str can't be a substring twice
			if len(nStr)%2 != 0 {
				continue
			}

			firstHalf, lastHalf := nStr[0:len(nStr)/2], nStr[len(nStr)/2:]

			if firstHalf == lastHalf {
				invalidInputs = append(invalidInputs, n)
			}
		}
	}

	result := 0
	for _, num := range invalidInputs {
		result += num
	}

	slog.Info("day2 part 1", "result", result)
}

func factors(n int) (out []int) {
	for i := range n {
		if i == 0 {
			continue
		}

		if n%i == 0 {
			out = append(out, i)
		}
	}
	return out
}

func Day2B() {
	input := utils.GetInput(2)

	invalidInputs := make([]int, 0)

	for interval := range strings.SplitSeq(input, ",") {
		parts := strings.Split(interval, "-")
		start, _ := strconv.Atoi(parts[0])
		end, _ := strconv.Atoi(parts[1])

		for n := range end - start + 1 {
			n += start

			nStr := strconv.Itoa(n)

			fac := factors(len(nStr))

			for _, f := range fac {
				repeated := strings.Repeat(nStr[0:f], len(nStr)/f)
				if repeated == nStr {
					invalidInputs = append(invalidInputs, n)
					break
				}
			}
		}
	}

	result := 0
	for _, num := range invalidInputs {
		result += num
	}

	slog.Info("day2 part 2", "result", result)
}
