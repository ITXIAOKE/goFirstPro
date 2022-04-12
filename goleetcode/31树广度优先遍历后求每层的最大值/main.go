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

func (root *TreeNode) largestValues() []int {
	res := []int{}
	if root == nil {
		return res
	}
	var q []*TreeNode
	q = append(q, root)
	for len(q) > 0 {
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
		maxValue := max(level)
		res = append(res, maxValue)
	}
	return res
}

func max(level []int) int {
	if len(level) == 0 {
		return -1
	}
	res := level[0]
	for _, v := range level {
		if v > res {
			res = v
		}
	}
	return res
}

func main() {
	root := &TreeNode{ //4 27  13
		Value: 4,
	}
	root.Left = NewTreeNode(2)
	root.Left.Left = NewTreeNode(1)
	root.Left.Right = NewTreeNode(3)

	root.Right = NewTreeNode(7)
	//root.Right.Left = NewTreeNode(8)
	//root.Right.Left.Left = NewTreeNode(10)

	root.reverse()
	fmt.Println()
	fmt.Println(root.largestValues())

}
