package solutions

import (
	"iter"
	"maps"
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/fredrikaugust/aoc25/utils"
)

type Vertex struct {
	X float64
	Y float64
	Z float64
}

func euclidianDistance3D(a, b *Vertex) float64 {
	return math.Sqrt(
		math.Pow(a.X-b.X, 2) + math.Pow(a.Y-b.Y, 2) + math.Pow(a.Z-b.Z, 2),
	)
}

type Edge struct {
	from     *Vertex
	to       *Vertex
	distance float64
}

func Day8A() int {
	input := utils.GetInput(8)

	lines := slices.Collect(strings.Lines(input))

	nodes := make([]Vertex, len(lines))
	for i, line := range lines {
		parts := strings.Split(strings.TrimSpace(line), ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		z, _ := strconv.Atoi(parts[2])
		nodes[i] = Vertex{float64(x), float64(y), float64(z)}
	}

	edges := make([]Edge, 0) // We do this to cache for later use
	for _, node1 := range nodes {
		for _, node2 := range nodes {
			if node1 == node2 {
				continue
			}
			dist := euclidianDistance3D(&node1, &node2)
			edges = append(edges, Edge{&node1, &node2, dist})
		}
	}
	slices.SortFunc(edges, func(e1, e2 Edge) int {
		if e1.distance < e2.distance {
			return -1
		} else if e1.distance > e2.distance {
			return 1
		}
		return 0
	})
	edges = slices.CompactFunc(edges, func(e1, e2 Edge) bool {
		return *e1.from == *e2.to && *e1.to == *e2.from
	})

	edges = edges[:1000]

	visitMap := make(map[Vertex]bool)
	for _, node := range nodes {
		visitMap[node] = false
	}

	components := make([]int, 0)
	for _, node := range nodes {
		if !visitMap[node] {
			visitMap[node] = true
			components = append(components, 1+DFS(node, &visitMap, edges))
		}
	}

	slices.Sort(components)
	slices.Reverse(components)

	return components[0] * components[1] * components[2]
}

func DFS(node Vertex, visitMap *map[Vertex]bool, edges []Edge) int {
	connectedNodes := make([]Vertex, 0)
	for _, edge := range edges {
		if *edge.from == node {
			if (*visitMap)[*edge.to] == false {
				connectedNodes = append(connectedNodes, *edge.to)
			}
		}

		if *edge.to == node {
			if (*visitMap)[*edge.from] == false {
				connectedNodes = append(connectedNodes, *edge.from)
			}
		}
	}

	sum := 1
	if (*visitMap)[node] == true {
		sum = 0
	}

	(*visitMap)[node] = true
	for _, node := range connectedNodes {
		sum += DFS(node, visitMap, edges)
	}
	return sum
}

type Graph struct {
	Vertices map[Vertex][]Vertex
}

func NewGraph() *Graph {
	return &Graph{
		Vertices: make(map[Vertex][]Vertex),
	}
}

func (g *Graph) AddVertex(v *Vertex) {
	g.Vertices[*v] = make([]Vertex, 0)
}

func (g *Graph) GetVertices() iter.Seq[Vertex] {
	return maps.Keys(g.Vertices)
}

func (g *Graph) AddEdge(e *Edge) {
	g.Vertices[*e.from] = append(g.Vertices[*e.from], *e.to)
	g.Vertices[*e.to] = append(g.Vertices[*e.to], *e.from)
}

func (g *Graph) GetNeighbours(v *Vertex) []Vertex {
	return g.Vertices[*v]
}

func (g *Graph) CalculateComponents() int {
	visited := make(map[Vertex]bool)
	for v := range g.GetVertices() {
		visited[v] = false
	}
	components := make([][]Vertex, 0)
	for v := range visited {
		if visited[v] == false {
			components = append(components, g.DFS(&v, &visited))
		}
	}
	return len(components)
}

func (g *Graph) DFS(v *Vertex, visited *map[Vertex]bool) []Vertex {
	(*visited)[*v] = true
	out := []Vertex{*v}

	for _, neighbour := range g.GetNeighbours(v) {
		if (*visited)[neighbour] {
			continue
		}

		out = append(out, g.DFS(&neighbour, visited)...)
	}

	return out
}

func Day8B() int {
	input := utils.GetInput(8)

	lines := slices.Collect(strings.Lines(input))

	g := NewGraph()

	for _, line := range lines {
		parts := strings.Split(strings.TrimSpace(line), ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		z, _ := strconv.Atoi(parts[2])

		g.AddVertex(&Vertex{float64(x), float64(y), float64(z)})
	}

	edges := make([]Edge, 0) // We do this to cache for later use
	for node1 := range g.GetVertices() {
		for node2 := range g.GetVertices() {
			if node1 == node2 {
				continue
			}
			dist := euclidianDistance3D(&node1, &node2)
			edges = append(edges, Edge{&node1, &node2, dist})
		}
	}
	slices.SortFunc(edges, func(e1, e2 Edge) int {
		if e1.distance < e2.distance {
			return -1
		} else if e1.distance > e2.distance {
			return 1
		}
		return 0
	})
	edges = slices.CompactFunc(edges, func(e1, e2 Edge) bool {
		return *e1.from == *e2.to && *e1.to == *e2.from
	})
	for {
		g.AddEdge(&edges[0])
		if g.CalculateComponents() == 1 {
			return int(edges[0].from.X * edges[0].to.X)
		}
		edges = edges[1:]

	}
}
