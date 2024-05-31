package main

import (
	"fmt"
)

type Graph struct {
	vertices map[int][]int
}

func NewGraph() *Graph {
	return &Graph{
		vertices: make(map[int][]int),
	}
}

func (g *Graph) AddEdge(u, v int) {
	g.vertices[u] = append(g.vertices[u], v)
	g.vertices[v] = append(g.vertices[v], u)
}

func (g *Graph) BFS(start int) {
	visited := make(map[int]bool)
	queue := make([]int, 0)

	visited[start] = true
	queue = append(queue, start)

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		fmt.Printf("%d ", node)

		for _, neighbor := range g.vertices[node] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}
}

func main() {
	graph := NewGraph()

	graph.AddEdge(0, 1)
	graph.AddEdge(0, 2)
	graph.AddEdge(0, 3)
	graph.AddEdge(1, 2)
	graph.AddEdge(2, 4)
	graph.AddEdge(3, 4)

	startNode := 0
	fmt.Printf("BFS traversal starting from node %d: ", startNode)
	graph.BFS(startNode)
}
