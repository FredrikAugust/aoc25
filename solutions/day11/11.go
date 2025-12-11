package day11

import (
	"iter"
	"log/slog"
	"maps"
	"strings"

	"github.com/fredrikaugust/aoc25/utils"
)

type Graph struct {
	Vertices map[string][]string
}

func NewGraph() *Graph {
	return &Graph{
		Vertices: make(map[string][]string),
	}
}

func (g *Graph) AddVertex(v string) {
	_, exists := g.Vertices[v]
	if !exists {
		g.Vertices[v] = make([]string, 0)
	}
}

func (g *Graph) GetVertices() iter.Seq[string] {
	return maps.Keys(g.Vertices)
}

func (g *Graph) AddEdge(from, to string) {
	g.Vertices[from] = append(g.Vertices[from], to)
	// This is a directed graph
	// g.Vertices[to] = append(g.Vertices[to], from)
}

func (g *Graph) GetNeighbours(v string) []string {
	return g.Vertices[v]
}

func Day11A() int {
	input := utils.GetInput(11)

	g := NewGraph()

	for line := range strings.Lines(input) {
		input, outputs, _ := strings.Cut(strings.TrimSpace(line), ": ")
		g.AddVertex(input)

		for o := range strings.SplitSeq(outputs, " ") {
			g.AddVertex(o)
			g.AddEdge(input, o)
		}
	}

	count := 0
	q := make([]string, 0)
	q = append(q, "you") // init

	for len(q) > 0 {
		head := q[0]
		q = q[1:]

		neighbours := g.GetNeighbours(head)

		for _, n := range neighbours {
			if n == "out" {
				count++
			} else {
				q = append(q, n)
			}
		}
	}

	return count
}

func Day11B() int {
	input := utils.GetInput(11)

	g := NewGraph()

	for line := range strings.Lines(input) {
		input, outputs, _ := strings.Cut(strings.TrimSpace(line), ": ")
		g.AddVertex(input)

		for o := range strings.SplitSeq(outputs, " ") {
			g.AddVertex(o)
			g.AddEdge(input, o)
		}
	}

	// just checked if dac->fft existed, and it didn't so we can just do this...
	p1 := dfs(g, "svr", "fft")
	clear(cache)
	p2 := dfs(g, "fft", "dac")
	clear(cache)
	p3 := dfs(g, "dac", "out")

	return p1 * p2 * p3
}

var cache = make(map[string]int)

func dfs(g *Graph, from, to string) int {
	if val, exists := cache[from]; exists {
		return val
	}

	paths := 0

	for _, neighbour := range g.GetNeighbours(from) {
		if neighbour == to {
			paths++
			if paths%100 == 0 {
				slog.Info("paths", "val", paths)
			}
			continue
		}

		paths += dfs(g, neighbour, to)
	}

	cache[from] = paths

	return paths
}
