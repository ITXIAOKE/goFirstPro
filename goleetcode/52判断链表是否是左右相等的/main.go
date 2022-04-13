package main

import (
	"fmt"
)

type LinkListNode struct {
	Val  int
	Next *LinkListNode
}

func NewLinkListNode(value int) *LinkListNode {
	return &LinkListNode{
		Val: value,
	}
}

func isPalindrome(head *LinkListNode) bool {
	//corner case
	if head == nil {
		return true
	}

	//find middle ,and reverse left  part
	slow, fast := head, head
	var pre *LinkListNode
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next

		next := slow.Next //保存头节点的下一个节点2
		slow.Next = pre   //将头节点指向前一个节点
		pre = slow        //更新前一个节点
		slow = next       //更新头节点
	}

	//define left and right pointers
	var left *LinkListNode
	var right *LinkListNode
	if fast.Next == nil {
		left = pre
		right = slow.Next
	} else {
		right = slow.Next
		slow.Next = pre
		left = slow
	}

	//compare
	for left != nil && right != nil {
		if left.Val != right.Val {
			return false
		}
		left = left.Next
		right = right.Next
	}

	return true
}

func main() {
	linkList := &LinkListNode{
		Val: 1,
	}
	linkList.Next = NewLinkListNode(2)
	linkList.Next.Next = NewLinkListNode(3)
	linkList.Next.Next.Next = NewLinkListNode(4)
	//fmt.Println(linkList.Val)
	//fmt.Println(linkList.Next.Val)
	//fmt.Println(linkList.Next.Next.Val)
	//fmt.Println(linkList.Next.Next.Next.Val)

	//fmt.Println(isPalindrome(linkList))
	head4 := linkList.reverseList()
	//head4 := linkList.reverseList2()
	fmt.Println(head4.Val)
	head3 := head4.Next
	fmt.Println(head3.Val)
	head2 := head3.Next
	fmt.Println(head2.Val)
	head1 := head2.Next
	fmt.Println(head1.Val)

}

//反转链表的实现
func (head *LinkListNode) reverseList() *LinkListNode {
	cur := head
	var pre *LinkListNode
	for cur != nil {
		pre, cur, cur.Next = cur, cur.Next, pre //这句话最重要
	}
	return pre
}


//反转链表的实现
func (head *LinkListNode) reverseList2() *LinkListNode {
	if head==nil{
		return nil
	}
	var pre *LinkListNode
	cur:=head
	for cur != nil {
		next := cur.Next //保存头节点的下一个节点
		cur.Next = pre   //将头节点指向前一个节点
		pre = cur        //更新前一个节点
		cur = next       //更新头节点
	}
	return pre
}
