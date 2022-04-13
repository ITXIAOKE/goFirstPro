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

func middleNode(head *LinkListNode) *LinkListNode {
	if head == nil {
		return nil
	}

	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil { //奇数
		slow = slow.Next
		fast = fast.Next.Next
	}

	if fast.Next != nil && fast.Next.Next == nil { //偶数
		slow = slow.Next
	}

	return slow
}

func main() {
	linkList := &LinkListNode{
		Val: 1,
	}
	linkList.Next = NewLinkListNode(2)
	linkList.Next.Next = NewLinkListNode(3)
	linkList.Next.Next.Next = NewLinkListNode(4)
	linkList.Next.Next.Next.Next = NewLinkListNode(5)

	fmt.Println(middleNode(linkList))

}
