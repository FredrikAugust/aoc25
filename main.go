package main

import (
	"log/slog"

	"github.com/fredrikaugust/aoc25/solutions"
)

func main() {
	slog.SetLogLoggerLevel(slog.LevelDebug)
	solutions.Day1A()
	solutions.Day1B()
}
