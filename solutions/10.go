package solutions

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/fredrikaugust/aoc25/utils"
)

type Machine struct {
	Lights   int
	Buttons  [][]int
	Joltages []int
}

type JoltageMeter struct {
	Values []int
}

func Press(old []int, buttons []int) []int {
	newVoltages := make([]int, len(old))
	copy(newVoltages, old)
	for _, button := range buttons {
		newVoltages[button] += 1
	}
	return newVoltages
}

func (j *JoltageMeter) String() string {
	return fmt.Sprintf("%v", j.Values)
}

func step(lights int, buttonPresses []int) int {
	for _, b := range buttonPresses {
		lights ^= 1 << b
	}
	return lights
}

func Day10A() int {
	input := utils.GetInput(10)

	machines := make([]Machine, 0)

	machineRegex := regexp.MustCompile(`\[([.#]+)\] ((\([\d,]+\) ?)+) \{([\d,]+)\}`)

	for line := range strings.Lines(input) {
		machine := Machine{}

		match := machineRegex.FindStringSubmatch(line)
		lights := match[1]
		buttons := match[2]
		// Not used in part 1
		// joltages := match[4]

		l := 0
		for i, li := range lights {
			if li == '#' {
				l |= 1 << i
			}
		}
		machine.Lights = l
		b := make([][]int, 0)
		for buttonGroup := range strings.FieldsSeq(buttons) {
			buttonGroup = buttonGroup[1 : len(buttonGroup)-1]
			b = append(b, utils.ListOfIntsFromString(buttonGroup))
		}
		machine.Buttons = b

		machines = append(machines, machine)
	}

	type G struct {
		Vertices map[int][]int
	}

	sum := 0
outer:
	for _, m := range machines {
		graph := G{
			Vertices: make(map[int][]int),
		}
		graph.Vertices[0] = make([]int, 0)

		queue := make([][]int, 0) // [val, depth]
		queue = append(queue, []int{0b0, 1})

		for len(queue) > 0 {
			head, depth := queue[0][0], queue[0][1]
			queue = queue[1:]

			for _, buttonSet := range m.Buttons {
				new := step(head, buttonSet)
				_, exists := graph.Vertices[new]
				if exists { // since bfs we can just skip in the case where it would create a loop and we'll still get shortest
					continue
				}
				if new == m.Lights {
					sum += depth
					continue outer
				}
				graph.Vertices[new] = []int{head}
				graph.Vertices[head] = append(graph.Vertices[head], new)
				queue = append(queue, []int{new, depth + 1})
			}
		}
	}

	return sum
}
