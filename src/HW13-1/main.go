package main

import (
	"HW13-1/bintree"
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
}

func main_notdirectedgraph() {
	fmt.Println("неориентированный граф")
}
