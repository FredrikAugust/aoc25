package main

import (
	"log/slog"
	"time"

	"github.com/fredrikaugust/aoc25/solutions"
)

func main() {
	slog.SetLogLoggerLevel(slog.LevelDebug)

	timeFunc(1, 1, solutions.Day1A)
	timeFunc(1, 2, solutions.Day1B)

	timeFunc(2, 1, solutions.Day2A)
	timeFunc(2, 2, solutions.Day2B)

	timeFunc(3, 1, solutions.Day3A)
	timeFunc(3, 2, solutions.Day3B)

	timeFunc(4, 1, solutions.Day4A)
	timeFunc(4, 2, solutions.Day4B)
}

func timeFunc(day, part int, f func() int) {
	pre := time.Now()
	res := f()
	slog.Info("finished", "day", day, "part", part, "duration", time.Since(pre), "result", res)
}
