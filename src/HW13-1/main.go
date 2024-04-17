package main

import (
	"HW13-1/bintree"
	"HW13-1/graph"
	"fmt"
)

func main() {
	main_bintree()
	main_notdirectedgraph()
	main_directedgraph()
}

func main_bintree() {
	fmt.Println("бинарное дерево")
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	t := bintree.CreateBinaryTree(arr)

	t.PrintInOrder()

	t.Add(15)
	t.Add(13)
	t.Add(11)

	t.PrintInOrder()

	fmt.Println(t.Find(3))

	t.DeleteNode(3)

	fmt.Println(t.Find(3))

	t.PrintInOrder()

	t.Free()
}

func main_directedgraph() {
	fmt.Println("ориентированный граф")
	// Структура: ориентированный граф
	gph := new(graph.Graph)
	gph.Init(6)
	gph.AddDirectedEdge(5, 2, 1)
	gph.AddDirectedEdge(5, 0, 1)
	gph.AddDirectedEdge(4, 0, 1)
	gph.AddDirectedEdge(4, 1, 1)
	gph.AddDirectedEdge(2, 3, 1)
	gph.AddDirectedEdge(3, 1, 1)
	gph.Print()
	// Функции: определения кратчайшего пути между любой парой вершин

}

func main_notdirectedgraph() {
	fmt.Println("неориентированный граф")
	// Структура: неориентированный граф.
	gph := new(graph.Graph)
	gph.Init(4)
	gph.AddUndirectedEdge(0, 1, 1)
	gph.AddUndirectedEdge(0, 2, 1)
	gph.AddUndirectedEdge(1, 2, 1)
	gph.AddUndirectedEdge(2, 3, 1)
	gph.Print()
	// Функции: поиск в ширину

}
