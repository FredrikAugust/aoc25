package utils

import (
	"fmt"
	"os"
	"strings"
)

func GetSample(day int) string {
	content, err := os.ReadFile(fmt.Sprintf("inputs/%d-sample.txt", day))
	if err != nil {
		panic("could not open sample input")
	}
	return strings.TrimSuffix(string(content), "\n")
}

func GetInput(day int) string {
	content, err := os.ReadFile(fmt.Sprintf("inputs/%d-input.txt", day))
	if err != nil {
		panic("could not open sample input")
	}
	return strings.TrimSuffix(string(content), "\n")
}
