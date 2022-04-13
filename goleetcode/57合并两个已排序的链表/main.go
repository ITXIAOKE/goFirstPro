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

func mergeTwoLinkedList(l1 *LinkListNode, l2 *LinkListNode) *LinkListNode {
	pre := &LinkListNode{0, nil}
	cur := pre
	for l1 != nil && l2 != nil {
		temp := l1.Val
		if l1.Val > l2.Val {
			temp = l2.Val
		}
		cur.Next = &LinkListNode{temp, nil}
		cur = cur.Next
		if l1.Val > l2.Val {
			l2 = l2.Next
		} else {
			l1 = l1.Next
		}
	}
	if l1 != nil {
		cur.Next = l1
	}
	if l2 != nil {
		cur.Next = l2
	}
	return pre.Next
}

func main() {
	linkList1 := &LinkListNode{
		Val: 1,
	}
	linkList1.Next = NewLinkListNode(2)
	linkList1.Next.Next = NewLinkListNode(4)

	linkList2 := &LinkListNode{
		Val: 1,
	}
	linkList2.Next = NewLinkListNode(3)
	linkList2.Next.Next = NewLinkListNode(4)

	h1 := mergeTwoLinkedList(linkList1, linkList2)
	fmt.Println(h1.Val)
	h2 := h1.Next
	fmt.Println(h2.Val)

	h3 := h2.Next
	fmt.Println(h3.Val)
	h4 := h3.Next
	fmt.Println(h4.Val)

	h5 := h4.Next
	fmt.Println(h5.Val)
	h6 := h5.Next
	fmt.Println(h6.Val)

}
