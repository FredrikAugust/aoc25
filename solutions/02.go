package solutions

import (
	"math"
	"strconv"
	"strings"

	"github.com/fredrikaugust/aoc25/utils"
)

func Day2A() int {
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

	return result
}

func factors(n int, cache *map[int][]int) (out []int) {
	val, exists := (*cache)[n]
	if exists {
		return val
	}

	for i := range n {
		if i == 0 {
			continue
		}

		if n%i == 0 {
			out = append(out, i)
		}
	}

	(*cache)[n] = out

	return out
}

func fn(n, x, l int) int {
	out := 0

	stem := int(n / int(math.Pow10(l-x)))

	for i := range l / x {
		out += stem * int(math.Pow10(i*x))
	}

	return out
}

func findInvalid(start, end int, factorsCache *map[int][]int, invalidCache *map[int]bool) []int {
	invalidInputs := make([]int, 0)
	for n := range end - start + 1 {
		n += start

		val, exists := (*invalidCache)[n]
		if exists && val {
			invalidInputs = append(invalidInputs, n)
			continue
		} else if exists && !val {
			continue
		}

		l := int(math.Ceil(math.Log10(float64(n))))
		fac := factors(l, factorsCache)

		found := false
		for _, f := range fac {
			repeated := fn(n, f, l)
			if repeated == n {
				found = true
				invalidInputs = append(invalidInputs, n)
				break
			}
		}

		(*invalidCache)[n] = found
	}
	return invalidInputs
}

func Day2B() int {
	input := utils.GetInput(2)

	invalidInputs := make([]int, 0)
	factorsCache := make(map[int][]int, 0)
	invalidCache := make(map[int]bool, 0)

	for interval := range strings.SplitSeq(input, ",") {
		s, e, _ := strings.Cut(interval, "-")
		start, _ := strconv.Atoi(s)
		end, _ := strconv.Atoi(e)

		invalid := findInvalid(start, end, &factorsCache, &invalidCache)
		invalidInputs = append(invalidInputs, invalid...)
	}

	result := 0
	for _, num := range invalidInputs {
		result += num
	}

	return result
}
