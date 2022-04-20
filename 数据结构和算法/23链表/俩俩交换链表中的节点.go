package _3链表

func swapPairs(head *ListNode) *ListNode {

	p := new(ListNode)
	q := p
	for head != nil && head.Next != nil {
		a := head
		b := head.Next
		head = b.Next
		a.Next = nil
		b.Next = a
		p.Next = b
		p = a
	}
	if head != nil {
		p.Next = head
	}
	return q.Next
}
