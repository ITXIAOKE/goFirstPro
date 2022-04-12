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

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Value == val { //改节点的根的值与期望的值相等，那么就返回整个节点
		return root
	} else if root.Value < val {
		return searchBST(root.Right, val)
	} else if root.Value > val {
		return searchBST(root.Left, val)
	}
	return nil

}

func main() {
	root := &TreeNode{
		Value: 4,
	}
	root.Left = NewTreeNode(2)
	root.Left.Left = NewTreeNode(1)
	root.Left.Right = NewTreeNode(3)

	root.Right = NewTreeNode(7)

	root.reverse()

	newNode := searchBST(root, 2)
	newNode.reverse()

}
