package _3链表

/**
使用哈希表，将其中一个链表的节点不断添加进去
遍历第二个链表，如果在哈希表中找到了节点，则返回节点的数据值
*/

func getIntersectionNode(headA, headB *ListNode) *ListNode {

	h := make(map[*ListNode]int, 0)

	for headA != nil {
		h[headA]++
		headA = headA.Next
	}

	for headB != nil {
		if h[headB] > 0 {
			return headB
		} else {
			headB = headB.Next
		}
	}

	return nil

}
