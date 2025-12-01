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

	slog.Info("done", "result", zeroCount)
}
