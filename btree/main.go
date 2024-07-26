package main

import "fmt"

var elements = []Node{}

func main() {
	t := Tree{
		Root: Node{
			val: 2,
			left: &Node{
				val: 1,
			},
			right: &Node{
				val: 4,
				left: &Node{
					val: 3,
				},
				right: &Node{
					val: 10,
				},
			},
		},
	} // balanced tree

	fmt.Println(t.SearchLargestN(5))
}

type Tree struct {
	Root Node
}

type Node struct {
	val   int
	left  *Node
	right *Node
}

func (t *Tree) PrintElements() {
	dfs(t.Root)
	for _, elem := range elements {
		fmt.Printf("%d ", elem.val)
	}
}

func (t *Tree) InvertBTree() {
	fmt.Println("not implemented")
}

func (t *Tree) SearchLargestN(n int) int {
	dfs(t.Root)
	if n > len(elements) {
		fmt.Println("largest n is greater than the number of elements in the tree")
		return -1
	}

	return elements[n-1].val
}

func dfs(node Node) {
	if node.left == nil && node.right == nil {
		elements = append(elements, node)
		return
	}

	if node.left != nil && node.right != nil {
		dfs(*node.left)
		elements = append(elements, node)
		dfs(*node.right)
	} else if node.left != nil {
		dfs(*node.left)
		elements = append(elements, node)
	} else {
		dfs(*node.right)
		elements = append(elements, node)
	}
}
