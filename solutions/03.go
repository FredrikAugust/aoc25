package solutions

import (
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/fredrikaugust/aoc25/utils"
)

func Day3A() int {
	banks := make([][]int, 0)

	input := utils.GetInput(3)
	for line := range strings.SplitSeq(input, "\n") {
		bank := make([]int, 0)
		for battery := range strings.SplitSeq(line, "") {
			val, _ := strconv.Atoi(battery)
			bank = append(bank, val)
		}
		banks = append(banks, bank)
	}

	sum := 0
	for _, bank := range banks {
		max := 0
		for i, battery1 := range bank[:len(bank)-1] {
			for _, battery2 := range bank[i+1:] {
				val := battery1*10 + battery2
				if val > max {
					max = val
				}
			}
		}
		sum += max
	}

	return sum
}

func Day3B() int {
	banks := make([][]int, 0)

	input := utils.GetInput(3)
	for line := range strings.SplitSeq(input, "\n") {
		bank := make([]int, 0)
		for battery := range strings.SplitSeq(line, "") {
			val, _ := strconv.Atoi(battery)
			bank = append(bank, val)
		}
		banks = append(banks, bank)
	}

	results := make([][]int, len(banks))
	for i, bank := range banks {
		solveLine(bank, 12, &results[i])
	}

	result := 0
	for _, line := range results {
		num := 0
		for i, n := range line {
			num += n * int(math.Pow10(12-i-1))
		}
		result += num
	}

	return result
}

func solveLine(bank []int, leftToChoose int, result *[]int) {
	if leftToChoose == 0 {
		return
	}

	searchSpace := bank[0 : len(bank)-leftToChoose+1]
	largestBattery := slices.Max(searchSpace)
	largestBatteryIx := slices.Index(searchSpace, largestBattery)

	*result = append(*result, largestBattery)
	solveLine(bank[largestBatteryIx+1:], leftToChoose-1, result)
}
