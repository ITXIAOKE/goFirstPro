package main

import "fmt"

type Hero struct {
	No    int
	Name  string
	Left  *Hero
	Right *Hero
}

func main() {

	//构建一个二叉树
	root := &Hero{
		No:   1,
		Name: "宋江",
	}

	left1 := &Hero{
		No:   2,
		Name: "吴用",
	}

	right1 := &Hero{
		No:   3,
		Name: "卢俊义",
	}

	root.Left = left1
	root.Right = right1

	right2 := &Hero{
		No:   4,
		Name: "林冲",
	}
	right1.Right = right2

	//吴用  宋江  卢俊义 林冲
	//PreOrder(root)
	//InfixOrder(root)
	PostOrder(root)

}

//前序遍历 根 左 右

func PreOrder(node *Hero) {
	if node != nil {
		fmt.Printf("no=%d name=%s\n", node.No, node.Name)
		PreOrder(node.Left)
		PreOrder(node.Right)
	}
}

//中序遍历 左 根 右

func InfixOrder(node *Hero) {
	if node != nil {
		InfixOrder(node.Left)
		fmt.Printf("no=%d name=%s\n", node.No, node.Name)
		InfixOrder(node.Right)
	}
}

//后序遍历 左 右 根

func PostOrder(node *Hero) {
	if node != nil {
		PostOrder(node.Left)
		PostOrder(node.Right)
		fmt.Printf("no=%d name=%s\n", node.No, node.Name)
	}
}
