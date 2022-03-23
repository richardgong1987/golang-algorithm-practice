package linked

func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// 虚拟头结点
	dummy := &ListNode{
		Val: -1,
	}
	p, p1, p2 := dummy, l1, l2
	for p1 != nil && p2 != nil {
		// 比较 p1 和 p2 两个指针
		// 将值较小的的节点接到 p 指针
		if p1.Val > p2.Val {
			p.Next = p2
			p2 = p2.Next
		} else {
			p.Next = p1
			p1 = p1.Next
		}
		// p 指针不断前进
		p = p.Next
	}
	if p1 != nil {
		p.Next = p1
	}
	if p2 != nil {
		p.Next = p2
	}
	return dummy.Next
}
