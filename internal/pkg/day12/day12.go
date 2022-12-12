package day12

import (
	"advent-of-code-2022/internal/pkg/utils"
	"sort"
	"strings"
)

/*
--- Day 12: Hill Climbing Algorithm ---

You try contacting the Elves using your handheld device, but the river you're following must be too low to get a decent signal.

You ask the device for a heightmap of the surrounding area (your puzzle input). The heightmap shows the local area from above broken into a grid; the elevation of each square of the grid is given by a single lowercase letter, where a is the lowest elevation, b is the next-lowest, and so on up to the highest elevation, z.

Also included on the heightmap are marks for your current position (S) and the location that should get the best signal (E). Your current position (S) has elevation a, and the location that should get the best signal (E) has elevation z.

You'd like to reach E, but to save energy, you should do it in as few steps as possible. During each step, you can move exactly one square up, down, left, or right. To avoid needing to get out your climbing gear, the elevation of the destination square can be at most one higher than the elevation of your current square; that is, if your current elevation is m, you could step to elevation n, but not to elevation o. (This also means that the elevation of the destination square can be much lower than the elevation of your current square.)

For example:

Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi

Here, you start in the top-left corner; your goal is near the middle. You could start by moving down or right, but eventually you'll need to head toward the e at the bottom. From there, you can spiral around to the goal:

v..v<<<<
>v.vv<<^
.>vv>E^^
..v>>>^^
..>>>>>^

In the above diagram, the symbols indicate whether the path exits each square moving up (^), down (v), left (<), or right (>). The location that should get the best signal is still E, and . marks unvisited squares.

This path reaches the goal in 31 steps, the fewest possible.

What is the fewest steps required to move from your current position to the location that should get the best signal?
*/

type VertexType int

const (
	Start  VertexType = 0
	End    VertexType = 1
	Normal VertexType = 2
)

type Position struct {
	X int
	Y int
}

type Vertex struct {
	Type   VertexType
	Parent *Vertex
	Height int
}

type Graph struct {
	vertices map[Position]*Vertex
	edges    map[Position]*utils.Set[*Vertex]
}

func (g *Graph) AddVertexAtPos(p Position, v *Vertex) {
	if g.vertices == nil {
		g.vertices = make(map[Position]*Vertex)
	}
	g.vertices[p] = v
}

func (g *Graph) AddEdge(p1 Position, p2 Position) {
	if g.edges == nil {
		g.edges = make(map[Position]*utils.Set[*Vertex])
	}

	_, found := g.edges[p1]
	if !found {
		g.edges[p1] = new(utils.Set[*Vertex])
	}
	v2 := g.vertices[p2]
	g.edges[p1].Add(v2)
}

func (g *Graph) GetEdgesAtPos(p Position) []*Vertex {
	return g.edges[p].ToSlice()
}

func (g *Graph) GetVertexAtPos(p Position) *Vertex {
	return g.vertices[p]
}

/*
 1  procedure BFS(G, root) is
 2      let Q be a queue
 3      label root as explored
 4      Q.enqueue(root)
 5      while Q is not empty do
 6          v := Q.dequeue()
 7          if v is the goal then
 8              return v
 9          for all edges from v to w in G.adjacentEdges(v) do
10              if w is not labeled as explored then
11                  label w as explored
12                  w.parent := v
13                  Q.enqueue(w)
*/

func (g *Graph) bfs(s *Vertex, e *Vertex) int {
	steps := 0

	q := utils.Queue[*Vertex]{}
	explored := utils.Set[*Vertex]{}

	q.Enqueue(s)
	explored.Add(s)

	for q.Len() > 0 {
		v := q.Dequeue()
		if v == e {
			for v != nil && v.Parent != nil {
				steps++
				v = v.Parent
			}
			for _, v := range g.vertices {
				v.Parent = nil
			}
			break
		}

		vpos := Position{}
		for ck, cv := range g.vertices {
			if cv == v {
				vpos = ck
				break
			}
		}

		for _, w := range g.GetEdgesAtPos(vpos) {
			if !explored.Contains(w) {
				explored.Add(w)
				w.Parent = v
				q.Enqueue(w)
			}
		}
	}

	return steps
}

func RuneToVertex(r rune) *Vertex {
	vType := Normal
	height := r - 96

	if r == 'S' {
		vType = Start
		height = 'a' - 96
	}

	if r == 'E' {
		vType = End
		height = 'z' - 96
	}

	return &Vertex{
		Type:   vType,
		Parent: nil,
		Height: int(height),
	}
}

func StarOne(input_path string) int {
	input := utils.MustReadFileToString(input_path)

	lines := strings.Split(input, "\n")

	heightMap := make([][]rune, len(lines))
	for i, line := range lines {
		for _, char := range line {
			heightMap[i] = append(heightMap[i], char)
		}
	}

	graph := Graph{}
	var start *Vertex = nil
	var end *Vertex = nil

	for y, row := range heightMap {
		for x, col := range row {
			v := RuneToVertex(col)

			if v.Type == Start {
				start = v
			}

			if v.Type == End {
				end = v
			}

			graph.AddVertexAtPos(Position{X: x, Y: y}, v)
		}
	}

	yMin := 0
	yMax := len(heightMap) - 1
	xMin := 0
	xMax := len(heightMap[yMin]) - 1

	for y, row := range heightMap {
		for x, _ := range row {
			p := Position{X: x, Y: y}
			v := graph.GetVertexAtPos(p)

			// node to the left of us
			if p.X-1 >= xMin {
				op := Position{Y: p.Y, X: p.X - 1}

				ov := graph.GetVertexAtPos(op)
				dh := ov.Height - v.Height

				if dh <= 1 {
					graph.AddEdge(p, op)
				}
			}

			// node to the right of us
			if p.X+1 <= xMax {
				op := Position{Y: p.Y, X: p.X + 1}

				ov := graph.GetVertexAtPos(op)
				dh := ov.Height - v.Height

				if dh <= 1 {
					graph.AddEdge(p, op)
				}
			}

			// node above us
			if p.Y-1 >= yMin {
				op := Position{Y: p.Y - 1, X: p.X}

				ov := graph.GetVertexAtPos(op)
				dh := ov.Height - v.Height

				if dh <= 1 {
					graph.AddEdge(p, op)
				}
			}

			// node below us
			if p.Y+1 <= yMax {
				op := Position{Y: p.Y + 1, X: p.X}

				ov := graph.GetVertexAtPos(op)
				dh := ov.Height - v.Height

				if dh <= 1 {
					graph.AddEdge(p, op)
				}
			}
		}
	}

	return graph.bfs(start, end)
}

func StarTwo(input_path string) int {
	input := utils.MustReadFileToString(input_path)

	lines := strings.Split(input, "\n")

	heightMap := make([][]rune, len(lines))
	for i, line := range lines {
		for _, char := range line {
			heightMap[i] = append(heightMap[i], char)
		}
	}

	graph := Graph{}
	startingPoints := make([]*Vertex, 0)
	var end *Vertex = nil

	for y, row := range heightMap {
		for x, col := range row {
			v := RuneToVertex(col)

			if v.Type == Start || v.Height == 1 {
				startingPoints = append(startingPoints, v)
			}

			if v.Type == End {
				end = v
			}

			graph.AddVertexAtPos(Position{X: x, Y: y}, v)
		}
	}

	yMin := 0
	yMax := len(heightMap) - 1
	xMin := 0
	xMax := len(heightMap[yMin]) - 1

	for y, row := range heightMap {
		for x, _ := range row {
			p := Position{X: x, Y: y}
			v := graph.GetVertexAtPos(p)

			// node to the left of us
			if p.X-1 >= xMin {
				op := Position{Y: p.Y, X: p.X - 1}

				ov := graph.GetVertexAtPos(op)
				dh := ov.Height - v.Height

				if dh <= 1 {
					graph.AddEdge(p, op)
				}
			}

			// node to the right of us
			if p.X+1 <= xMax {
				op := Position{Y: p.Y, X: p.X + 1}

				ov := graph.GetVertexAtPos(op)
				dh := ov.Height - v.Height

				if dh <= 1 {
					graph.AddEdge(p, op)
				}
			}

			// node above us
			if p.Y-1 >= yMin {
				op := Position{Y: p.Y - 1, X: p.X}

				ov := graph.GetVertexAtPos(op)
				dh := ov.Height - v.Height

				if dh <= 1 {
					graph.AddEdge(p, op)
				}
			}

			// node below us
			if p.Y+1 <= yMax {
				op := Position{Y: p.Y + 1, X: p.X}

				ov := graph.GetVertexAtPos(op)
				dh := ov.Height - v.Height

				if dh <= 1 {
					graph.AddEdge(p, op)
				}
			}
		}
	}

	results := make([]int, 0)
	for _, start := range startingPoints {
		result := graph.bfs(start, end)

		if result != 0 {
			results = append(results, result)
		}
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i] < results[j]
	})

	return results[0]
}
