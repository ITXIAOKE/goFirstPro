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

//广度优先遍历
func (node *TreeNode) widthSearch() {
	if node == nil {
		return
	}
	result := []int{}
	nodes := []*TreeNode{node}
	for len(nodes) > 0 {
		curNode := nodes[0]
		nodes = nodes[1:]
		result = append(result, curNode.Value)

		if curNode.Left != nil {
			nodes = append(nodes, curNode.Left)
		}
		if curNode.Right != nil {
			nodes = append(nodes, curNode.Right)
		}
	}
	for _, v := range result {
		fmt.Print(v, " ")
	}
}

//返回2叉树的高度

func (node *TreeNode) GetHight(n *TreeNode) int {
	if n == nil {
		return 0
	} else {
		l := node.GetHight(n.Left)
		r := node.GetHight(n.Right)
		if l > r {
			return l + 1
		} else {
			return r + 1
		}
	}
}

func main() {
	root := &TreeNode{ //4 27  13
		Value: 4,
		//Left: &TreeNode{
		//	Value: 2,
		//	Left:  NewTreeNode(1),
		//	Right: NewTreeNode(3),
		//},
		//Right: &TreeNode{
		//	Value: 7,
		//},
	}
	root.Left = NewTreeNode(2)
	root.Left.Left = NewTreeNode(1)
	root.Left.Right = NewTreeNode(3)

	root.Right = NewTreeNode(7)

	root.reverse()
	root.widthSearch()
	fmt.Println()
	fmt.Println(root.GetHight(root))

}
