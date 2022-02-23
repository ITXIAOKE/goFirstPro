package main

import "fmt"

// Node 二叉树节点，采用链表的形式表示
type Node struct {
	Left  *Node
	Data  interface{}
	Right *Node
}

/**
 * 接口定义包括：
 * 新节点的创建、初始化，
 * 二叉树的输出、度的计算、叶子结点的统计，
 * 二叉树的前中后序遍历、广度遍历
 */

// Initer 接口提供了SetData()方法可以对节点的Data字段进行初始化
type Initer interface {
	SetData(data interface{})
}

// Operater 对生成的二叉树，接口提供了PrintBT(), Depth()和LeafCount()对二叉树进行输出、深度计算和叶子统计
type Operater interface {
	PrintBT()
	Depth() int
	LeafCount() int
}

// Order 接口提供了2种二叉树遍历的方式
// 深度遍历：先序遍历、中序遍历、后序遍历
// 广度遍历
type Order interface {
	PreOrder()      // 先序遍历
	InOrder()       // 中序遍历
	PostOrder()     // 后序遍历
	BreadthTravel() // 广度遍历
}

/**
 * 接口方法的实现，通过底层函数实现
 */

func (n *Node) SetData(data interface{}) { n.Data = data }
func (n *Node) PrintBT()                 { PrintBT(n) }
func (n *Node) Depth() int               { return Depth(n) }
func (n *Node) LeafCount() int           { return LeafCount(n) }
func (n *Node) PreOrder()                { PreOrder(n) }
func (n *Node) InOrder()                 { InOrder(n) }
func (n *Node) PostOrder()               { PostOrder(n) }
func (n *Node) BreadthTravel()           { BreadthTravel(n) }

/**
 * 底层函数的实现
 */

// NewNode 创建一个新的节点
func NewNode(left, right *Node) *Node {
	return &Node{left, nil, right}
}

// PrintBT 用于输出一个给定二叉树的嵌套括号表示，采用递归的算法
// 根节点-->左子树-->右子树
func PrintBT(n *Node) {
	if n != nil {
		fmt.Printf("%v ", n.Data)
		if n.Left != nil || n.Right != nil {
			fmt.Printf("(")
			PrintBT(n.Left)
			if n.Right != nil {
				fmt.Printf(",")
			}
			PrintBT(n.Right)
			fmt.Printf(")")
		}
	}
}

// Depth 用于计算二叉树的深度，采用递归算法
// 若一个二叉树为空，则其深度为0；
// 否则其深度等于左子树或右子树的最大深度加1
func Depth(n *Node) int {
	var depthleft, depthright int
	if n == nil {
		return 0
	}
	depthleft = Depth(n.Left)
	depthright = Depth(n.Right)
	if depthleft > depthright {
		return depthleft + 1
	}
	return depthright + 1
}

// LeafCount 用于统计一个二叉树叶子节点数，采用递归算法
// 若一个二叉树为空，则其叶子节点数为 0；
// 若一棵二叉树的左右字数均为空，则其叶子节点数为 1；
// 否则叶子节点数等于左子树和右子树叶子节点总数之和
func LeafCount(n *Node) int {
	if n == nil {
		return 0
	} else if n.Left == nil && n.Right == nil {
		return 1
	} else {
		return LeafCount(n.Left) + LeafCount(n.Right)
	}
}

// PreOrder 先序遍历，采用递归算法
// 根节点-->左子树-->右子树
func PreOrder(n *Node) {
	if n != nil {
		fmt.Printf("%v ", n.Data)
		PreOrder(n.Left)
		PreOrder(n.Right)
	}
}

// InOrder 中序遍历，采用递归算法
// 左子树-->根节点-->右子树
func InOrder(n *Node) {
	if n != nil {
		InOrder(n.Left)
		fmt.Printf("%v ", n.Data)
		InOrder(n.Right)
	}
}

// PostOrder 后序遍历，采用递归算法
// 左子树-->右子树-->根节点
func PostOrder(n *Node) {
	if n != nil {
		PostOrder(n.Left)
		PostOrder(n.Right)
		fmt.Printf("%v ", n.Data)
	}
}

// BreadthTravel 广度遍历
func BreadthTravel(n *Node) {
	var queue []*Node
	queue = []*Node{n}

	for len(queue) > 0 {
		root := queue[0]
		fmt.Printf("%v ", root.Data)
		queue = queue[1:]
		if root.Left != nil {
			queue = append(queue, root.Left)
		}
		if root.Right != nil {
			queue = append(queue, root.Right)
		}
	}
}

func main() {
	//创建二叉树
	root := NewNode(nil, nil)
	root.SetData("root")

	a := NewNode(nil, nil)
	a.SetData("left")

	al := NewNode(nil, nil)
	al.SetData(100)

	ar := NewNode(nil, nil)
	ar.SetData(3.14)

	a.Left = al
	a.Right = ar

	b := NewNode(nil, nil)
	b.SetData("right")

	root.Left = a
	root.Right = b

	root.PrintBT()
	fmt.Println()

	// 使用 Order 接口实现对二叉树的基本操作
	var it Order  // 定义接口类型
	it = root     // 将*Node类型变量赋值给接口，*Node实现了接口的所有方法
	it.PreOrder() // 先序遍历
	fmt.Println()
	it.InOrder() // 中序遍历
	fmt.Println()
	it.PostOrder() // 后序遍历
	fmt.Println()
	it.BreadthTravel() // 广度遍历
}
