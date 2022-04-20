package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (head *ListNode) reverseListNode() *ListNode {
	cur := head

	var pre *ListNode = nil
	for cur != nil {
		pre, cur, cur.Next = cur, cur.Next, pre
	}
	return pre
}

func main() {
	head := new(ListNode)
	head.Val = 1
	ln2 := new(ListNode)
	ln2.Val = 2
	ln3 := new(ListNode)
	ln3.Val = 3
	ln4 := new(ListNode)
	ln4.Val = 4
	ln5 := new(ListNode)
	ln5.Val = 5

	head.Next = ln2
	ln2.Next = ln3
	ln3.Next = ln4
	ln4.Next = ln5

	listNode := head.reverseListNode()
	fmt.Println(listNode.Val)
	fmt.Println(listNode.Next.Val)
	fmt.Println(listNode.Next.Next.Val)

}
