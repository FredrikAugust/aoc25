package solutions_test

import (
	"testing"

	"github.com/fredrikaugust/aoc25/solutions"
)

func Test_Day05(t *testing.T) {
	t.Run("one span", func(t *testing.T) {
		r := []solutions.Range{
			solutions.Range{
				Start: 0,
				End:   0,
			},
		}

		v := solutions.SolveDay05(r)

		if v != 1 {
			t.Error("wrong result")
		}
	})
	t.Run("fully nested", func(t *testing.T) {
		r := []solutions.Range{
			solutions.Range{5, 10},
			solutions.Range{6, 8},
		}

		v := solutions.SolveDay05(r)

		if v != 6 {
			t.Error("wrong result")
		}
	})
	t.Run("partially nested", func(t *testing.T) {
		r := []solutions.Range{
			solutions.Range{5, 10},
			solutions.Range{8, 12},
		}

		v := solutions.SolveDay05(r)

		if v != 8 {
			t.Error("wrong result")
		}
	})
}
