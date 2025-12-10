package utils

import (
	"strconv"
	"strings"
)

func ListOfIntsFromString(i string) []int {
	out := make([]int, 0)
	for n := range strings.SplitSeq(i, ",") {
		num, _ := strconv.Atoi(n)
		out = append(out, num)
	}
	return out
}
