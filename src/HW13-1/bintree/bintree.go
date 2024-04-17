package bintree

import "fmt"

// Структура: двоичное дерево
type Node struct {
	value       int
	left, right *Node
}

type Tree struct {
	root *Node
}

// Освобождение памяти
func (t *Tree) Free() {
	t.root = nil
}

// Функции: вставка элемента
func (t *Tree) Add(value int) {
	t.root = addUtil(t.root, value)
}

func addUtil(n *Node, value int) *Node {
	if n == nil {
		n = new(Node)
		n.value = value
		return n
	}
	if value < n.value {
		n.left = addUtil(n.left, value)
	} else {
		n.right = addUtil(n.right, value)
	}
	return n
}

// Функции: удаление элемента
func (t *Tree) DeleteNode(value int) {
	t.root = DeleteNode(t.root, value)
}

func DeleteNode(node *Node, value int) *Node {
	var temp *Node = nil
	if node != nil {
		if node.value == value {
			if node.left == nil && node.right == nil {
				return nil
			}
			if node.left == nil {
				temp = node.right
				return temp
			}
			if node.right == nil {
				temp = node.left
				return temp
			}
			maxNode := FindMax(node.left)
			maxValue := maxNode.value
			node.value = maxValue
			node.left = DeleteNode(node.left, maxValue)
		} else {
			if node.value > value {
				node.left = DeleteNode(node.left, value)
			} else {
				node.right = DeleteNode(node.right, value)
			}
		}
	}
	return node
}

func FindMax(curr *Node) *Node {
	var node *Node = curr
	if node == nil {
		fmt.Println("EmptyTreeException")
		return nil
	}

	for node.right != nil {
		node = node.right
	}
	return node
}

// Функции: поиск элемента
func (t *Tree) Find(value int) bool {
	var curr *Node = t.root
	for curr != nil {
		if curr.value == value {
			return true
		} else if curr.value > value {
			curr = curr.left
		} else {
			curr = curr.right
		}
	}
	return false
}

// Функции: печать элементов дерева
func (t *Tree) PrintPreOrder() {
	printPreOrder(t.root)
	fmt.Println()
}

func printPreOrder(n *Node) {
	if n == nil {
		return
	}
	fmt.Print(n.value, " ")
	printPreOrder(n.left)
	printPreOrder(n.right)
}

// Функции: печать элементов дерева
func (t *Tree) PrintPostOrder() {
	printPostOrder(t.root)
	fmt.Println()
}

func printPostOrder(n *Node) {
	if n == nil {
		return
	}
	printPostOrder(n.left)
	printPostOrder(n.right)
	fmt.Print(n.value, " ")
}

// Функции: печать элементов дерева
func (t *Tree) PrintInOrder() {
	printInOrder(t.root)
	fmt.Println()
}

func printInOrder(n *Node) {
	if n == nil {
		return
	}
	printInOrder(n.left)
	fmt.Print(n.value, " ")
	printInOrder(n.right)
}

// Создать бинарное дерево
func CreateBinaryTree(arr []int) *Tree {
	t := new(Tree)
	size := len(arr)
	t.root = createBinaryTreeUtil(arr, 0, size-1)
	return t
}

// создаем бинарное дерево из сортированого массива
func createBinaryTreeUtil(arr []int, start int, end int) *Node {
	if start > end {
		return nil
	}
	mid := (start + end) / 2
	curr := new(Node)
	curr.value = arr[mid]
	curr.left = createBinaryTreeUtil(arr, start, mid-1)
	curr.right = createBinaryTreeUtil(arr, mid+1, end)
	return curr
}
