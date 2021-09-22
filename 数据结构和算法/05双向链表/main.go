package main

import (
	"fmt"
)

type HeroNode struct {
	no       int
	name     string
	nickname string
	next     *HeroNode //指向下一个结点
	pre      *HeroNode //指向前一个结点
}

//给双向链表插入一个结点
//编写第一种插入方法，在双向链表的最后插入

func InsertHeroNode(head *HeroNode, newHeroNode *HeroNode) {
	temp := head
	for {
		if temp.next == nil { //找到最后
			break
		}
		temp = temp.next //让temp不断的指向下一个结点
	}
	//将newHeroNode加入到链表的最后
	temp.next = newHeroNode
	newHeroNode.pre = temp
}

//第二中插入方法，根据no的编号从小到大进行插入

func InsertHeroNode2(head *HeroNode, newHeroNode *HeroNode) {
	temp := head
	flagMy := true
	for {
		if temp.next == nil {
			break
		} else if temp.next.no >= newHeroNode.no { //按照no从小到大排序
			//说明newheronode应该插入到temp后面
			break
		} else if temp.next.no == newHeroNode.no {
			flagMy = false
			break
		}
		temp = temp.next
	}

	if !flagMy {
		fmt.Println("对不起，已经存在no=", newHeroNode.no)
		return
	} else {
		newHeroNode.next = temp.next
		newHeroNode.pre = temp
		if temp.next != nil {
			temp.next.pre = newHeroNode
		}
		temp.next = newHeroNode
	}

}

//显示链表的所有结点信息

func ListHeroNode(head *HeroNode) {
	temp := head

	if temp.next == nil {
		fmt.Println("空链表")
		return
	}
	for { //遍历这个链表
		fmt.Printf("链表信息如下：【%d，%s，%s】==>", temp.next.no, temp.next.name, temp.next.nickname)
		temp = temp.next
		if temp.next == nil {
			break
		}
	}
}

//逆序，遍历双向链表

func ListHeroNode2(head *HeroNode) {
	temp := head

	if temp.next == nil {
		fmt.Println("空链表...")
		return
	}

	for {
		if temp.next == nil {
			break
		}
		temp = temp.next
	}

	for { //遍历这个链表
		fmt.Printf("链表信息如下：【%d，%s，%s】==>", temp.no, temp.name, temp.nickname)
		temp = temp.pre

		if temp.pre == nil {
			break
		}
	}
}

//删除一个结点

func DelHeroNode(head *HeroNode, id int) {
	temp := head
	flag := false
	for {
		if temp.next == nil {
			break
		} else if temp.next.no == id {
			flag = true
			break
		}
		temp = temp.next
	}
	if flag {
		temp.next = temp.next.next
		if temp.next != nil {
			temp.next.pre = temp
		}
	} else {
		fmt.Println("sorry,要删除的id不存在")
	}
}

func main() {
	/**
	双向链表：
	*/
	//1,创建一个头结点
	head := &HeroNode{}
	//2,创建一个新的heronode
	hero1 := &HeroNode{
		no:       1,
		name:     "宋江",
		nickname: "及时雨",
	}

	hero2 := &HeroNode{
		no:       2,
		name:     "吕俊逸",
		nickname: "玉麒麟",
	}

	hero3 := &HeroNode{
		no:       3,
		name:     "林冲",
		nickname: "豹子头",
	}

	InsertHeroNode(head, hero1)
	InsertHeroNode(head, hero2)
	InsertHeroNode(head, hero3)

	fmt.Println("从前往后打印如下：---------")
	ListHeroNode(head)

	fmt.Println()
	fmt.Println("逆序打印如下:---------")
	ListHeroNode2(head)

	fmt.Println()
	fmt.Println("删除指定节点打印如下:---------")
	DelHeroNode(head, -1)
	ListHeroNode(head)
}
