package main

import (
	"fmt"
)

type Node struct {
	Value string
	Left  *Node
	Right *Node
}

func preorder(node *Node) {
	if node == nil {
		return
	}
	fmt.Print(node.Value + " ")
	preorder(node.Left)
	preorder(node.Right)
}

func postorder(node *Node) {
	if node == nil {
		return
	}
	postorder(node.Left)
	postorder(node.Right)
	fmt.Print(node.Value + " ")
}

func main() {
	a := &Node{Value: "a"}
	b := &Node{Value: "b"}
	c := &Node{Value: "c"}
	minus := &Node{Value: "-", Left: b, Right: c}
	plus := &Node{Value: "+", Left: a, Right: minus}

	root := plus

	fmt.Println("Preorder Traversal:")
	preorder(root)
	fmt.Println()

	fmt.Println("Postorder Traversal:")
	postorder(root)
	fmt.Println()
}
