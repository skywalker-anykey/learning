package graph

import "fmt"

type Edge struct {
	destination int
	cost        int
	next        *Edge
}

type Graph struct {
	count int
	Edges [](*Edge)
}

func (gph *Graph) Init(cnt int) {
	gph.count = cnt
	gph.Edges = make([]*Edge, cnt)
}

func (gph *Graph) AddDirectedEdge(source int, destination int, cost int) {
	edge := &Edge{destination, cost, gph.Edges[source]}
	gph.Edges[source] = edge
}

func (gph *Graph) AddDirectedEdgeUnweighted(source int, destination int) {
	gph.AddDirectedEdge(source, destination, 1)
}

func (gph *Graph) AddUndirectedEdge(source int, destination int, cost int) {
	gph.AddDirectedEdge(source, destination, cost)
	gph.AddDirectedEdge(destination, source, cost)
}

func (gph *Graph) AddUndirectedEdgeUnweghted(source int, destination int) {
	gph.AddUndirectedEdge(source, destination, 1)
}

func (gph *Graph) Print() {
	for i := 0; i < gph.count; i++ {
		ad := gph.Edges[i]
		fmt.Print("Vertex ", i, " is connected to : ")
		for ad != nil {
			fmt.Print(ad.destination, "(cost:", ad.cost, ") ")
			ad = ad.next
		}
		fmt.Println()
	}
}
