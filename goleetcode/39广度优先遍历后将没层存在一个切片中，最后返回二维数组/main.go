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
	//var res [][]int
	res := [][]int{}
	if root == nil {
		return res
	}

	var q []*TreeNode
	q = append(q, root)

	for len(q) != 0 {
		length := len(q)
		println(length)  //长度是3，代表整个树的深度是3
		level := []int{} //存放每个层节点上的值

		for i := 0; i < length; i++ {
			node := q[0] //取第一层
			level = append(level, node.Value)

			q = q[1:]

			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}

		//println(level)
		if len(level) > 0 {
			res = append(res, level) //stack 将一维切片里面的数据存放进二位切片中
		}
	}

	return res
}

func main() {
	root := &TreeNode{
		Value: 8,
	}

	root.Right = NewTreeNode(6)
	root.Left = NewTreeNode(5)

	root.Left.Right = NewTreeNode(2)
	root.Left.Left = NewTreeNode(1)

	root.Right.Left = NewTreeNode(3)

	//root.reverse()

	fmt.Println("===========")
	fmt.Println(width(root))
}
