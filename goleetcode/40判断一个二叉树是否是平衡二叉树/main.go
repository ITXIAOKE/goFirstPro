package main

import (
	"fmt"
)

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(value int) *TreeNode {
	return &TreeNode{
		Value: value,
	}
}

func (node *TreeNode) reverse() {
	if node == nil {
		return
	}
	fmt.Println(node.Value)
	node.Left.reverse()
	node.Right.reverse()
}

func isBalances(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return abs(heightHelper(root.Left, 0)-heightHelper(root.Right, 0)) < 2 && isBalances(root.Left) && isBalances(root.Right)
}

func heightHelper(root *TreeNode, height int) int {
	if root == nil {
		return height
	}
	return max(heightHelper(root.Left, height+1), heightHelper(root.Right, height+1))
}

func abs(a int) int {
	if a < 0 {
		return 0 - a
	}
	return a
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {
	root := &TreeNode{
		Value: 3,
	}

	root.Right = NewTreeNode(20)
	root.Left = NewTreeNode(9)

	root.Right.Left = NewTreeNode(15)
	root.Right.Right = NewTreeNode(7)
	root.Right.Right.Right = NewTreeNode(70)
	root.Right.Right.Right.Right = NewTreeNode(700)
	root.Right.Right.Right.Right.Right = NewTreeNode(7000)

	root.reverse()

	fmt.Println()
	fmt.Println(isBalances(root))
}
