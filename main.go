package main

import (
	"log/slog"
	"time"

	"github.com/fredrikaugust/aoc25/solutions"
)

func main() {
	slog.SetLogLoggerLevel(slog.LevelDebug)

	timeFunc(solutions.Day1A)
	timeFunc(solutions.Day1B)

	timeFunc(solutions.Day2A)
	timeFunc(solutions.Day2B)
}

func timeFunc(f func()) {
	pre := time.Now()
	f()
	slog.Info("finished", "duration", time.Since(pre))
}
