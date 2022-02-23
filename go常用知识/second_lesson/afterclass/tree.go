package main

import "fmt"

// 定义节点属性
type Node2 struct {
	item  int
	left  *Node2
	right *Node2
}

// 二叉树
type Tree struct {
	root *Node2
}

// 生成二叉树
func (tree *Tree) createTree(value int) {
	node := &Node2{item: value}
	fmt.Println("增加节点", value)
	if tree.root == nil {
		tree.root = node
		return
	} else {
		var queue []*Node2 = []*Node2{tree.root}
		for len(queue) != 0 {
			cur := queue[0]
			queue = queue[1:]
			if cur.left == nil {
				cur.left = node
				return
			} else if cur.right == nil {
				cur.right = node
				return
			} else {
				queue = append(queue, cur.left)
				queue = append(queue, cur.right)
			}
		}
	}
}




// 层次遍历：即是广度优先方式搜索，使用队列方式实现
func (tree *Tree) selectTree() {
	if tree.root == nil {
		fmt.Println("not root")
		return
	}
	var queue []*Node2 = []*Node2{tree.root}
	for len(queue) != 0 {
		cur := queue[0]
		queue = queue[1:]
		fmt.Printf("%v ", cur.item)
		if cur.left != nil {
			queue = append(queue, cur.left)
		}
		if cur.right != nil {
			queue = append(queue, cur.right)
		}
	}

}

// 先序遍历：即是深度优先方式搜索，使用递归方式实现
// 根节点 --> 左节点 --> 右节点
func (tree *Tree) selectFront(node *Node2) {
	if node == nil {
		return
	}
	fmt.Printf("%v ", node.item)
	tree.selectFront(node.left)
	tree.selectFront(node.right)
}

// 中序遍历：即是深度优先方式搜索，使用递归方式实现
// 左节点 --> 根节点 --> 右节点
func (tree *Tree) selectMid(node *Node2) {
	if node == nil {
		return
	}
	tree.selectMid(node.left)
	fmt.Printf("%v ", node.item)
	tree.selectMid(node.right)
}

// 后序遍历：即是深度优先方式搜索，使用递归方式实现
// 左节点 --> 右节点 --> 根节点
func (tree *Tree) selectBack(node *Node2) {
	if node == nil {
		return
	}
	tree.selectBack(node.left)
	tree.selectBack(node.right)
	fmt.Printf("%v ", node.item)
}

func main() {
	var t Tree
	t.createTree(2)
	t.createTree(10)
	t.createTree(3)
	t.createTree(5)
	// 层次遍历
	fmt.Println("层次遍历")
	t.selectTree()
	fmt.Printf("\n")
	// 先序遍历
	fmt.Println("先序遍历")
	t.selectFront(t.root)
	fmt.Printf("\n")
	// 中序遍历
	fmt.Println("中序遍历")
	t.selectMid(t.root)
	fmt.Printf("\n")
	// 后序遍历
	fmt.Println("后序遍历")
	t.selectBack(t.root)
	fmt.Printf("\n")
}