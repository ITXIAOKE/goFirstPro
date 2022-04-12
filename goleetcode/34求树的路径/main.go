package main

import (
	"fmt"
	"strconv"
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


func binaryTreePaths(root *TreeNode) []string {
	res := []string{}
	if root == nil {
		return res
	}
	helper(root, "", &res)
	return res
}

func helper(root *TreeNode, path string, res *[]string) {
	path += strconv.Itoa(root.Value)

	if root.Left == nil && root.Right == nil {
		*res = append(*res, path)
		return
	}

	path += "->"
	if root.Left != nil {
		helper(root.Left, path, res)
	}
	if root.Right != nil {
		helper(root.Right, path, res)
	}

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
	fmt.Println(binaryTreePaths(root))
}
