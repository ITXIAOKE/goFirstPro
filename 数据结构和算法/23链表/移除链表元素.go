package _3链表

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {

	p := &ListNode{}

	p.Next = head
	q := p

	for p != nil && p.Next != nil {
		if p.Next.Val == val {
			p.Next = p.Next.Next
		} else {
			p = p.Next
		}

	}
	return q.Next

}
