package main

import (
	"fmt"
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

	timeFunc(5, 1, solutions.Day5A)
	timeFunc(5, 2, solutions.Day5B)

	timeFunc(6, 1, solutions.Day6A)
	timeFunc(6, 2, solutions.Day6B)

	timeFunc(7, 1, solutions.Day7A)
	timeFunc(7, 2, solutions.Day7B)

	timeFunc(8, 1, solutions.Day8A)
	timeFunc(8, 2, solutions.Day8B2)

	timeFunc(9, 1, solutions.Day9A)
	timeFunc(9, 2, solutions.Day9B)
}

func timeFunc(day, part int, f func() int) {
	pre := time.Now()
	res := f()
	fmt.Printf("Finished Day %d.%d in %14s. Result = %v\n", day, part, time.Since(pre), res)
}
