package main

import "fmt"

type NodeImpl interface {
	SetValue(value int)
	Traverse() //遍历方法
}

type TreeNode struct {
	Value     int
	LeftNode  *TreeNode
	RightNode *TreeNode
}

func NewTreeNode(value int) *TreeNode {
	return &TreeNode{
		Value: value,
	}
}

func (tree *TreeNode) SetValue(value int) {
	tree.Value = value
}

func (tree *TreeNode) Traverse() {
	if tree == nil {
		return
	}

	fmt.Println(tree.Value) //前序遍历  根左右
	tree.LeftNode.Traverse()
	//fmt.Println(tree.Value) //中  左根右
	tree.RightNode.Traverse()
	//fmt.Println(tree.Value) //后  左右根

}

func main() {

	root := new(TreeNode)
	root.SetValue(3)

	root.LeftNode = NewTreeNode(2)
	root.RightNode = NewTreeNode(4)

	root.Traverse()

}
