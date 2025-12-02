package solutions

import (
	"log/slog"
	"math"
	"strconv"
	"strings"

	"github.com/fredrikaugust/aoc25/utils"
)

func Day1A() {
	input := utils.GetInput(1)

	pos := 50
	zeroCount := 0

	for line := range strings.SplitSeq(input, "\n") {
		if line == "" {
			continue
		}

		switch line[0] {
		case 'L':
			value, _ := strconv.Atoi(line[1:])
			pos -= value
		case 'R':
			value, _ := strconv.Atoi(line[1:])
			pos += value
		}

		for pos < 0 {
			pos += 100
		}

		for pos > 99 {
			pos -= 100
		}

		if pos == 0 {
			zeroCount++
		}
	}

	slog.Info("day 1 done part1", "result", zeroCount)
}

func Day1B() {
	input := utils.GetInput(1)

	pos := 50
	zeroCount := 0

	for line := range strings.SplitSeq(input, "\n") {
		if line == "" {
			continue
		}

		switch line[0] {
		case 'L':
			value, _ := strconv.Atoi(line[1:])
			rotateLeft(&pos, &zeroCount, value)
		case 'R':
			value, _ := strconv.Atoi(line[1:])
			rotateRight(&pos, &zeroCount, value)
		}
	}

	slog.Info("day 1 done part2", "result", zeroCount)
}

func rotateLeft(pos, zeroCount *int, value int) {
	fullRounds := int(math.Floor(float64(value) / 100.0))
	*zeroCount += fullRounds
	value %= 100

	oldValue := *pos
	*pos -= value

	if *pos < 0 {
		if oldValue != 0 {
			*zeroCount++
		}
		*pos += 100
	} else if *pos == 0 {
		*zeroCount++
	}
}

func rotateRight(pos, zeroCount *int, value int) {
	fullRounds := int(math.Floor(float64(value) / 100.0))
	*zeroCount += fullRounds
	value %= 100

	*pos += value

	if *pos > 99 {
		*zeroCount++
		*pos -= 100
	}
}
