package utils

import (
	"math"
	"strconv"
	"strings"
)

func ListOfIntsFromString(i string, sep string) []int {
	out := make([]int, 0)
	for n := range strings.SplitSeq(i, sep) {
		num, _ := strconv.Atoi(n)
		out = append(out, num)
	}
	return out
}

func BigIntFromString(i string) int {
	out := 0
	for i, n := range strings.Split(i, ",") {
		num, _ := strconv.Atoi(n)
		out += num * int(math.Pow10(len(n)-i-1))
	}
	return out
}
