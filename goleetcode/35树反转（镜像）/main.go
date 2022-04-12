package main

import (
	"fmt"
)

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

func (n *TreeNode) reverse() {
	if n == nil {
		return
	}
	fmt.Println(n.Value)
	n.Left.reverse()
	n.Right.reverse()
}

func NewTreeNode(value int) *TreeNode {
	return &TreeNode{
		Value: value,
	}
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	temp := root.Left
	root.Left = root.Right
	root.Right = temp

	invertTree(root.Left)
	invertTree(root.Right)

	return root
}

func main() { //73题
	root := &TreeNode{ //4 27  13
		Value: 4,
	}

	root.Right = NewTreeNode(7)
	root.Left = NewTreeNode(2)

	root.Left.Right = NewTreeNode(3)
	root.Left.Left = NewTreeNode(1)

	root.reverse()

	fmt.Println()

	no:=invertTree(root)
	no.reverse()
}
