package main

import "fmt"

type CatNode struct {
	no   int
	name string
	next *CatNode
}

func InsertCatNode(head *CatNode, newCatNode *CatNode) {
	if head.next == nil {
		head.no = newCatNode.no
		head.name = newCatNode.name
		head.next = head //构成一个环形
		fmt.Println(newCatNode, "加入到环形的链表")
		return
	}
	temp := head //临时变量，帮忙找到环形的最后节点
	for {
		if temp.next == head {
			break
		}
		temp = temp.next
	}
	//加入到链表中
	temp.next = newCatNode
	newCatNode.next = head
}

func ListCircleLink(head *CatNode) {
	fmt.Println("环形链表的基本情况如下：")
	temp := head
	if temp.next == nil {
		fmt.Println("空空如也的环形链表")
		return
	}
	for {
		fmt.Printf("猫的信息为=[id=%d name=%s]-->\n", temp.no, temp.name)
		if temp.next == head {
			break
		}
		temp = temp.next
	}
}

//删除节点

func DelCatNode(head *CatNode, id int) *CatNode {
	temp := head
	helper := head
	if temp.next == nil {
		fmt.Println("这个是一个空的环形链表，不能删除")
		return head
	}
	if temp.next == head { //只有一个节点
		temp.next = nil
		return head
	}

	//将helper定位到环形链表的最后
	for {
		if helper.next == head {
			break
		}
		helper = helper.next
	}

	flag := true
	for {
		if temp.next == head { //比较到最后一个
			break
		}
		if temp.no == id { //找到
			if temp == head { //删除的是头节点
				head = head.next
			}
			helper.next = temp.next
			fmt.Printf("猫猫= %d\n", id)
			flag = false
			break
		}
		temp = temp.next     //移动---比较
		helper = helper.next //移动--找到删除
	}
	if flag { //上面没有删除
		if temp.no == id {
			helper.next = temp.next
			fmt.Printf("猫猫=%d \n", id)
		} else {
			fmt.Printf("对不起，没有no=%d\n", id)
		}
	}
	return head
}

func main() {
	head := &CatNode{}
	cat1 := &CatNode{
		no:   1,
		name: "tome",
	}

	cat2 := &CatNode{
		no:   2,
		name: "jimy",
	}

	cat3 := &CatNode{
		no:   3,
		name: "xixk",
	}
	InsertCatNode(head, cat1)
	InsertCatNode(head, cat2)
	InsertCatNode(head, cat3)
	ListCircleLink(head)

	head1:=DelCatNode(head, 30)
	ListCircleLink(head1)
}
