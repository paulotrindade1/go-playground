package main

import "fmt"

type Node struct {
	val   int
	left  *Node
	right *Node
}

func tree() []Node {
	return []Node{
		{
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
					val: 5,
				},
			},
		},
	}
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

var elements = []Node{}

func main() {
	t := tree()

	dfs(t[0])
	for _, elem := range elements {
		fmt.Println(elem.val)
	}
}
