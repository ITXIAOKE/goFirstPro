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

//找祖先

func lowestCommonAncestor(root, p, q *TreeNode) int {
	if root.Value > p.Value && root.Value > q.Value {
		return lowestCommonAncestor(root.Left, p, q)
	}

	if root.Value < p.Value && root.Value < q.Value {
		return lowestCommonAncestor(root.Right, p, q)
	}
	return root.Value //返回祖先的值
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
	fmt.Println(lowestCommonAncestor(root, root.Left, root.Right))
}
