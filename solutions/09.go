package solutions

import (
	"fmt"
	"math"
	"strings"

	"github.com/fredrikaugust/aoc25/utils"
)

type Coord struct {
	x int
	y int
}

func size(a, b Coord) int {
	return int((math.Abs(float64(a.x-b.x)) + 1) * (math.Abs(float64(a.y-b.y)) + 1))
}

func Day9A() int {
	input := utils.GetInput(9)
	coords := make([]Coord, 0)
	for line := range strings.Lines(input) {
		var x, y int
		fmt.Sscanf(line, "%d,%d", &x, &y)
		coords = append(coords, Coord{x, y})
	}

	biggest := 0

	for _, c1 := range coords {
		for _, c2 := range coords {
			if size(c1, c2) > biggest {
				biggest = size(c1, c2)
			}
		}
	}

	return biggest
}

func Day9B() int {
	input := utils.GetInput(9)
	coords := make([]Coord, 0)
	for line := range strings.Lines(input) {
		var x, y int
		fmt.Sscanf(line, "%d,%d", &x, &y)
		coords = append(coords, Coord{x, y})
	}
	edges := make([][]Coord, 0)
	for i, c := range coords {
		var last Coord
		if i == 0 {
			last = coords[len(coords)-1]
		} else {
			last = coords[i-1]
		}
		edges = append(edges, []Coord{c, last})
	}
	// this was borrowed in large part from https://github.com/blfuentes/AdventOfCode_AllYears/blob/main/AdventOfCode_2025_Go/day09/day09_2.go
	// which is again inspired by https://kishimotostudios.com/articles/aabb_collision/
	// ended up having to doodle this out on a piece of paper to understand it
	intersects := func(coord1, coord2 Coord) bool {
		var minX, maxX, minY, maxY int
		if coord1.x < coord2.x {
			minX = coord1.x
			maxX = coord2.x
		} else {
			minX = coord2.x
			maxX = coord1.x
		}
		if coord1.y < coord2.y {
			minY = coord1.y
			maxY = coord2.y
		} else {
			minY = coord2.y
			maxY = coord1.y
		}

		for _, edge := range edges {
			v1, v2 := edge[0], edge[1]
			var edgeMinX, edgeMaxX, edgeMinY, edgeMaxY int
			if v1.x < v2.x {
				edgeMinX = v1.x
				edgeMaxX = v2.x
			} else {
				edgeMinX = v2.x
				edgeMaxX = v1.x
			}
			if v1.y < v2.y {
				edgeMinY = v1.y
				edgeMaxY = v2.y
			} else {
				edgeMinY = v2.y
				edgeMaxY = v1.y
			}

			if minX < edgeMaxX && maxX > edgeMinX && minY < edgeMaxY && maxY > edgeMinY {
				return true
			}
		}

		return false
	}

	biggest := 0

	for _, c1 := range coords {
		for _, c2 := range coords {
			if intersects(c1, c2) {
				continue
			}
			if size(c1, c2) > biggest {
				biggest = size(c1, c2)
			}
		}
	}

	return biggest
}
