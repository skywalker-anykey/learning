package main

import (
	"fmt"
	"math"
)

type Graph struct {
	vertices map[int]map[int]int
}

func NewGraph() *Graph {
	return &Graph{
		vertices: make(map[int]map[int]int),
	}
}

func (g *Graph) AddEdge(u, v, weight int) {
	if g.vertices[u] == nil {
		g.vertices[u] = make(map[int]int)
	}
	g.vertices[u][v] = weight
}

func (g *Graph) ShortestPath(start, end int) []int {
	dist := make(map[int]int)
	visited := make(map[int]bool)
	prev := make(map[int]int)
	queue := make([]int, 0)

	for vertex := range g.vertices {
		dist[vertex] = math.MaxInt32
		visited[vertex] = false
		prev[vertex] = -1
	}
	dist[start] = 0
	queue = append(queue, start)

	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]
		visited[u] = true

		for v, weight := range g.vertices[u] {
			if dist[u]+weight < dist[v] {
				dist[v] = dist[u] + weight
				prev[v] = u
				if !visited[v] {
					queue = append(queue, v)
				}
			}
		}
	}

	path := []int{end}
	for prev[end] != -1 {
		path = append([]int{prev[end]}, path...)
		end = prev[end]
	}

	return path
}

func main() {
	graph := NewGraph()

	graph.AddEdge(0, 1, 4)
	graph.AddEdge(0, 2, 1)
	graph.AddEdge(1, 3, 1)
	graph.AddEdge(2, 1, 2)
	graph.AddEdge(2, 3, 3)

	startNode := 0
	endNode := 3
	shortestPath := graph.ShortestPath(startNode, endNode)

	fmt.Printf("The shortest path between nodes %d and %d is: %v\n", startNode, endNode, shortestPath)
}
