package _3链表

/**
快慢指针，fast 与 slow。它们起始都位于链表的头部。随后，slow 指针每次向后移动一个位置，而 fast 指针向后移动两个位置。
如果链表中存在环，则fast 指针最终将再次与 slow 指针在环中相遇。

因此，当发现 slow 与 fast 相遇时，我们再额外使用一个指针ptr。起始，它指向链表头部；
随后，它和 slow 每次向后移动一个位置。最终，它们会在入环点相遇。
*/
func detectCycle(head *ListNode) *ListNode {

	p, q := head, head
	for q != nil && q.Next != nil {
		p = p.Next
		q = q.Next.Next
		if p == q {
			for head != q {
				head = head.Next
				q = q.Next
			}
			return head
		}
	}
	return nil
}
