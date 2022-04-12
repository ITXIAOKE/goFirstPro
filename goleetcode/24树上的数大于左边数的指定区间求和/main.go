package main

import (
	"fmt"
)

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(val int) *TreeNode {
	return &TreeNode{
		Value: val,
	}
}

func rangeSumBST(root *TreeNode, L int, R int) int {
	arr := []int{}
	inorder(root, &arr)
	res := 0
	for _, v := range arr {
		if v >= L && v <= R {
			res += v
		}
	}
	return res
}

func inorder(root *TreeNode, arr *[]int) {
	if root == nil {
		return
	}
	inorder(root.Left, arr)
	*arr = append(*arr, root.Value)
	inorder(root.Right, arr)
}

func main() {
	//root:=new(TreeNode)
	root := &TreeNode{
		Value: 10,
	}
	root.Left = NewTreeNode(5)
	root.Right = NewTreeNode(15)
	root.Left.Left = NewTreeNode(3)
	root.Left.Right = NewTreeNode(7)
	root.Right.Right = NewTreeNode(18)

	//root.MyTraverse()
	fmt.Println(rangeSumBST(root, 7, 15))
}

func (tree *TreeNode) MyTraverse() {
	if tree == nil {
		return
	}
	fmt.Println(tree.Value) //前序遍历  根左右
	tree.Left.MyTraverse()
	//fmt.Println(tree.Value) //中  左根右
	tree.Right.MyTraverse()
	//fmt.Println(tree.Value) //后  左右根

}
