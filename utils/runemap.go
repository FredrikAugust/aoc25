package utils

import "strings"

func InputTo2DRuneMap(input string) [][]rune {
	m := make([][]rune, 0)
	for l := range strings.Lines(input) {
		l = strings.TrimSpace(l)
		m = append(m, []rune(l))
	}
	return m
}
