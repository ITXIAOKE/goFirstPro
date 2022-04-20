package _3链表

func removeNthFromEnd(head *ListNode, n int) *ListNode {

	p := &ListNode{}
	p.Next = head
	q := p
	t := p
	for i := 0; i < n; i++ {
		q = q.Next
	}

	for q.Next != nil {
		q = q.Next
		p = p.Next
	}
	p.Next = p.Next.Next

	return t.Next
}
