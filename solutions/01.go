package solutions

import (
	"log/slog"
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

		slog.Debug("update", "dir", string(line[0]), "amount", line[1:], "new pos", pos)
	}

	slog.Info("done part1", "result", zeroCount)
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

	slog.Info("done part2", "result", zeroCount)
}

func rotateLeft(pos, zeroCount *int, value int) {
	for _ = range value {
		(*pos)--
		if *pos == 0 {
			(*zeroCount)++
		} else if *pos < 0 {
			*pos = 99
		}
	}
}

func rotateRight(pos, zeroCount *int, value int) {
	for _ = range value {
		(*pos)++
		if *pos == 100 {
			(*zeroCount)++
			*pos = 0
		}
	}
}
