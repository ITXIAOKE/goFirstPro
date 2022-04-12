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

func (node *TreeNode) SumofLeftLeaves() int {
	res := 0
	if node == nil {
		return res
	}
	SumofLeftLeavesHelper(node, &res)
	return res
}

func SumofLeftLeavesHelper(node *TreeNode, res *int) {
	if node == nil {
		return
	}
	if node.Left != nil && node.Left.Left == nil && node.Left.Right == nil {
		*res += node.Left.Value
	}
	SumofLeftLeavesHelper(node.Left, res)
	SumofLeftLeavesHelper(node.Right, res)
}

func main() {
	root := &TreeNode{ //4 27  13
		Value: 4,
	}
	root.Left = NewTreeNode(2)
	root.Left.Left = NewTreeNode(1)
	root.Left.Right = NewTreeNode(3)

	root.Right = NewTreeNode(7)
	root.Right.Left = NewTreeNode(8)
	root.Right.Left.Left = NewTreeNode(10)

	root.reverse()
	fmt.Println()
	fmt.Println(root.SumofLeftLeaves())

}
