package main

import "fmt"

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

func (root *TreeNode) findMode() []int {
	res := []int{}
	if root == nil {
		return res
	}
	dic := make(map[int]int)
	inorder(root, &dic)
	max := 0
	for _, v := range dic {
		if v > max {
			max = v
		}
	}
	for k, v := range dic {
		if v == max {
			res = append(res, k)
		}
	}
	return res
}

func inorder(root *TreeNode, dic *map[int]int) {
	if root == nil {
		return
	}
	inorder(root.Left, dic)
	(*dic)[root.Value]++
	inorder(root.Right, dic)
}

func main() {
	root := &TreeNode{ //4 27  13
		Value: 1,
	}

	root.Right = NewTreeNode(2)
	root.Left = NewTreeNode(33)
	root.Left.Right = NewTreeNode(33)
	root.Right.Left = NewTreeNode(2)
	root.Right.Right = NewTreeNode(33)

	root.reverse()
	fmt.Println()

	fmt.Println(root.findMode())

}
