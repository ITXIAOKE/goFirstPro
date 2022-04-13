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

func mergeTwoLinkedListSum(l1 *LinkListNode, l2 *LinkListNode) *LinkListNode {
	pre := &LinkListNode{0, nil}
	cur := pre
	carry := 0
	for l1 != nil && l2 != nil {
		a, b := 0, 0
		a = l1.Val
		b = l2.Val

		sum := a + b + carry //2+8==1
		carry = sum / 10
		sum = sum % 10

		cur.Next = &LinkListNode{sum, nil}
		cur = cur.Next

		l1 = l1.Next
		l2 = l2.Next
	}
	if carry == 1 { //99+56-->6+9=15-->最后一位是1
		cur.Next = &LinkListNode{carry, nil}
	}
	return pre.Next
}

func main() {
	linkList1 := &LinkListNode{
		Val: 4,
	}
	linkList1.Next = NewLinkListNode(5)
	linkList1.Next.Next = NewLinkListNode(6)

	linkList2 := &LinkListNode{
		Val: 9,
	}
	linkList2.Next = NewLinkListNode(9)
	linkList2.Next.Next = NewLinkListNode(9)

	sum := mergeTwoLinkedListSum(linkList1, linkList2)
	fmt.Println(sum.Val)
	sum1 := sum.Next
	fmt.Println(sum1.Val)
	sum2 := sum1.Next
	fmt.Println(sum2.Val)

}
