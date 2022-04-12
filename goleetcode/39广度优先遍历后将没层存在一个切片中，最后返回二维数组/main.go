package main

import "fmt"

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

func width(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}

	var q []*TreeNode
	q = append(q, root)

	for len(q) != 0 {
		length := len(q)
		level := []int{}
		for i := 0; i < length; i++ {
			node := q[0]
			q = q[1:]
			level = append(level, node.Value)

			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		if len(level) > 0 {
			res = append(res, level) //stack
		}
	}

	result := [][]int{}
	//for i := len(res) - 1; i >= 0; i-- {//倒叙打印，从上往上
	//	result = append(result, res[i])
	//}

	for i := 0; i < len(res); i++ {//从上到下，
		result = append(result, res[i])
	}
	return result
}

func main() {
	root := &TreeNode{
		Value: 7,
	}

	root.Right = NewTreeNode(5)
	root.Left = NewTreeNode(4)

	root.Left.Right = NewTreeNode(2)
	root.Left.Left = NewTreeNode(1)

	root.Right.Left = NewTreeNode(3)

	root.reverse()

	fmt.Println()
	fmt.Println(width(root))
}
